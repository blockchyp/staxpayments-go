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
	"errors"
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

// Ping tests connectivity with a payment terminal.
func (client *Client) Ping(request staxpayments.PingRequest) (*staxpayments.PingResponse, error) {
	var response staxpayments.PingResponse

	err := client.core.RouteTransaction(request.TerminalName, "/api/test", "/api/terminal-test", "POST", &request, &response, request.Test, request.Timeout)

	if errors.Is(err, core.ErrUnknownTerminal) {
		response.ResponseDescription = staxpayments.ResponseUnknownTerminal
	} else if timeout, ok := err.(net.Error); ok && timeout.Timeout() {
		response.ResponseDescription = staxpayments.ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	return &response, err
}

// Locate returns routing and location data for a payment terminal.
func (client *Client) Locate(request staxpayments.LocateRequest) (*staxpayments.LocateResponse, error) {
	var response staxpayments.LocateResponse

	err := client.core.GatewayRequest("/api/terminal-locate", "POST", request, &response, request.Test, request.Timeout)

	if err, ok := err.(net.Error); ok && err.Timeout() {
		response.ResponseDescription = staxpayments.ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	return &response, err
}

// Clear clears the line item display and any in progress transaction.
func (client *Client) Clear(request staxpayments.ClearTerminalRequest) (*staxpayments.Acknowledgement, error) {
	var response staxpayments.Acknowledgement

	err := client.core.RouteTransaction(request.TerminalName, "/api/clear", "/api/terminal-clear", "POST", &request, &response, request.Test, request.Timeout)

	if errors.Is(err, core.ErrUnknownTerminal) {
		response.ResponseDescription = staxpayments.ResponseUnknownTerminal
	} else if timeout, ok := err.(net.Error); ok && timeout.Timeout() {
		response.ResponseDescription = staxpayments.ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	return &response, err
}

// TerminalStatus returns the current status of a terminal.
func (client *Client) TerminalStatus(request staxpayments.TerminalStatusRequest) (*staxpayments.TerminalStatusResponse, error) {
	var response staxpayments.TerminalStatusResponse

	err := client.core.RouteTransaction(request.TerminalName, "/api/terminal-status", "/api/terminal-status", "POST", &request, &response, request.Test, request.Timeout)

	if errors.Is(err, core.ErrUnknownTerminal) {
		response.ResponseDescription = staxpayments.ResponseUnknownTerminal
	} else if timeout, ok := err.(net.Error); ok && timeout.Timeout() {
		response.ResponseDescription = staxpayments.ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	return &response, err
}

// CaptureSignature captures and returns a signature.
func (client *Client) CaptureSignature(request staxpayments.CaptureSignatureRequest) (*staxpayments.CaptureSignatureResponse, error) {
	var response staxpayments.CaptureSignatureResponse

	err := client.core.RouteTransaction(request.TerminalName, "/api/capture-signature", "/api/capture-signature", "POST", &request, &response, request.Test, request.Timeout)

	if errors.Is(err, core.ErrUnknownTerminal) {
		response.ResponseDescription = staxpayments.ResponseUnknownTerminal
	} else if timeout, ok := err.(net.Error); ok && timeout.Timeout() {
		response.ResponseDescription = staxpayments.ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	return &response, err
}

// NewTransactionDisplay displays a new transaction on the terminal.
func (client *Client) NewTransactionDisplay(request staxpayments.TransactionDisplayRequest) (*staxpayments.Acknowledgement, error) {
	var response staxpayments.Acknowledgement

	err := client.core.RouteTransaction(request.TerminalName, "/api/txdisplay", "/api/terminal-txdisplay", "POST", &request, &response, request.Test, request.Timeout)

	if errors.Is(err, core.ErrUnknownTerminal) {
		response.ResponseDescription = staxpayments.ResponseUnknownTerminal
	} else if timeout, ok := err.(net.Error); ok && timeout.Timeout() {
		response.ResponseDescription = staxpayments.ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	return &response, err
}

// UpdateTransactionDisplay appends items to an existing transaction display.
// Subtotal, Tax, and Total are overwritten by the request. Items with the
// same description are combined into groups.
func (client *Client) UpdateTransactionDisplay(request staxpayments.TransactionDisplayRequest) (*staxpayments.Acknowledgement, error) {
	var response staxpayments.Acknowledgement

	err := client.core.RouteTransaction(request.TerminalName, "/api/txdisplay", "/api/terminal-txdisplay", "PUT", &request, &response, request.Test, request.Timeout)

	if errors.Is(err, core.ErrUnknownTerminal) {
		response.ResponseDescription = staxpayments.ResponseUnknownTerminal
	} else if timeout, ok := err.(net.Error); ok && timeout.Timeout() {
		response.ResponseDescription = staxpayments.ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	return &response, err
}

// Message displays a short message on the terminal.
func (client *Client) Message(request staxpayments.MessageRequest) (*staxpayments.Acknowledgement, error) {
	var response staxpayments.Acknowledgement

	err := client.core.RouteTransaction(request.TerminalName, "/api/message", "/api/message", "POST", &request, &response, request.Test, request.Timeout)

	if errors.Is(err, core.ErrUnknownTerminal) {
		response.ResponseDescription = staxpayments.ResponseUnknownTerminal
	} else if timeout, ok := err.(net.Error); ok && timeout.Timeout() {
		response.ResponseDescription = staxpayments.ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	return &response, err
}

// BooleanPrompt asks the consumer a yes/no question.
func (client *Client) BooleanPrompt(request staxpayments.BooleanPromptRequest) (*staxpayments.BooleanPromptResponse, error) {
	var response staxpayments.BooleanPromptResponse

	err := client.core.RouteTransaction(request.TerminalName, "/api/boolean-prompt", "/api/boolean-prompt", "POST", &request, &response, request.Test, request.Timeout)

	if errors.Is(err, core.ErrUnknownTerminal) {
		response.ResponseDescription = staxpayments.ResponseUnknownTerminal
	} else if timeout, ok := err.(net.Error); ok && timeout.Timeout() {
		response.ResponseDescription = staxpayments.ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	return &response, err
}

// TextPrompt asks the consumer a text based question.
func (client *Client) TextPrompt(request staxpayments.TextPromptRequest) (*staxpayments.TextPromptResponse, error) {
	var response staxpayments.TextPromptResponse

	err := client.core.RouteTransaction(request.TerminalName, "/api/text-prompt", "/api/text-prompt", "POST", &request, &response, request.Test, request.Timeout)

	if errors.Is(err, core.ErrUnknownTerminal) {
		response.ResponseDescription = staxpayments.ResponseUnknownTerminal
	} else if timeout, ok := err.(net.Error); ok && timeout.Timeout() {
		response.ResponseDescription = staxpayments.ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	return &response, err
}

// Terminals returns all terminals associated with the merchant account.
func (client *Client) Terminals(request staxpayments.TerminalProfileRequest) (*staxpayments.TerminalProfileResponse, error) {
	var response staxpayments.TerminalProfileResponse

	err := client.core.DashboardRequest("/api/terminals", "GET", request, &response, request.Timeout)

	if err, ok := err.(net.Error); ok && err.Timeout() {
		response.ResponseDescription = staxpayments.ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	return &response, err
}

// DeactivateTerminal deactivates a terminal.
func (client *Client) DeactivateTerminal(request staxpayments.TerminalDeactivationRequest) (*staxpayments.Acknowledgement, error) {
	var response staxpayments.Acknowledgement

	err := client.core.DashboardRequest("/api/terminal/"+request.TerminalID, "DELETE", request, &response, request.Timeout)

	if err, ok := err.(net.Error); ok && err.Timeout() {
		response.ResponseDescription = staxpayments.ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	return &response, err
}

// ActivateTerminal activates a terminal.
func (client *Client) ActivateTerminal(request staxpayments.TerminalActivationRequestV2) (*staxpayments.Acknowledgement, error) {
	var response staxpayments.Acknowledgement

	err := client.core.CoreRequest("/internal/terminal", "POST", request, &response, request.Timeout)

	if err, ok := err.(net.Error); ok && err.Timeout() {
		response.ResponseDescription = staxpayments.ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	return &response, err
}

// Reboot reboot a payment terminal.
func (client *Client) Reboot(request staxpayments.PingRequest) (*staxpayments.Acknowledgement, error) {
	var response staxpayments.Acknowledgement

	err := client.core.RouteTransaction(request.TerminalName, "/api/reboot", "/api/terminal-reboot", "POST", &request, &response, request.Test, request.Timeout)

	if errors.Is(err, core.ErrUnknownTerminal) {
		response.ResponseDescription = staxpayments.ResponseUnknownTerminal
	} else if timeout, ok := err.(net.Error); ok && timeout.Timeout() {
		response.ResponseDescription = staxpayments.ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	return &response, err
}
