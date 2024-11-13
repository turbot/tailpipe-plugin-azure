package azure

import (
	"log/slog"

	"github.com/turbot/go-kit/helpers"
	"github.com/turbot/tailpipe-plugin-azure/config"
	_ "github.com/turbot/tailpipe-plugin-azure/sources"
	"github.com/turbot/tailpipe-plugin-sdk/plugin"
	"github.com/turbot/tailpipe-plugin-sdk/table"

	// reference the table package to ensure that the tables are registered by the init functions
	_ "github.com/turbot/tailpipe-plugin-azure/tables"
)

type Plugin struct {
	plugin.PluginImpl
}

func NewPlugin() (_ plugin.TailpipePlugin, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = helpers.ToError(r)
		}
	}()
	slog.Info("Azure Plugin starting")
	//time.Sleep(10 * time.Second)
	slog.Info("Azure Plugin started")

	p := &Plugin{
		PluginImpl: plugin.NewPluginImpl("azure", config.NewAzureConnection),
	}

	// initialise table factory
	if err := table.Factory.Init(); err != nil {
		return nil, err
	}

	return p, nil
}
