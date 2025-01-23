//go:build !docker

package docker

import (
	"time"

	ptypes "github.com/traefik/paerser/types"

	"github.com/traefik/traefik/v3/pkg/provider/noop"
)

type Provider struct {
	noop.Provider
	HTTPClientTimeout time.Duration
}

type SwarmProvider struct {
	noop.Provider
	RefreshSeconds    ptypes.Duration
	HTTPClientTimeout time.Duration
}
