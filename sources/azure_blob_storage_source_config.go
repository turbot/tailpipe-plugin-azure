package sources

import (
	"fmt"

	"github.com/hashicorp/hcl/v2"

	"github.com/turbot/tailpipe-plugin-sdk/artifact_source_config"
)

type AzureBlobStorageSourceConfig struct {
	artifact_source_config.ArtifactSourceConfigImpl

	// required to allow partial decoding
	Remain hcl.Body `hcl:",remain" json:"-"`

	AccountName string `hcl:"account_name"`
	Container   string `hcl:"container"`

	// TODO: determine if these are required
	Prefix *string `hcl:"prefix,optional"`
}

func (a *AzureBlobStorageSourceConfig) Validate() error {
	if a.AccountName == "" {
		return fmt.Errorf("account_name is required and cannot be empty")
	}

	if a.Container == "" {
		return fmt.Errorf("container is required and cannot be empty")
	}

	return nil
}

func (a *AzureBlobStorageSourceConfig) Identifier() string {
	return AzureBlobStorageSourceIdentifier
}
