package printer

import (
	"fmt"
	"math"
	"os"
	"runtime"
	"unicode/utf8"

	"github.com/aleksaan/diskusage/analyzer"
	"github.com/aleksaan/diskusage/config"
	filespkg "github.com/aleksaan/diskusage/files"
)

const (
	AppTitle   = "github/aleksaan/diskusage"
	AppVersion = "2.1.0"
	AppAuthor  = "Anufriev Alexander"
	AppYear    = "2020"
)

var f *os.File
var cfg *config.Config
var initFlag = false
var files *filespkg.TFiles
var overallInfo *analyzer.TOverallInfo
var filesToPrint = &filespkg.TFiles{}

//Load -
func Load() {
	cfg = config.Cfg
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
func Run() {
	files = analyzer.FinalFiles
	overallInfo = analyzer.OverallInfo
	files.Sort(*cfg.Printer.Sort)
	prepareData()
	printConfig()
	printFiles()
	//overallInfo.totalTime = totalTime
	//prepareOverallInfo(files, totalTime)
	printOverall(overallInfo)
	printSystemReport()
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

func printConfig() {
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

func printFiles() {
	fmt.Printf(es())
	fmt.Printf("Results:%s", es())
	maxlen := calculateMaxLenFilename()
	var strfmt = "   %3d.| %-7s %-" + fmt.Sprintf("%d", maxlen+2) + "s | SIZE: %8.2f %-4s | DEPTH: %d %s"
	var dirorfile = "PATH:"
	for i, f := range *filesToPrint {
		dirorfile = "PATH:"
		if !f.IsDir {
			dirorfile = "PATH:"
		}
		fmt.Printf(strfmt, i+1, dirorfile, f.RelativePath+f.Name, f.AdaptedSize, f.AdaptedUnit, f.Depth, es())
	}
}

func printOverall(overallInfo *analyzer.TOverallInfo) {
	fmt.Printf(es())
	fmt.Printf("Overall info:%s", es())
	fmt.Printf("   Total time: %s%s", overallInfo.TotalTime, es())
	fmt.Printf("   Total dirs: %d%s", overallInfo.TotalDirs, es())
	fmt.Printf("   Total files: %d%s", overallInfo.TotalFiles, es())
	fmt.Printf("   Total links: %d%s", overallInfo.TotalLinks, es())
	fmt.Printf("   Total size: %.2f %s%s", overallInfo.TotalAdaptedSize, overallInfo.TotalAdaptedUnit, es())
	fmt.Printf("   Total size (bytes): %d%s", overallInfo.TotalSize, es())
	fmt.Printf("   Unaccessible dirs & files: %d%s", overallInfo.TotalNotAccessibleFiles, es())
}

func printSystemReport() {
	fmt.Printf(es())
	fmt.Printf("System resources:%s", es())
	// f := &filespkg.TFile{}
	// var sizeoff = int(unsafe.Sizeof(f))
	// sizeInBytes := int64(len(*analyzer.FinalFiles) * sizeoff)
	var units = ""
	mTotal, _ := getMemoryUsage()
	adaptedSize, adaptedUnits := filespkg.GetAdaptedSize(int64(mTotal), &units)
	fmt.Printf("   Total used memory*: %.2f %s%s", adaptedSize, adaptedUnits, es())
	//adaptedSizeC, adaptedUnitsC := filespkg.GetAdaptedSize(int64(mCurrent), &units)
	//fmt.Printf("   Current used memory: %.2f %s%s", adaptedSizeC, adaptedUnitsC, es())

}

func es() string {
	switch runtime.GOOS {
	case "windows":
		return "\r\n"
	default:
		return "\n"
	}
}

func prepareData() {
	var c = 0
	for _, f := range *analyzer.FinalFiles {
		if f.Depth <= *cfg.Analyzer.Depth {
			c++
			//break if we up to defined limit
			if isExceedLimit(c, cfg.Printer.Limit) {
				break
			}
			*filesToPrint = append(*filesToPrint, f)
		}
	}
}

func isExceedLimit(checkedValue int, limit *int) bool {
	return checkedValue > *limit && *limit != 0
}

//calculateMaxLenFilename -
func calculateMaxLenFilename() int {
	var maxlen = 0
	for _, f := range *files {
		strlen := utf8.RuneCountInString(f.RelativePath) + 1 + utf8.RuneCountInString(f.Name)
		maxlen = int(math.Max(float64(maxlen), float64(strlen)))
	}
	return maxlen
}

// PrintMemUsage outputs the current, total and OS memory being used. As well as the number
// of garage collection cycles completed.
func getMemoryUsage() (uint64, uint64) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return m.Sys, m.Alloc
}
