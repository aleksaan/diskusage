package main

import (
	"github.com/aleksaan/diskusage/analyzer"
	"github.com/aleksaan/diskusage/config"

	"github.com/aleksaan/diskusage/console"
	"github.com/aleksaan/diskusage/printer"
)

//-----------------------------------------------------------------------------------------
//main function
func main() {
	//var cfg = config.Cfg

	//cfg.Load()
	printer.Load()
	printer.PrintAbout()
	analyzer.Run()
	printer.Run()
	console.WaitExit(config.Cfg.Printer.ToTextFile == "")
	printer.Close()

}
