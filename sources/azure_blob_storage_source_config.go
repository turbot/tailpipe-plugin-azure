package sources

import (
	"fmt"
	"strings"

	"github.com/hashicorp/hcl/v2"

	"github.com/turbot/tailpipe-plugin-sdk/artifact_source_config"
)

type AzureBlobStorageSourceConfig struct {
	artifact_source_config.ArtifactSourceConfigBase

	// required to allow partial decoding
	Remain hcl.Body `hcl:",remain" json:"-"`

	AccountName string `hcl:"account_name"`
	Container   string `hcl:"container"`

	// TODO: determine if these are required
	Prefix     string   `hcl:"prefix"`
	Extensions []string `hcl:"extensions"`
}

func (a *AzureBlobStorageSourceConfig) Validate() error {
	if a.AccountName == "" {
		return fmt.Errorf("account_name is required and cannot be empty")
	}

	if a.Container == "" {
		return fmt.Errorf("container is required and cannot be empty")
	}

	if len(a.Extensions) > 0 {
		var invalidExtensions []string
		for _, e := range a.Extensions {
			if len(e) == 0 {
				invalidExtensions = append(invalidExtensions, "<empty>")
			} else if e[0] != '.' {
				invalidExtensions = append(invalidExtensions, e)
			}
		}
		if len(invalidExtensions) > 0 {
			return fmt.Errorf("invalid extensions: %s", strings.Join(invalidExtensions, ","))
		}
	}

	return nil
}

func (a *AzureBlobStorageSourceConfig) Identifier() string {
	return AzureBlobStorageSourceIdentifier
}
