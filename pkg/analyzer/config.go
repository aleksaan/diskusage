package analyzer

import (
	"flag"
	"fmt"
	"os"

	"github.com/aleksaan/diskusage/pkg/models"
)

var cfg = &models.TAnalyserConfig{}
var defaultPath, _ = os.Getwd()
var defaultDepth uint = 0
var defaultHr = false
var defaultHrRows uint = 50

var path *string
var depth *uint
var hr *bool
var hrRows *uint

func init() {
	path = flag.String("path", defaultPath, "Starting point (path) to analyse (default value is a current path)")
	depth = flag.Uint("depth", defaultDepth, "Depth of analysis (how many levels of directories will be outputed)")
	hr = flag.Bool("hr", defaultHr, "Switch to human friendly mode of representation outputed results (default - JSON)")
	hrRows = flag.Uint("hrrows", defaultHrRows, "Number rows will be printed in the hr mode (default - "+fmt.Sprintf("%d", uint64(defaultHrRows))+" rows) ")
}

func CfgInit() {
	flag.Parse()
	chk, err := isDirectory(*path)
	if err != nil {
		fmt.Println("Error: Path '" + *path + "' cannot be opened or not exists")
		os.Exit(1)
	}
	if !chk {
		fmt.Println("Error: Path '" + *path + "' is not a directory")
		os.Exit(1)
	}
	cfg.Path = *path
	cfg.Depth = *depth
	cfg.Hr = *hr
	cfg.HrRows = *hrRows
}

func isDirectory(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false, err
	}

	return fileInfo.IsDir(), err
}
