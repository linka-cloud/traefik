//go:build !ecs

package ecs

import (
	"github.com/traefik/traefik/v3/pkg/provider/noop"
)

type Provider struct {
	noop.Provider
}
