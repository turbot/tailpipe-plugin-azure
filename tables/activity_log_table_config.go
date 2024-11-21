package tables

type ActivityLogTableConfig struct{}

func (a *ActivityLogTableConfig) Validate() error {
	return nil
}

func (a *ActivityLogTableConfig) Identifier() string {
	return ActivityLogTableIdentifier
}
