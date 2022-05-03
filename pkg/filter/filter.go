package filter

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/aleksaan/diskusage/pkg/models"
)

// 	"encoding/json"
// 	"fmt"

var aresults = &models.TAnalyserResult{}
var workfiles = &models.TFiles{}
var fresults = &models.TFilterResult{}

func Start() {
	fresults.Config = *cfg
	getInput()
	doFiltering()
	doSorting()
	doTop()
	if cfg.Hr {
		WriteHumanReadableToConsole()
	} else {
		WriteJSONToConsole()
	}
}

func getInput() {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	err := json.Unmarshal([]byte(line), &aresults)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Input data is in not compatible format: ", err)
	}
}

func doFiltering() {
	var isDir = true
	if cfg.Filter == "f" {
		isDir = false
	}
	var isAll = cfg.Filter == "df"

	for _, f := range aresults.Files {
		if ((f.Depth <= cfg.Depth) || cfg.Depth == 0) && (f.IsDir == isDir || isAll) && strings.Contains(aresults.Config.Path+f.RelativePath, cfg.Path) {
			*workfiles = append(*workfiles, f)
		}
	}
}

func doSorting() {
	workfiles.Sort(cfg.Size)
}

func doTop() {
	for i, f := range *workfiles {
		if i == int(cfg.Top) {
			break
		}
		fresults.Files = append(fresults.Files, f)
	}
}
