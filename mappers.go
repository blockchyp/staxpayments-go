// Copyright 2019-2026 BlockChyp, Inc. All rights reserved. Use of this code
// is governed by a license that can be found in the LICENSE file.
//
// This file was generated automatically by the BlockChyp SDK Generator.
// Changes to this file will be lost every time the code is regenerated.

package staxpayments

// AuthRequestMapper maps a Stax Payments charge or preauth request onto the
// BlockChyp wire model. Fields not listed here carry the same name on both
// sides and are matched by name; the generated mapper spells every one of
// them out.
func AuthRequestMapper(src AuthRequest) AuthorizationRequest {
	var dst AuthorizationRequest

	dst.TipAmount = src.TipAmount
	dst.TaxAmount = src.TaxAmount
	dst.TerminalName = src.TerminalName
	dst.Amount = src.Amount
	dst.CurrencyCode = src.CurrencyCode
	dst.PromptForTip = src.PromptForTip
	dst.Test = src.Test
	dst.Enroll = src.Enroll
	dst.ExternalPartnerMetadata = src.ExternalPartnerMetadata

	return dst
}

// AuthResponseMapper maps a BlockChyp charge or preauth response onto the
// Stax Payments model. Fields not listed here are matched by name.
func AuthResponseMapper(src AuthorizationResponse) AuthResponse {
	var dst AuthResponse

	dst.Success = src.Success
	dst.Error = src.Error
	dst.ResponseDescription = src.ResponseDescription
	// The Stax transaction id comes back in the transaction ref. BlockChyp's
	// transactionId is BlockChyp's own id and is deliberately not surfaced.
	dst.TransactionID = src.TransactionRef
	dst.Approved = src.Approved
	dst.AuthCode = src.AuthCode
	dst.RequestedAmount = src.RequestedAmount
	dst.AuthorizedAmount = src.AuthorizedAmount
	dst.CurrencyCode = src.CurrencyCode
	dst.EntryMethod = src.EntryMethod
	dst.MaskedPAN = src.MaskedPAN
	dst.Network = src.Network
	dst.Timestamp = src.Timestamp
	dst.Status = src.Status
	dst.CardMetadata = src.CardMetadata
	dst.ReceiptSuggestions = src.ReceiptSuggestions
	dst.Test = src.Test

	return dst
}
