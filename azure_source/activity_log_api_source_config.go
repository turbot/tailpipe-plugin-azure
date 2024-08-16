package azure_source

type ActivityLogAPISourceConfig struct {
	// TODO: #config determine if we can support other authentication types and configuration requirements for these
	TenantId       string `hcl:"tenant_id"`
	SubscriptionId string `hcl:"subscription_id"`
	ClientId       string `hcl:"client_id"`
	ClientSecret   string `hcl:"client_secret"`
}

func (a *ActivityLogAPISourceConfig) Validate() error {
	return nil
}
