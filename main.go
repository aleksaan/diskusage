package main

import (
	"github.com/aleksaan/diskusage/analyzer"
	"github.com/aleksaan/diskusage/config"

	"github.com/aleksaan/diskusage/console"
	"github.com/aleksaan/diskusage/printer"
)

//Cfg saves program configuration
//var cfg *config.Config = &config.Config{}

//-----------------------------------------------------------------------------------------
//main function
func main() {
	var cfg = config.Cfg


	cfg.Load()
	printer.Load()
	printer.PrintAbout()
	analyzer.Run()
	printer.Run()
	console.WaitExit(*config.Cfg.Printer.ToFile == "")
	printer.Close()
	
}
