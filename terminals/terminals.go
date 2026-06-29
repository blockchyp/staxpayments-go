// Copyright 2019-2026 BlockChyp, Inc. All rights reserved. Use of this code
// is governed by a license that can be found in the LICENSE file.
//
// This file was generated automatically by the BlockChyp SDK Generator.
// Changes to this file will be lost every time the code is regenerated.

// Package terminals is the Terminal Management Endpoints namespace of the Stax
// Payments Go SDK. It is its own module so it can be imported and versioned
// independently of the other namespaces; it depends only on the shared core
// module github.com/blockchyp/staxpayments-go.
package terminals

import (
	"net"

	staxpayments "github.com/blockchyp/staxpayments-go"
)

// Client is the Terminal Management Endpoints client. It wraps the shared core client for
// transport and credentials via an unexported field, so only this namespace's
// own operations are part of its public API (the core transport methods are
// deliberately not exposed). The embedded *HostConfig exposes the configurable
// API hosts (GatewayHost, TestGatewayHost, DashboardHost, CoreHost), shared
// with the underlying core client.
type Client struct {
	*staxpayments.HostConfig
	core *staxpayments.Client
}

// NewClient returns a Terminal Management Endpoints client configured with the given Stax
// bearer credentials.
func NewClient(creds staxpayments.StaxAPICredentials) *Client {
	c := staxpayments.NewClient(creds)
	return &Client{HostConfig: c.HostConfig, core: &c}
}

// Terminals returns all terminals associated with the merchant account.
func (client *Client) Terminals(request TerminalProfileRequest) (*TerminalProfileResponse, error) {
	var response TerminalProfileResponse

	err := client.core.DashboardRequest("/api/terminals", "GET", request, &response, request.Timeout)

	if err, ok := err.(net.Error); ok && err.Timeout() {
		response.ResponseDescription = staxpayments.ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	return &response, err
}
