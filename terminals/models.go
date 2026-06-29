// Copyright 2019-2026 BlockChyp, Inc. All rights reserved. Use of this code
// is governed by a license that can be found in the LICENSE file.
//
// This file was generated automatically by the BlockChyp SDK Generator.
// Changes to this file will be lost every time the code is regenerated.

// Package terminals request/response models for the Terminal Management Endpoints
// namespace. Types shared across namespaces (acknowledgements, enums,
// credentials) live in the core module and are referenced with the staxpayments
// package prefix.
package terminals

// TerminalProfileRequest models a terminal profile request.
type TerminalProfileRequest struct {
	// Timeout is the request timeout in seconds.
	Timeout int `json:"timeout"`

	// Test specifies whether or not to route transaction to the test gateway.
	Test bool `json:"test"`
}

// TerminalProfileResponse models a terminal profile response.
type TerminalProfileResponse struct {
	// Success indicates whether or not the request succeeded.
	Success bool `json:"success"`

	// Error is the error, if an error occurred.
	Error string `json:"error"`

	// ResponseDescription contains a narrative description of the transaction
	// result.
	ResponseDescription string `json:"responseDescription"`

	// Results enumerates all terminal profiles in the response.
	Results []TerminalProfile `json:"results"`
}

// TerminalProfile contains details about a merchant board platform
// configuration.
type TerminalProfile struct {
	// ID primary identifier for a given terminal.
	ID string `json:"id"`

	// IPAddress is the terminal's local IP address.
	IPAddress string `json:"ipAddress"`

	// TerminalName is the name assigned to the terminal during activation.
	TerminalName string `json:"terminalName"`

	// TerminalType is the terminal type.
	TerminalType string `json:"terminalType"`

	// TerminalTypeDisplayString is the terminal type display string.
	TerminalTypeDisplayString string `json:"terminalTypeDisplayString"`

	// BlockChypFirmwareVersion is the current firmware version deployed on the
	// terminal.
	BlockChypFirmwareVersion string `json:"blockChypFirmwareVersion"`

	// CloudBased indicates whether or not the terminal is configured for cloud
	// relay.
	CloudBased bool `json:"cloudBased"`

	// PublicKey is the terminal's elliptic curve public key.
	PublicKey string `json:"publicKey"`

	// SerialNumber is the manufacturer's serial number.
	SerialNumber string `json:"serialNumber"`

	// Online indicates whether or not the terminal is currently online.
	Online bool `json:"online"`

	// Since the date and time the terminal was first brought online.
	Since string `json:"since"`

	// TotalMemory is the total memory on the terminal.
	TotalMemory int `json:"totalMemory"`

	// TotalStorage is the storage on the terminal.
	TotalStorage int `json:"totalStorage"`

	// AvailableMemory is the available (unused) memory on the terminal.
	AvailableMemory int `json:"availableMemory"`

	// AvailableStorage is the available (unused) storage on the terminal.
	AvailableStorage int `json:"availableStorage"`

	// UsedMemory is the memory currently in use on the terminal.
	UsedMemory int `json:"usedMemory"`

	// UsedStorage is the storage currently in use on the terminal.
	UsedStorage int `json:"usedStorage"`

	// BrandingPreview is the branding asset currently displayed on the terminal.
	BrandingPreview string `json:"brandingPreview"`

	// GroupID the id of the terminal group to which the terminal belongs, if
	// any.
	GroupID string `json:"groupId"`

	// GroupName the name of the terminal group to which the terminal belongs, if
	// any.
	GroupName string `json:"groupName"`
}
