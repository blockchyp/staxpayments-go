// Copyright 2019-2026 BlockChyp, Inc. All rights reserved. Use of this code
// is governed by a license that can be found in the LICENSE file.
//
// This file was generated automatically by the BlockChyp SDK Generator.
// Changes to this file will be lost every time the code is regenerated.

package staxpayments

import (
	"reflect"
	"time"
)

// APICredentials models gateway credentials.
type APICredentials struct {
	APIKey      string `json:"apiKey"`
	BearerToken string `json:"bearerToken"`
	SigningKey  string `json:"signingKey"`
}

// StaxAPICredentials is the Stax bearer token used to construct a client. It is
// the only credential an integrator supplies; the BlockChyp gateway credentials
// are obtained and managed internally by the SDK.
type StaxAPICredentials struct {
	BearerToken string `json:"bearerToken"`
}

// CardType is used to differentiate credit, debit, and EBT.
type CardType int

// CardTypes.
const (
	CardTypeCredit CardType = iota
	CardTypeDebit
	CardTypeEBT
	CardTypeBlockchainGift
	CardTypeHealthcare
)

// SignatureFormat is used to specify the output format for customer
// signature images.
type SignatureFormat string

// SignatureFormats.
const (
	SignatureFormatNone = ""
	SignatureFormatPNG  = "png"
	SignatureFormatJPG  = "jpg"
	SignatureFormatGIF  = "gif"
)

// CVMType designates a customer verification method.
type CVMType string

// CVMs defined.
const (
	CVMTypeSignature  CVMType = "Signature"
	CVMTypeOfflinePIN CVMType = "Offline PIN"
	CVMTypeOnlinePIN  CVMType = "Online PIN"
	CVMTypeCDCVM      CVMType = "CDCVM"
	CVMTypeNoCVM      CVMType = "No CVM"
)

// PromptType is used to specify the type of text input data being requested
// from a customer.
type PromptType string

// PromptTypes.
const (
	PromptTypeAmount         = "amount"
	PromptTypeEmail          = "email"
	PromptTypePhone          = "phone"
	PromptTypeCustomerNumber = "customer-number"
	PromptTypeRewardsNumber  = "rewards-number"
	PromptTypeFirstName      = "first-name"
	PromptTypeLastName       = "last-name"
)

// AVSResponse indicates the result of address verification.
type AVSResponse string

// AVSResponse types.
const (
	AVSResponseNotApplicable             AVSResponse = ""
	AVSResponseNotSupported                          = "not_supported"
	AVSResponseRetry                                 = "retry"
	AVSResponseNoMatch                               = "no_match"
	AVSResponseAddressMatch                          = "address_match"
	AVSResponsePostalCodeMatch                       = "zip_match"
	AVSResponseAddressAndPostalCodeMatch             = "match"
)

// HealthcareType is a category of healthcare.
type HealthcareType string

// HealthcareTypes.
const (
	HealthcareTypeHealthcare   = "healthcare"
	HealthcareTypePrescription = "prescription"
	HealthcareTypeVision       = "vision"
	HealthcareTypeClinic       = "clinic"
	HealthcareTypeDental       = "dental"
)

// RoundingMode indicates how partial penny rounding operations should work
type RoundingMode string

// RoundingMode types
const (
	RoundingModeUp      = "up"
	RoundingModeNearest = "nearest"
	RoundingModeDown    = "down"
)

// ReceiptSuggestions contains EMV fields we recommend developers put on their
// receipts.
type ReceiptSuggestions struct {
	// AID is the EMV Application Identifier.
	AID string `json:"aid,omitempty"`

	// ARQC is the EMV Application Request Cryptogram.
	ARQC string `json:"arqc,omitempty"`

	// IAD is the EMV Issuer Application Data.
	IAD string `json:"iad,omitempty"`

	// ARC is the EMV Authorization Response Code.
	ARC string `json:"arc,omitempty"`

	// TC is the EMV Transaction Certificate.
	TC string `json:"tc,omitempty"`

	// TVR is the EMV Terminal Verification Response.
	TVR string `json:"tvr,omitempty"`

	// TSI is the EMV Transaction Status Indicator.
	TSI string `json:"tsi,omitempty"`

	// TerminalID is the ID of the payment terminal.
	TerminalID string `json:"terminalId,omitempty"`

	// MerchantName is the name of the merchant's business.
	MerchantName string `json:"merchantName,omitempty"`

	// MerchantID is the ID of the merchant.
	MerchantID string `json:"merchantId,omitempty"`

	// MerchantKey is the partially masked merchant key required on EMV receipts.
	MerchantKey string `json:"merchantKey,omitempty"`

	// ApplicationLabel is a description of the selected AID.
	ApplicationLabel string `json:"applicationLabel,omitempty"`

	// RequestSignature indicates that the receipt should contain a signature
	// line.
	RequestSignature bool `json:"requestSignature"`

	// MaskedPAN is the masked primary account number of the payment card, as
	// required.
	MaskedPAN string `json:"maskedPan,omitempty"`

	// AuthorizedAmount is the amount authorized by the payment network. Could be
	// less than the requested amount for partial auth.
	AuthorizedAmount string `json:"authorizedAmount"`

	// TransactionType is the type of transaction performed (CHARGE, PREAUTH,
	// REFUND, etc).
	TransactionType string `json:"transactionType"`

	// EntryMethod is the method by which the payment card was entered (MSR,
	// CHIP, KEYED, etc.).
	EntryMethod string `json:"entryMethod,omitempty"`

	// PINVerified indicates that PIN verification was performed.
	PINVerified bool `json:"pinVerified,omitempty"`

	// CVMUsed indicates the customer verification method used for the
	// transaction.
	CVMUsed CVMType `json:"cvmUsed,omitempty"`

	// Fallback indicates that a chip read failure caused the transaction to fall
	// back to the magstripe.
	Fallback bool `json:"fallback,omitempty"`

	// BatchSequence is the sequence of the transaction in the batch.
	BatchSequence int `json:"batchSequence,omitempty"`

	// CashBackAmount is the amount of cash back that was approved.
	CashBackAmount string `json:"cashBackAmount,omitempty"`

	// Surcharge is the amount added to the transaction to cover eligible credit
	// card fees.
	Surcharge string `json:"surcharge,omitempty"`

	// CashDiscount is the discount applied to the transaction for payment
	// methods ineligible for surcharges.
	CashDiscount string `json:"cashDiscount,omitempty"`

	// ServiceFeeAmount is the amount added to the transaction to cover
	// processing fees.
	ServiceFeeAmount string `json:"serviceFeeAmount,omitempty"`
}

// Acknowledgement contains a basic api acknowledgement.
type Acknowledgement struct {
	// Success indicates whether or not the request succeeded.
	Success bool `json:"success"`

	// Error is the error, if an error occurred.
	Error string `json:"error"`

	// ResponseDescription contains a narrative description of the transaction
	// result.
	ResponseDescription string `json:"responseDescription"`
}

// CaptureSignatureRequest contains a request for customer signature data.
type CaptureSignatureRequest struct {
	// Timeout is the request timeout in seconds.
	Timeout int `json:"timeout"`

	// Test specifies whether or not to route transaction to the test gateway.
	Test bool `json:"test"`

	// TransactionRef contains a user-assigned reference that can be used to
	// recall or reverse transactions.
	TransactionRef string `json:"transactionRef,omitempty"`

	// AutogeneratedRef indicates that the transaction reference was
	// autogenerated and should be ignored for the purposes of duplicate
	// detection.
	AutogeneratedRef bool `json:"autogeneratedRef"`

	// Async defers the response to the transaction and returns immediately.
	// Callers should retrive the transaction result using the Transaction Status
	// API.
	Async bool `json:"async"`

	// Queue adds the transaction to the queue and returns immediately. Callers
	// should retrive the transaction result using the Transaction Status API.
	Queue bool `json:"queue"`

	// WaitForRemovedCard specifies whether or not the request should block until
	// all cards have been removed from the card reader.
	WaitForRemovedCard bool `json:"waitForRemovedCard,omitempty"`

	// Force causes a transaction to override any in-progress transactions.
	Force bool `json:"force,omitempty"`

	// OrderRef is an identifier from an external point of sale system.
	OrderRef string `json:"orderRef,omitempty"`

	// DestinationAccount is the settlement account for merchants with split
	// settlements.
	DestinationAccount string `json:"destinationAccount,omitempty"`

	// TestCase can include a code used to trigger simulated conditions for the
	// purposes of testing and certification. Valid for test merchant accounts
	// only.
	TestCase string `json:"testCase,omitempty"`

	// SigFile is a location on the filesystem which a customer signature should
	// be written to.
	SigFile string `json:"sigFile,omitempty"`

	// SigFormat specifies the image format to be used for returning signatures.
	SigFormat SignatureFormat `json:"sigFormat,omitempty"`

	// SigWidth is the width that the signature image should be scaled to,
	// preserving the aspect ratio. If not provided, the signature is returned in
	// the terminal's max resolution.
	SigWidth int `json:"sigWidth,omitempty"`

	// DisableSignature specifies whether or not signature prompt should be
	// skipped on the terminal. The terminal will indicate whether or not a
	// signature is required by the card in the receipt suggestions response.
	DisableSignature bool `json:"disableSignature,omitempty"`

	// TerminalName is the name of the target payment terminal.
	TerminalName string `json:"terminalName,omitempty"`

	// ResetConnection forces the terminal cloud connection to be reset while a
	// transactions is in flight. This is a diagnostic settings that can be used
	// only for test transactions.
	ResetConnection bool `json:"resetConnection"`
}

// CaptureSignatureResponse contains customer signature data.
type CaptureSignatureResponse struct {
	// Success indicates whether or not the request succeeded.
	Success bool `json:"success"`

	// Error is the error, if an error occurred.
	Error string `json:"error"`

	// ResponseDescription contains a narrative description of the transaction
	// result.
	ResponseDescription string `json:"responseDescription"`

	// SigFile contains the hex encoded signature data.
	SigFile string `json:"sigFile,omitempty"`
}

// PingRequest contains information needed to test connectivity with a
// terminal.
type PingRequest struct {
	// Timeout is the request timeout in seconds.
	Timeout int `json:"timeout"`

	// Test specifies whether or not to route transaction to the test gateway.
	Test bool `json:"test"`

	// TransactionRef contains a user-assigned reference that can be used to
	// recall or reverse transactions.
	TransactionRef string `json:"transactionRef,omitempty"`

	// AutogeneratedRef indicates that the transaction reference was
	// autogenerated and should be ignored for the purposes of duplicate
	// detection.
	AutogeneratedRef bool `json:"autogeneratedRef"`

	// Async defers the response to the transaction and returns immediately.
	// Callers should retrive the transaction result using the Transaction Status
	// API.
	Async bool `json:"async"`

	// Queue adds the transaction to the queue and returns immediately. Callers
	// should retrive the transaction result using the Transaction Status API.
	Queue bool `json:"queue"`

	// WaitForRemovedCard specifies whether or not the request should block until
	// all cards have been removed from the card reader.
	WaitForRemovedCard bool `json:"waitForRemovedCard,omitempty"`

	// Force causes a transaction to override any in-progress transactions.
	Force bool `json:"force,omitempty"`

	// OrderRef is an identifier from an external point of sale system.
	OrderRef string `json:"orderRef,omitempty"`

	// DestinationAccount is the settlement account for merchants with split
	// settlements.
	DestinationAccount string `json:"destinationAccount,omitempty"`

	// TestCase can include a code used to trigger simulated conditions for the
	// purposes of testing and certification. Valid for test merchant accounts
	// only.
	TestCase string `json:"testCase,omitempty"`

	// TerminalName is the name of the target payment terminal.
	TerminalName string `json:"terminalName,omitempty"`

	// ResetConnection forces the terminal cloud connection to be reset while a
	// transactions is in flight. This is a diagnostic settings that can be used
	// only for test transactions.
	ResetConnection bool `json:"resetConnection"`
}

// PingResponse contains the response to a ping request.
type PingResponse struct {
	// Success indicates whether or not the request succeeded.
	Success bool `json:"success"`

	// Error is the error, if an error occurred.
	Error string `json:"error"`

	// ResponseDescription contains a narrative description of the transaction
	// result.
	ResponseDescription string `json:"responseDescription"`

	// TransactionID is the ID assigned to the transaction.
	TransactionID string `json:"transactionId"`

	// BatchID is the ID assigned to the batch.
	BatchID string `json:"batchId,omitempty"`

	// TransactionRef is the transaction reference string assigned to the
	// transaction request. If no transaction ref was assiged on the request,
	// then the gateway will randomly generate one.
	TransactionRef string `json:"transactionRef,omitempty"`

	// TransactionType is the type of transaction.
	TransactionType string `json:"transactionType"`

	// Timestamp is the timestamp of the transaction.
	Timestamp string `json:"timestamp"`

	// TickBlock is the hash of the last tick block.
	TickBlock string `json:"tickBlock"`

	// Test indicates that the transaction was processed on the test gateway.
	Test bool `json:"test"`

	// DestinationAccount is the settlement account for merchants with split
	// settlements.
	DestinationAccount string `json:"destinationAccount,omitempty"`

	// Sig is the ECC signature of the response. Can be used to ensure that it
	// was signed by the terminal and detect man-in-the middle attacks.
	Sig string `json:"sig,omitempty"`
}

// LocateRequest contains information needed to retrieve location information
// for a terminal.
type LocateRequest struct {
	// Timeout is the request timeout in seconds.
	Timeout int `json:"timeout"`

	// Test specifies whether or not to route transaction to the test gateway.
	Test bool `json:"test"`

	// TransactionRef contains a user-assigned reference that can be used to
	// recall or reverse transactions.
	TransactionRef string `json:"transactionRef,omitempty"`

	// AutogeneratedRef indicates that the transaction reference was
	// autogenerated and should be ignored for the purposes of duplicate
	// detection.
	AutogeneratedRef bool `json:"autogeneratedRef"`

	// Async defers the response to the transaction and returns immediately.
	// Callers should retrive the transaction result using the Transaction Status
	// API.
	Async bool `json:"async"`

	// Queue adds the transaction to the queue and returns immediately. Callers
	// should retrive the transaction result using the Transaction Status API.
	Queue bool `json:"queue"`

	// WaitForRemovedCard specifies whether or not the request should block until
	// all cards have been removed from the card reader.
	WaitForRemovedCard bool `json:"waitForRemovedCard,omitempty"`

	// Force causes a transaction to override any in-progress transactions.
	Force bool `json:"force,omitempty"`

	// OrderRef is an identifier from an external point of sale system.
	OrderRef string `json:"orderRef,omitempty"`

	// DestinationAccount is the settlement account for merchants with split
	// settlements.
	DestinationAccount string `json:"destinationAccount,omitempty"`

	// TestCase can include a code used to trigger simulated conditions for the
	// purposes of testing and certification. Valid for test merchant accounts
	// only.
	TestCase string `json:"testCase,omitempty"`

	// TerminalName is the name of the target payment terminal.
	TerminalName string `json:"terminalName,omitempty"`

	// ResetConnection forces the terminal cloud connection to be reset while a
	// transactions is in flight. This is a diagnostic settings that can be used
	// only for test transactions.
	ResetConnection bool `json:"resetConnection"`
}

// LocateResponse contains the response to a locate request.
type LocateResponse struct {
	// Success indicates whether or not the request succeeded.
	Success bool `json:"success"`

	// Error is the error, if an error occurred.
	Error string `json:"error"`

	// ResponseDescription contains a narrative description of the transaction
	// result.
	ResponseDescription string `json:"responseDescription"`

	// TransactionID is the ID assigned to the transaction.
	TransactionID string `json:"transactionId"`

	// BatchID is the ID assigned to the batch.
	BatchID string `json:"batchId,omitempty"`

	// TransactionRef is the transaction reference string assigned to the
	// transaction request. If no transaction ref was assiged on the request,
	// then the gateway will randomly generate one.
	TransactionRef string `json:"transactionRef,omitempty"`

	// TransactionType is the type of transaction.
	TransactionType string `json:"transactionType"`

	// Timestamp is the timestamp of the transaction.
	Timestamp string `json:"timestamp"`

	// TickBlock is the hash of the last tick block.
	TickBlock string `json:"tickBlock"`

	// Test indicates that the transaction was processed on the test gateway.
	Test bool `json:"test"`

	// DestinationAccount is the settlement account for merchants with split
	// settlements.
	DestinationAccount string `json:"destinationAccount,omitempty"`

	// Sig is the ECC signature of the response. Can be used to ensure that it
	// was signed by the terminal and detect man-in-the middle attacks.
	Sig string `json:"sig,omitempty"`

	// TerminalName is the name assigned to the terminal at activation.
	TerminalName string `json:"terminalName"`

	// IPAddress is the local IP address of the terminal.
	IPAddress string `json:"ipAddress"`

	// CloudRelay indicates whether or not the terminal is running in cloud relay
	// mode.
	CloudRelay bool `json:"cloudRelay"`

	// PublicKey is the terminal's public key.
	PublicKey string `json:"publicKey"`
}

// MessageRequest contains a message to be displayed on the terminal screen.
type MessageRequest struct {
	// Timeout is the request timeout in seconds.
	Timeout int `json:"timeout"`

	// Test specifies whether or not to route transaction to the test gateway.
	Test bool `json:"test"`

	// TransactionRef contains a user-assigned reference that can be used to
	// recall or reverse transactions.
	TransactionRef string `json:"transactionRef,omitempty"`

	// AutogeneratedRef indicates that the transaction reference was
	// autogenerated and should be ignored for the purposes of duplicate
	// detection.
	AutogeneratedRef bool `json:"autogeneratedRef"`

	// Async defers the response to the transaction and returns immediately.
	// Callers should retrive the transaction result using the Transaction Status
	// API.
	Async bool `json:"async"`

	// Queue adds the transaction to the queue and returns immediately. Callers
	// should retrive the transaction result using the Transaction Status API.
	Queue bool `json:"queue"`

	// WaitForRemovedCard specifies whether or not the request should block until
	// all cards have been removed from the card reader.
	WaitForRemovedCard bool `json:"waitForRemovedCard,omitempty"`

	// Force causes a transaction to override any in-progress transactions.
	Force bool `json:"force,omitempty"`

	// OrderRef is an identifier from an external point of sale system.
	OrderRef string `json:"orderRef,omitempty"`

	// DestinationAccount is the settlement account for merchants with split
	// settlements.
	DestinationAccount string `json:"destinationAccount,omitempty"`

	// TestCase can include a code used to trigger simulated conditions for the
	// purposes of testing and certification. Valid for test merchant accounts
	// only.
	TestCase string `json:"testCase,omitempty"`

	// TerminalName is the name of the target payment terminal.
	TerminalName string `json:"terminalName,omitempty"`

	// ResetConnection forces the terminal cloud connection to be reset while a
	// transactions is in flight. This is a diagnostic settings that can be used
	// only for test transactions.
	ResetConnection bool `json:"resetConnection"`

	// Message is the message to be displayed on the terminal.
	Message string `json:"message"`
}

// BooleanPromptRequest contains a simple yes no prompt request.
type BooleanPromptRequest struct {
	// Timeout is the request timeout in seconds.
	Timeout int `json:"timeout"`

	// Test specifies whether or not to route transaction to the test gateway.
	Test bool `json:"test"`

	// TransactionRef contains a user-assigned reference that can be used to
	// recall or reverse transactions.
	TransactionRef string `json:"transactionRef,omitempty"`

	// AutogeneratedRef indicates that the transaction reference was
	// autogenerated and should be ignored for the purposes of duplicate
	// detection.
	AutogeneratedRef bool `json:"autogeneratedRef"`

	// Async defers the response to the transaction and returns immediately.
	// Callers should retrive the transaction result using the Transaction Status
	// API.
	Async bool `json:"async"`

	// Queue adds the transaction to the queue and returns immediately. Callers
	// should retrive the transaction result using the Transaction Status API.
	Queue bool `json:"queue"`

	// WaitForRemovedCard specifies whether or not the request should block until
	// all cards have been removed from the card reader.
	WaitForRemovedCard bool `json:"waitForRemovedCard,omitempty"`

	// Force causes a transaction to override any in-progress transactions.
	Force bool `json:"force,omitempty"`

	// OrderRef is an identifier from an external point of sale system.
	OrderRef string `json:"orderRef,omitempty"`

	// DestinationAccount is the settlement account for merchants with split
	// settlements.
	DestinationAccount string `json:"destinationAccount,omitempty"`

	// TestCase can include a code used to trigger simulated conditions for the
	// purposes of testing and certification. Valid for test merchant accounts
	// only.
	TestCase string `json:"testCase,omitempty"`

	// TerminalName is the name of the target payment terminal.
	TerminalName string `json:"terminalName,omitempty"`

	// ResetConnection forces the terminal cloud connection to be reset while a
	// transactions is in flight. This is a diagnostic settings that can be used
	// only for test transactions.
	ResetConnection bool `json:"resetConnection"`

	// YesCaption is the preferred caption for the 'yes' button.
	YesCaption string `json:"yesCaption"`

	// NoCaption is the preferred caption for the 'no' button.
	NoCaption string `json:"noCaption"`

	// Prompt is the text to be displayed on the terminal.
	Prompt string `json:"prompt"`
}

// TextPromptRequest contains a text prompt request.
type TextPromptRequest struct {
	// Timeout is the request timeout in seconds.
	Timeout int `json:"timeout"`

	// Test specifies whether or not to route transaction to the test gateway.
	Test bool `json:"test"`

	// TransactionRef contains a user-assigned reference that can be used to
	// recall or reverse transactions.
	TransactionRef string `json:"transactionRef,omitempty"`

	// AutogeneratedRef indicates that the transaction reference was
	// autogenerated and should be ignored for the purposes of duplicate
	// detection.
	AutogeneratedRef bool `json:"autogeneratedRef"`

	// Async defers the response to the transaction and returns immediately.
	// Callers should retrive the transaction result using the Transaction Status
	// API.
	Async bool `json:"async"`

	// Queue adds the transaction to the queue and returns immediately. Callers
	// should retrive the transaction result using the Transaction Status API.
	Queue bool `json:"queue"`

	// WaitForRemovedCard specifies whether or not the request should block until
	// all cards have been removed from the card reader.
	WaitForRemovedCard bool `json:"waitForRemovedCard,omitempty"`

	// Force causes a transaction to override any in-progress transactions.
	Force bool `json:"force,omitempty"`

	// OrderRef is an identifier from an external point of sale system.
	OrderRef string `json:"orderRef,omitempty"`

	// DestinationAccount is the settlement account for merchants with split
	// settlements.
	DestinationAccount string `json:"destinationAccount,omitempty"`

	// TestCase can include a code used to trigger simulated conditions for the
	// purposes of testing and certification. Valid for test merchant accounts
	// only.
	TestCase string `json:"testCase,omitempty"`

	// TerminalName is the name of the target payment terminal.
	TerminalName string `json:"terminalName,omitempty"`

	// ResetConnection forces the terminal cloud connection to be reset while a
	// transactions is in flight. This is a diagnostic settings that can be used
	// only for test transactions.
	ResetConnection bool `json:"resetConnection"`

	// PromptType is the prompt type (email, phone, etc).
	PromptType PromptType `json:"promptType"`
}

// Customer models a customer record.
type Customer struct {
	// ID BlockChyp assigned customer id.
	ID string `json:"id"`

	// CustomerRef optional customer ref that can be used for the client's
	// system's customer id.
	CustomerRef string `json:"customerRef"`

	// FirstName customer's first name.
	FirstName string `json:"firstName"`

	// LastName customer's last name.
	LastName string `json:"lastName"`

	// CompanyName customer's company name.
	CompanyName string `json:"companyName"`

	// EmailAddress customer's email address.
	EmailAddress string `json:"emailAddress"`

	// SmsNumber customer's SMS or mobile number.
	SmsNumber string `json:"smsNumber"`

	// PaymentMethods model saved payment methods associated with a customer.
	PaymentMethods []CustomerToken `json:"paymentMethods"`
}

// CustomerToken models a customer token.
type CustomerToken struct {
	// Token BlockChyp assigned customer id.
	Token string `json:"token"`

	// MaskedPAN masked primary account number.
	MaskedPAN string `json:"maskedPan"`

	// ExpiryMonth expiration month.
	ExpiryMonth string `json:"expiryMonth"`

	// ExpiryYear expiration month.
	ExpiryYear string `json:"expiryYear"`

	// PaymentType payment type.
	PaymentType string `json:"paymentType"`

	// AccountType bank account type (checking, saving).
	AccountType string `json:"accountType"`

	// AccountHolderType bank account holder type (personal, business).
	AccountHolderType string `json:"accountHolderType"`

	// BankName bank name.
	BankName string `json:"bankName"`

	// RoutingNumber routing number.
	RoutingNumber string `json:"routingNumber"`

	// TokenHash token hash (generated with a static salt, Merchant ID,
	// Registration Date and PAN.
	TokenHash string `json:"tokenHash"`

	// Bin card bin.
	Bin string `json:"bin"`

	// PostalCode is the card postal code.
	PostalCode string `json:"postalCode,omitempty"`

	// Address is the card address.
	Address string `json:"address,omitempty"`

	// Country is the card country.
	Country string `json:"country,omitempty"`

	// CardHolderName is the card holder name.
	CardHolderName string `json:"cardHolderName,omitempty"`

	// HasCVV indicates whether the token was enrolled with a CVV value present.
	HasCVV bool `json:"hasCvv"`

	// Customers models customer records associated with a payment token.
	Customers []Customer `json:"customers"`
}

// TextPromptResponse contains the response to a text prompt request.
type TextPromptResponse struct {
	// Success indicates whether or not the request succeeded.
	Success bool `json:"success"`

	// Error is the error, if an error occurred.
	Error string `json:"error"`

	// ResponseDescription contains a narrative description of the transaction
	// result.
	ResponseDescription string `json:"responseDescription"`

	// Response is the text prompt response.
	Response string `json:"response"`
}

// BooleanPromptResponse contains the response to a boolean prompt request.
type BooleanPromptResponse struct {
	// Success indicates whether or not the request succeeded.
	Success bool `json:"success"`

	// Error is the error, if an error occurred.
	Error string `json:"error"`

	// ResponseDescription contains a narrative description of the transaction
	// result.
	ResponseDescription string `json:"responseDescription"`

	// Response is the boolean prompt response.
	Response bool `json:"response"`
}

// WhiteListedCard shows details about a white listed card.
type WhiteListedCard struct {
	// Bin is the card BIN.
	Bin string `json:"bin"`

	// Track1 is the track 1 data from the card.
	Track1 string `json:"track1"`

	// Track2 is the track 2 data from the card.
	Track2 string `json:"track2"`

	// PAN is the card primary account number.
	PAN string `json:"pan"`
}

// AuthorizationRequest contains an authorization request for a charge,
// preauth, or reverse transaction.
type AuthorizationRequest struct {
	// Timeout is the request timeout in seconds.
	Timeout int `json:"timeout"`

	// Test specifies whether or not to route transaction to the test gateway.
	Test bool `json:"test"`

	// TransactionRef contains a user-assigned reference that can be used to
	// recall or reverse transactions.
	TransactionRef string `json:"transactionRef,omitempty"`

	// AutogeneratedRef indicates that the transaction reference was
	// autogenerated and should be ignored for the purposes of duplicate
	// detection.
	AutogeneratedRef bool `json:"autogeneratedRef"`

	// Async defers the response to the transaction and returns immediately.
	// Callers should retrive the transaction result using the Transaction Status
	// API.
	Async bool `json:"async"`

	// Queue adds the transaction to the queue and returns immediately. Callers
	// should retrive the transaction result using the Transaction Status API.
	Queue bool `json:"queue"`

	// WaitForRemovedCard specifies whether or not the request should block until
	// all cards have been removed from the card reader.
	WaitForRemovedCard bool `json:"waitForRemovedCard,omitempty"`

	// Force causes a transaction to override any in-progress transactions.
	Force bool `json:"force,omitempty"`

	// OrderRef is an identifier from an external point of sale system.
	OrderRef string `json:"orderRef,omitempty"`

	// DestinationAccount is the settlement account for merchants with split
	// settlements.
	DestinationAccount string `json:"destinationAccount,omitempty"`

	// TestCase can include a code used to trigger simulated conditions for the
	// purposes of testing and certification. Valid for test merchant accounts
	// only.
	TestCase string `json:"testCase,omitempty"`

	// Token is the payment token to be used for this transaction. This should be
	// used for recurring transactions. The /enroll endpoint ignores this field.
	Token string `json:"token,omitempty"`

	// Track1 contains track 1 magnetic stripe data.
	Track1 string `json:"track1,omitempty"`

	// Track2 contains track 2 magnetic stripe data.
	Track2 string `json:"track2,omitempty"`

	// PAN contains the primary account number. We recommend using the terminal
	// or e-commerce tokenization libraries instead of passing account numbers in
	// directly, as this would put your application in PCI scope.
	PAN string `json:"pan,omitempty"`

	// RoutingNumber is the ACH routing number for ACH transactions.
	RoutingNumber string `json:"routingNumber,omitempty"`

	// CardholderName is the cardholder name. Only required if the request
	// includes a primary account number or track data.
	CardholderName string `json:"cardholderName,omitempty"`

	// ExpMonth is the card expiration month for use with PAN based transactions.
	ExpMonth string `json:"expMonth,omitempty"`

	// ExpYear is the card expiration year for use with PAN based transactions.
	ExpYear string `json:"expYear,omitempty"`

	// CVV is the card CVV for use with PAN based transactions.
	CVV string `json:"cvv,omitempty"`

	// Address is the cardholder address for use with address verification.
	Address string `json:"address,omitempty"`

	// PostalCode is the cardholder postal code for use with address
	// verification.
	PostalCode string `json:"postalCode,omitempty"`

	// Country is the cardholder country.
	Country string `json:"country,omitempty"`

	// ManualEntry specifies that the payment entry method is a manual keyed
	// transaction. If this is true, no other payment method will be accepted.
	ManualEntry bool `json:"manualEntry,omitempty"`

	// KSN is the key serial number used for DUKPT encryption.
	KSN string `json:"ksn,omitempty"`

	// PINBlock is the encrypted pin block.
	PINBlock string `json:"pinBlock,omitempty"`

	// CardType designates categories of cards: credit, debit, EBT.
	CardType CardType `json:"cardType,omitempty"`

	// PaymentType designates brands of payment methods: Visa, Discover, etc.
	PaymentType string `json:"paymentType,omitempty"`

	// CurrencyCode indicates the transaction currency code.
	CurrencyCode string `json:"currencyCode"`

	// Amount is the requested amount.
	Amount string `json:"amount"`

	// TaxExempt indicates that the request is tax exempt. Only required for tax
	// exempt level 2 processing.
	TaxExempt bool `json:"taxExempt"`

	// Surcharge is a flag to add a surcharge to the transaction to cover credit
	// card fees, if permitted.
	Surcharge bool `json:"surcharge"`

	// CashDiscount is a flag that applies a discount to negate the surcharge for
	// debit transactions or other surcharge ineligible payment methods.
	CashDiscount bool `json:"cashDiscount"`

	// SigFile is a location on the filesystem which a customer signature should
	// be written to.
	SigFile string `json:"sigFile,omitempty"`

	// SigFormat specifies the image format to be used for returning signatures.
	SigFormat SignatureFormat `json:"sigFormat,omitempty"`

	// SigWidth is the width that the signature image should be scaled to,
	// preserving the aspect ratio. If not provided, the signature is returned in
	// the terminal's max resolution.
	SigWidth int `json:"sigWidth,omitempty"`

	// DisableSignature specifies whether or not signature prompt should be
	// skipped on the terminal. The terminal will indicate whether or not a
	// signature is required by the card in the receipt suggestions response.
	DisableSignature bool `json:"disableSignature,omitempty"`

	// TipAmount is the tip amount.
	TipAmount string `json:"tipAmount,omitempty"`

	// TaxAmount is the tax amount.
	TaxAmount string `json:"taxAmount,omitempty"`

	// TerminalName is the name of the target payment terminal.
	TerminalName string `json:"terminalName,omitempty"`

	// ResetConnection forces the terminal cloud connection to be reset while a
	// transactions is in flight. This is a diagnostic settings that can be used
	// only for test transactions.
	ResetConnection bool `json:"resetConnection"`

	// TransactionID can be used to update a pre-auth to a new amount, sometimes
	// called incremental auth.
	TransactionID string `json:"transactionId,omitempty"`

	// OnlineAuthCode is used to validate online gift card authorizations.
	OnlineAuthCode string `json:"onlineAuthCode,omitempty"`

	// Enroll indicates that the payment method should be added to the token
	// vault alongside the authorization.
	Enroll bool `json:"enroll,omitempty"`

	// BypassDupeFilter indicates duplicate detection should be bypassed.
	BypassDupeFilter bool `json:"bypassDupeFilter,omitempty"`

	// Description contains a narrative description of the transaction.
	Description string `json:"description,omitempty"`

	// PromptForTip indicates that the terminal should request a tip from the
	// user before starting the transaction.
	PromptForTip bool `json:"promptForTip,omitempty"`

	// CashBackEnabled indicates that cash back should be enabled for supported
	// cards.
	CashBackEnabled bool `json:"cashBackEnabled,omitempty"`

	// CardOnFile indicates that this transaction should be treated as MOTO with
	// a card on file.
	CardOnFile bool `json:"cardOnFile,omitempty"`

	// Recurring indicates that this transaction should be treated as a recurring
	// transaction.
	Recurring bool `json:"recurring,omitempty"`

	// Cit manually sets the CIT (Customer Initiated Transaction) flag.
	Cit bool `json:"cit,omitempty"`

	// Mit manually sets the MIT (Merchant Initiated Transaction) flag.
	Mit bool `json:"mit,omitempty"`

	// Subscription indicates that this transaction should be treated as a
	// subscription recurring transaction.
	Subscription bool `json:"subscription,omitempty"`

	// PurchaseOrderNumber is the purchase order number, if known.
	PurchaseOrderNumber string `json:"purchaseOrderNumber,omitempty"`

	// SupplierReferenceNumber is the supplier reference number, if known.
	SupplierReferenceNumber string `json:"supplierReferenceNumber,omitempty"`

	// LineItems is an item to display. Can be overwritten or appended, based on
	// the request type.
	LineItems []*TransactionDisplayItem `json:"lineItems"`

	// AltPrices is a map of alternate currencies and the price in each currency.
	// Use only if you want to set your own exchange rate for a crypto
	// transaction.
	AltPrices map[string]string `json:"altPrices,omitempty"`

	// Customer contains customer information.
	Customer *Customer `json:"customer"`

	// RoundingMode indicates how partial pennies should be rounded for
	// calculated values like surcharges. Rounding up is the default behavior.
	RoundingMode *RoundingMode `json:"roundingMode"`

	// HealthcareMetadata contains details for HSA/FSA transactions.
	HealthcareMetadata *HealthcareMetadata `json:"healthcareMetadata,omitempty"`

	// Cryptocurrency indicates that the transaction should be a cryptocurrency
	// transaction. Value should be a crypto currency code (ETH, BTC) or ANY to
	// prompt the user to choose from supported cryptocurrencies.
	Cryptocurrency *string `json:"cryptocurrency,omitempty"`

	// CryptoNetwork is an optional parameter that can be used to force a crypto
	// transaction onto a level one or level two network. Valid values are L1 and
	// L2. Defaults to L1.
	CryptoNetwork *string `json:"cryptoNetwork,omitempty"`

	// CryptoReceiveAddress can be used to specify a specific receive address for
	// a crypto transaction. Disabled by default. This should only be used by
	// sophisticated users with access to properly configured hot wallets.
	CryptoReceiveAddress *string `json:"cryptoReceiveAddress,omitempty"`

	// PaymentRequestLabel can optionally add a label to the payment request if
	// the target cryptocurrency supports labels. Defaults to the merchant's DBA
	// Name.
	PaymentRequestLabel *string `json:"paymentRequestLabel,omitempty"`

	// PaymentRequestMessage can optionally add a message to the payment request
	// if the target cryptocurrency supports labels. Defaults to empty.
	PaymentRequestMessage *string `json:"paymentRequestMessage,omitempty"`

	// SimulateChipRejection instructs the terminal to simulate a post auth chip
	// rejection that would trigger an automatic reversal.
	SimulateChipRejection bool `json:"simulateChipRejection,omitempty"`

	// SimulateOutOfOrderReversal instructs the terminal to simulate an out of
	// order automatic reversal.
	SimulateOutOfOrderReversal bool `json:"simulateOutOfOrderReversal,omitempty"`

	// AsyncReversals causes auto-reversals on the terminal to be executed
	// asyncronously. Use with caution and in conjunction with the transaction
	// status API.
	AsyncReversals bool `json:"asyncReversals,omitempty"`

	// PassthroughSurcharge is a passthrough surcharge amount. This surcharge
	// amount will be passed directly to the gateway and is not directly
	// calculated.
	PassthroughSurcharge string `json:"passthroughSurcharge,omitempty"`

	// Healthcare marks a transaction as HSA/FSA.
	Healthcare bool `json:"healthcare,omitempty"`

	// HealthcareTotal is the total amount to process as healthcare.
	HealthcareTotal string `json:"healthcareTotal,omitempty"`

	// EBTTotal is the total amount to process as ebt.
	EBTTotal string `json:"ebtTotal,omitempty"`

	// CardMetadataLookup indicates that this transaction will include a card
	// metadata lookup.
	CardMetadataLookup bool `json:"cardMetadataLookup,omitempty"`

	// TotalDiscountAmount is the total discount amount for the transaction, and
	// will overide additive logic for line item discounts.
	TotalDiscountAmount string `json:"totalDiscountAmount,omitempty"`

	// ShippingAmount indicates the shipping cost associated with the
	// transaction.
	ShippingAmount string `json:"shippingAmount,omitempty"`

	// DutyAmount indicates the duty amount associated with the transaction.
	DutyAmount string `json:"dutyAmount,omitempty"`

	// ProcessorID indicates the processor ID associated with the transaction.
	ProcessorID string `json:"processorId,omitempty"`

	// ExternalCustomerID indicates the external customer ID associated with the
	// transaction.
	ExternalCustomerID string `json:"externalCustomerId,omitempty"`

	// DestinationCountryCode three character, numeric, ship-to country code.
	// Defaults to '840' (USA) if not specified.
	DestinationCountryCode string `json:"destinationCountryCode,omitempty"`

	// ShipFromPostalCode nine character postal code for shipping origin
	// addresses. For US addresses, this is a 5+4 ZIP or five digit ZIP.
	ShipFromPostalCode string `json:"shipFromPostalCode,omitempty"`

	// ShipToPostalCode nine character postal code for shipping destination
	// addresses. For US addresses, this is a 5+4 ZIP or five digit ZIP.
	ShipToPostalCode string `json:"shipToPostalCode,omitempty"`

	// OrderDate indicates the purchase order date.
	OrderDate *time.Time `json:"orderDate,omitempty"`

	// ShipmentCount indicates the number of shipments the original authorization
	// will be broken into.
	ShipmentCount int `json:"shipmentCount"`

	// ShipmentNumber indicates which shipment this particular capture is for.
	ShipmentNumber int `json:"shipmentNumber"`

	// ExternalPartnerMetadata is an optional field that can be used to pass
	// through data to external partners.
	ExternalPartnerMetadata string `json:"externalPartnerMetadata,omitempty"`

	// ExternalCustomerEmail is the external customer's email address.
	ExternalCustomerEmail string `json:"externalCustomerEmail,omitempty"`

	// ExternalCustomerPhone is the external customer's phone number.
	ExternalCustomerPhone string `json:"externalCustomerPhone,omitempty"`

	// ExternalCustomerCompany is the external customer's company name.
	ExternalCustomerCompany string `json:"externalCustomerCompany,omitempty"`
}

// CardMetadata contains essential information about a payment card derived
// from its BIN/IIN.
type CardMetadata struct {
	// CardBrand is the brand or network of the card (e.g., Visa, Mastercard,
	// Amex).
	CardBrand string `json:"cardBrand"`

	// IssuerName is the name of the financial institution that issued the card.
	IssuerName string `json:"issuerName"`

	// L3 indicates whether the card supports Level 3 processing for detailed
	// transaction data.
	L3 bool `json:"l3"`

	// L2 indicates whether the card supports Level 2 processing for additional
	// transaction data.
	L2 bool `json:"l2"`

	// ProductType is the general category or type of the card product.
	ProductType string `json:"productType"`

	// ProductName is the specific name or designation of the card product.
	ProductName string `json:"productName"`

	// EBT indicates whether the card is an Electronic Benefit Transfer (EBT)
	// card.
	EBT bool `json:"ebt"`

	// Debit indicates whether the card is a debit card.
	Debit bool `json:"debit"`

	// Healthcare indicates whether the card is a healthcare-specific payment
	// card.
	Healthcare bool `json:"healthcare"`

	// Prepaid indicates whether the card is a prepaid card.
	Prepaid bool `json:"prepaid"`

	// Region is the geographical region associated with the card's issuer.
	Region string `json:"region"`

	// Country is the country associated with the card's issuer.
	Country string `json:"country"`
}

// RefundRequest contains a refund request.
type RefundRequest struct {
	// Timeout is the request timeout in seconds.
	Timeout int `json:"timeout"`

	// Test specifies whether or not to route transaction to the test gateway.
	Test bool `json:"test"`

	// TransactionRef contains a user-assigned reference that can be used to
	// recall or reverse transactions.
	TransactionRef string `json:"transactionRef,omitempty"`

	// AutogeneratedRef indicates that the transaction reference was
	// autogenerated and should be ignored for the purposes of duplicate
	// detection.
	AutogeneratedRef bool `json:"autogeneratedRef"`

	// Async defers the response to the transaction and returns immediately.
	// Callers should retrive the transaction result using the Transaction Status
	// API.
	Async bool `json:"async"`

	// Queue adds the transaction to the queue and returns immediately. Callers
	// should retrive the transaction result using the Transaction Status API.
	Queue bool `json:"queue"`

	// WaitForRemovedCard specifies whether or not the request should block until
	// all cards have been removed from the card reader.
	WaitForRemovedCard bool `json:"waitForRemovedCard,omitempty"`

	// Force causes a transaction to override any in-progress transactions.
	Force bool `json:"force,omitempty"`

	// OrderRef is an identifier from an external point of sale system.
	OrderRef string `json:"orderRef,omitempty"`

	// DestinationAccount is the settlement account for merchants with split
	// settlements.
	DestinationAccount string `json:"destinationAccount,omitempty"`

	// TestCase can include a code used to trigger simulated conditions for the
	// purposes of testing and certification. Valid for test merchant accounts
	// only.
	TestCase string `json:"testCase,omitempty"`

	// Token is the payment token to be used for this transaction. This should be
	// used for recurring transactions. The /enroll endpoint ignores this field.
	Token string `json:"token,omitempty"`

	// Track1 contains track 1 magnetic stripe data.
	Track1 string `json:"track1,omitempty"`

	// Track2 contains track 2 magnetic stripe data.
	Track2 string `json:"track2,omitempty"`

	// PAN contains the primary account number. We recommend using the terminal
	// or e-commerce tokenization libraries instead of passing account numbers in
	// directly, as this would put your application in PCI scope.
	PAN string `json:"pan,omitempty"`

	// RoutingNumber is the ACH routing number for ACH transactions.
	RoutingNumber string `json:"routingNumber,omitempty"`

	// CardholderName is the cardholder name. Only required if the request
	// includes a primary account number or track data.
	CardholderName string `json:"cardholderName,omitempty"`

	// ExpMonth is the card expiration month for use with PAN based transactions.
	ExpMonth string `json:"expMonth,omitempty"`

	// ExpYear is the card expiration year for use with PAN based transactions.
	ExpYear string `json:"expYear,omitempty"`

	// CVV is the card CVV for use with PAN based transactions.
	CVV string `json:"cvv,omitempty"`

	// Address is the cardholder address for use with address verification.
	Address string `json:"address,omitempty"`

	// PostalCode is the cardholder postal code for use with address
	// verification.
	PostalCode string `json:"postalCode,omitempty"`

	// Country is the cardholder country.
	Country string `json:"country,omitempty"`

	// ManualEntry specifies that the payment entry method is a manual keyed
	// transaction. If this is true, no other payment method will be accepted.
	ManualEntry bool `json:"manualEntry,omitempty"`

	// KSN is the key serial number used for DUKPT encryption.
	KSN string `json:"ksn,omitempty"`

	// PINBlock is the encrypted pin block.
	PINBlock string `json:"pinBlock,omitempty"`

	// CardType designates categories of cards: credit, debit, EBT.
	CardType CardType `json:"cardType,omitempty"`

	// PaymentType designates brands of payment methods: Visa, Discover, etc.
	PaymentType string `json:"paymentType,omitempty"`

	// TransactionID is the ID of the previous transaction being referenced.
	TransactionID string `json:"transactionId"`

	// CurrencyCode indicates the transaction currency code.
	CurrencyCode string `json:"currencyCode"`

	// Amount is the requested amount.
	Amount string `json:"amount"`

	// TaxExempt indicates that the request is tax exempt. Only required for tax
	// exempt level 2 processing.
	TaxExempt bool `json:"taxExempt"`

	// Surcharge is a flag to add a surcharge to the transaction to cover credit
	// card fees, if permitted.
	Surcharge bool `json:"surcharge"`

	// CashDiscount is a flag that applies a discount to negate the surcharge for
	// debit transactions or other surcharge ineligible payment methods.
	CashDiscount bool `json:"cashDiscount"`

	// SigFile is a location on the filesystem which a customer signature should
	// be written to.
	SigFile string `json:"sigFile,omitempty"`

	// SigFormat specifies the image format to be used for returning signatures.
	SigFormat SignatureFormat `json:"sigFormat,omitempty"`

	// SigWidth is the width that the signature image should be scaled to,
	// preserving the aspect ratio. If not provided, the signature is returned in
	// the terminal's max resolution.
	SigWidth int `json:"sigWidth,omitempty"`

	// DisableSignature specifies whether or not signature prompt should be
	// skipped on the terminal. The terminal will indicate whether or not a
	// signature is required by the card in the receipt suggestions response.
	DisableSignature bool `json:"disableSignature,omitempty"`

	// TipAmount is the tip amount.
	TipAmount string `json:"tipAmount,omitempty"`

	// TaxAmount is the tax amount.
	TaxAmount string `json:"taxAmount,omitempty"`

	// TerminalName is the name of the target payment terminal.
	TerminalName string `json:"terminalName,omitempty"`

	// ResetConnection forces the terminal cloud connection to be reset while a
	// transactions is in flight. This is a diagnostic settings that can be used
	// only for test transactions.
	ResetConnection bool `json:"resetConnection"`

	// HealthcareMetadata contains details for HSA/FSA transactions.
	HealthcareMetadata *HealthcareMetadata `json:"healthcareMetadata,omitempty"`

	// SimulateChipRejection instructs the terminal to simulate a post auth chip
	// rejection that would trigger an automatic reversal.
	SimulateChipRejection bool `json:"simulateChipRejection,omitempty"`

	// SimulateOutOfOrderReversal instructs the terminal to simulate an out of
	// order automatic reversal.
	SimulateOutOfOrderReversal bool `json:"simulateOutOfOrderReversal,omitempty"`

	// AsyncReversals causes auto-reversals on the terminal to be executed
	// asyncronously. Use with caution and in conjunction with the transaction
	// status API.
	AsyncReversals bool `json:"asyncReversals,omitempty"`

	// Cit manually sets the CIT (Customer Initiated Transaction) flag.
	Cit bool `json:"cit,omitempty"`

	// Mit manually sets the MIT (Merchant Initiated Transaction) flag.
	Mit bool `json:"mit,omitempty"`

	// CardMetadataLookup indicates that this transaction will include a card
	// metadata lookup.
	CardMetadataLookup bool `json:"cardMetadataLookup,omitempty"`
}

// ClearTerminalRequest contains the information needed to enroll a new
// payment method in the token vault.
type ClearTerminalRequest struct {
	// Timeout is the request timeout in seconds.
	Timeout int `json:"timeout"`

	// Test specifies whether or not to route transaction to the test gateway.
	Test bool `json:"test"`

	// TransactionRef contains a user-assigned reference that can be used to
	// recall or reverse transactions.
	TransactionRef string `json:"transactionRef,omitempty"`

	// AutogeneratedRef indicates that the transaction reference was
	// autogenerated and should be ignored for the purposes of duplicate
	// detection.
	AutogeneratedRef bool `json:"autogeneratedRef"`

	// Async defers the response to the transaction and returns immediately.
	// Callers should retrive the transaction result using the Transaction Status
	// API.
	Async bool `json:"async"`

	// Queue adds the transaction to the queue and returns immediately. Callers
	// should retrive the transaction result using the Transaction Status API.
	Queue bool `json:"queue"`

	// WaitForRemovedCard specifies whether or not the request should block until
	// all cards have been removed from the card reader.
	WaitForRemovedCard bool `json:"waitForRemovedCard,omitempty"`

	// Force causes a transaction to override any in-progress transactions.
	Force bool `json:"force,omitempty"`

	// OrderRef is an identifier from an external point of sale system.
	OrderRef string `json:"orderRef,omitempty"`

	// DestinationAccount is the settlement account for merchants with split
	// settlements.
	DestinationAccount string `json:"destinationAccount,omitempty"`

	// TestCase can include a code used to trigger simulated conditions for the
	// purposes of testing and certification. Valid for test merchant accounts
	// only.
	TestCase string `json:"testCase,omitempty"`

	// TerminalName is the name of the target payment terminal.
	TerminalName string `json:"terminalName,omitempty"`

	// ResetConnection forces the terminal cloud connection to be reset while a
	// transactions is in flight. This is a diagnostic settings that can be used
	// only for test transactions.
	ResetConnection bool `json:"resetConnection"`
}

// AuthorizationResponse contains the response to an authorization request.
type AuthorizationResponse struct {
	// Success indicates whether or not the request succeeded.
	Success bool `json:"success"`

	// Error is the error, if an error occurred.
	Error string `json:"error"`

	// ResponseDescription contains a narrative description of the transaction
	// result.
	ResponseDescription string `json:"responseDescription"`

	// Approved indicates that the transaction was approved.
	Approved bool `json:"approved"`

	// AuthCode is the auth code from the payment network.
	AuthCode string `json:"authCode,omitempty"`

	// AuthResponseCode is the code returned by the terminal or the card issuer
	// to indicate the disposition of the message.
	AuthResponseCode string `json:"authResponseCode,omitempty"`

	// TransactionID is the ID assigned to the transaction.
	TransactionID string `json:"transactionId"`

	// BatchID is the ID assigned to the batch.
	BatchID string `json:"batchId,omitempty"`

	// TransactionRef is the transaction reference string assigned to the
	// transaction request. If no transaction ref was assiged on the request,
	// then the gateway will randomly generate one.
	TransactionRef string `json:"transactionRef,omitempty"`

	// TransactionType is the type of transaction.
	TransactionType string `json:"transactionType"`

	// Timestamp is the timestamp of the transaction.
	Timestamp string `json:"timestamp"`

	// TickBlock is the hash of the last tick block.
	TickBlock string `json:"tickBlock"`

	// Test indicates that the transaction was processed on the test gateway.
	Test bool `json:"test"`

	// DestinationAccount is the settlement account for merchants with split
	// settlements.
	DestinationAccount string `json:"destinationAccount,omitempty"`

	// Sig is the ECC signature of the response. Can be used to ensure that it
	// was signed by the terminal and detect man-in-the middle attacks.
	Sig string `json:"sig,omitempty"`

	// PartialAuth indicates whether or not the transaction was approved for a
	// partial amount.
	PartialAuth bool `json:"partialAuth"`

	// AltCurrency indicates whether or not an alternate currency was used.
	AltCurrency bool `json:"altCurrency"`

	// FSAAuth indicates whether or not a request was settled on an FSA card.
	FSAAuth bool `json:"fsaAuth"`

	// CurrencyCode is the currency code used for the transaction.
	CurrencyCode string `json:"currencyCode"`

	// RequestedAmount is the requested amount.
	RequestedAmount string `json:"requestedAmount"`

	// AuthorizedAmount is the authorized amount. May not match the requested
	// amount in the event of a partial auth.
	AuthorizedAmount string `json:"authorizedAmount"`

	// RemainingBalance is the remaining balance on the payment method.
	RemainingBalance string `json:"remainingBalance"`

	// TipAmount is the tip amount.
	TipAmount string `json:"tipAmount"`

	// TaxAmount is the tax amount.
	TaxAmount string `json:"taxAmount"`

	// RequestedCashBackAmount is the cash back amount the customer requested
	// during the transaction.
	RequestedCashBackAmount string `json:"requestedCashBackAmount"`

	// AuthorizedCashBackAmount is the amount of cash back authorized by the
	// gateway. This amount will be the entire amount requested, or zero.
	AuthorizedCashBackAmount string `json:"authorizedCashBackAmount"`

	// Confirmed indicates that the transaction has met the standard criteria for
	// confirmation on the network. (For example, 6 confirmations for level one
	// bitcoin.)
	Confirmed bool `json:"confirmed"`

	// CryptoAuthorizedAmount is the amount submitted to the blockchain.
	CryptoAuthorizedAmount string `json:"cryptoAuthorizedAmount"`

	// CryptoNetworkFee is the network level fee assessed for the transaction
	// denominated in cryptocurrency. This fee goes to channel operators and
	// crypto miners, not BlockChyp.
	CryptoNetworkFee string `json:"cryptoNetworkFee"`

	// Cryptocurrency is the three letter cryptocurrency code used for the
	// transactions.
	Cryptocurrency string `json:"cryptocurrency"`

	// CryptoNetwork indicates whether or not the transaction was processed on
	// the level one or level two network.
	CryptoNetwork string `json:"cryptoNetwork"`

	// CryptoReceiveAddress the address on the crypto network the transaction was
	// sent to.
	CryptoReceiveAddress string `json:"cryptoReceiveAddress"`

	// CryptoBlock hash or other identifier that identifies the block on the
	// cryptocurrency network, if available or relevant.
	CryptoBlock string `json:"cryptoBlock"`

	// CryptoTransactionID hash or other transaction identifier that identifies
	// the transaction on the cryptocurrency network, if available or relevant.
	CryptoTransactionID string `json:"cryptoTransactionId"`

	// CryptoPaymentRequest is the payment request URI used for the transaction,
	// if available.
	CryptoPaymentRequest string `json:"cryptoPaymentRequest"`

	// CryptoStatus is used for additional status information related to crypto
	// transactions.
	CryptoStatus string `json:"cryptoStatus"`

	// Token is the payment token, if the payment was enrolled in the vault.
	Token string `json:"token,omitempty"`

	// EntryMethod is the entry method for the transaction (CHIP, MSR, KEYED,
	// etc).
	EntryMethod string `json:"entryMethod,omitempty"`

	// PaymentType is the card brand (VISA, MC, AMEX, DEBIT, etc).
	PaymentType string `json:"paymentType,omitempty"`

	// Network provides network level detail on how a transaction was routed,
	// especially for debit transactions.
	Network string `json:"network,omitempty"`

	// Logo identifies the card association based on bin number. Used primarily
	// used to indicate the major logo on a card, even when debit transactions
	// are routed on a different network.
	Logo string `json:"logo,omitempty"`

	// MaskedPAN is the masked primary account number.
	MaskedPAN string `json:"maskedPan,omitempty"`

	// PublicKey is the BlockChyp public key if the user presented a BlockChyp
	// payment card.
	PublicKey string `json:"publicKey,omitempty"`

	// ScopeAlert indicates that the transaction did something that would put the
	// system in PCI scope.
	ScopeAlert bool `json:"ScopeAlert,omitempty"`

	// CardHolder is the cardholder name.
	CardHolder string `json:"cardHolder,omitempty"`

	// ExpMonth is the card expiration month in MM format.
	ExpMonth string `json:"expMonth,omitempty"`

	// ExpYear is the card expiration year in YY format.
	ExpYear string `json:"expYear,omitempty"`

	// PostalCode is the card postal code.
	PostalCode string `json:"postalCode,omitempty"`

	// Address is the card address.
	Address string `json:"address,omitempty"`

	// Country is the card country.
	Country string `json:"country,omitempty"`

	// AVSResponse contains address verification results if address information
	// was submitted.
	AVSResponse AVSResponse `json:"avsResponse"`

	// CVVResponse contains the CVV verification result if CVV was submitted.
	CVVResponse string `json:"cvvResponse,omitempty"`

	// ReceiptSuggestions contains suggested receipt fields.
	ReceiptSuggestions ReceiptSuggestions `json:"receiptSuggestions"`

	// Customer contains customer data, if any. Preserved for reverse
	// compatibility.
	Customer *Customer `json:"customer"`

	// Customers contains customer data, if any.
	Customers []Customer `json:"customers"`

	// SigFile contains the hex encoded signature data.
	SigFile string `json:"sigFile,omitempty"`

	// WhiteListedCard contains card BIN ranges can be whitelisted so that they
	// are read instead of being processed directly. This is useful for
	// integration with legacy gift card systems.
	WhiteListedCard *WhiteListedCard `json:"whiteListedCard"`

	// StoreAndForward indicates that the transaction was flagged for store and
	// forward due to network problems.
	StoreAndForward bool `json:"storeAndForward"`

	// Status indicates the current status of a transaction.
	Status string `json:"status"`

	// CardMetadata contains details about a payment card derived from its
	// BIN/IIN.
	CardMetadata *CardMetadata `json:"cardMetadata,omitempty"`

	// EnhancedDataPassed indicates whether enhanced data was passed for the
	// transaction.
	EnhancedDataPassed *bool `json:"enhancedDataPassed,omitempty"`
}

// TransactionDisplayDiscount is an item level discount for transaction
// display. Discounts never combine.
type TransactionDisplayDiscount struct {
	// Description is the discount description.
	Description string `json:"description"`

	// Amount is the amount of the discount.
	Amount string `json:"amount"`
}

// TransactionDisplayItem is an item category in a transaction display. Groups
// combine if their descriptions match. Calculated subtotal amounts are
// rounded to two decimal places of precision. Quantity is a floating point
// number that is not rounded at all.
type TransactionDisplayItem struct {
	// ID is a unique value identifying the item. This is not required, but
	// recommended since it is required to update or delete line items.
	ID string `json:"id"`

	// Description is a description of the line item.
	Description string `json:"description"`

	// Price is the price of the line item.
	Price string `json:"price"`

	// Quantity is the quantity of the line item.
	Quantity float64 `json:"quantity"`

	// Extended is an item category in a transaction display. Groups combine if
	// their descriptions match. Calculated subtotal amounts are rounded to two
	// decimal places of precision. Quantity is a floating point number that is
	// not rounded at all.
	Extended string `json:"extended"`

	// UnitCode is an alphanumeric code for units of measurement as used in
	// international trade.
	UnitCode string `json:"unitCode"`

	// CommodityCode is an international description code of the item.
	CommodityCode string `json:"commodityCode"`

	// ProductCode is a merchant-defined description code of the item.
	ProductCode string `json:"productCode"`

	// Discounts are displayed under their corresponding item.
	Discounts []*TransactionDisplayDiscount `json:"discounts"`

	// TaxAmount is the amount of any value added taxes which apply to the item.
	TaxAmount string `json:"taxAmount,omitempty"`

	// TaxRate is the tax rate as a percentage. Example: '8.5' for 8.5% tax rate.
	TaxRate string `json:"taxRate,omitempty"`

	// DiscountCode is how tax was applied to discounted items. '0' = no
	// discount, '1' = tax calculated after discount, '2' = taxcalculated before
	// discount.
	DiscountCode string `json:"discountCode,omitempty"`
}

// TransactionDisplayTransaction contains the items to display on a terminal.
type TransactionDisplayTransaction struct {
	// Subtotal is the subtotal to display.
	Subtotal string `json:"subtotal"`

	// Tax is the tax to display.
	Tax string `json:"tax"`

	// Total is the total to display.
	Total string `json:"total"`

	// Items is an item to display. Can be overwritten or appended, based on the
	// request type.
	Items []*TransactionDisplayItem `json:"items"`
}

// TransactionDisplayRequest is used to start or update a transaction line
// item display on a terminal.
type TransactionDisplayRequest struct {
	// Timeout is the request timeout in seconds.
	Timeout int `json:"timeout"`

	// Test specifies whether or not to route transaction to the test gateway.
	Test bool `json:"test"`

	// TransactionRef contains a user-assigned reference that can be used to
	// recall or reverse transactions.
	TransactionRef string `json:"transactionRef,omitempty"`

	// AutogeneratedRef indicates that the transaction reference was
	// autogenerated and should be ignored for the purposes of duplicate
	// detection.
	AutogeneratedRef bool `json:"autogeneratedRef"`

	// Async defers the response to the transaction and returns immediately.
	// Callers should retrive the transaction result using the Transaction Status
	// API.
	Async bool `json:"async"`

	// Queue adds the transaction to the queue and returns immediately. Callers
	// should retrive the transaction result using the Transaction Status API.
	Queue bool `json:"queue"`

	// WaitForRemovedCard specifies whether or not the request should block until
	// all cards have been removed from the card reader.
	WaitForRemovedCard bool `json:"waitForRemovedCard,omitempty"`

	// Force causes a transaction to override any in-progress transactions.
	Force bool `json:"force,omitempty"`

	// OrderRef is an identifier from an external point of sale system.
	OrderRef string `json:"orderRef,omitempty"`

	// DestinationAccount is the settlement account for merchants with split
	// settlements.
	DestinationAccount string `json:"destinationAccount,omitempty"`

	// TestCase can include a code used to trigger simulated conditions for the
	// purposes of testing and certification. Valid for test merchant accounts
	// only.
	TestCase string `json:"testCase,omitempty"`

	// TerminalName is the name of the target payment terminal.
	TerminalName string `json:"terminalName,omitempty"`

	// ResetConnection forces the terminal cloud connection to be reset while a
	// transactions is in flight. This is a diagnostic settings that can be used
	// only for test transactions.
	ResetConnection bool `json:"resetConnection"`

	// Transaction transaction to display on the terminal.
	Transaction *TransactionDisplayTransaction `json:"transaction"`
}

// HeartbeatResponse contains the response to a basic API health check. If the
// security context permits it, the response may also include the public key
// of the current merchant.
type HeartbeatResponse struct {
	// Success indicates whether or not the request succeeded.
	Success bool `json:"success"`

	// Error is the error, if an error occurred.
	Error string `json:"error"`

	// ResponseDescription contains a narrative description of the transaction
	// result.
	ResponseDescription string `json:"responseDescription"`

	// Timestamp is the timestamp of the heartbeat.
	Timestamp time.Time `json:"timestamp"`

	// Clockchain is the public key of the clockchain. This is blockchain stuff
	// that you don't really need to worry about. It is a base 58 encoded and
	// compressed eliptic curve public key. For the production clockchain, this
	// will always be: '3cuhsckVUd9HzMjbdUSW17aY5kCcm1d6YAphJMUwmtXRj7WLyU'.
	Clockchain string `json:"clockchain"`

	// LatestTick is the hash of the last tick block.
	LatestTick string `json:"latestTick"`

	// MerchantPublicKey is the public key for the merchant's blockchain.
	MerchantPublicKey string `json:"merchantPk"`
}

// TerminalStatusRequest contains a request for the status of a terminal.
type TerminalStatusRequest struct {
	// Timeout is the request timeout in seconds.
	Timeout int `json:"timeout"`

	// Test specifies whether or not to route transaction to the test gateway.
	Test bool `json:"test"`

	// TransactionRef contains a user-assigned reference that can be used to
	// recall or reverse transactions.
	TransactionRef string `json:"transactionRef,omitempty"`

	// AutogeneratedRef indicates that the transaction reference was
	// autogenerated and should be ignored for the purposes of duplicate
	// detection.
	AutogeneratedRef bool `json:"autogeneratedRef"`

	// Async defers the response to the transaction and returns immediately.
	// Callers should retrive the transaction result using the Transaction Status
	// API.
	Async bool `json:"async"`

	// Queue adds the transaction to the queue and returns immediately. Callers
	// should retrive the transaction result using the Transaction Status API.
	Queue bool `json:"queue"`

	// WaitForRemovedCard specifies whether or not the request should block until
	// all cards have been removed from the card reader.
	WaitForRemovedCard bool `json:"waitForRemovedCard,omitempty"`

	// Force causes a transaction to override any in-progress transactions.
	Force bool `json:"force,omitempty"`

	// OrderRef is an identifier from an external point of sale system.
	OrderRef string `json:"orderRef,omitempty"`

	// DestinationAccount is the settlement account for merchants with split
	// settlements.
	DestinationAccount string `json:"destinationAccount,omitempty"`

	// TestCase can include a code used to trigger simulated conditions for the
	// purposes of testing and certification. Valid for test merchant accounts
	// only.
	TestCase string `json:"testCase,omitempty"`

	// TerminalName is the name of the target payment terminal.
	TerminalName string `json:"terminalName,omitempty"`

	// ResetConnection forces the terminal cloud connection to be reset while a
	// transactions is in flight. This is a diagnostic settings that can be used
	// only for test transactions.
	ResetConnection bool `json:"resetConnection"`
}

// TerminalStatusResponse contains the current status of a terminal.
type TerminalStatusResponse struct {
	// Success indicates whether or not the request succeeded.
	Success bool `json:"success"`

	// Error is the error, if an error occurred.
	Error string `json:"error"`

	// ResponseDescription contains a narrative description of the transaction
	// result.
	ResponseDescription string `json:"responseDescription"`

	// Idle indicates that the terminal is idle.
	Idle bool `json:"idle"`

	// CardInSlot indicates whether or not a card is currently in the card slot.
	CardInSlot bool `json:"cardInSlot"`

	// Status contains the operation that the terminal is performing.
	Status string `json:"status"`

	// TransactionRef contains the transaction reference for an ongoing
	// transaction, if one was specified at request time.
	TransactionRef string `json:"transactionRef"`

	// Since is the timestamp of the last status change.
	Since time.Time `json:"since"`
}

// HealthcareMetadata contains fields for HSA/FSA transactions.
type HealthcareMetadata struct {
	// Types is a list of healthcare categories in the transaction.
	Types []HealthcareGroup `json:"types"`

	// IIASVerified indicates that the purchased items were verified against an
	// Inventory Information Approval System (IIAS).
	IIASVerified bool `json:"iiasVerified"`

	// IIASExempt indicates that the transaction is exempt from IIAS
	// verification.
	IIASExempt bool `json:"iiasExempt"`
}

// HealthcareGroup is a group of fields for a specific type of healthcare.
type HealthcareGroup struct {
	// Type the type of healthcare cost.
	Type HealthcareType `json:"type"`

	// Amount is the amount of this type.
	Amount string `json:"amount"`

	// ProviderID the provider ID used for Mastercard and Discover IIAS requests.
	ProviderID string `json:"providerId"`

	// ServiceTypeCode the service type code used for Mastercard and Discover
	// IIAS requests.
	ServiceTypeCode string `json:"serviceTypeCode"`

	// PayerOrCarrierID thr payer ID/carrier ID used for Mastercard and Discover
	// IIAS requests.
	PayerOrCarrierID string `json:"payerOrCarrierId"`

	// ApprovalRejectReasonCode the approval or reject reason code used for
	// Mastercard and Discover IIAS requests.
	ApprovalRejectReasonCode string `json:"approvalRejectReasonCode"`
}

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

// TerminalDeactivationRequest models a terminal deactivation request.
type TerminalDeactivationRequest struct {
	// Timeout is the request timeout in seconds.
	Timeout int `json:"timeout"`

	// Test specifies whether or not to route transaction to the test gateway.
	Test bool `json:"test"`

	// TerminalName is the terminal name assigned to the terminal.
	TerminalName string `json:"terminalName"`

	// TerminalID is the id assigned by BlockChyp to the terminal.
	TerminalID string `json:"terminalId"`
}

// TerminalActivationRequestV2 models a terminal activation request. The
// merchant is derived from the Stax bearer token, so no merchant id is
// accepted.
type TerminalActivationRequestV2 struct {
	// Timeout is the request timeout in seconds.
	Timeout int `json:"timeout"`

	// Test specifies whether or not to route transaction to the test gateway.
	Test bool `json:"test"`

	// TerminalName is the name to be assigned to the terminal. Must be unique
	// for the merchant account.
	TerminalName string `json:"terminalName"`

	// ActivationCode is the terminal activation code displayed on the terminal.
	ActivationCode string `json:"activationCode"`
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

// UploadMetadata models information needed to process a file upload.
type UploadMetadata struct {
	// Timeout is the request timeout in seconds.
	Timeout int `json:"timeout"`

	// Test specifies whether or not to route transaction to the test gateway.
	Test bool `json:"test"`

	// UploadID optional id used to track status and progress of an upload while
	// in progress.
	UploadID string `json:"uploadId"`

	// FileSize is the size of the file to be uploaded in bytes.
	FileSize int64 `json:"fileSize"`

	// FileName is the name of file to be uploaded.
	FileName string `json:"fileName"`
}

// AuthRequest models a Stax Payments charge or preauth request.
type AuthRequest struct {
	// TipAmount is the tip amount.
	TipAmount string `json:"tipAmount,omitempty"`

	// TaxAmount is the tax amount.
	TaxAmount string `json:"taxAmount,omitempty"`

	// TerminalName is the name of the terminal that should run the transaction.
	TerminalName string `json:"terminalName,omitempty"`

	// Amount is the requested amount.
	Amount string `json:"amount"`

	// CurrencyCode is the transaction currency code.
	CurrencyCode string `json:"currencyCode"`

	// PromptForTip indicates that the terminal should prompt for a tip.
	PromptForTip bool `json:"promptForTip,omitempty"`

	// Test specifies whether or not to route the transaction to the test
	// gateway.
	Test bool `json:"test"`

	// Enroll indicates that the payment method should be added to the token
	// vault alongside the authorization.
	Enroll bool `json:"enroll,omitempty"`

	// ExternalPartnerMetadata is an optional field that can be used to pass
	// through data to external partners.
	ExternalPartnerMetadata string `json:"externalPartnerMetadata,omitempty"`
}

// AuthResponse models a Stax Payments charge or preauth response.
type AuthResponse struct {
	// Success indicates whether or not the request succeeded.
	Success bool `json:"success"`

	// Error is the error, if an error occurred.
	Error string `json:"error"`

	// ResponseDescription contains a narrative description of the transaction
	// result.
	ResponseDescription string `json:"responseDescription"`

	// TransactionID is the Stax transaction id.
	TransactionID string `json:"transactionId,omitempty"`

	// Approved indicates that the transaction was approved.
	Approved bool `json:"approved"`

	// AuthCode is the auth code from the payment network.
	AuthCode string `json:"authCode,omitempty"`

	// RequestedAmount is the requested amount.
	RequestedAmount string `json:"requestedAmount"`

	// AuthorizedAmount is the authorized amount, which may differ from the
	// requested amount on a partial auth.
	AuthorizedAmount string `json:"authorizedAmount"`

	// CurrencyCode is the transaction currency code.
	CurrencyCode string `json:"currencyCode"`

	// EntryMethod is the entry method, such as CHIP, MSR or KEYED.
	EntryMethod string `json:"entryMethod,omitempty"`

	// MaskedPAN is the masked primary account number.
	MaskedPAN string `json:"maskedPan,omitempty"`

	// Network is the card network.
	Network string `json:"network,omitempty"`

	// Timestamp is the timestamp of the transaction.
	Timestamp string `json:"timestamp"`

	// Status indicates the current status of a transaction.
	Status string `json:"status"`

	// CardMetadata contains details about a payment card derived from its
	// BIN/IIN.
	CardMetadata *CardMetadata `json:"cardMetadata,omitempty"`

	// ReceiptSuggestions contains suggested receipt fields.
	ReceiptSuggestions ReceiptSuggestions `json:"receiptSuggestions"`

	// Test indicates that the transaction ran on the test gateway.
	Test bool `json:"test"`
}

// AbstractAcknowledgement contains fields which should be returned with
// standard requests.
type AbstractAcknowledgement struct {
	// Success indicates whether or not the request succeeded.
	Success bool

	// Error is the error, if an error occurred.
	Error string

	// ResponseDescription contains a narrative description of the transaction
	// result.
	ResponseDescription string
}

// From creates an instance of AbstractAcknowledgement with values
// from a generic type.
func (r AbstractAcknowledgement) From(raw interface{}) (result AbstractAcknowledgement, ok bool) {
	ok = copyTo(raw, &r)
	return r, ok
}

// TerminalReference contains a reference to a terminal name.
type TerminalReference struct {
	// TerminalName is the name of the target payment terminal.
	TerminalName string

	// ResetConnection forces the terminal cloud connection to be reset while a
	// transactions is in flight. This is a diagnostic settings that can be used
	// only for test transactions.
	ResetConnection bool
}

// From creates an instance of TerminalReference with values
// from a generic type.
func (r TerminalReference) From(raw interface{}) (result TerminalReference, ok bool) {
	ok = copyTo(raw, &r)
	return r, ok
}

// SignatureResponse contains customer signature data.
type SignatureResponse struct {
	// SigFile contains the hex encoded signature data.
	SigFile string
}

// From creates an instance of SignatureResponse with values
// from a generic type.
func (r SignatureResponse) From(raw interface{}) (result SignatureResponse, ok bool) {
	ok = copyTo(raw, &r)
	return r, ok
}

// SignatureRequest contains a request for customer signature data.
type SignatureRequest struct {
	// SigFile is a location on the filesystem which a customer signature should
	// be written to.
	SigFile string

	// SigFormat specifies the image format to be used for returning signatures.
	SigFormat SignatureFormat

	// SigWidth is the width that the signature image should be scaled to,
	// preserving the aspect ratio. If not provided, the signature is returned in
	// the terminal's max resolution.
	SigWidth int

	// DisableSignature specifies whether or not signature prompt should be
	// skipped on the terminal. The terminal will indicate whether or not a
	// signature is required by the card in the receipt suggestions response.
	DisableSignature bool
}

// From creates an instance of SignatureRequest with values
// from a generic type.
func (r SignatureRequest) From(raw interface{}) (result SignatureRequest, ok bool) {
	ok = copyTo(raw, &r)
	return r, ok
}

// ApprovalResponse contains response fields for an approved transaction.
type ApprovalResponse struct {
	// Approved indicates that the transaction was approved.
	Approved bool

	// AuthCode is the auth code from the payment network.
	AuthCode string

	// AuthResponseCode is the code returned by the terminal or the card issuer
	// to indicate the disposition of the message.
	AuthResponseCode string
}

// From creates an instance of ApprovalResponse with values
// from a generic type.
func (r ApprovalResponse) From(raw interface{}) (result ApprovalResponse, ok bool) {
	ok = copyTo(raw, &r)
	return r, ok
}

// TimeoutRequest models a low level request with a timeout and test flag.
type TimeoutRequest struct {
	// Timeout is the request timeout in seconds.
	Timeout int

	// Test specifies whether or not to route transaction to the test gateway.
	Test bool
}

// From creates an instance of TimeoutRequest with values
// from a generic type.
func (r TimeoutRequest) From(raw interface{}) (result TimeoutRequest, ok bool) {
	ok = copyTo(raw, &r)
	return r, ok
}

// CoreRequest contains core request fields for a transaction.
type CoreRequest struct {
	// TransactionRef contains a user-assigned reference that can be used to
	// recall or reverse transactions.
	TransactionRef string

	// AutogeneratedRef indicates that the transaction reference was
	// autogenerated and should be ignored for the purposes of duplicate
	// detection.
	AutogeneratedRef bool

	// Async defers the response to the transaction and returns immediately.
	// Callers should retrive the transaction result using the Transaction Status
	// API.
	Async bool

	// Queue adds the transaction to the queue and returns immediately. Callers
	// should retrive the transaction result using the Transaction Status API.
	Queue bool

	// WaitForRemovedCard specifies whether or not the request should block until
	// all cards have been removed from the card reader.
	WaitForRemovedCard bool

	// Force causes a transaction to override any in-progress transactions.
	Force bool

	// OrderRef is an identifier from an external point of sale system.
	OrderRef string

	// DestinationAccount is the settlement account for merchants with split
	// settlements.
	DestinationAccount string

	// TestCase can include a code used to trigger simulated conditions for the
	// purposes of testing and certification. Valid for test merchant accounts
	// only.
	TestCase string
}

// From creates an instance of CoreRequest with values
// from a generic type.
func (r CoreRequest) From(raw interface{}) (result CoreRequest, ok bool) {
	ok = copyTo(raw, &r)
	return r, ok
}

// PaymentMethodResponse contains response details about a payment method.
type PaymentMethodResponse struct {
	// Token is the payment token, if the payment was enrolled in the vault.
	Token string

	// EntryMethod is the entry method for the transaction (CHIP, MSR, KEYED,
	// etc).
	EntryMethod string

	// PaymentType is the card brand (VISA, MC, AMEX, DEBIT, etc).
	PaymentType string

	// Network provides network level detail on how a transaction was routed,
	// especially for debit transactions.
	Network string

	// Logo identifies the card association based on bin number. Used primarily
	// used to indicate the major logo on a card, even when debit transactions
	// are routed on a different network.
	Logo string

	// MaskedPAN is the masked primary account number.
	MaskedPAN string

	// PublicKey is the BlockChyp public key if the user presented a BlockChyp
	// payment card.
	PublicKey string

	// ScopeAlert indicates that the transaction did something that would put the
	// system in PCI scope.
	ScopeAlert bool

	// CardHolder is the cardholder name.
	CardHolder string

	// ExpMonth is the card expiration month in MM format.
	ExpMonth string

	// ExpYear is the card expiration year in YY format.
	ExpYear string

	// PostalCode is the card postal code.
	PostalCode string

	// Address is the card address.
	Address string

	// Country is the card country.
	Country string

	// AVSResponse contains address verification results if address information
	// was submitted.
	AVSResponse AVSResponse

	// CVVResponse contains the CVV verification result if CVV was submitted.
	CVVResponse string

	// ReceiptSuggestions contains suggested receipt fields.
	ReceiptSuggestions ReceiptSuggestions

	// Customer contains customer data, if any. Preserved for reverse
	// compatibility.
	Customer *Customer

	// Customers contains customer data, if any.
	Customers []Customer
}

// From creates an instance of PaymentMethodResponse with values
// from a generic type.
func (r PaymentMethodResponse) From(raw interface{}) (result PaymentMethodResponse, ok bool) {
	ok = copyTo(raw, &r)
	return r, ok
}

// CryptocurrencyResponse contains response details for a cryptocurrency
// transaction.
type CryptocurrencyResponse struct {
	// Confirmed indicates that the transaction has met the standard criteria for
	// confirmation on the network. (For example, 6 confirmations for level one
	// bitcoin.)
	Confirmed bool

	// CryptoAuthorizedAmount is the amount submitted to the blockchain.
	CryptoAuthorizedAmount string

	// CryptoNetworkFee is the network level fee assessed for the transaction
	// denominated in cryptocurrency. This fee goes to channel operators and
	// crypto miners, not BlockChyp.
	CryptoNetworkFee string

	// Cryptocurrency is the three letter cryptocurrency code used for the
	// transactions.
	Cryptocurrency string

	// CryptoNetwork indicates whether or not the transaction was processed on
	// the level one or level two network.
	CryptoNetwork string

	// CryptoReceiveAddress the address on the crypto network the transaction was
	// sent to.
	CryptoReceiveAddress string

	// CryptoBlock hash or other identifier that identifies the block on the
	// cryptocurrency network, if available or relevant.
	CryptoBlock string

	// CryptoTransactionID hash or other transaction identifier that identifies
	// the transaction on the cryptocurrency network, if available or relevant.
	CryptoTransactionID string

	// CryptoPaymentRequest is the payment request URI used for the transaction,
	// if available.
	CryptoPaymentRequest string

	// CryptoStatus is used for additional status information related to crypto
	// transactions.
	CryptoStatus string
}

// From creates an instance of CryptocurrencyResponse with values
// from a generic type.
func (r CryptocurrencyResponse) From(raw interface{}) (result CryptocurrencyResponse, ok bool) {
	ok = copyTo(raw, &r)
	return r, ok
}

// PaymentAmounts contains response details about tender amounts.
type PaymentAmounts struct {
	// PartialAuth indicates whether or not the transaction was approved for a
	// partial amount.
	PartialAuth bool

	// AltCurrency indicates whether or not an alternate currency was used.
	AltCurrency bool

	// FSAAuth indicates whether or not a request was settled on an FSA card.
	FSAAuth bool

	// CurrencyCode is the currency code used for the transaction.
	CurrencyCode string

	// RequestedAmount is the requested amount.
	RequestedAmount string

	// AuthorizedAmount is the authorized amount. May not match the requested
	// amount in the event of a partial auth.
	AuthorizedAmount string

	// RemainingBalance is the remaining balance on the payment method.
	RemainingBalance string

	// TipAmount is the tip amount.
	TipAmount string

	// TaxAmount is the tax amount.
	TaxAmount string

	// RequestedCashBackAmount is the cash back amount the customer requested
	// during the transaction.
	RequestedCashBackAmount string

	// AuthorizedCashBackAmount is the amount of cash back authorized by the
	// gateway. This amount will be the entire amount requested, or zero.
	AuthorizedCashBackAmount string
}

// From creates an instance of PaymentAmounts with values
// from a generic type.
func (r PaymentAmounts) From(raw interface{}) (result PaymentAmounts, ok bool) {
	ok = copyTo(raw, &r)
	return r, ok
}

// PaymentMethod contains request details about a payment method.
type PaymentMethod struct {
	// Token is the payment token to be used for this transaction. This should be
	// used for recurring transactions. The /enroll endpoint ignores this field.
	Token string

	// Track1 contains track 1 magnetic stripe data.
	Track1 string

	// Track2 contains track 2 magnetic stripe data.
	Track2 string

	// PAN contains the primary account number. We recommend using the terminal
	// or e-commerce tokenization libraries instead of passing account numbers in
	// directly, as this would put your application in PCI scope.
	PAN string

	// RoutingNumber is the ACH routing number for ACH transactions.
	RoutingNumber string

	// CardholderName is the cardholder name. Only required if the request
	// includes a primary account number or track data.
	CardholderName string

	// ExpMonth is the card expiration month for use with PAN based transactions.
	ExpMonth string

	// ExpYear is the card expiration year for use with PAN based transactions.
	ExpYear string

	// CVV is the card CVV for use with PAN based transactions.
	CVV string

	// Address is the cardholder address for use with address verification.
	Address string

	// PostalCode is the cardholder postal code for use with address
	// verification.
	PostalCode string

	// Country is the cardholder country.
	Country string

	// ManualEntry specifies that the payment entry method is a manual keyed
	// transaction. If this is true, no other payment method will be accepted.
	ManualEntry bool

	// KSN is the key serial number used for DUKPT encryption.
	KSN string

	// PINBlock is the encrypted pin block.
	PINBlock string

	// CardType designates categories of cards: credit, debit, EBT.
	CardType CardType

	// PaymentType designates brands of payment methods: Visa, Discover, etc.
	PaymentType string
}

// From creates an instance of PaymentMethod with values
// from a generic type.
func (r PaymentMethod) From(raw interface{}) (result PaymentMethod, ok bool) {
	ok = copyTo(raw, &r)
	return r, ok
}

// RequestAmount contains request details about tender amounts.
type RequestAmount struct {
	// CurrencyCode indicates the transaction currency code.
	CurrencyCode string

	// Amount is the requested amount.
	Amount string

	// TaxExempt indicates that the request is tax exempt. Only required for tax
	// exempt level 2 processing.
	TaxExempt bool

	// Surcharge is a flag to add a surcharge to the transaction to cover credit
	// card fees, if permitted.
	Surcharge bool

	// CashDiscount is a flag that applies a discount to negate the surcharge for
	// debit transactions or other surcharge ineligible payment methods.
	CashDiscount bool
}

// From creates an instance of RequestAmount with values
// from a generic type.
func (r RequestAmount) From(raw interface{}) (result RequestAmount, ok bool) {
	ok = copyTo(raw, &r)
	return r, ok
}

// Subtotals contains request subtotals.
type Subtotals struct {
	// TipAmount is the tip amount.
	TipAmount string

	// TaxAmount is the tax amount.
	TaxAmount string
}

// From creates an instance of Subtotals with values
// from a generic type.
func (r Subtotals) From(raw interface{}) (result Subtotals, ok bool) {
	ok = copyTo(raw, &r)
	return r, ok
}

// PreviousTransaction contains a reference to a previous transaction.
type PreviousTransaction struct {
	// TransactionID is the ID of the previous transaction being referenced.
	TransactionID string
}

// From creates an instance of PreviousTransaction with values
// from a generic type.
func (r PreviousTransaction) From(raw interface{}) (result PreviousTransaction, ok bool) {
	ok = copyTo(raw, &r)
	return r, ok
}

// CoreResponse contains core response fields for a transaction.
type CoreResponse struct {
	// TransactionID is the ID assigned to the transaction.
	TransactionID string

	// BatchID is the ID assigned to the batch.
	BatchID string

	// TransactionRef is the transaction reference string assigned to the
	// transaction request. If no transaction ref was assiged on the request,
	// then the gateway will randomly generate one.
	TransactionRef string

	// TransactionType is the type of transaction.
	TransactionType string

	// Timestamp is the timestamp of the transaction.
	Timestamp string

	// TickBlock is the hash of the last tick block.
	TickBlock string

	// Test indicates that the transaction was processed on the test gateway.
	Test bool

	// DestinationAccount is the settlement account for merchants with split
	// settlements.
	DestinationAccount string

	// Sig is the ECC signature of the response. Can be used to ensure that it
	// was signed by the terminal and detect man-in-the middle attacks.
	Sig string
}

// From creates an instance of CoreResponse with values
// from a generic type.
func (r CoreResponse) From(raw interface{}) (result CoreResponse, ok bool) {
	ok = copyTo(raw, &r)
	return r, ok
}

func copyTo(from, to interface{}) (ok bool) {
	fromV := reflect.ValueOf(from)
	if fromV.Kind() == reflect.Ptr {
		fromV = fromV.Elem()
	}

	toV := reflect.ValueOf(to).Elem()

	for i := 0; i < toV.NumField(); i++ {
		val := fromV.FieldByName(toV.Type().Field(i).Name)
		if !val.IsValid() {
			continue
		}

		ok = true
		toV.Field(i).Set(val)
	}

	return
}

func clearField(ptr interface{}, field string) {
	v := reflect.ValueOf(ptr)
	if v.Kind() != reflect.Ptr {
		panic("must pass by pointer")
	}
	v = v.Elem()

	if f := v.FieldByName(field); f.IsValid() {
		f.Set(reflect.Zero(f.Type()))
	}
}
