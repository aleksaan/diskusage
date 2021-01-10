package config

//Config - utility configuration
type Config struct {
	Analyzer AnalyzerConfig `yaml:"analyzer"`
	Printer  PrinterConfig  `yaml:"printer"`
}

//AnalyzerConfig section
type AnalyzerConfig struct {
	Path  *string `yaml:"path,omitempty"`
	Depth *int    `yaml:"depth,omitempty"`
}

//PrinterConfig section
type PrinterConfig struct {
	Limit      *int    `yaml:"limit,omitempty"`
	Units      *string `yaml:"units,omitempty"`
	ToTextFile *string `yaml:"toTextFile,omitempty"`
	ToYamlFile *string `yaml:"toYamlFile,omitempty"`
	Sort       *string
}
