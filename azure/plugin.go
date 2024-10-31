package azure

import (
	"github.com/turbot/tailpipe-plugin-azure/sources"
	"github.com/turbot/tailpipe-plugin-azure/tables"
	"github.com/turbot/tailpipe-plugin-sdk/plugin"
	"github.com/turbot/tailpipe-plugin-sdk/row_source"
	"github.com/turbot/tailpipe-plugin-sdk/table"
)

func NewPlugin() (plugin.TailpipePlugin, error) {
	p := plugin.NewPlugin("azure")

	err := p.RegisterResources(
		&plugin.ResourceFunctions{
			Tables:  []func() table.Table{tables.NewActivityLogTable},
			Sources: []func() row_source.RowSource{sources.NewActivityLogAPISource},
		})

	if err != nil {
		return nil, err
	}

	return p, nil
}
