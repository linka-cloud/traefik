//go:build !nomad

package nomad

import (
	"github.com/traefik/traefik/v3/pkg/provider/noop"
)

type ProviderBuilder struct {
	noop.ProviderBuilder
}
