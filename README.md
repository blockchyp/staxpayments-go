# Stax Payments Go SDK

[![Build Status](https://github.com/blockchyp/staxpayments-go/actions/workflows/main.yml/badge.svg)](https://github.com/blockchyp/staxpayments-go/actions/workflows/main.yml)
[![Release](https://img.shields.io/github/release/blockchyp/staxpayments-go/all.svg?style=shield)](https://github.com/blockchyp/staxpayments-go/releases/latest)
[![Go Report Card](https://goreportcard.com/badge/github.com/blockchyp/staxpayments-go)](https://goreportcard.com/report/github.com/blockchyp/staxpayments-go)
[![GoDoc](https://godoc.org/github.com/blockchyp/staxpayments-go?status.svg)](https://godoc.org/github.com/blockchyp/staxpayments-go)
[![License: MIT](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/blockchyp/staxpayments-go/blob/master/LICENSE)

This is the Go SDK for Stax Payments. Like all Stax Payments SDKs, it provides a full
Go client for the Stax Payments gateway and Stax Payments payment terminals.

This project also contains a command line interface for Windows, Linux, and
Mac OS developers working in languages or on platforms for which Stax Payments doesn't
currently provide a supported SDK.

## Command Line Interface

In addition to the standard Go SDK, the Makefile includes special targets for
Windows and Linux command line binaries.

These binaries are intended for unique situations where using an SDK or doing
a direct REST integration aren't practical.

Check out the [CLI Reference](docs/cli.md) for more information.

## Go Installation

For Go developers, you can install Stax Payments in the usual way with `go get`.

```
go get github.com/blockchyp/staxpayments-go
```

## A Simple Example

Running your first terminal transaction is easy. Make sure you have a Stax Payments
terminal, activate it, and obtain a Stax bearer token.

```
package main

import (
    "encoding/json"
    "fmt"
    "log"

    staxpayments "github.com/blockchyp/staxpayments-go"
    "github.com/blockchyp/staxpayments-go/client"
    "github.com/blockchyp/staxpayments-go/payments"
)

func main() {

    // Construct the root client with your Stax bearer token. It builds one
    // shared transport that transparently exchanges the token for short-lived
    // transient credentials and reuses them across every namespace, so a single
    // credential cache is shared by client.Payments, client.Terminals, etc.
    staxClient := client.New(staxpayments.StaxAPICredentials{
        BearerToken: "<your-stax-bearer-token>",
    })

    req := payments.AuthorizationRequest{
        Test: true,
        TerminalName: "Test Terminal",
        Amount: "55.00",
    }

    response, err := staxClient.Payments.Charge(req)

    if err != nil {
        log.Fatal(err)
    }

    if response.Approved {
        fmt.Println("Approved")
        fmt.Println(response.AuthCode)
        fmt.Println(response.AuthorizedAmount)
        fmt.Println(response.ReceiptSuggestions.AID)
    } else {
        fmt.Println(response.ResponseDescription)
    }

    b, err := json.Marshal(response)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Print(string(b))
}
```

The response contains all the information you'll need to complete processing
a transaction. Of particular importance is the ReceiptSuggestions struct, which
contains all the fields that are required or recommended for PCI or EMV compliance.



## Additional Documentation

Complete documentation can be found on our [Developer Documentation Portal].

[Developer Documentation Portal]: https://docs.blockchyp.com/

## Authentication

This SDK authenticates with a **Stax bearer token**. Construct a client with your
bearer token and the SDK transparently exchanges it for short-lived BlockChyp
transient credentials — cached and refreshed automatically — whenever you run a
terminal transaction such as `charge` or `refund`. Listing terminals is served
directly by the Stax core API using your bearer token.

## Getting a Developer Kit

In order to test your integration with real terminals, you'll need a BlockChyp
Developer Kit. Our kits include a fully functioning payment terminal with
test pin encryption keys. Every kit includes a comprehensive set of test
cards with test cards for every major card brand and entry method, including
Contactless and Contact EMV and mag stripe cards. Each kit also includes
test gift cards for our blockchain gift card system.

Access to BlockChyp's developer program is currently invite only, but you
can request an invitation by contacting our engineering team at **nerds@blockchyp.com**.

You can also view a number of long form demos and learn more about us on our [YouTube Channel](https://www.youtube.com/channel/UCE-iIVlJic_XArs_U65ZcJg).

## Transaction Code Examples

You don't want to read words. You want examples. Here's a quick rundown of the
stuff you can do with the Stax Payments Go SDK and a few basic examples.

### Payment Endpoints


These are the core payment APIs used to execute and work with payment transactions in BlockChyp.



#### Charge



* **API Credential Types:** Merchant
* **Required Role:** Payment API Access

Our most popular transaction executes a standard authorization and capture.
This is the most basic of
basic payment transactions, typically used in conventional retail.

Charge transactions can use a payment terminal to capture a payment or
use a previously enrolled payment token.

**Terminal Transactions**

For terminal transactions, make sure you pass in the terminal name using the `terminalName` property.

**Token Transactions**

If you have a payment token, omit the `terminalName` property and pass in the token with the `token`
property instead.

**Card Numbers and Mag Stripes**

You can also pass in PANs and Mag Stripes, but you probably shouldn't, as this will
put you in PCI scope and the most common vector for POS breaches is keylogging.
If you use terminals for manual card entry, you'll bypass any keyloggers that
might be maliciously running on the point-of-sale system.

**Common Variations**

* **Gift Card Redemption**:  There's no special API for gift card redemption in BlockChyp. Simply execute a plain charge transaction and if the customer swipes a gift card, our terminals will identify the gift card and run a gift card redemption. Also note that if for some reason the gift card's original purchase transaction is associated with fraud or a chargeback, the transaction will be rejected.
* **EBT**: Set the `CardType` field to `staxpayments.CardTypeEBT` to process an EBT SNAP transaction. Note that test EBT transactions always assume a balance of $100.00, so test EBT transactions over that amount may be declined.
* **Cash Back**: To enable cash back for debit transactions, set the `CashBack` field. If the card presented isn't a debit card, the `CashBack` field will be ignored.
* **Manual Card Entry**: Set the `ManualEntry` field to enable manual card entry. Good as a backup when chips and MSR's don't work or for more secure phone orders. You can even combine the `ManualEntry` field with the `CardType` field set to `staxpayments.CardTypeEBT` for manual EBT card entry.
* **Inline Tokenization**: You can enroll the payment method in the token vault inline with a charge transaction by setting the `Enroll` field. You'll get a token back in the response. You can even bind the token to a customer record if you also pass in customer data.
* **Prompting for Tips**: Set the `PromptForTip` field if you'd like to prompt the customer for a tip before authorization. Good for pay-at-the-table and other service related scenarios.
* **Cash Discounting and Surcharging**:  The `Surcharge` and `CashDiscount` fields can be used together to support cash discounting or surcharge problems. Consult the Cash Discount documentation for more details.
* **Cryptocurrency** The `Cryptocurrency` field can be used to switch the standard present card screen to a cryptocurrency screen.  The field value can be `ANY` to enable any supported cryptocurrency or a single currency code such as `BTC` for Bitcoin.



```go
package main

import (
    "fmt"
    "log"

    staxpayments "github.com/blockchyp/staxpayments-go"
    "github.com/blockchyp/staxpayments-go/client"
    "github.com/blockchyp/staxpayments-go/payments"
)

func chargeExample() {
    // construct the root client with your Stax bearer token; it builds one shared
    // transport that transparently exchanges the token for short-lived transient
    // credentials and reuses them across every namespace
    staxClient := client.New(staxpayments.StaxAPICredentials{
        BearerToken: "<your-stax-bearer-token>",
    })

    // setup request object
    request := payments.AuthorizationRequest{
        Test:         true,
        TerminalName: "Test Terminal",
        Amount:       "55.00",
    }

    response, err := staxClient.Payments.Charge(request)

    if err != nil {
        log.Fatal(err)
    }

    //process the result
    if response.Approved {
        fmt.Println("approved")
    }

    fmt.Printf("Response: %+v\n", response)
}

```

#### Preauthorization



* **API Credential Types:** Merchant
* **Required Role:** Payment API Access

A preauthorization puts a hold on funds and must be captured later.  This is used
in scenarios where the final transaction amount might change.  A common example is 
fine dining, where a tip adjustment is required before final settlement.

Another use case for preauthorization is e-commerce.  Typically, an online order
is preauthorized at the time of the order and then captured when the order ships.

Preauthorizations can use a payment terminal to capture a payment or
use a previously enrolled payment token.

**Terminal Transactions**

For terminal transactions, make sure you pass in the terminal name using the `terminalName` property.

**Token Transactions**

If you have a payment token, omit the `terminalName` property and pass in the token with the `token`
property instead.

**Card Numbers and Mag Stripes**

You can also pass in PANs and Mag Stripes, but you probably shouldn't, as this will
put you in PCI scope and the most common vector for POS breaches is key logging.
If you use terminals for manual card entry, you'll bypass any key loggers that
might be maliciously running on the point-of-sale system.

**Cryptocurrency**

Note that preauths are not supported for cryptocurrency.

**Common Variations**

* **Manual Card Entry**: Set the `ManualEntry` field to enable manual card entry. Good as a backup when chips and MSR's don't work or for more secure phone orders. You can even combine the `ManualEntry` field with `CardType` set to `staxpayments.CardTypeEBT` for manual EBT card entry.
* **Inline Tokenization**: You can enroll the payment method in the token vault in line with a charge transaction by setting the `Enroll` field. You'll get a token back in the response. You can even bind the token to a customer record if you also pass in customer data.
* **Prompting for Tips**: Set the `PromptForTip` field if you'd like to prompt the customer for a tip before authorization. You can prompt for tips as part of a preauthorization, although it's not a very common approach.
* **Cash Discounting and Surcharging**: The `Surcharge` and `CashDiscount` fields can be used together to support cash discounting or surcharge problems. Consult the Cash Discount documentation for more details.




```go
package main

import (
    "fmt"
    "log"

    staxpayments "github.com/blockchyp/staxpayments-go"
    "github.com/blockchyp/staxpayments-go/client"
    "github.com/blockchyp/staxpayments-go/payments"
)

func preauthExample() {
    // construct the root client with your Stax bearer token; it builds one shared
    // transport that transparently exchanges the token for short-lived transient
    // credentials and reuses them across every namespace
    staxClient := client.New(staxpayments.StaxAPICredentials{
        BearerToken: "<your-stax-bearer-token>",
    })

    // setup request object
    request := payments.AuthorizationRequest{
        Test:         true,
        TerminalName: "Test Terminal",
        Amount:       "27.00",
    }

    response, err := staxClient.Payments.Preauth(request)

    if err != nil {
        log.Fatal(err)
    }

    //process the result
    if response.Approved {
        fmt.Println("approved")
    }

    fmt.Printf("Response: %+v\n", response)
}

```

#### Refund



* **API Credential Types:** Merchant
* **Required Role:** Payment API Access

It's not ideal, but sometimes customers want their money back.

Our refund API allows you to confront this unpleasant reality by executing refunds in a few different scenarios.

The most fraud resistent method is to execute refunds in the context of a previous transaction.  You should always keep track of the Transaction ID
returned in a BlockChyp response.  To refund the full amount of the previous transaction, just pass in the original Transaction ID with the refund requests.

**Partial Refunds**

For a partial refund, just pass in an amount along with the Transaction ID.
The only rule is that the amount must be equal to or less than the original
transaction.  You can execute multiple partial refunds against the same
original transaction as long as the total refunded amount doesn't exceed the original amount.

**Tokenized Refunds**

You can also use a token to execute a refund.  Pass in a token instead
of the Transaction ID and the desired refund amount.

**Free Range Refunds**

When you execute a refund without referencing a previous transaction, we
call this a *free range refund*.

We don't recommend this type of refund, but it is permitted.  If you absolutely insist on
doing it, pass in a Terminal Name and an amount.

You can execute a manual or keyed refund by passing the `ManualEntry` field
to a free range refund request.

**Gift Card Refunds**

Gift card refunds are allowed in the context of a previous transaction, but
free range gift card refunds are not allowed.  Use the gift card activation
API if you need to add more funds to a gift card.

**Store and Forward Support**

Refunds are not permitted when a terminal falls back to store and forward mode.

**Auto Voids**

If a refund referencing a previous transaction is executed for the full amount
before the original transaction's batch is closed, the refund is automatically
converted to a void.  This saves the merchant a little bit of money.

**Cryptocurrency**

Note that refunds are not supported for cryptocurrency.  You must refund crypto transactions
manually from your cryptocurrency wallet.




```go
package main

import (
    "fmt"
    "log"

    staxpayments "github.com/blockchyp/staxpayments-go"
    "github.com/blockchyp/staxpayments-go/client"
    "github.com/blockchyp/staxpayments-go/payments"
)

func refundExample() {
    // construct the root client with your Stax bearer token; it builds one shared
    // transport that transparently exchanges the token for short-lived transient
    // credentials and reuses them across every namespace
    staxClient := client.New(staxpayments.StaxAPICredentials{
        BearerToken: "<your-stax-bearer-token>",
    })

    // setup request object
    request := payments.RefundRequest{
        TransactionID: "<PREVIOUS TRANSACTION ID>",

        // Optional amount for partial refunds.
        Amount: "5.00",
    }

    response, err := staxClient.Payments.Refund(request)

    if err != nil {
        log.Fatal(err)
    }

    //process the result
    if response.Approved {
        fmt.Println("approved")
    }

    fmt.Printf("Response: %+v\n", response)
}

```

### Terminal Management Endpoints


These APIs support terminal management functions and additional terminal 
features such as line item display, messages, and interactive prompts.  
These features can be used to extend a point of sale system's functionality.



#### List Terminals



* **API Credential Types:** Merchant & Partner
* **Required Role:** Terminal Management

This API returns details about terminals associated with a merchant account.

Status and resource information is returned for all terminals along with a preview of the 
current branding image displayed on the terminal




```go
package main

import (
    "fmt"
    "log"

    staxpayments "github.com/blockchyp/staxpayments-go"
    "github.com/blockchyp/staxpayments-go/client"
    "github.com/blockchyp/staxpayments-go/terminals"
)

func terminalsExample() {
    // construct the root client with your Stax bearer token; it builds one shared
    // transport that transparently exchanges the token for short-lived transient
    // credentials and reuses them across every namespace
    staxClient := client.New(staxpayments.StaxAPICredentials{
        BearerToken: "<your-stax-bearer-token>",
    })

    // setup request object
    request := terminals.TerminalProfileRequest{}

    response, err := staxClient.Terminals.Terminals(request)

    if err != nil {
        log.Fatal(err)
    }

    //process the result
    if response.Success {
        fmt.Println("Success")
    }

    fmt.Printf("Response: %+v\n", response)
}

```

### Terms & Conditions Endpoints


Developers can use BlockChyp to display and capture acceptance of contracts or agreements related to transactions.
These agreements can be any long-form contract ranging from rental agreements to HIPPA disclosures.

There are two basic approaches to terms and conditions capture.  Merchants can store contract templates in 
BlockChyp or they can send the full agreement text as part of every API call.  The right approach will largely 
depend on whether or not the system being integrated with BlockChyp already has a mechanism for organizing 
and managing agreements.  For systems that already have this feature built in, it's probably not necessary 
to use Terms and Conditions.

When agreements are displayed on a terminal, the consumer can scroll through and read the entire agreement,
and provide a signature.  Results are returned as part of the API response, but BlockChyp also stores a 
record of the agreement including the signature image, timestamp, and the full text of the agreement that was 
agreed to.

The Terms and Conditions Log APIs can be used to search and retrieve acceptance records.  Those records
can also be linked to a transaction if a transaction id is provided with the original API request.



### Token Management


BlockChyp supports saved payments and recurring payments through the use of tokens.  Tokens can be created
via the Enroll API or the web tokenizer.  Once created, these tokens can be used for subsequent payments 
or associated with customer records as saved payment methods.

Tokens are limited to a single merchant by default, but can be shared across an organization for multi-location 
merchants by special arrangement with BlockChyp.  Contact your BlockChyp rep to setup token sharing.



### Customer Endpoints


These APIs allow developers to create and manage customer records in BlockChyp.  Developers who wish to use
BlockChyp for tokenized recurring payments can use tokens directly if they have their own customer management
system.  However, BlockChyp provides additional tools for managing customers and keeping track of a customer's saved
payment tokens.

In addition, if customer features are used, BlockChyp can detect a payment method associated with an existing
customer, and return customer data with payment transactions.  This can be used as a passive method to detect
repeat customers.



### Survey Reference


These APIs are used to work with post-transaction surveys and survey data.

Merchants can optionally configure scaled (1-5) or yes/no questions that can be presented to consumers
after every approved Charge and Preauth transaction.  Surveys do not require any custom programming and
merchants can simply configure them without the point-of-sale system needing any additional customization.

However, these APIs allow point-of-sale or third-party system developers to integrate survey question configuration
or result visualization into their own systems.



### Media and Branding Control


BlockChyp has a sophisticated terminal media and branding control platform.  Terminals can be configured to
display logos, images, videos, and slide shows when a terminal is idle.  Branding assets can be configured
at the partner, organization, and merchant level with fine-grained hour-by-hour schedules, if desired. 

Conceptually, all branding and media start with the media library.  Merchants, Partners, and Organizations can
upload images or video and build branding assets from uploaded media.

Slide shows can combine images from the media library into a timed loop of repeating images.

Branding Assets can then be used to combine media or slide shows with priority and timing rules to create what 
we call the Terminal Branding Stack.

We call a group of branding assets the *Terminal Branding Stack* because there are implicit rules about which 
branding assets take priority. For example, a merchant with no branding assets configured will inherit the 
branding rules from any organization to which the merchant may belong.  If the merchant doesn't belong to an organization 
or the organization has no branding rules configured, then the system will defer to branding defaults established 
by the point-of-sale or software partner that owns the merchant.

This feature enables partners and organizations (multi-store operators and large national chains) to configure branding
for potentially thousands of terminals from a single interface.

Terminal Branding can also be configured at the individual terminal level and a merchant's terminal fleet 
can be broken into groups and branding configured at the group level.  Branding configured at the terminal
level will always override branding from any higher level group.

The order of priority for the Terminal Branding Stack is given below.

* Terminal
* Terminal Group
* Merchant
* Organization (Region, Chain, etc)
* Partner
* BlockChyp Default Logo



### Merchant Management


These APIs allow partners to manage and configure their merchant portfolios.

Use of these APIs (other than the Merchant Profile API) requires partner scoped API credentials
with special roles and permissions that may require a special arrangement with BlockChyp.

For example, Partners usually can't board merchants directly, but must board merchants using
the standard underwriting process via offer codes and invitations.



### Partner Utilities


These partner only APIs give ISV partners advanced reporting and tools for managing their portfolio.

Most of the APIs below are for portfolio reporting and range from basic partner commission statements
to individual statements with all underlying card brand data.

We also provide a pricing policy API that enables partners to pull down the current pricing rules
in force for any merchant in their portfolio.

<aside class="info">
<b>Currency Data</b>
<p>
All partner APIs return currency and percentage values in two formats: floating point and formatted strings.
</p>
<p>
It's recommended that all developers use the formatted string as this will ensure the most precise values.
Floating point numbers are usually not appropriate for currency or fixed point decimal numbers as
the underlying binary encoding can lead to errors in precision.  We provide floating point values
only as a convenience for developers want to save development time and can live with approximated
values in their use case.
</p>
</aside>







## Running Integration Tests

If you'd like to run the integration tests, create a new file on your system
called `sdk-itest-config.json` with the API credentials you'll be using as
shown in the example below.

```
{
 "gatewayHost": "https://api.blockchyp.com",
 "testGatewayHost": "https://test.blockchyp.com",
 "apiKey": "PZZNEFK7HFULCB3HTLA7HRQDJU",
 "bearerToken": "QUJCHIKNXOMSPGQ4QLT2UJX5DI",
 "signingKey": "f88a72d8bc0965f193abc7006bbffa240663c10e4d1dc3ba2f81e0ca10d359f5"
}
```

This file can be located in a few different places, but is usually located
at `<USER_HOME>/.config/blockchyp/sdk-itest-config.json`. All BlockChyp SDKs
use the same configuration file.

To run the integration test suite via `make`, type the following command:

`make integration`

## Running Regression Tests

The regression package contains interactive tests that can be run to test the
entire stack from end to end.

### Setup

Create a default test merchant on the SIM plugin.

Change these settings:

* Enable partial auth
* Enable PINs
* Enable Missing Signature Reversal
* Enable cash back
* Enable JCB and Union Pay
* Whitelist the BIN range for a chosen MSR test card
* Add a pricing policy:
  * Flat rate: 350 basis points
  * Transaction fee: $0.50

Create a staxpayments.json file with credentials for the test merchant.

### Running

To execute the tests, run:

`make regression`

Follow the prompts.

## Contributions

BlockChyp welcomes contributions from the open source community, but bear in mind
that this repository has been generated by our internal SDK Generator tool. If
we choose to accept a PR or contribution, your code will be moved into our SDK
Generator project, which is a private repository.

## License

Copyright BlockChyp, Inc., 2019

Distributed under the terms of the [MIT] license, staxpayments-go is free and open source software.

[MIT]: https://github.com/blockchyp/staxpayments-go/blob/master/LICENSE

## Other SDKs

BlockChyp has officially supported SDKs for eight different development platforms and counting.
Here's the full list with links to their GitHub repositories.

[Go SDK](https://github.com/blockchyp/blockchyp-go)

[Node.js/JavaScript SDK](https://github.com/blockchyp/blockchyp-js)

[Typescript SDK](https://github.com/blockchyp/blockchyp-ts)

[Java SDK](https://github.com/blockchyp/blockchyp-java)

[.net/C# SDK](https://github.com/blockchyp/blockchyp-csharp)

[Ruby SDK](https://github.com/blockchyp/blockchyp-ruby)

[PHP SDK](https://github.com/blockchyp/blockchyp-php)

[Python SDK](https://github.com/blockchyp/blockchyp-python)

[iOS (Objective-C/Swift) SDK](https://github.com/blockchyp/blockchyp-ios)

[Rust SDK](https://github.com/blockchyp/blockchyp-rust)
