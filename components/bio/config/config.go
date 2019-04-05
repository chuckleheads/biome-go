package config

import (
	"net/url"
)

// Config struct for Hab CLI
type Config struct {
	Origin            string  `mapstructure:"origin"`
	AuthToken         string  `mapstructure:"auth_token"`
	CTLSecret         string  `mapstructure:"ctl_secret"`
	BldrURL           url.URL `mapstructure:"bldr_url"`
	ArtifactCachePath string  `mapstructure:"artifact_cache"`
}
