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
