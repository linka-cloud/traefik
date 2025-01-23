//go:build !plugins

package middleware

import (
	"context"
	"net/http"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
	"github.com/traefik/traefik/v3/pkg/config/static"
	"github.com/traefik/traefik/v3/pkg/plugins"
	"github.com/traefik/traefik/v3/pkg/provider"
	"github.com/traefik/traefik/v3/pkg/provider/noop"
)

type PluginsBuilder interface {
	Build(pName string, config map[string]interface{}, middlewareName string) (plugins.Constructor, error)
	BuildProvider(string, static.PluginConf) (provider.Provider, error)
}

type NoopPluginsBuilder struct{}

func (NoopPluginsBuilder) Build(pName string, config map[string]interface{}, middlewareName string) (interface{}, error) {
	return nil, nil
}

func (NoopPluginsBuilder) BuildProvider(string, static.PluginConf) (provider.Provider, error) {
	return noop.Provider{}, nil
}

func findPluginConfig(rawConfig map[string]dynamic.PluginConf) (string, map[string]interface{}, error) {
	return "", map[string]interface{}{}, nil
}

func newTraceablePlugin(ctx context.Context, name string, plug plugins.Constructor, next http.Handler) (http.Handler, error) {
	return next, nil
}
