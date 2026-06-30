// Copyright 2019-2026 BlockChyp, Inc. All rights reserved. Use of this code
// is governed by a license that can be found in the LICENSE file.
//
// This file was generated automatically by the BlockChyp SDK Generator.
// Changes to this file will be lost every time the code is regenerated.

// Package client provides the root Stax Payments SDK client. It builds the
// shared core transport once and exposes each API namespace (e.g. Payments,
// Terminals) as a field, so a single set of transient credentials is fetched,
// cached, and transparently refreshed across every namespace rather than per
// namespace.
package client

import (
	staxpayments "github.com/blockchyp/staxpayments-go"
	"github.com/blockchyp/staxpayments-go/internal/core"
	"github.com/blockchyp/staxpayments-go/payments"
	"github.com/blockchyp/staxpayments-go/terminals"
)

// Client is the root Stax Payments client. It owns the shared core transport
// (credential cache, HTTP clients, and host configuration) and wires that one
// instance into every namespace client, so all namespaces share a single
// transient-credential cache.
type Client struct {
	core *core.Client

	// Payments exposes the Payment Endpoints.
	Payments *payments.Client

	// Terminals exposes the Terminal Management Endpoints.
	Terminals *terminals.Client
}

// New returns a root client configured with the given Stax bearer credentials.
// The shared core transport is constructed exactly once here; every namespace
// client receives the same instance, so transient credentials are exchanged a
// single time and reused (and refreshed on expiry) across all namespaces.
func New(creds staxpayments.StaxAPICredentials) *Client {
	c := core.New(creds)

	return &Client{
		core:      c,
		Payments:  payments.NewClient(c),
		Terminals: terminals.NewClient(c),
	}
}

// SetGatewayHost overrides the gateway host used for live transactions across
// every namespace. An empty value restores the default
// (https://api.blockchyp.com).
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

// SetDashboardHost overrides the dashboard host used across every namespace. An
// empty value restores the default (https://dashboard.blockchyp.com).
func (client *Client) SetDashboardHost(host string) {
	client.core.DashboardHost = host
}

// DashboardHost returns the dashboard host.
func (client *Client) DashboardHost() string {
	return client.core.DashboardHost
}

// SetCoreHost overrides the Stax core API host used across every namespace. An
// empty value restores the default (https://apiprod.fattlabs.com).
func (client *Client) SetCoreHost(host string) {
	client.core.CoreHost = host
}

// CoreHost returns the Stax core API host.
func (client *Client) CoreHost() string {
	return client.core.CoreHost
}
