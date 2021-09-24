package config

import (
	"fmt"
	"io/ioutil"
	"log"

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
		c.setDefaultValues()
		c.createDefaultConfigYamlFile()
	} else {
		_ = c.readConfigFromYamlFile(opt.ConfigFile)
		c.setDefaultValues()
	}
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

	fmt.Fprintf(f, ConfigTemplate, *c.Analyzer.Path)

	// fmt.Fprintf(f, "analyzer:")
	// files.PrintEndOfLine(f)
	// fmt.Fprintf(f, "  path:       %s", *c.Analyzer.Path)
	// files.PrintEndOfLine(f)
	// fmt.Fprintf(f, "  depth:      %d", *c.Analyzer.Depth)
	// files.PrintEndOfLine(f)
	// fmt.Fprintf(f, "  includeNestedFolders:  %s", *c.Analyzer.SizeCalculatingMethod)
	// files.PrintEndOfLine(f)
	// fmt.Fprintf(f, "printer:")
	// files.PrintEndOfLine(f)
	// fmt.Fprintf(f, "  limit:      %d", *c.Printer.Limit)
	// files.PrintEndOfLine(f)
	// fmt.Fprintf(f, "  units:      %s", *c.Printer.Units)
	// files.PrintEndOfLine(f)
	// fmt.Fprintf(f, "  filterBy:  %s", *c.Printer.FilterByObjectType)
	// files.PrintEndOfLine(f)
	// fmt.Fprintf(f, "  toTextFile: %s", *c.Printer.ToTextFile)
	// files.PrintEndOfLine(f)
	// fmt.Fprintf(f, "  toYamlFile: %s", *c.Printer.ToYamlFile)
}
