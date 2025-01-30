package activity_log_api

type ActivityLogAPISourceConfig struct {
}

func (a *ActivityLogAPISourceConfig) Validate() error {
	return nil
}

func (a *ActivityLogAPISourceConfig) Identifier() string {
	return ActivityLogAPISourceIdentifier
}
