package config

import (
	"flag"
	"fmt"
)

var Cfg *Config

func init() {
	initDefaultsValues()
	cfgInit()
}

func cfgInit() {
	Cfg = &Config{}
	pathToScan := flag.String("pathToScan", defaultPath, "Start point (path) for analysis (default value is current path)")
	sizeCalculatingMethod := flag.String("sizeCalculatingMethod", defaultSizeCalculatingMethod, "Method of calculating sizes (plain | cumulative)")
	depth := flag.Int("depth", defaultDepth, fmt.Sprintf("Depth of analysed levels of folder tree (0 - full depth, %d - default value)", defaultDepth))
	filterByObjectType := flag.String("filterByObjectType", defaultFilterByObject, "Type of objects that will be analyzed (folders | files | folders&files)")
	limit := flag.Int("limit", defaultLimit, fmt.Sprintf("Number of biggest folders/files will be outputed (0 - all folders, %d - default value)", defaultLimit))
	units := flag.String("units", defaultUnits, "Type of objects that will be analyzed (folders | files | folders&files)")

	flag.Parse()

	Cfg.Path = *pathToScan
	Cfg.SizeCalculatingMethod = *sizeCalculatingMethod
	Cfg.Depth = *depth
	Cfg.FilterByObjectType = *filterByObjectType
	Cfg.Limit = *limit
	Cfg.Units = *units

	if Cfg.SizeCalculatingMethod != "plain" && Cfg.SizeCalculatingMethod != "cumulative" {
		Cfg.SizeCalculatingMethod = defaultSizeCalculatingMethod
	}

	if Cfg.FilterByObjectType != "folders" && Cfg.FilterByObjectType != "files" && Cfg.FilterByObjectType != "folders&files" {
		Cfg.FilterByObjectType = defaultFilterByObject
	}

	Cfg.ToTextFile = defaultToTextFile

	Cfg.Sort = defaultSort
}
