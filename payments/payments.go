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
)

// Client is the Payment Endpoints client. It embeds the shared
// *staxpayments.Client for transport and credentials and exposes only this
// namespace's methods.
type Client struct {
	*staxpayments.Client
}

// New returns a Payment Endpoints client configured with the given Stax bearer
// credentials.
func New(creds staxpayments.StaxAPICredentials) *Client {
	c := staxpayments.NewClient(creds)
	return &Client{&c}
}

// Charge executes a standard direct preauth and capture.
func (client *Client) Charge(request AuthorizationRequest) (*AuthorizationResponse, error) {
	var response AuthorizationResponse

	err := client.RouteTerminalRequest(request.TerminalName, "/api/charge", "/api/charge", "POST", &request, &response, request.Test, request.Timeout)

	if errors.Is(err, staxpayments.ErrUnknownTerminal) {
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

	err := client.RouteTerminalRequest(request.TerminalName, "/api/preauth", "/api/preauth", "POST", &request, &response, request.Test, request.Timeout)

	if errors.Is(err, staxpayments.ErrUnknownTerminal) {
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

	err := client.RouteTerminalRequest(request.TerminalName, "/api/refund", "/api/refund", "POST", &request, &response, request.Test, request.Timeout)

	if errors.Is(err, staxpayments.ErrUnknownTerminal) {
		response.ResponseDescription = staxpayments.ResponseUnknownTerminal
	} else if timeout, ok := err.(net.Error); ok && timeout.Timeout() {
		response.ResponseDescription = staxpayments.ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	return &response, err
}
