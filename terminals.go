// Copyright 2019-2026 BlockChyp, Inc. All rights reserved. Use of this code
// is governed by a license that can be found in the LICENSE file.
//
// This file was generated automatically by the BlockChyp SDK Generator.
// Changes to this file will be lost every time the code is regenerated.

package staxpayments

import (
	"net"
)

// TerminalsClient exposes the Terminal Management Endpoints for the Stax Payments
// API. It embeds *Client for the shared transport (host configuration,
// credential exchange, and request plumbing) and adds only this namespace's
// methods.
type TerminalsClient struct {
	*Client
}

// Terminals returns a client scoped to the Terminal Management Endpoints
// namespace. The returned client shares this client's transport and
// credentials.
func (client *Client) Terminals() *TerminalsClient {
	return &TerminalsClient{client}
}

// Terminals returns all terminals associated with the merchant account.
func (client *TerminalsClient) Terminals(request TerminalProfileRequest) (*TerminalProfileResponse, error) {
	var response TerminalProfileResponse

	err := client.DashboardRequest("/api/terminals", "GET", request, &response, request.Timeout)

	if err, ok := err.(net.Error); ok && err.Timeout() {
		response.ResponseDescription = ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	return &response, err
}
