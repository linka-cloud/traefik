//go:build !influxdb2

package metrics

import (
	"context"

	"github.com/traefik/traefik/v3/pkg/observability/types"
)

func RegisterInfluxDB2(ctx context.Context, config *types.InfluxDB2) Registry {
	return &standardRegistry{}
}

func StopInfluxDB2() {}
