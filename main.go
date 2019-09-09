package main

import (
	"time"

	"github.com/aleksaan/diskusage/analyzer"
	"github.com/aleksaan/diskusage/config"
	"github.com/aleksaan/diskusage/console"
	"github.com/aleksaan/diskusage/printer"
)

//-----------------------------------------------------------------------------------------
//main function
func main() {
	//start timer
	start := time.Now()
	printer.PrintAbout()

	//gets command line arguments
	cfg, _ := config.LoadConfig()

	//get files
	analyzer.Run(cfg)

	//print files results to console
	printer.Run(cfg, analyzer.Files, time.Since(start))

	console.WaitExit()

}
