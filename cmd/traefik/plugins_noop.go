//go:build !plugins

package main

import (
	"github.com/traefik/traefik/v3/pkg/config/static"
	"github.com/traefik/traefik/v3/pkg/server/middleware"
)

func createPluginBuilder(staticConfiguration *static.Configuration) (middleware.PluginsBuilder, error) {
	return nil, nil
}
