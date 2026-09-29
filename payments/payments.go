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
//
// It takes and returns the Stax Payments models. The BlockChyp wire models are
// an implementation detail: the request is mapped on the way out and the reply
// on the way back. The wire response is described before mapping, so a
// transport error still reaches the caller's ResponseDescription.
func (client *Client) Charge(request staxpayments.AuthRequest) (*staxpayments.AuthResponse, error) {
	wire := staxpayments.AuthRequestMapper(request)

	var wireResponse staxpayments.AuthorizationResponse

	err := client.core.RouteTransaction(wire.TerminalName, "/api/charge", "/api/charge", "POST", &wire, &wireResponse, wire.Test, wire.Timeout)

	if errors.Is(err, core.ErrUnknownTerminal) {
		wireResponse.ResponseDescription = staxpayments.ResponseUnknownTerminal
	} else if timeout, ok := err.(net.Error); ok && timeout.Timeout() {
		wireResponse.ResponseDescription = staxpayments.ResponseTimedOut
	} else if err != nil {
		wireResponse.ResponseDescription = err.Error()
	}

	response := staxpayments.AuthResponseMapper(wireResponse)

	return &response, err
}

// Preauth executes a preauthorization intended to be captured later.
//
// It takes and returns the Stax Payments models. The BlockChyp wire models are
// an implementation detail: the request is mapped on the way out and the reply
// on the way back. The wire response is described before mapping, so a
// transport error still reaches the caller's ResponseDescription.
func (client *Client) Preauth(request staxpayments.AuthRequest) (*staxpayments.AuthResponse, error) {
	wire := staxpayments.AuthRequestMapper(request)

	var wireResponse staxpayments.AuthorizationResponse

	err := client.core.RouteTransaction(wire.TerminalName, "/api/preauth", "/api/preauth", "POST", &wire, &wireResponse, wire.Test, wire.Timeout)

	if errors.Is(err, core.ErrUnknownTerminal) {
		wireResponse.ResponseDescription = staxpayments.ResponseUnknownTerminal
	} else if timeout, ok := err.(net.Error); ok && timeout.Timeout() {
		wireResponse.ResponseDescription = staxpayments.ResponseTimedOut
	} else if err != nil {
		wireResponse.ResponseDescription = err.Error()
	}

	response := staxpayments.AuthResponseMapper(wireResponse)

	return &response, err
}
