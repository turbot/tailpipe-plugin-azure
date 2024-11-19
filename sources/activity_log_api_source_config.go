package sources

import "fmt"

type ActivityLogAPISourceConfig struct {
	TenantId       *string `hcl:"tenant_id"`
	SubscriptionId *string `hcl:"subscription_id"`
	ClientId       *string `hcl:"client_id"`
	ClientSecret   *string `hcl:"client_secret"`
}

func (a *ActivityLogAPISourceConfig) Validate() error {
	if a.TenantId == nil {
		return fmt.Errorf("tenant_id is required")
	}
	if a.SubscriptionId == nil {
		return fmt.Errorf("subscription_id is required")
	}
	if a.ClientId == nil {
		return fmt.Errorf("client_id is required")
	}
	if a.ClientSecret == nil {
		return fmt.Errorf("client_secret is required")
	}
	return nil
}
