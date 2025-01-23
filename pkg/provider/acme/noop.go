//go:build !acme

package acme

import (
	"net/http"

	"github.com/go-acme/lego/v4/challenge"
	ptypes "github.com/traefik/paerser/types"
	"github.com/traefik/traefik/v3/pkg/config/dynamic"
	"github.com/traefik/traefik/v3/pkg/provider/noop"
	"github.com/traefik/traefik/v3/pkg/safe"
	traefiktls "github.com/traefik/traefik/v3/pkg/tls"
	"github.com/traefik/traefik/v3/pkg/types"
)

func NewChallengeHTTP() Provider {
	return Provider{}
}

func NewChallengeTLSALPN() Provider {
	return Provider{}
}

type Store interface{}

type Provider struct {
	noop.Provider
	*Configuration
	ResolverName          string
	Store                 Store
	TLSChallengeProvider  challenge.Provider
	HTTPChallengeProvider challenge.Provider
}

func (Provider) SetTLSManager(*traefiktls.Manager)                {}
func (Provider) SetConfigListenerChan(chan dynamic.Configuration) {}
func (Provider) ListenConfiguration(dynamic.Configuration)        {}

func (Provider) Present(domain, token, keyAuth string) error {
	return nil
}

func (Provider) CleanUp(domain, token, keyAuth string) error {
	return nil
}

func (Provider) ServeHTTP(http.ResponseWriter, *http.Request) {}

type LocalStore struct{}
type Storage struct{}

type Configuration struct {
	Email                string `description:"Email address used for registration." json:"email,omitempty" toml:"email,omitempty" yaml:"email,omitempty"`
	CAServer             string `description:"CA server to use." json:"caServer,omitempty" toml:"caServer,omitempty" yaml:"caServer,omitempty"`
	PreferredChain       string `description:"Preferred chain to use." json:"preferredChain,omitempty" toml:"preferredChain,omitempty" yaml:"preferredChain,omitempty" export:"true"`
	Storage              string `description:"Storage to use." json:"storage,omitempty" toml:"storage,omitempty" yaml:"storage,omitempty" export:"true"`
	KeyType              string `description:"KeyType used for generating certificate private key. Allow value 'EC256', 'EC384', 'RSA2048', 'RSA4096', 'RSA8192'." json:"keyType,omitempty" toml:"keyType,omitempty" yaml:"keyType,omitempty" export:"true"`
	EAB                  *EAB   `description:"External Account Binding to use." json:"eab,omitempty" toml:"eab,omitempty" yaml:"eab,omitempty"`
	CertificatesDuration int    `description:"Certificates' duration in hours." json:"certificatesDuration,omitempty" toml:"certificatesDuration,omitempty" yaml:"certificatesDuration,omitempty" export:"true"`

	CACertificates   []string `description:"Specify the paths to PEM encoded CA Certificates that can be used to authenticate an ACME server with an HTTPS certificate not issued by a CA in the system-wide trusted root list." json:"caCertificates,omitempty" toml:"caCertificates,omitempty" yaml:"caCertificates,omitempty"`
	CASystemCertPool bool     `description:"Define if the certificates pool must use a copy of the system cert pool." json:"caSystemCertPool,omitempty" toml:"caSystemCertPool,omitempty" yaml:"caSystemCertPool,omitempty" export:"true"`
	CAServerName     string   `description:"Specify the CA server name that can be used to authenticate an ACME server with an HTTPS certificate not issued by a CA in the system-wide trusted root list." json:"caServerName,omitempty" toml:"caServerName,omitempty" yaml:"caServerName,omitempty" export:"true"`

	DNSChallenge  *DNSChallenge  `description:"Activate DNS-01 Challenge." json:"dnsChallenge,omitempty" toml:"dnsChallenge,omitempty" yaml:"dnsChallenge,omitempty" label:"allowEmpty" file:"allowEmpty" export:"true"`
	HTTPChallenge *HTTPChallenge `description:"Activate HTTP-01 Challenge." json:"httpChallenge,omitempty" toml:"httpChallenge,omitempty" yaml:"httpChallenge,omitempty" label:"allowEmpty" file:"allowEmpty" export:"true"`
	TLSChallenge  *TLSChallenge  `description:"Activate TLS-ALPN-01 Challenge." json:"tlsChallenge,omitempty" toml:"tlsChallenge,omitempty" yaml:"tlsChallenge,omitempty" label:"allowEmpty" file:"allowEmpty" export:"true"`
}

type Certificate struct {
	Domain      types.Domain `json:"domain,omitempty" toml:"domain,omitempty" yaml:"domain,omitempty"`
	Certificate []byte       `json:"certificate,omitempty" toml:"certificate,omitempty" yaml:"certificate,omitempty"`
	Key         []byte       `json:"key,omitempty" toml:"key,omitempty" yaml:"key,omitempty"`
}

// EAB contains External Account Binding configuration.
type EAB struct {
	Kid         string `description:"Key identifier from External CA." json:"kid,omitempty" toml:"kid,omitempty" yaml:"kid,omitempty" loggable:"false"`
	HmacEncoded string `description:"Base64 encoded HMAC key from External CA." json:"hmacEncoded,omitempty" toml:"hmacEncoded,omitempty" yaml:"hmacEncoded,omitempty" loggable:"false"`
}

// DNSChallenge contains DNS challenge configuration.
type DNSChallenge struct {
	Provider    string       `description:"Use a DNS-01 based challenge provider rather than HTTPS." json:"provider,omitempty" toml:"provider,omitempty" yaml:"provider,omitempty" export:"true"`
	Resolvers   []string     `description:"Use following DNS servers to resolve the FQDN authority." json:"resolvers,omitempty" toml:"resolvers,omitempty" yaml:"resolvers,omitempty"`
	Propagation *Propagation `description:"DNS propagation checks configuration" json:"propagation,omitempty" toml:"propagation,omitempty" yaml:"propagation,omitempty"  label:"allowEmpty" file:"allowEmpty" export:"true"`

	// Deprecated: please use Propagation.DelayBeforeChecks instead.
	DelayBeforeCheck ptypes.Duration `description:"(Deprecated) Assume DNS propagates after a delay in seconds rather than finding and querying nameservers." json:"delayBeforeCheck,omitempty" toml:"delayBeforeCheck,omitempty" yaml:"delayBeforeCheck,omitempty" export:"true"`
	// Deprecated: please use Propagation.DisableChecks instead.
	DisablePropagationCheck bool `description:"(Deprecated) Disable the DNS propagation checks before notifying ACME that the DNS challenge is ready. [not recommended]" json:"disablePropagationCheck,omitempty" toml:"disablePropagationCheck,omitempty" yaml:"disablePropagationCheck,omitempty" export:"true"`
}

type Propagation struct {
	DisableChecks     bool            `description:"Disables the challenge TXT record propagation checks (not recommended)." json:"disableChecks,omitempty" toml:"disableChecks,omitempty" yaml:"disableChecks,omitempty" export:"true"`
	DisableANSChecks  bool            `description:"Disables the challenge TXT record propagation checks against authoritative nameservers." json:"disableANSChecks,omitempty" toml:"disableANSChecks,omitempty" yaml:"disableANSChecks,omitempty" export:"true"`
	RequireAllRNS     bool            `description:"Requires the challenge TXT record to be propagated to all recursive nameservers." json:"requireAllRNS,omitempty" toml:"requireAllRNS,omitempty" yaml:"requireAllRNS,omitempty" export:"true"`
	DelayBeforeChecks ptypes.Duration `description:"Defines the delay before checking the challenge TXT record propagation." json:"delayBeforeChecks,omitempty" toml:"delayBeforeChecks,omitempty" yaml:"delayBeforeChecks,omitempty" export:"true"`
}

// HTTPChallenge contains HTTP challenge configuration.
type HTTPChallenge struct {
	EntryPoint string `description:"HTTP challenge EntryPoint" json:"entryPoint,omitempty" toml:"entryPoint,omitempty" yaml:"entryPoint,omitempty" export:"true"`
}

// TLSChallenge contains TLS challenge configuration.
type TLSChallenge struct{}

// NewLocalStore initializes a new LocalStore with a file name.
func NewLocalStore(string, *safe.Pool) *LocalStore {
	store := &LocalStore{}
	return store
}
