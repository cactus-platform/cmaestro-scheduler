package config

type RuntimeConfig struct {
}

func (c *StaticConfig) ResolveRuntime() (*RuntimeConfig, error) {
	return &RuntimeConfig{}, nil
}
