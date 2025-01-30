package azure

import (
	"github.com/turbot/go-kit/helpers"
	"github.com/turbot/tailpipe-plugin-azure/config"
	"github.com/turbot/tailpipe-plugin-azure/sources/activity_log_api"
	"github.com/turbot/tailpipe-plugin-azure/sources/blob_storage"
	"github.com/turbot/tailpipe-plugin-azure/tables/activity_log"
	"github.com/turbot/tailpipe-plugin-sdk/plugin"
	"github.com/turbot/tailpipe-plugin-sdk/row_source"
	"github.com/turbot/tailpipe-plugin-sdk/table"
)

type Plugin struct {
	plugin.PluginImpl
}

func init() {
	// Register tables, with type parameters:
	// 1. row struct
	// 2. table implementation
	table.RegisterTable[*activity_log.ActivityLog, *activity_log.ActivityLogTable]()

	// register sources
	row_source.RegisterRowSource[*activity_log_api.ActivityLogAPISource]()
	row_source.RegisterRowSource[*blob_storage.AzureBlobStorageSource]()
}

func NewPlugin() (_ plugin.TailpipePlugin, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = helpers.ToError(r)
		}
	}()

	p := &Plugin{
		PluginImpl: plugin.NewPluginImpl(config.PluginName),
	}

	return p, nil
}
