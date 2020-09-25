package printer

import (
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/aleksaan/diskusage/config"
	"github.com/aleksaan/diskusage/files"
)

const (
	AppTitle   = "github/aleksaan/diskusage"
	AppVersion = "2.0.3"
	AppAuthor  = "Anufriev Alexander"
	AppYear    = "2019"
)

var f *os.File

var initFlag = false

var cfg *config.Config

//Init -
func Init(cfg *config.Config) {
	if *cfg.Printer.ToFile != "" {
		f = createResultsFile(cfg.Printer.ToFile)
		os.Stdout = f
	}
}

//Close -
func Close() {
	f.Close()
}

//Run - run print process
func Run(cfg *config.Config, files *files.TFiles, totalTime time.Duration) {
	files.Sort(*cfg.Printer.Sort)
	prepareData(cfg, files)
	printConfig(cfg)
	printFiles(cfg, preparedFiles)
	prepareOverallInfo(files, totalTime)
	printOverall()
}

func createResultsFile(filename *string) *os.File {
	// open output file
	f, err := os.Create(*filename)
	if err != nil {
		panic(err)
	}
	return f
}

//PrintAbout -
func PrintAbout() {
	fmt.Printf("About:%s   %s, %s, %s, %s%s", es(), AppTitle, AppVersion, AppAuthor, AppYear, es())
}

func printConfig(cfg *config.Config) {
	fmt.Printf(es())
	fmt.Println("Arguments:")
	fmt.Printf("   %-10s %s%s", "path:", *cfg.Analyzer.Path, es())
	fmt.Printf("   %-10s %d%s", "limit:", *cfg.Printer.Limit, es())
	units := *cfg.Printer.Units
	if units == "" {
		units = "<dynamic>"
	}
	fmt.Printf("   %-10s %s%s", "units:", units, es())
	fmt.Printf("   %-10s %d%s", "depth:", *cfg.Analyzer.Depth, es())
	//fmt.Printf("   %-10s %s%s", "sort:", *cfg.Printer.Sort, es())
	tofile := *cfg.Printer.ToFile
	if tofile == "" {
		tofile = "<no file>"
	}
	fmt.Printf("   %-10s %s\n", "tofile:", tofile)
}

func printFiles(cfg *config.Config, files *files.TFiles) {
	fmt.Printf(es())
	fmt.Printf("Results:%s", es())
	maxlen := calculateMaxLenFilename()
	var strfmt = "   %3d.| %-7s %-" + fmt.Sprintf("%d", maxlen+2) + "s | SIZE: %8.2f %-4s | DEPTH: %d %s"
	var dirorfile = "PATH:"
	for i, f := range *preparedFiles {
		dirorfile = "PATH:"
		if !f.IsDir {
			dirorfile = "PATH:"
		}
		fmt.Printf(strfmt, i+1, dirorfile, f.RelativePath+f.Name, f.AdaptedSize, f.AdaptedUnit, f.Depth, es())
	}
}

func printOverall() {
	fmt.Printf(es())
	fmt.Printf("Overall info:%s", es())
	fmt.Printf("   Total time: %s%s", overallInfo.totalTime, es())
	fmt.Printf("   Total dirs: %d%s", overallInfo.totalDirs, es())
	fmt.Printf("   Total files: %d%s", overallInfo.totalFiles, es())
	fmt.Printf("   Total links: %d%s", overallInfo.totalLinks, es())
	fmt.Printf("   Total size: %.2f %s%s", overallInfo.totalAdaptedSize, overallInfo.totalAdaptedUnit, es())
	fmt.Printf("   Total size (bytes): %d%s", overallInfo.totalSize, es())
	fmt.Printf("   Unaccessible dirs & files: %d%s", overallInfo.totalNotAccessibleFiles, es())
}

func es() string {
	switch runtime.GOOS {
	case "windows":
		return "\r\n"
	default:
		return "\n"
	}
}
