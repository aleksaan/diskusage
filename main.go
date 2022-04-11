package main

import (
	"github.com/aleksaan/diskusage/analyzer"
	"github.com/aleksaan/diskusage/printer"
)

//-----------------------------------------------------------------------------------------
//main function
func main() {
	//var cfg = config.Cfg

	//cfg.Load()
	//printer.Load()
	//printer.PrintAbout()
	analyzer.Run()
	printer.Run()
	//console.WaitExit(config.Cfg.ToTextFile == "")
	//printer.Close()
}
