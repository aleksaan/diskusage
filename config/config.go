package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

var Cfg *Config

func init() {
	e := godotenv.Load() //Загрузить файл .env
	if e != nil {
		fmt.Print(e)
	}
	initDefaultsValues()
	cfgInit()
}

func cfgInit() {
	Cfg = &Config{}
	Cfg.Analyzer.Path = getEnvAsStr("pathToScan", defaultPath)

	Cfg.Analyzer.SizeCalculatingMethod = getEnvAsStr("sizeCalculatingMethod", defaultSizeCalculatingMethod)
	if Cfg.Analyzer.SizeCalculatingMethod != "plain" && Cfg.Analyzer.SizeCalculatingMethod != "cumulative" {
		Cfg.Analyzer.SizeCalculatingMethod = defaultSizeCalculatingMethod
	}

	Cfg.Filter.Depth = getEnvAsInt("depth", defaultDepth)

	Cfg.Filter.FilterByObjectType = getEnvAsStr("filterByObjectType", defaultFilterByObject)
	if Cfg.Filter.FilterByObjectType != "folders" && Cfg.Filter.FilterByObjectType != "files" && Cfg.Filter.FilterByObjectType != "folders&files" {
		Cfg.Filter.FilterByObjectType = defaultFilterByObject
	}

	Cfg.Filter.Limit = getEnvAsInt("limit", defaultLimit)

	Cfg.Printer.Units = getEnvAsStr("units", defaultUnits)

	Cfg.Printer.ToTextFile = getEnvAsStr("toTextFile", defaultToTextFile)

	Cfg.Printer.ToYamlFile = getEnvAsStr("toYamlFile", defaultToYamlFile)

	Cfg.Printer.Sort = defaultSort
}

// Simple helper function to read an environment or return a default value
func getEnvAsStr(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

// Simple helper function to read an environment variable into integer or return a default value
func getEnvAsInt(key string, defaultVal int) int {
	if value, exists := os.LookupEnv(key); exists {
		if valueInt, err := strconv.Atoi(value); err == nil {
			return valueInt
		}
	}
	return defaultVal
}

// Helper to read an environment variable into a bool or return default value
func getEnvAsBool(key string, defaultVal bool) bool {
	if value, exists := os.LookupEnv(key); exists {
		if valueBool, err := strconv.ParseBool(value); err == nil {
			return valueBool
		}
	}
	return defaultVal
}

// Helper to read an environment variable into a string slice or return default value
func getEnvAsSlice(key string, defaultVal []string, sep string) []string {
	if value, exists := os.LookupEnv(key); exists && value != "" {

		valueSlice := strings.Split(value, sep)
		return valueSlice
	}

	return defaultVal
}
