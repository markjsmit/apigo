package config

import "github.com/maxpower89/gotroller/pkg/config"

type Config struct {
	DefaultEntityConfig *EntityConfig
	Gotroller           *config.Config
	Docs                *DocsConfig
}

func NewConfig() *Config {
	return &Config{
		DefaultEntityConfig: NewEntityConfig(),
		Gotroller:           config.NewConfig(),
		Docs:                NewDocsConfig(),
	};
}
