package azure

import (
	"github.com/turbot/tailpipe-plugin-azure/azure_partition"
	"github.com/turbot/tailpipe-plugin-sdk/partition"
	"log/slog"
	"time"

	"github.com/turbot/tailpipe-plugin-azure/azure_source"
	"github.com/turbot/tailpipe-plugin-sdk/plugin"
	"github.com/turbot/tailpipe-plugin-sdk/row_source"
)

type Plugin struct {
	plugin.PluginBase
}

func NewPlugin() (plugin.TailpipePlugin, error) {
	p := &Plugin{}

	slog.Info("Azure Plugin starting")
	time.Sleep(10 * time.Second)
	slog.Info("Running...")

	err := p.RegisterResources(
		&plugin.ResourceFunctions{
			Partitions: []func() partition.Partition{azure_partition.NewActivityLogPartition},
			Sources:    []func() row_source.RowSource{azure_source.NewActivityLogAPISource},
		})

	if err != nil {
		return nil, err
	}

	return p, nil
}

func (t *Plugin) Identifier() string {
	return "azure"
}
