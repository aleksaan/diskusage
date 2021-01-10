package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/aleksaan/diskusage/files"
	"gopkg.in/yaml.v3"
)

var defaultConfigFile = "diskusage_config.yaml"

//set of configuration options
var opt Options

//Cfg -
var Cfg = &Config{}

//Load - load config from a yaml file & creates it if not exist
func (c *Config) Load() {

	if opt.ConfigFile == nil {
		opt.ConfigFile = &defaultConfigFile
	}

	if !files.CheckFileIsExist(*opt.ConfigFile) {
		c.createDefaultConfigYamlFile()
	}

	_ = c.readConfigFromYamlFile(opt.ConfigFile)
	c.setDefaultValues()
}

func (c *Config) readConfigFromYamlFile(location *string) error {

	yamlFile, err := ioutil.ReadFile(*location)
	if err != nil {
		log.Fatalf("Failed to read config file: %s", *location)
	}

	err = yaml.Unmarshal(yamlFile, c)

	if err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}

	return err
}

/* func (c *Config) saveConfigToYamlFile(config *Config, fileName *string) {
	file, err := yaml.Marshal(config)
	if err != nil {
		log.Fatalf("Failed to marshal config file: %v", err)
	}
	err = ioutil.WriteFile(*fileName, file, os.ModePerm)
	if err != nil {
		log.Fatalf("Failed to save config file: %v", err)
	}
} */

func (c *Config) createDefaultConfigYamlFile() {
	f := files.CreateFile(&defaultConfigFile)
	defer f.Close()
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(f, "analyzer:")
	files.PrintEndOfLine(f)
	fmt.Fprintf(f, "  path: %s", dir)
	files.PrintEndOfLine(f)
	fmt.Fprintf(f, "  depth: 5")
	files.PrintEndOfLine(f)
	fmt.Fprintf(f, "printer:")
	files.PrintEndOfLine(f)
	fmt.Fprintf(f, "  limit: 20")
	files.PrintEndOfLine(f)
	fmt.Fprintf(f, "  units:")
	files.PrintEndOfLine(f)
	fmt.Fprintf(f, "  toTextFile: diskusage_out.txt")
	files.PrintEndOfLine(f)
	fmt.Fprintf(f, "  toYamlFile: diskusage_out.yaml")
}
