// Copyright 2019-2026 BlockChyp, Inc. All rights reserved. Use of this code
// is governed by a license that can be found in the LICENSE file.
//
// This file was generated automatically by the BlockChyp SDK Generator.
// Changes to this file will be lost every time the code is regenerated.

package staxpayments

import (
	"errors"
	"time"
)

// Default client configuration constants. These are consumed by the shared
// core client (see internal/core); they live in the public package so callers
// can reference the defaults.
const (
	DefaultGatewayHost     = "https://api.blockchyp.com"
	DefaultDashboardHost   = "https://dashboard.blockchyp.com"
	DefaultTestGatewayHost = "https://test.blockchyp.com"
	DefaultCoreHost        = "https://apiprod.fattlabs.com"
	DefaultHTTPS           = true
	DefaultRouteCacheTTL   = 60 * time.Minute
	DefaultGatewayTimeout  = 20 * time.Second
	DefaultTerminalTimeout = 2 * time.Minute
)

// Default filesystem configuration.
const (
	ConfigDir  = "staxpayments"
	ConfigFile = "staxpayments.json"
)

// Clientside response constants.
const (
	ResponseUnknownTerminal = "Unknown Terminal"
	ResponseTimedOut        = "Request Timed Out"
)

// ErrInvalidAsyncRequest is returned when a request cannot be called
// asynchronously.
var ErrInvalidAsyncRequest = errors.New("async requests must be terminal requests")

// Version contains the version at build time
var Version string
