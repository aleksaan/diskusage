package analyzer

import (
	"time"

	"github.com/aleksaan/diskusage/files"
)

//TOverallInfo -
type TOverallInfo struct {
	TotalTime               time.Duration
	TotalDirs               int64
	TotalFiles              int64
	TotalLinks              int64
	TotalSize               int64
	TotalAdaptedSize        float64
	TotalAdaptedUnit        string
	TotalNotAccessibleFiles int64
}

var startTime time.Time

//OverallInfo - summary information about entire proccess
var OverallInfo = &TOverallInfo{}

func addToOverallInfo(file *files.TFile) {

	if file.Depth == 1 {
		OverallInfo.TotalSize += file.Size
	}

	if file.IsNotAccessible {
		OverallInfo.TotalNotAccessibleFiles++
	}

	if file.IsDir {
		OverallInfo.TotalDirs++
	} else {
		OverallInfo.TotalFiles++
	}

	if file.IsLink {
		OverallInfo.TotalLinks++
	}

}

func calcAdaptedSizeInOverallInfo() {
	x := ""
	OverallInfo.TotalAdaptedSize, OverallInfo.TotalAdaptedUnit = files.GetAdaptedSize(OverallInfo.TotalSize, &x)
}
