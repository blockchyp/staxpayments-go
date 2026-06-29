// Copyright 2019-2026 BlockChyp, Inc. All rights reserved. Use of this code
// is governed by a license that can be found in the LICENSE file.
//
// This file was generated automatically by the BlockChyp SDK Generator.
// Changes to this file will be lost every time the code is regenerated.

package staxpayments

import (
	"crypto/tls"
	"encoding/hex"
	"errors"
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
)

// Default client configuration constants.
const (
	DefaultGatewayHost     = "https://api.blockchyp.com"
	DefaultDashboardHost   = "https://dashboard.blockchyp.com"
	DefaultTestGatewayHost = "https://test.blockchyp.com"
	DefaultCoreHost        = "https://apiprod.fattlabs.com"
	DefaultHTTPS           = true
	DefaultRouteCacheTTL   = 60 * time.Minute
	DefaultGatewayTimeout  = 20 * time.Second
	DefaultTerminalTimeout = 2 * time.Minute
)

// Default filesystem configuration.
const (
	ConfigDir  = "staxpayments"
	ConfigFile = "staxpayments.json"
)

// terminalCN is the common name on a terminal certificate.
const terminalCN = "blockchyp-terminal"

// Clientside response constants.
const (
	ResponseUnknownTerminal = "Unknown Terminal"
	ResponseTimedOut        = "Request Timed Out"
)

// ErrInvalidAsyncRequest is returned when a request cannot be called
// asynchronously.
var ErrInvalidAsyncRequest = errors.New("async requests must be terminal requests")

// Version contains the version at build time
var Version string

// Client is the main interface used by application developers.
type Client struct {
	Credentials         StaxAPICredentials
	bcCredentials       APICredentials
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

// NewClient returns a default Client configured with the given credentials.
func NewClient(creds StaxAPICredentials) Client {
	userAgent := BuildUserAgent()

	return Client{
		Credentials:     creds,
		GatewayHost:     DefaultGatewayHost,
		TestGatewayHost: DefaultTestGatewayHost,
		DashboardHost:   DefaultDashboardHost,
		CoreHost:        DefaultCoreHost,
		HTTPS:           DefaultHTTPS,
		RouteCache:      filepath.Join(os.TempDir(), ".staxpayments_routes"),

		GatewayTimeout:  DefaultGatewayTimeout,
		TerminalTimeout: DefaultTerminalTimeout,

		routeCacheTTL: DefaultRouteCacheTTL,
		gatewayHTTPClient: &http.Client{
			Transport: AddUserAgent(
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
			Transport: AddUserAgent(
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
			Transport: AddUserAgent(
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

// ExpireRouteCache invalidates the route cache to for testing.
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
	APICredentials
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
	sigOpts, ok := (SignatureRequest{}).From(request)
	if !ok {
		return nil
	}

	if sigOpts.SigFile != "" && sigOpts.SigFormat == "" {
		x := strings.Split(sigOpts.SigFile, ".")
		sigOpts.SigFormat = SignatureFormat(strings.ToLower(x[len(x)-1]))
	}

	switch sigOpts.SigFormat {
	case SignatureFormatNone, SignatureFormatPNG, SignatureFormatJPG, SignatureFormatGIF:
	default:
		return fmt.Errorf("invalid signature format: %s", sigOpts.SigFormat)
	}

	copyTo(sigOpts, request)

	return nil
}

func handleSignature(request, response interface{}) error {
	requestOpts, ok := (SignatureRequest{}).From(request)
	if !ok {
		return nil
	}

	responseOpts, ok := (SignatureResponse{}).From(response)
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
