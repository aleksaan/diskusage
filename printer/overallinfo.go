package printer

import (
	"time"
)

//OverallInfo -
type TOverallInfo struct {
	totalTime               time.Duration
	totalDirs               int64
	totalFiles              int64
	totalLinks              int64
	totalSize               int64
	totalAdaptedSize        float64
	totalAdaptedUnit        string
	totalNotAccessibleFiles int64
}
