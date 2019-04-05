package types

import "github.com/biome-sh/biome-go/components/core/ident"

type PagedSearchResp struct {
	RangeStart int           `mapstructure:"range_start"`
	RangeEnd   int           `mapstructure:"range_end"`
	TotalCount int           `mapstructure:"total_count"`
	Data       []ident.Ident `mapstructure:"data"`
}

type RDepsResp struct {
	RDeps []string `mapstructure:"rdeps"`
}

type ChannelName struct {
	Name string `mapstructure:"name"`
}

type ShowPackageResp struct {
	Channels []string    `mapstructure:"channels"`
	Ident    ident.Ident `mapstructure:"ident"`
}
