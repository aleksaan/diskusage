package config

import "os"

const (
	defaultDepth      = 5
	defaultLimit      = 20
	defaultUnits      = ""
	defaultSort       = "size_desc"
	defaultPrintOnly  = "folders&files"
	defaultToTextFile = "diskusage_out.txt"
	defaultToYamlFile = ""
	//DefaultToFile = "<no file>"
)

func (c *Config) setDefaultValues() {

	if c.Analyzer.Depth == nil {
		d := defaultDepth
		c.Analyzer.Depth = &d
	}
	if c.Analyzer.Path == nil {
		dir, _ := os.Getwd()
		c.Analyzer.Path = &dir
	}

	if c.Printer.Limit == nil {
		l := defaultLimit
		c.Printer.Limit = &l
	}

	if c.Printer.Units == nil {
		u := defaultUnits
		c.Printer.Units = &u
	}

	if c.Printer.PrintOnly == nil {
		u := defaultPrintOnly
		c.Printer.PrintOnly = &u
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
