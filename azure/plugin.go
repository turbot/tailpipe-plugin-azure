package azure

import (
	"github.com/turbot/tailpipe-plugin-azure/config"
	"github.com/turbot/tailpipe-plugin-azure/sources"
	"github.com/turbot/tailpipe-plugin-azure/tables"
	"github.com/turbot/tailpipe-plugin-sdk/plugin"
	"github.com/turbot/tailpipe-plugin-sdk/row_source"
	"github.com/turbot/tailpipe-plugin-sdk/table"
)

type Plugin struct {
	plugin.PluginBase
}

func NewPlugin() (plugin.TailpipePlugin, error) {

	p := &Plugin{
		PluginBase: plugin.NewPluginBase("azure", config.NewAzureConnection),
	}

	// register the tables that we provide
	resources := &plugin.ResourceFunctions{
		Tables:  []func() table.Table{tables.NewActivityLogTable},
		Sources: []func() row_source.RowSource{sources.NewActivityLogAPISource},
	}

	if err := p.RegisterResources(resources); err != nil {
		return nil, err
	}

	return p, nil
}
