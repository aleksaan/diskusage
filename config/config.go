package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"runtime"

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

	if !fileExists(*opt.ConfigFile) {
		createDefaultConfig()
	}

	cfg, err := readFile(opt.ConfigFile)
	cfg.setDefaults()

	return cfg, opt
}

func readFile(location *string) (*Config, error) {

	yamlFile, err := ioutil.ReadFile(*location)
	if err != nil {
		log.Fatalf("Failed to read config file: %s", *location)
	}

	cfg := new(Config)
	err = yaml.Unmarshal(yamlFile, cfg)

	if err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}

	return cfg, err
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

func createDefaultConfig() {
	f := createFile(&defaultConfigFile)
	defer f.Close()
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(f, "analyzer:")
	printEndOfLine(f)
	fmt.Fprintf(f, "  path: %s", dir)
	printEndOfLine(f)
	fmt.Fprintf(f, "  depth: 5")
	printEndOfLine(f)
	fmt.Fprintf(f, "printer:")
	printEndOfLine(f)
	fmt.Fprintf(f, "  limit: 20")
	printEndOfLine(f)
	fmt.Fprintf(f, "  fixunit:")
	printEndOfLine(f)
	fmt.Fprintf(f, "  tofile: out.txt")
}

func createFile(filename *string) *os.File {
	// open output file
	f, err := os.Create(*filename)
	if err != nil {
		panic(err)
	}
	return f
}

func fileExists(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}

func printEndOfLine(f *os.File) {
	fmt.Fprintf(f, "%s", es())
}

func es() string {
	switch runtime.GOOS {
	case "windows":
		return "\r\n"
	default:
		return "\n"
	}
}
