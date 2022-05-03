package filter

import (
	"flag"
	"fmt"
	"os"

	"github.com/aleksaan/diskusage/pkg/models"
)

var cfg = &models.TFilterConfig{}

var defaultPath = ""
var defaultDepth = uint(0)
var defaultHr = false
var defaultTop = uint(20)
var defaultFilter = "df"
var defaultSize = "p"

var path *string
var top *uint
var depth *uint
var hr *bool
var filter *string
var size *string

func init() {
	top = flag.Uint("top", defaultTop, "How many biggest directories will be founded (default - 20)")
	depth = flag.Uint("depth", defaultDepth, "Depth of analysed directory's levels")
	hr = flag.Bool("hr", defaultHr, "Switch to human friendly mode of representation outputed results (default - JSON)")
	filter = flag.String("filter", defaultFilter, "What objects will be analysed: d - directories only, f - files only, df - both of them")
	size = flag.String("size", defaultSize, "Ordering by: с - clean size (without subdirectories), f - full size (with subdirectories)")
	path = flag.String("path", defaultPath, "Path or its part (filter results by including of this string")
	flag.Parse()
	cfg.Depth = *depth
	cfg.Filter = *filter
	if cfg.Filter != "d" && cfg.Filter != "f" && cfg.Filter != "df" {
		fmt.Printf("ERROR: Wrong value of agrument '-filter=%s'\n", cfg.Filter)
		os.Exit(1)
	}
	cfg.Hr = *hr
	cfg.Size = *size
	if cfg.Size != "c" && cfg.Size != "f" {
		fmt.Printf("ERROR: Wrong value of agrument '-size=%s'\n", cfg.Filter)
		os.Exit(1)
	}
	cfg.Top = *top
	cfg.Path = *path
}
