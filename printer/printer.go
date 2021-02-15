package printer

import (
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"runtime"
	"unicode/utf8"

	"github.com/aleksaan/diskusage/analyzer"
	"github.com/aleksaan/diskusage/config"
	filespkg "github.com/aleksaan/diskusage/files"
	"gopkg.in/yaml.v3"
)

const (
	AppTitle   = "https://github.com/aleksaan/diskusage"
	AppVersion = "2.4.0"
	AppAuthor  = "Anufriev Alexander"
	AppYear    = "2021"
)

var writerToText io.Writer = os.Stdout
var writerToYaml io.Writer = nil
var resultsTextFile *os.File
var resultsYamlFile *os.File
var cfg *config.Config
var initFlag = false
var files *filespkg.TFiles
var overallInfo *analyzer.TOverallInfo
var filesToPrint = &filespkg.TFiles{}

//Load -
func Load() {
	cfg = config.Cfg
	if *cfg.Printer.ToTextFile != "" {
		resultsTextFile = createTextResultsFile(cfg.Printer.ToTextFile)
		writerToText = resultsTextFile
		//os.Stdout = f
	}

	if *cfg.Printer.ToYamlFile != "" {
		resultsYamlFile = createYamlResultsFile(cfg.Printer.ToYamlFile)
		writerToYaml = resultsYamlFile
		//os.Stdout = f
	}

}

//Close -
func Close() {
	resultsTextFile.Close()
	resultsYamlFile.Close()
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
	writeToYamlFile()
}

func createTextResultsFile(filename *string) *os.File {
	// open output file
	f, err := os.Create(*filename)
	if err != nil {
		panic(err)
	}
	return f
}

func createYamlResultsFile(filename *string) *os.File {
	// open output file
	f, err := os.Create(*filename)
	if err != nil {
		panic(err)
	}
	return f
}

//PrintAbout -
func PrintAbout() {
	fmt.Fprintf(writerToText, "About:%s   %s, %s, %s, %s%s", es(), AppTitle, AppVersion, AppAuthor, AppYear, es())
}

func printConfig() {
	fmt.Fprintf(writerToText, es())
	fmt.Fprintln(writerToText, "Arguments:")
	fmt.Fprintf(writerToText, "   %-10s %s%s", "path:", *cfg.Analyzer.Path, es())
	fmt.Fprintf(writerToText, "   %-10s %d%s", "limit:", *cfg.Printer.Limit, es())
	units := *cfg.Printer.Units
	// if units == "" {
	// 	units = "<dynamic>"
	// }
	fmt.Fprintf(writerToText, "   %-10s %s%s", "units:", units, es())
	fmt.Fprintf(writerToText, "   %-10s %d%s", "depth:", *cfg.Analyzer.Depth, es())
	fmt.Fprintf(writerToText, "   %-10s %s%s", "printonly:", *cfg.Printer.PrintOnly, es())
	//fmt.Printf("   %-10s %s%s", "sort:", *cfg.Printer.Sort, es())
	toTextFile := *cfg.Printer.ToTextFile
	// if toTextFile == "" {
	// 	toTextFile = "<no text file>"
	// }
	fmt.Fprintf(writerToText, "   %-10s %s\n", "toTextFile:", toTextFile)
	toYamlFile := *cfg.Printer.ToYamlFile
	// if toYamlFile == "" {
	// 	toYamlFile = "<no yaml file>"
	// }
	fmt.Fprintf(writerToText, "   %-10s %s\n", "toYamlFile:", toYamlFile)
}

func printFiles() {
	fmt.Fprintf(writerToText, es())
	fmt.Fprintf(writerToText, "Results:%s", es())
	maxlen := calculateMaxLenFilename()
	var strfmt = "   %3d.| %-7s %-" + fmt.Sprintf("%d", maxlen+2) + "s | SIZE: %8.2f %-4s | DEPTH: %d %s"
	var dirorfile = "PATH:"
	for i, fi := range *filesToPrint {
		dirorfile = "PATH:"
		if !fi.IsDir {
			dirorfile = "PATH:"
		}
		fmt.Fprintf(writerToText, strfmt, i+1, dirorfile, fi.RelativePath+fi.Name, fi.AdaptedSize, fi.AdaptedUnit, fi.Depth, es())
	}
}

func printOverall(overallInfo *analyzer.TOverallInfo) {
	fmt.Fprintf(writerToText, es())
	fmt.Fprintf(writerToText, "Overall info:%s", es())
	fmt.Fprintf(writerToText, "   Total time: %s%s", overallInfo.TotalTime, es())
	fmt.Fprintf(writerToText, "   Total dirs: %d%s", overallInfo.TotalDirs, es())
	fmt.Fprintf(writerToText, "   Total files: %d%s", overallInfo.TotalFiles, es())
	fmt.Fprintf(writerToText, "   Total links: %d%s", overallInfo.TotalLinks, es())
	fmt.Fprintf(writerToText, "   Total size: %.2f %s%s", overallInfo.TotalAdaptedSize, overallInfo.TotalAdaptedUnit, es())
	fmt.Fprintf(writerToText, "   Total size (bytes): %d%s", overallInfo.TotalSize, es())
	fmt.Fprintf(writerToText, "   Unaccessible dirs & files: %d%s", overallInfo.TotalNotAccessibleFiles, es())
}

func printSystemReport() {
	fmt.Fprintf(writerToText, es())
	fmt.Fprintf(writerToText, "System resources:%s", es())
	var units = ""
	mTotal, _ := getMemoryUsage()
	adaptedSize, adaptedUnits := filespkg.GetAdaptedSize(int64(mTotal), &units)
	fmt.Fprintf(writerToText, "   Total used memory*: %.2f %s%s", adaptedSize, adaptedUnits, es())

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
	var isDir = true
	if *cfg.Printer.PrintOnly == "files" {
		isDir = false
	}

	//are files & folders both we need (or not)
	var isAll = *cfg.Printer.PrintOnly == "folders&files"

	for _, f := range *analyzer.FinalFiles {
		if f.Depth <= *cfg.Analyzer.Depth && (f.IsDir == isDir || isAll) {
			c++
			//break if we up to defined limit
			if isExceedLimit(c, cfg.Printer.Limit) {
				break
			}
			f.FullPath = *cfg.Analyzer.Path + "\\" + f.RelativePath + f.Name
			f.Number = c
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

func writeToYamlFile() error {
	var err error
	if writerToYaml != nil {
		d, err := yaml.Marshal(filesToPrint)

		if err != nil {
			log.Fatalf("error: %v", err)
		}

		fmt.Fprintf(writerToYaml, "%s", string(d))
	}
	return err
}
