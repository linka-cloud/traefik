package noop

import (
	"github.com/traefik/traefik/v3/pkg/config/dynamic"
	"github.com/traefik/traefik/v3/pkg/provider"
	"github.com/traefik/traefik/v3/pkg/safe"
)

var _ provider.Provider = (*Provider)(nil)

type Provider struct{}

func (Provider) Provide(chan<- dynamic.Message, *safe.Pool) error {
	return nil
}

func (Provider) Init() error {
	return nil
}

func (Provider) HandleConfigUpdate(dynamic.Configuration) {}

type ProviderBuilder struct{}

func (ProviderBuilder) BuildProviders() []provider.Provider {
	return nil
}
