package config

//Config - utility configuration
type Config struct {
	Path                  string `json:"path,omitempty"`
	SizeCalculatingMethod string `json:"sizeCalculatingMethod,omitempty"`
	Depth                 int    `json:"depth,omitempty"`
	Limit                 int    `json:"limit,omitempty"`
	FilterByObjectType    string `json:"filterByObjectType,omitempty"`
	Units                 string `json:"units,omitempty"`
	ToTextFile            string `json:"toTextFile,omitempty"`
	Sort                  string
}
