package main

import (
	"time"

	"github.com/aleksaan/diskusage/config"
	"github.com/aleksaan/diskusage/console"
	"github.com/aleksaan/diskusage/files"
	"github.com/aleksaan/diskusage/printer"
)

//-----------------------------------------------------------------------------------------
//main function
func main() {

	//start timer
	start := time.Now()

	//gets command line arguments
	cfg, opt := config.LoadConfig()

	console.Init()
	defer console.Close()

	printer.PrintConfig(cfg)

	//get files	
	analyzer.Files := &files.TFiles{}
	analyzer.Cfg := Cfg
	analyzer.ScanDir(cfg.Main.Path, 1)

	//sort files by size
	analyzer.Files.Sort(diskusage.InputArgs.Sort)

	//print files results to console
	files.PrintFilesSizes()
	//finish work and calculate elapsed time
	elapsed := time.Since(start)

	files.SaveToCsv()

	//print overall info
	total := files.GetOverallInfo(elapsed)
	total.PrintOverallInfo()

}
