package printer

import (
	"fmt"
	"runtime"
	"time"

	"github.com/aleksaan/diskusage/config"
	"github.com/aleksaan/diskusage/files"
)

const (
	AppTitle   = "github/aleksaan/diskusage"
	AppVersion = "2.0.1"
	AppAuthor  = "Anufriev Alexander"
	AppYear    = "2019"
)

//Run - run print process
func Run(cfg *config.Config, files *files.TFiles, totalTime time.Duration) {
	files.Sort(*cfg.Printer.Sort)
	prepareData(cfg, files)
	printConfig(cfg)
	printFiles(cfg, preparedFiles)
	prepareOverallInfo(files, totalTime)
	printOverall()
}

//PrintAbout -
func PrintAbout() {
	fmt.Printf("About:\n   %s, %s, %s, %s\n", AppTitle, AppVersion, AppAuthor, AppYear)
}

func printConfig(cfg *config.Config) {
	fmt.Printf("\n")
	fmt.Println("Arguments:")
	fmt.Printf("   %-10s %s\n", "path:", *cfg.Analyzer.Path)
	fmt.Printf("   %-10s %d\n", "limit:", *cfg.Printer.Limit)
	units := *cfg.Printer.Units
	if units == "" {
		units = "<dynamic>"
	}
	fmt.Printf("   %-10s %s\n", "units:", units)
	fmt.Printf("   %-10s %d\n", "depth:", *cfg.Analyzer.Depth)
	fmt.Printf("   %-10s %s\n", "sort:", *cfg.Printer.Sort)
	fmt.Printf("   %-10s %s\n", "tofile:", *cfg.Printer.ToFile)
}

func printFiles(cfg *config.Config, files *files.TFiles) {
	fmt.Printf("\n")
	fmt.Printf("Results:\n")
	maxlen := calculateMaxLenFilename()
	var strfmt = "   %3d.| %-7s %-" + fmt.Sprintf("%d", maxlen+2) + "s | SIZE: %8.2f %-4s | DEPTH: %d %s"
	var dirorfile = "PATH:"
	for i, f := range *preparedFiles {
		fmt.Printf(strfmt, i+1, dirorfile, f.RelativePath+f.Name, f.AdaptedSize, f.AdaptedUnit, f.Depth, es())
	}
}

func printOverall() {
	fmt.Printf("\n")
	fmt.Printf("Overall info:\n")
	fmt.Printf("   Total time: %s\n", overallInfo.totalTime)
	fmt.Printf("   Total dirs: %d\n", overallInfo.totalDirs)
	fmt.Printf("   Total files: %d\n", overallInfo.totalFiles)
	fmt.Printf("   Total links: %d\n", overallInfo.totalLinks)
	fmt.Printf("   Total size: %.2f %s\n", overallInfo.totalAdaptedSize, overallInfo.totalAdaptedUnit)
	fmt.Printf("   Total size (bytes): %d\n", overallInfo.totalSize)
	fmt.Printf("   Unaccessible dirs & files: %d\n", overallInfo.totalNotAccessibleFiles)
}

func es() string {
	switch runtime.GOOS {
	case "windows":
		return "\r\n"
	default:
		return "\n"
	}
}
