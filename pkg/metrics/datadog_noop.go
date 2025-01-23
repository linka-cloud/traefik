//go:build !datadog

package metrics

import (
	"context"

	"github.com/traefik/traefik/v3/pkg/types"
)

func RegisterDatadog(ctx context.Context, config *types.Datadog) Registry {
	return &standardRegistry{}
}

func StopDatadog() {}
