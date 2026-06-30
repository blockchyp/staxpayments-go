// Copyright 2019-2026 BlockChyp, Inc. All rights reserved. Use of this code
// is governed by a license that can be found in the LICENSE file.
//
// This file was generated automatically by the BlockChyp SDK Generator.
// Changes to this file will be lost every time the code is regenerated.

// Package payments is the Payment Endpoints namespace of the Stax
// Payments Go SDK. Construct it through the root client (client.New), which
// builds the shared core transport once and injects it here so every namespace
// shares a single transient-credential cache.
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
// deliberately not exposed). Host configuration lives on the root client.
type Client struct {
	core *core.Client
}

// NewClient returns a Payment Endpoints client backed by the given shared core
// transport. It is wired by the root client (client.New); callers construct the
// SDK through that root rather than instantiating namespaces directly, so every
// namespace shares one transient-credential cache.
func NewClient(core *core.Client) *Client {
	return &Client{
		core: core,
	}
}

// Charge executes a standard direct preauth and capture.
func (client *Client) Charge(request staxpayments.AuthorizationRequest) (*staxpayments.AuthorizationResponse, error) {
	var response staxpayments.AuthorizationResponse

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
func (client *Client) Preauth(request staxpayments.AuthorizationRequest) (*staxpayments.AuthorizationResponse, error) {
	var response staxpayments.AuthorizationResponse

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
func (client *Client) Refund(request staxpayments.RefundRequest) (*staxpayments.AuthorizationResponse, error) {
	var response staxpayments.AuthorizationResponse

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
