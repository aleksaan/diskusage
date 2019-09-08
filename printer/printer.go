package printer

import (
	"fmt"
	"time"

	"github.com/aleksaan/diskusage/config"
	"github.com/aleksaan/diskusage/files"
)

//Run - run print process
func Run(cfg *config.Config, files *files.TFiles, totalTime time.Duration) {
	prepareData(cfg, files)
	printConfig(cfg)
	printFiles(cfg, preparedFiles)
	prepareOverallInfo(files, totalTime)
	printOverall()
}

func printConfig(cfg *config.Config) {
	fmt.Println("\nArguments:")
	fmt.Printf("   path: %s\n", *cfg.Analyzer.Path)
	fmt.Printf("   limit: %d\n", *cfg.Printer.Limit)
	fmt.Printf("   units: %s\n", *cfg.Printer.Units)
	fmt.Printf("   depth: %d\n", *cfg.Analyzer.Depth)
	//fmt.Printf("   sort: %s\n", cfg.Printer.Sort)
	fmt.Printf("   tofile: %s\n", *cfg.Printer.ToFile)
}

func printFiles(cfg *config.Config, files *files.TFiles) {
	var strfmt = "%3d.| %-7s %-" + fmt.Sprintf("%d", calculateMaxLenFilename()+2) + "s | SIZE: %8.2f %-4s | DEPTH: %d %s"
	var dirorfile = "PATH:"
	c := 0
	for _, f := range *preparedFiles {
		c++
		fmt.Printf(strfmt, c, dirorfile, f.RelativePath+f.Name, f.AdaptedSize, f.AdaptedUnit, f.Depth)
	}
}

func printOverall() {
	fmt.Printf("\nOverall info:\n")
	fmt.Printf("   Total time: %s\n", overallInfo.totalTime)
	fmt.Printf("   Total dirs: %d\n", overallInfo.totalDirs)
	fmt.Printf("   Total files: %d\n", overallInfo.totalFiles)
	fmt.Printf("   Total links: %d\n", overallInfo.totalLinks)
	fmt.Printf("   Total size: %.2f %s\n", overallInfo.totalAdaptedSize, overallInfo.totalAdaptedUnit)
	fmt.Printf("   Total size (bytes): %d\n", overallInfo.totalSize)
	fmt.Printf("   Unaccessible dirs & files: %d\n", overallInfo.totalNotAccessibleFiles)
}
