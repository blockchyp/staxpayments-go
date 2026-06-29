// Copyright 2019-2026 BlockChyp, Inc. All rights reserved. Use of this code
// is governed by a license that can be found in the LICENSE file.
//
// This file was generated automatically by the BlockChyp SDK Generator.
// Changes to this file will be lost every time the code is regenerated.

package staxpayments

import (
	"errors"
	"log"
	"net"
)

// PaymentsClient exposes the Payment Endpoints for the Stax Payments
// API. It embeds *Client for the shared transport (host configuration,
// credential exchange, and request plumbing) and adds only this namespace's
// methods.
type PaymentsClient struct {
	*Client
}

// Payments returns a client scoped to the Payment Endpoints
// namespace. The returned client shares this client's transport and
// credentials.
func (client *Client) Payments() *PaymentsClient {
	return &PaymentsClient{client}
}

// Charge executes a standard direct preauth and capture.
func (client *PaymentsClient) Charge(request AuthorizationRequest) (*AuthorizationResponse, error) {
	var response AuthorizationResponse
	var err error

	if err = client.ensure(); err != nil {
		return &response, err
	}

	if err := populateSignatureOptions(&request); err != nil {
		return nil, err
	}

	if request.TerminalName != "" {
		var route TerminalRoute
		route, err = client.resolveTerminalRoute(request.TerminalName)
		if err != nil {
			if errors.Is(err, ErrUnknownTerminal) {
				response.ResponseDescription = ResponseUnknownTerminal
				return &response, err
			}

			return nil, err
		}

		if route.CloudRelayEnabled {
			err = client.RelayRequest("/api/charge", "POST", request, &response, request.Test, request.Timeout)
		} else {
			authRequest := TerminalAuthorizationRequest{
				APICredentials: route.TransientCredentials,
				Request:        request,
			}
			err = client.terminalRequest(route, "/api/charge", "POST", authRequest, &response, request.Timeout)
		}
	} else {
		err = client.GatewayRequest("/api/charge", "POST", request, &response, request.Test, request.Timeout)
	}

	if timeout, ok := err.(net.Error); ok && timeout.Timeout() {
		response.ResponseDescription = ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	if err := handleSignature(request, &response); err != nil {
		log.Printf("Failed to write signature: %+v", err)
	}

	return &response, err
}

// Preauth executes a preauthorization intended to be captured later.
func (client *PaymentsClient) Preauth(request AuthorizationRequest) (*AuthorizationResponse, error) {
	var response AuthorizationResponse
	var err error

	if err = client.ensure(); err != nil {
		return &response, err
	}

	if err := populateSignatureOptions(&request); err != nil {
		return nil, err
	}

	if request.TerminalName != "" {
		var route TerminalRoute
		route, err = client.resolveTerminalRoute(request.TerminalName)
		if err != nil {
			if errors.Is(err, ErrUnknownTerminal) {
				response.ResponseDescription = ResponseUnknownTerminal
				return &response, err
			}

			return nil, err
		}

		if route.CloudRelayEnabled {
			err = client.RelayRequest("/api/preauth", "POST", request, &response, request.Test, request.Timeout)
		} else {
			authRequest := TerminalAuthorizationRequest{
				APICredentials: route.TransientCredentials,
				Request:        request,
			}
			err = client.terminalRequest(route, "/api/preauth", "POST", authRequest, &response, request.Timeout)
		}
	} else {
		err = client.GatewayRequest("/api/preauth", "POST", request, &response, request.Test, request.Timeout)
	}

	if timeout, ok := err.(net.Error); ok && timeout.Timeout() {
		response.ResponseDescription = ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	if err := handleSignature(request, &response); err != nil {
		log.Printf("Failed to write signature: %+v", err)
	}

	return &response, err
}

// Refund executes a refund.
func (client *PaymentsClient) Refund(request RefundRequest) (*AuthorizationResponse, error) {
	var response AuthorizationResponse
	var err error

	if err = client.ensure(); err != nil {
		return &response, err
	}

	if err := populateSignatureOptions(&request); err != nil {
		return nil, err
	}

	if request.TerminalName != "" {
		var route TerminalRoute
		route, err = client.resolveTerminalRoute(request.TerminalName)
		if err != nil {
			if errors.Is(err, ErrUnknownTerminal) {
				response.ResponseDescription = ResponseUnknownTerminal
				return &response, err
			}

			return nil, err
		}

		if route.CloudRelayEnabled {
			err = client.RelayRequest("/api/refund", "POST", request, &response, request.Test, request.Timeout)
		} else {
			authRequest := TerminalRefundRequest{
				APICredentials: route.TransientCredentials,
				Request:        request,
			}
			err = client.terminalRequest(route, "/api/refund", "POST", authRequest, &response, request.Timeout)
		}
	} else {
		err = client.GatewayRequest("/api/refund", "POST", request, &response, request.Test, request.Timeout)
	}

	if timeout, ok := err.(net.Error); ok && timeout.Timeout() {
		response.ResponseDescription = ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	if err := handleSignature(request, &response); err != nil {
		log.Printf("Failed to write signature: %+v", err)
	}

	return &response, err
}
