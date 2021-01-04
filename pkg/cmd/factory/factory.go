package factory

import (
	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/config"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/connection"
)

// Factory is an abstract type which provides access to
// the root configuration and connections for the CLI
type Factory struct {
	// Interface to read/write to the config
	Config config.IConfig
	// Creates a connection to the API
	Connection func() (connection.IConnection, error)
}
