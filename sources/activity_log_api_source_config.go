package sources

import "fmt"

type ActivityLogAPISourceConfig struct {
	// TODO: #config determine if we can support other authentication types and configuration requirements for these
	TenantId       string `hcl:"tenant_id"`
	SubscriptionId string `hcl:"subscription_id"`
	ClientId       string `hcl:"client_id"`
	ClientSecret   string `hcl:"client_secret"`
}

func (a *ActivityLogAPISourceConfig) Validate() error {
	if a.TenantId == "" {
		return fmt.Errorf("tenant_id is required")
	}
	if a.SubscriptionId == "" {
		return fmt.Errorf("subscription_id is required")
	}
	if a.ClientId == "" {
		return fmt.Errorf("client_id is required")
	}
	if a.ClientSecret == "" {
		return fmt.Errorf("client_secret is required")
	}
	return nil
}
