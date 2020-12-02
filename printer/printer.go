package printer

import (
	"fmt"
	"io"
	"math"
	"os"
	"runtime"
	"unicode/utf8"

	"github.com/aleksaan/diskusage/analyzer"
	"github.com/aleksaan/diskusage/config"
	filespkg "github.com/aleksaan/diskusage/files"
)

const (
	AppTitle   = "https://github.com/aleksaan/diskusage"
	AppVersion = "2.2.0"
	AppAuthor  = "Anufriev Alexander"
	AppYear    = "2020"
)

var writer io.Writer = os.Stdout
var resultsFile *os.File
var cfg *config.Config
var initFlag = false
var files *filespkg.TFiles
var overallInfo *analyzer.TOverallInfo
var filesToPrint = &filespkg.TFiles{}

//Load -
func Load() {
	cfg = config.Cfg
	if *cfg.Printer.ToFile != "" {
		resultsFile = createResultsFile(cfg.Printer.ToFile)
		writer = resultsFile
		//os.Stdout = f
	}

}

//Close -
func Close() {
	resultsFile.Close()
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
	fmt.Fprintf(writer, "About:%s   %s, %s, %s, %s%s", es(), AppTitle, AppVersion, AppAuthor, AppYear, es())
}

func printConfig() {
	fmt.Fprintf(writer, es())
	fmt.Fprintln(writer, "Arguments:")
	fmt.Fprintf(writer, "   %-10s %s%s", "path:", *cfg.Analyzer.Path, es())
	fmt.Fprintf(writer, "   %-10s %d%s", "limit:", *cfg.Printer.Limit, es())
	units := *cfg.Printer.Units
	if units == "" {
		units = "<dynamic>"
	}
	fmt.Fprintf(writer, "   %-10s %s%s", "units:", units, es())
	fmt.Fprintf(writer, "   %-10s %d%s", "depth:", *cfg.Analyzer.Depth, es())
	//fmt.Printf("   %-10s %s%s", "sort:", *cfg.Printer.Sort, es())
	tofile := *cfg.Printer.ToFile
	if tofile == "" {
		tofile = "<no file>"
	}
	fmt.Fprintf(writer, "   %-10s %s\n", "tofile:", tofile)
}

func printFiles() {
	fmt.Fprintf(writer, es())
	fmt.Fprintf(writer, "Results:%s", es())
	maxlen := calculateMaxLenFilename()
	var strfmt = "   %3d.| %-7s %-" + fmt.Sprintf("%d", maxlen+2) + "s | SIZE: %8.2f %-4s | DEPTH: %d %s"
	var dirorfile = "PATH:"
	for i, fi := range *filesToPrint {
		dirorfile = "PATH:"
		if !fi.IsDir {
			dirorfile = "PATH:"
		}
		fmt.Fprintf(writer, strfmt, i+1, dirorfile, fi.RelativePath+fi.Name, fi.AdaptedSize, fi.AdaptedUnit, fi.Depth, es())
	}
}

func printOverall(overallInfo *analyzer.TOverallInfo) {
	fmt.Fprintf(writer, es())
	fmt.Fprintf(writer, "Overall info:%s", es())
	fmt.Fprintf(writer, "   Total time: %s%s", overallInfo.TotalTime, es())
	fmt.Fprintf(writer, "   Total dirs: %d%s", overallInfo.TotalDirs, es())
	fmt.Fprintf(writer, "   Total files: %d%s", overallInfo.TotalFiles, es())
	fmt.Fprintf(writer, "   Total links: %d%s", overallInfo.TotalLinks, es())
	fmt.Fprintf(writer, "   Total size: %.2f %s%s", overallInfo.TotalAdaptedSize, overallInfo.TotalAdaptedUnit, es())
	fmt.Fprintf(writer, "   Total size (bytes): %d%s", overallInfo.TotalSize, es())
	fmt.Fprintf(writer, "   Unaccessible dirs & files: %d%s", overallInfo.TotalNotAccessibleFiles, es())
}

func printSystemReport() {
	fmt.Fprintf(writer, es())
	fmt.Fprintf(writer, "System resources:%s", es())
	var units = ""
	mTotal, _ := getMemoryUsage()
	adaptedSize, adaptedUnits := filespkg.GetAdaptedSize(int64(mTotal), &units)
	fmt.Fprintf(writer, "   Total used memory*: %.2f %s%s", adaptedSize, adaptedUnits, es())

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
	for _, f := range *filesToPrint {
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
