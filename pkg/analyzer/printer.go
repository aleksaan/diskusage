package analyzer

import (
	"encoding/json"
	"fmt"

	"github.com/aleksaan/diskusage/pkg/models"
	"github.com/aleksaan/diskusage/pkg/printer"
)

func WriteJSONToConsole() {
	r, err := json.Marshal(Result)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("%s", string(r))
}

func WriteHumanReadableToConsole() {
	printConfig(cfg)
	printer.PrintFiles(&Result.Files, "")
	printOverall(&Result.Overall)
}

func printConfig(cfg *models.TAnalyserConfig) {
	fmt.Println("-------------------\nArguments:")
	fmt.Printf("   %-10s %s\n", "path:", cfg.Path)
	fmt.Printf("   %-10s %d\n", "depth:", cfg.Depth)
	fmt.Printf("   %-10s %v\n", "hr:", cfg.Hr)
}

func printOverall(overallInfo *models.TOverallInfo) {
	fmt.Printf("-------------------\nOverall info:\n")
	fmt.Printf("   Total time: %s\n", overallInfo.TotalTime)
	fmt.Printf("   Total dirs: %d\n", overallInfo.TotalDirs)
	fmt.Printf("   Total files: %d\n", overallInfo.TotalFiles)
	fmt.Printf("   Total links: %d\n", overallInfo.TotalLinks)
	fmt.Printf("   Total size: %.2f %s\n", overallInfo.TotalAdaptedSize, overallInfo.TotalAdaptedUnit)
	fmt.Printf("   Total size (bytes): %d\n", overallInfo.TotalSize)
	fmt.Printf("   Unaccessible dirs & files: %d\n", overallInfo.TotalNotAccessibleFiles)
	fmt.Printf("-------------------\n")
}

// func printSystemReport() {
// 	fmt.Println("System resources:")
// 	var units = ""
// 	mTotal, _ := getMemoryUsage()
// 	adaptedSize, adaptedUnits := files.GetAdaptedSize(int64(mTotal), &units)
// 	fmt.Printf("   Total used memory*: %.2f %s%s", adaptedSize, adaptedUnits)
// }

// PrintMemUsage outputs the current, total and OS memory being used. As well as the number
// of garage collection cycles completed.
// func getMemoryUsage() (uint64, uint64) {
// 	var m runtime.MemStats
// 	runtime.ReadMemStats(&m)
// 	return m.Sys, m.Alloc
// }
