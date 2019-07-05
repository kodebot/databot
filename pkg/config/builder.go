package config

import (
	"errors"
	"os"
	"strings"
)

// Builder provides options to build config from various sources
type Builder struct {
	sources []string
}

// NewBuilder returns a new builder to build configs
func NewBuilder() *Builder {
	return &Builder{sources: []string{}}
}

// UseEnv adds environment variables as one of the source of config
func (b *Builder) UseEnv() *Builder {
	b.sources = append(b.sources, "env")
	return b
}

// UseTomlFile adds the given toml file as one of the source of config
func (b *Builder) UseTomlFile(path string) *Builder {
	panic(errors.New("toml config sources is not supported yet"))
	b.sources = append(b.sources, path)
	return b
}

// Build sets the config options from the sources configured
func (b *Builder) Build() {
	buildWithDefaults()
	// todo: make sure the order of loading is in the same order as sources were requested to be added
	updateFromEnv()
}

func buildWithDefaults() {
	options = &Options{
		useDBCache:       defaultUseDBCache,
		cacheDBType:      defaultCacheDBType,
		cacheDBConStr:    defaultCacheConStr,
		exportToDBConStr: defaultExportToDBConStr,
	}
}

func updateFromEnv() {
	lines := os.Environ()
	for _, line := range lines {
		fields := strings.SplitN(line, "=", 2)
		key := fields[0]
		val := fields[1]

		switch key {
		case "DATABOT_USE_DB_CACHE":
			options.useDBCache = parseBool(val, false)
		case "DATABOT_CACHE_DB_TYPE":
			options.cacheDBType = val
		case "DATABOT_CACHE_DB_CON_STR":
			options.cacheDBConStr = val
		case "DATABOT_EXPORT_TO_DB_CON_STR":
			options.exportToDBConStr = val
		}
	}
}

func parseBool(value string, fallback bool) bool {
	if value == "" {
		return fallback
	}

	value = strings.ToLower(value)
	if value == "1" || value == "yes" || value == "true" || value == "on" {
		return true
	}

	return false
}
