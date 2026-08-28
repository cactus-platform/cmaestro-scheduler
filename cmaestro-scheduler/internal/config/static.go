package config

type StaticConfig struct {
}

func Load() *StaticConfig {
	return &StaticConfig{}
}
