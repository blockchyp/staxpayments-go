// Copyright 2019-2026 BlockChyp, Inc. All rights reserved. Use of this code
// is governed by a license that can be found in the LICENSE file.
//
// This file was generated automatically by the BlockChyp SDK Generator.
// Changes to this file will be lost every time the code is regenerated.

// Package payments is the Payment Endpoints namespace of the Stax
// Payments Go SDK. It is its own module so it can be imported and versioned
// independently of the other namespaces; it depends only on the shared core
// module github.com/blockchyp/staxpayments-go.
package payments

import (
	"errors"
	"net"

	staxpayments "github.com/blockchyp/staxpayments-go"
	"github.com/blockchyp/staxpayments-go/internal/core"
)

// Client is the Payment Endpoints client. It wraps the shared core client for
// transport and credentials via an unexported field, so only this namespace's
// own operations are part of its public API (the core transport methods are
// deliberately not exposed).
type Client struct {
	core *core.Client
}

// NewClient returns a Payment Endpoints client configured with the given Stax
// bearer credentials. The shared core client can only be constructed through a
// namespace like this one; the public staxpayments package exposes no client
// constructor.
func NewClient(creds staxpayments.StaxAPICredentials) *Client {
	return &Client{
		core: core.New(creds),
	}
}

// SetGatewayHost overrides the gateway host used for live transactions. Set it
// to point the SDK at a non-production environment. An empty value restores the
// default (https://api.blockchyp.com).
func (client *Client) SetGatewayHost(host string) {
	client.core.GatewayHost = host
}

// GatewayHost returns the gateway host used for live transactions.
func (client *Client) GatewayHost() string {
	return client.core.GatewayHost
}

// SetTestGatewayHost overrides the gateway host used for test transactions. An
// empty value restores the default (https://test.blockchyp.com).
func (client *Client) SetTestGatewayHost(host string) {
	client.core.TestGatewayHost = host
}

// TestGatewayHost returns the gateway host used for test transactions.
func (client *Client) TestGatewayHost() string {
	return client.core.TestGatewayHost
}

// SetDashboardHost overrides the dashboard host. An empty value restores the
// default (https://dashboard.blockchyp.com).
func (client *Client) SetDashboardHost(host string) {
	client.core.DashboardHost = host
}

// DashboardHost returns the dashboard host.
func (client *Client) DashboardHost() string {
	return client.core.DashboardHost
}

// SetCoreHost overrides the Stax core API host. An empty value restores the
// default (https://apiprod.fattlabs.com).
func (client *Client) SetCoreHost(host string) {
	client.core.CoreHost = host
}

// CoreHost returns the Stax core API host.
func (client *Client) CoreHost() string {
	return client.core.CoreHost
}

// Charge executes a standard direct preauth and capture.
func (client *Client) Charge(request AuthorizationRequest) (*AuthorizationResponse, error) {
	var response AuthorizationResponse

	err := client.core.RouteTerminalRequest(request.TerminalName, "/api/charge", "/api/charge", "POST", &request, &response, request.Test, request.Timeout)

	if errors.Is(err, core.ErrUnknownTerminal) {
		response.ResponseDescription = staxpayments.ResponseUnknownTerminal
	} else if timeout, ok := err.(net.Error); ok && timeout.Timeout() {
		response.ResponseDescription = staxpayments.ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	return &response, err
}

// Preauth executes a preauthorization intended to be captured later.
func (client *Client) Preauth(request AuthorizationRequest) (*AuthorizationResponse, error) {
	var response AuthorizationResponse

	err := client.core.RouteTerminalRequest(request.TerminalName, "/api/preauth", "/api/preauth", "POST", &request, &response, request.Test, request.Timeout)

	if errors.Is(err, core.ErrUnknownTerminal) {
		response.ResponseDescription = staxpayments.ResponseUnknownTerminal
	} else if timeout, ok := err.(net.Error); ok && timeout.Timeout() {
		response.ResponseDescription = staxpayments.ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	return &response, err
}

// Refund executes a refund.
func (client *Client) Refund(request RefundRequest) (*AuthorizationResponse, error) {
	var response AuthorizationResponse

	err := client.core.RouteTerminalRequest(request.TerminalName, "/api/refund", "/api/refund", "POST", &request, &response, request.Test, request.Timeout)

	if errors.Is(err, core.ErrUnknownTerminal) {
		response.ResponseDescription = staxpayments.ResponseUnknownTerminal
	} else if timeout, ok := err.(net.Error); ok && timeout.Timeout() {
		response.ResponseDescription = staxpayments.ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	return &response, err
}
