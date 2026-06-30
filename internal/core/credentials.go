package core

import (
	"sync"
	"time"

	staxpayments "github.com/blockchyp/staxpayments-go"
)

// transientCredentialsSkew refreshes transient credentials this long before
// their stated expiry to avoid a credential lapsing mid-request (clock-skew
// buffer).
const transientCredentialsSkew = 30 * time.Second

// credentialStore is an in-memory, concurrency-safe cache for the short-lived
// Stax Payments transient credentials exchanged from the Stax bearer token. A
// single store is owned by the shared core client and therefore shared by every
// namespace client, so the credentials are exchanged once and reused across all
// namespaces until they near expiry.
type credentialStore struct {
	mu     sync.RWMutex
	creds  staxpayments.APICredentials
	expiry time.Time
	skew   time.Duration
}

// newCredentialStore returns an empty store that treats cached credentials as
// expired skew before their stated expiry.
func newCredentialStore(skew time.Duration) *credentialStore {
	return &credentialStore{skew: skew}
}

// current returns the cached credentials. Callers must have made them valid via
// ensure earlier in the request; current is the read path used for signing and
// route-cache keys.
func (s *credentialStore) current() staxpayments.APICredentials {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.creds
}

// validLocked reports whether non-expired credentials are cached, with the skew
// buffer applied. The caller must hold the lock.
func (s *credentialStore) validLocked() bool {
	return !s.expiry.IsZero() && time.Now().Add(s.skew).Before(s.expiry)
}

// ensure guarantees the store holds valid credentials, invoking refresh to
// exchange new ones when the cache is empty or near expiry. The write lock is
// held across refresh so only one exchange is in flight at a time
// (single-flight): concurrent callers block, then observe the freshly stored
// credentials and return without exchanging again.
func (s *credentialStore) ensure(refresh func() (staxpayments.APICredentials, time.Time, error)) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.validLocked() {
		return nil
	}

	creds, expiry, err := refresh()
	if err != nil {
		return err
	}

	s.creds = creds
	s.expiry = expiry
	return nil
}
