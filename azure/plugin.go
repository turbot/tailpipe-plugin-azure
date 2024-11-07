package azure

import (
	"log/slog"

	"github.com/turbot/go-kit/helpers"
	"github.com/turbot/tailpipe-plugin-azure/config"
	"github.com/turbot/tailpipe-plugin-azure/sources"
	"github.com/turbot/tailpipe-plugin-azure/tables"
	"github.com/turbot/tailpipe-plugin-sdk/plugin"
	"github.com/turbot/tailpipe-plugin-sdk/row_source"
	"github.com/turbot/tailpipe-plugin-sdk/table"
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

	// register the tables that we provide
	resources := &plugin.ResourceFunctions{
		Tables: []func() table.Table{
			tables.NewActivityLogTable,
		},
		Sources: []func() row_source.RowSource{sources.NewActivityLogAPISource},
	}

	if err := p.RegisterResources(resources); err != nil {
		return nil, err
	}

	return p, nil
}
