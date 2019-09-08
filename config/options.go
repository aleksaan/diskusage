package config

// Options with cli flags
type Options struct {
	ConfigFile *string `short:"c" long:"config" description:"Path to YAML config file"`
	Version    bool    `short:"v" long:"version" description:"Print version"`
}
