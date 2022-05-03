package analyzer

import (
	"time"

	"github.com/aleksaan/diskusage/pkg/models"
)

var startTime time.Time

func addToOverallInfo(file *models.TFile) {

	// if file.Depth == 1 {
	// 	Result.Overall.TotalSize += file.Size
	// }

	if file.IsNotAccessible {
		Result.Overall.TotalNotAccessibleFiles++
	}

	if file.IsDir {
		Result.Overall.TotalDirs++
	} else {
		Result.Overall.TotalFiles++
	}

	if file.IsLink {
		Result.Overall.TotalLinks++
	}

}

// func calcAdaptedSizeInOverallInfo() {
// 	x := ""
// 	Result.Overall.TotalAdaptedSize, Result.Overall.TotalAdaptedUnit = files.GetAdaptedSize(Result.Overall.TotalSize, &x)
// }
