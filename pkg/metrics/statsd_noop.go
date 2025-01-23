//go:build !statsd

package metrics

import (
	"context"

	"github.com/traefik/traefik/v3/pkg/types"
)

func RegisterStatsd(ctx context.Context, config *types.Statsd) Registry {
	return &standardRegistry{}
}

func StopStatsd() {}
