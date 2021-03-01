package config

//Config - utility configuration
type Config struct {
	Analyzer AnalyzerOptions `yaml:"analyzerOptions"`
	Filter   FilterOptions   `yaml:"filterOptions"`
	Printer  PrinterOptions  `yaml:"printerOptions"`
}

//AnalyzerOptions section
type AnalyzerOptions struct {
	Path                  *string `yaml:"path,omitempty"`
	SizeCalculatingMethod *string `yaml:"sizeCalculatingMethod,omitempty"`
}

//FilterOptions section
type FilterOptions struct {
	Depth              *int    `yaml:"depth,omitempty"`
	Limit              *int    `yaml:"limit,omitempty"`
	FilterByObjectType *string `yaml:"filterByObjectType,omitempty"`
}

//PrinterOptions section
type PrinterOptions struct {
	Units      *string `yaml:"units,omitempty"`
	ToTextFile *string `yaml:"toTextFile,omitempty"`
	ToYamlFile *string `yaml:"toYamlFile,omitempty"`
	Sort       *string
}
