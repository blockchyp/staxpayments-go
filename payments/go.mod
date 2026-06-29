module github.com/blockchyp/staxpayments-go/payments

go 1.23

require github.com/blockchyp/staxpayments-go v0.0.0

// In-repo builds resolve the shared core module from the parent directory.
// Published releases pin a tagged core version via the require above.
replace github.com/blockchyp/staxpayments-go => ../
