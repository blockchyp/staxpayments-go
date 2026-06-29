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
