package config

import "github.com/turbot/tailpipe-plugin-sdk/parse"

type AzureConnection struct {
	TenantId         *string  `hcl:"tenant_id"`
	SubscriptionId   *string  `hcl:"subscription_id"`
	ClientId         *string  `hcl:"client_id"`
	ClientSecret     *string  `hcl:"client_secret"`
}

func NewAzureConnection() parse.Config {
	return &AzureConnection{}
}

func (c *AzureConnection) Validate() error {
	return nil
}
