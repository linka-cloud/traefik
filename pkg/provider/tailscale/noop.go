//go:build !tailscale

package tailscale

import (
	"github.com/traefik/traefik/v3/pkg/provider/noop"
)

type Provider struct {
	noop.Provider
	ResolverName string `export:"true"`
}
