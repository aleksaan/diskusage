package printer

import (
	"math"
	"time"

	"github.com/aleksaan/diskusage/analyzer"
	"github.com/aleksaan/diskusage/config"
	"github.com/aleksaan/diskusage/files"
)

var preparedFiles *files.TFiles

var overallInfo = &TOverallInfo{}

func prepareData(cfg *config.Config, files *files.TFiles) {
	var c = 0
	for _, f := range *files {
		if f.Depth <= *cfg.Analyzer.Depth {
			c++
			//break if we up to defined limit
			if isExceedLimit(c+1, cfg.Printer.Limit) {
				break
			}
			*preparedFiles = append(*preparedFiles, f)
		}
	}
}

func isExceedLimit(checkedValue int, limit *int) bool {
	return checkedValue > *limit && *limit != 0
}

func calculateMaxLenFilename() int {
	var maxlen = 0
	for _, f := range *preparedFiles {
		maxlen = int(math.Max(float64(maxlen), float64(len(f.RelativePath)+1+len(f.Name))))
	}
	return maxlen
}

func prepareOverallInfo(files *files.TFiles, totalTime time.Duration) {

	overallInfo.totalTime = totalTime

	for _, file := range *files {
		if file.Depth == 1 {
			overallInfo.totalSize += file.Size
		}
		if file.IsNotAccessible {
			overallInfo.totalNotAccessibleFiles++
		}
		if file.IsDir {
			overallInfo.totalDirs++
		} else {
			overallInfo.totalFiles++
		}

		if file.IsLink {
			overallInfo.totalLinks++
		}
	}

	x := ""
	overallInfo.totalAdaptedSize, overallInfo.totalAdaptedUnit = analyzer.GetAdaptedSize(overallInfo.totalSize, &x)

}
