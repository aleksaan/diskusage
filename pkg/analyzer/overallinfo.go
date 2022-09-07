package analyzer

import (
	"github.com/aleksaan/diskusage/pkg/models"
)

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

	if !file.IsNotAccessible && !file.IsDir {
		Result.Overall.TotalSize = Result.Overall.TotalSize + file.SizeSubFoldersExcludes
	}

}

// func calcAdaptedSizeInOverallInfo() {
// 	x := ""
// 	Result.Overall.TotalAdaptedSize, Result.Overall.TotalAdaptedUnit = files.GetAdaptedSize(Result.Overall.TotalSize, &x)
// }
