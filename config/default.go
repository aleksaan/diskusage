package config

import "os"

const (
	defaultDepth                 = 5
	defaultLimit                 = 20
	defaultSizeCalculatingMethod = "cumulative"
	defaultUnits                 = ""
	defaultSort                  = "size_desc"
	defaultFilterByObject        = "folders&files"
	defaultToTextFile            = "diskusage_out.txt"
	defaultToYamlFile            = ""
	//DefaultToFile = "<no file>"
)

func (c *Config) setDefaultValues() {

	if c.Filter.Depth == nil {
		d := defaultDepth
		c.Filter.Depth = &d
	}

	if c.Analyzer.Path == nil || *c.Analyzer.Path == "" {
		dir, _ := os.Getwd()
		c.Analyzer.Path = &dir
	}

	if c.Analyzer.SizeCalculatingMethod == nil {
		h := defaultSizeCalculatingMethod
		c.Analyzer.SizeCalculatingMethod = &h
	} else if *c.Analyzer.SizeCalculatingMethod != "plain" && *c.Analyzer.SizeCalculatingMethod != "cumulative" {
		h := defaultSizeCalculatingMethod
		c.Analyzer.SizeCalculatingMethod = &h
	}

	if c.Filter.Limit == nil {
		l := defaultLimit
		c.Filter.Limit = &l
	}

	if c.Printer.Units == nil {
		u := defaultUnits
		c.Printer.Units = &u
	}

	if c.Filter.FilterByObjectType == nil {
		u := defaultFilterByObject
		c.Filter.FilterByObjectType = &u
	}

	if c.Printer.ToTextFile == nil {
		u := defaultToTextFile
		c.Printer.ToTextFile = &u
	}

	if c.Printer.ToYamlFile == nil {
		u := defaultToYamlFile
		c.Printer.ToYamlFile = &u
	}

	if c.Printer.Sort == nil {
		s := defaultSort
		c.Printer.Sort = &s
	}
}
