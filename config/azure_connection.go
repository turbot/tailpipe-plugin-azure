package config

import "github.com/turbot/tailpipe-plugin-sdk/parse"

type AzureConnection struct {
}

func NewAzureConnection() parse.Config {
	return &AzureConnection{}
}

func (c *AzureConnection) Validate() error {
	return nil
}
