package staxpayments

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

// DefaultCoreHost is the default Stax core API host used when none is supplied.
const DefaultCoreHost = "https://api.staxpayments.com"

// transientCredentialsPath is the core API path that exchanges a Stax bearer
// token for short-lived BlockChyp transient credentials. This is an internal
// mechanism of the SDK and is never exposed as a public method.
const transientCredentialsPath = "/terminal/transient-credentials"

// expirySkew refreshes transient credentials this long before their stated
// expiry to avoid a credential lapsing mid-request (clock-skew buffer).
const expirySkew = 30 * time.Second

// transientCredentialsResponse models the core API exchange response.
type transientCredentialsResponse struct {
	APIKey      string `json:"apiKey"`
	BearerToken string `json:"bearerToken"`
	SigningKey  string `json:"signingKey"`
	// ExpiresAt is supplied by the Stax core API. It is assumed here and will be
	// populated once the core endpoint is implemented.
	ExpiresAt time.Time `json:"expiresAt"`
}

// staxAuth owns every piece of Stax-specific authentication state: the Stax
// bearer token, the core host, and the cached merchant-scoped transient
// credentials with their expiry. The generated client touches it through
// exactly two hooks: ensure() and coreRequest().
type staxAuth struct {
	bearerToken string
	coreHost    string
	gatewayHost string
	httpClient  *http.Client

	mu        sync.Mutex
	cached    APICredentials
	expiresAt time.Time
	valid     bool
}

func newStaxAuth(bearerToken, coreHost, gatewayHost string) *staxAuth {
	return &staxAuth{
		bearerToken: bearerToken,
		coreHost:    coreHost,
		gatewayHost: gatewayHost,
		httpClient:  &http.Client{Timeout: 20 * time.Second},
	}
}

// ensure returns valid merchant-scoped transient credentials, exchanging the
// Stax bearer token via the core API when the cache is empty or expired. The
// mutex serializes concurrent callers so only one exchange is in flight at a
// time (single-flight).
func (s *staxAuth) ensure() (APICredentials, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.valid && time.Now().Add(expirySkew).Before(s.expiresAt) {
		return s.cached, nil
	}

	s.valid = false

	var result transientCredentialsResponse
	if err := s.coreRequest(http.MethodPost, transientCredentialsPath, struct{}{}, &result); err != nil {
		return APICredentials{}, err
	}

	s.cached = APICredentials{
		APIKey:      result.APIKey,
		BearerToken: result.BearerToken,
		SigningKey:  result.SigningKey,
	}
	if result.ExpiresAt.IsZero() {
		s.expiresAt = time.Now().Add(5 * time.Minute)
	} else {
		s.expiresAt = result.ExpiresAt
	}
	s.valid = true

	return s.cached, nil
}

// coreRequest issues a Stax-bearer-authenticated call directly to the core API.
// Used by core-routed endpoints such as listing terminals.
func (s *staxAuth) coreRequest(method, path string, request, response interface{}) error {
	var body io.Reader
	if method != http.MethodGet && request != nil {
		payload, err := json.Marshal(request)
		if err != nil {
			return err
		}
		body = bytes.NewReader(payload)
	}

	req, err := http.NewRequest(method, s.coreHost+path, body)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+s.bearerToken)
	req.Header.Set("Content-Type", "application/json")

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("stax core request failed: %s", resp.Status)
	}

	if response != nil {
		return json.NewDecoder(resp.Body).Decode(response)
	}

	return nil
}
