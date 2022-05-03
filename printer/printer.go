package printer

import (
	"encoding/json"
	"fmt"

	"github.com/aleksaan/diskusage/analyzer"
	"github.com/aleksaan/diskusage/config"
	"github.com/aleksaan/diskusage/models"
)

var files = &models.TFiles{}

//Run - run print process
func Run() {
	files = analyzer.FinalFiles
	files.Sort()
	prepareData()
	writeResultToConsole()
}

func prepareData() {
	var c = 0
	var isDir = true
	if config.Cfg.FilterByObjectType == "files" {
		isDir = false
	}

	//are files & folders both we need (or not)
	var isAll = config.Cfg.FilterByObjectType == "folders&files"

	for _, f := range *analyzer.FinalFiles {
		if f.Depth <= config.Cfg.Depth && (f.IsDir == isDir || isAll) {
			c++
			//break if we up to defined limit
			if isExceedLimit(c, &config.Cfg.Limit) {
				break
			}
			f.FullPath = config.Cfg.Path + "\\" + f.RelativePath + f.Name
			f.Number = c
			//*filesToPrint = append(*filesToPrint, f)
			analyzer.Result.Files = append(analyzer.Result.Files, f)
		}
	}
}

func isExceedLimit(checkedValue int, limit *int) bool {
	return checkedValue > *limit && *limit != 0
}

func writeResultToConsole() {
	b, err := json.Marshal(analyzer.Result)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
	fmt.Println(string(b))
}
