// Copyright 2019-2026 BlockChyp, Inc. All rights reserved. Use of this code
// is governed by a license that can be found in the LICENSE file.
//
// This file was generated automatically by the BlockChyp SDK Generator.
// Changes to this file will be lost every time the code is regenerated.

// Package terminals is the Terminal Management Endpoints namespace of the Stax
// Payments Go SDK. Construct it through the root client (client.New), which
// builds the shared core transport once and injects it here so every namespace
// shares a single transient-credential cache.
package terminals

import (
	"net"

	staxpayments "github.com/blockchyp/staxpayments-go"
	"github.com/blockchyp/staxpayments-go/internal/core"
)

// Client is the Terminal Management Endpoints client. It wraps the shared core client for
// transport and credentials via an unexported field, so only this namespace's
// own operations are part of its public API (the core transport methods are
// deliberately not exposed). Host configuration lives on the root client.
type Client struct {
	core *core.Client
}

// NewClient returns a Terminal Management Endpoints client backed by the given shared core
// transport. It is wired by the root client (client.New); callers construct the
// SDK through that root rather than instantiating namespaces directly, so every
// namespace shares one transient-credential cache.
func NewClient(core *core.Client) *Client {
	return &Client{
		core: core,
	}
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
