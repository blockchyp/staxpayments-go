// Package core holds the shared Stax Payments client used by the namespace SDK
// modules (e.g. payments, terminals). It lives under internal/ so that only
// packages within github.com/blockchyp/staxpayments-go may construct a client;
// external callers must go through a namespace client such as
// payments.NewClient. The public staxpayments package intentionally exposes no
// client constructor.
package core

import (
	"crypto/tls"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	staxpayments "github.com/blockchyp/staxpayments-go"
)

// terminalCN is the common name on a terminal certificate.
const terminalCN = "blockchyp-terminal"

// Client is the shared Stax Payments client. It carries credentials, host
// configuration, and the HTTP transport used by every namespace client.
type Client struct {
	Credentials         staxpayments.StaxAPICredentials
	bcCredentials       staxpayments.APICredentials
	bcCredentialsExpiry time.Time
	bcCredentialsMu     sync.Mutex
	GatewayHost         string
	DashboardHost       string
	TestGatewayHost     string
	CoreHost            string
	HTTPS               bool
	RouteCache          string

	GatewayTimeout  time.Duration
	TerminalTimeout time.Duration

	routeCacheTTL      time.Duration
	gatewayHTTPClient  *http.Client
	terminalHTTPClient *http.Client
	coreHTTPClient     *http.Client

	LogRequests bool
}

// New returns a default Client configured with the given credentials.
func New(creds staxpayments.StaxAPICredentials) *Client {
	userAgent := staxpayments.BuildUserAgent()

	return &Client{
		Credentials:     creds,
		GatewayHost:     staxpayments.DefaultGatewayHost,
		TestGatewayHost: staxpayments.DefaultTestGatewayHost,
		DashboardHost:   staxpayments.DefaultDashboardHost,
		CoreHost:        staxpayments.DefaultCoreHost,
		HTTPS:           staxpayments.DefaultHTTPS,
		RouteCache:      filepath.Join(os.TempDir(), ".staxpayments_routes"),

		GatewayTimeout:  staxpayments.DefaultGatewayTimeout,
		TerminalTimeout: staxpayments.DefaultTerminalTimeout,

		routeCacheTTL: staxpayments.DefaultRouteCacheTTL,
		gatewayHTTPClient: &http.Client{
			Transport: staxpayments.AddUserAgent(
				&http.Transport{
					Dial: (&net.Dialer{
						Timeout: 5 * time.Second,
					}).Dial,
					TLSHandshakeTimeout: 5 * time.Second,
				},
				userAgent,
			),
		}, // Timeout is set per request
		terminalHTTPClient: &http.Client{
			Transport: staxpayments.AddUserAgent(
				&http.Transport{
					Dial: (&net.Dialer{
						Timeout: 5 * time.Second,
					}).Dial,
					TLSHandshakeTimeout: 5 * time.Second,
					TLSClientConfig: &tls.Config{
						RootCAs:    terminalCertPool(),
						ServerName: terminalCN,
					},
				},
				userAgent,
			),
		},
		coreHTTPClient: &http.Client{
			Transport: staxpayments.AddUserAgent(
				&http.Transport{
					Dial: (&net.Dialer{
						Timeout: 5 * time.Second,
					}).Dial,
					TLSHandshakeTimeout: 5 * time.Second,
				},
				userAgent,
			),
		}, // Timeout is set per request
	}
}

// ExpireRouteCache invalidates the route cache for testing.
func (client *Client) ExpireRouteCache() {
	for key, value := range routeCache {
		value.TTL = time.Now()
		routeCache[key] = value
	}

	offlineCache := client.readOfflineCache()

	if offlineCache != nil {
		for _, route := range offlineCache.Routes {
			route.TTL = time.Now()
			client.updateOfflineCache(&route)
		}
	}
}

// terminalRequestEnvelope wraps a request with the transient credentials used
// to authenticate directly against a terminal. It mirrors the JSON shape of the
// generated Terminal*Request types but is request-model agnostic so the shared
// terminal routing can build it for any namespace.
type terminalRequestEnvelope struct {
	staxpayments.APICredentials
	Request interface{} `json:"request"`
}

// RouteTerminalRequest runs a terminal-routed transaction end to end: it ensures
// valid transient credentials, populates signature options, resolves the
// terminal route, and dispatches to the terminal, cloud relay, or gateway as
// appropriate. It is the shared entry point the namespace clients (e.g. the
// payments client) use for terminal transactions, so the terminal routing lives
// in one place rather than being duplicated per namespace.
func (client *Client) RouteTerminalRequest(terminalName, terminalPath, gatewayPath, method string, request, response interface{}, testTx bool, requestTimeout interface{}) error {
	if err := client.ensure(); err != nil {
		return err
	}

	if err := populateSignatureOptions(request); err != nil {
		return err
	}

	var err error
	if terminalName != "" {
		var route TerminalRoute
		route, err = client.resolveTerminalRoute(terminalName)
		if err != nil {
			return err
		}

		if route.CloudRelayEnabled {
			err = client.RelayRequest(gatewayPath, method, request, response, testTx, requestTimeout)
		} else {
			authRequest := terminalRequestEnvelope{
				APICredentials: route.TransientCredentials,
				Request:        request,
			}
			err = client.terminalRequest(route, terminalPath, method, authRequest, response, requestTimeout)
		}
	} else {
		err = client.GatewayRequest(gatewayPath, method, request, response, testTx, requestTimeout)
	}

	if err != nil {
		return err
	}

	if err := handleSignature(request, response); err != nil {
		log.Printf("Failed to write signature: %+v", err)
	}

	return nil
}

func getTimeout(requestTimeout interface{}, defaultTimeout time.Duration) time.Duration {
	var requestTimeoutDuration time.Duration
	switch v := requestTimeout.(type) {
	case int:
		requestTimeoutDuration = time.Duration(v) * time.Second
	case time.Duration:
		requestTimeoutDuration = v
	case nil:
	default:
		panic("must be int or time.Duration")
	}

	if requestTimeoutDuration <= 0 {
		return defaultTimeout
	}
	return requestTimeoutDuration
}

func populateSignatureOptions(request interface{}) error {
	sigOpts, ok := (staxpayments.SignatureRequest{}).From(request)
	if !ok {
		return nil
	}

	if sigOpts.SigFile != "" && sigOpts.SigFormat == "" {
		x := strings.Split(sigOpts.SigFile, ".")
		sigOpts.SigFormat = staxpayments.SignatureFormat(strings.ToLower(x[len(x)-1]))
	}

	switch sigOpts.SigFormat {
	case staxpayments.SignatureFormatNone, staxpayments.SignatureFormatPNG, staxpayments.SignatureFormatJPG, staxpayments.SignatureFormatGIF:
	default:
		return fmt.Errorf("invalid signature format: %s", sigOpts.SigFormat)
	}

	copyTo(sigOpts, request)

	return nil
}

func handleSignature(request, response interface{}) error {
	requestOpts, ok := (staxpayments.SignatureRequest{}).From(request)
	if !ok {
		return nil
	}

	responseOpts, ok := (staxpayments.SignatureResponse{}).From(response)
	if !ok {
		return nil
	}

	if requestOpts.SigFile == "" || responseOpts.SigFile == "" {
		return nil
	}

	clearField(response, "SigFile")

	content, err := hex.DecodeString(responseOpts.SigFile)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(requestOpts.SigFile, content, 0600)
}
