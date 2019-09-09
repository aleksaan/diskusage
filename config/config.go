package config

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/jessevdk/go-flags"
	"gopkg.in/yaml.v3"
)

var defaultConfigFile = "config.yaml"

// LoadConfig - load configuration file
func LoadConfig() (*Config, Options) {

	var opt Options
	_, err := flags.Parse(&opt)

	if err != nil {
		log.Fatalln("Error! Wrong options")
	}

	if opt.ConfigFile == nil {
		opt.ConfigFile = &defaultConfigFile
	}

	cfg := readFile(opt.ConfigFile)
	cfg.setDefaults()

	return cfg, opt
}

func readFile(location *string) *Config {

	yamlFile, err := ioutil.ReadFile(*location)
	if err != nil {
		log.Fatalf("Failed to read config file: %s", *location)
	}

	cfg := new(Config)
	err = yaml.Unmarshal(yamlFile, cfg)

	if err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}

	return cfg
}

func saveFile(config *Config, fileName *string) {
	file, err := yaml.Marshal(config)
	if err != nil {
		log.Fatalf("Failed to marshal config file: %v", err)
	}
	err = ioutil.WriteFile(*fileName, file, os.ModePerm)
	if err != nil {
		log.Fatalf("Failed to save config file: %v", err)
	}
}
