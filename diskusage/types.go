package diskusage

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"time"
)

//TFile - struct for file object
type TFile struct {
	RelativePath           string
	Name                   string
	Size                   int64
	IsDir                  bool
	IsLink                 bool
	LinkedDirPath          string
	Depth                  int
	IsNotAccessible        bool
	IsNotAccessibleMessage string
	AdaptedSize            float64
	AdaptedUnit            string
}

//OverallInfo -
type OverallInfo struct {
	totalTime               time.Duration
	totalDirs               int64
	totalFiles              int64
	totalLinks              int64
	totalSize               int64
	totalAdaptedSize        float64
	totalAdaptedUnit        string
	totalNotAccessibleFiles int64
}

//TFiles - struct for files array object
type TFiles []TFile

type sizeSorter []TFile

func (a sizeSorter) Len() int           { return len(a) }
func (a sizeSorter) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a sizeSorter) Less(i, j int) bool { return a[i].Size < a[j].Size }

type sizeDescSorter []TFile

func (a sizeDescSorter) Len() int           { return len(a) }
func (a sizeDescSorter) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a sizeDescSorter) Less(i, j int) bool { return a[i].Size > a[j].Size }

type sizeAndNameSorter []TFile

func (a sizeAndNameSorter) Len() int      { return len(a) }
func (a sizeAndNameSorter) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a sizeAndNameSorter) Less(i, j int) bool {
	return a[i].Size > a[j].Size || (a[i].Size == a[j].Size && (a[i].RelativePath+a[i].Name < a[j].RelativePath+a[j].Name))
}

type nameSorter []TFile

func (a nameSorter) Len() int      { return len(a) }
func (a nameSorter) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a nameSorter) Less(i, j int) bool {
	switch {
	case (a[i].IsDir == false && a[j].IsDir == true):
		return false
	case (a[i].IsDir == true && a[j].IsDir == false):
		return true
	default:
		return a[i].RelativePath+strings.ToLower(a[i].Name) < a[j].RelativePath+strings.ToLower(a[j].Name)
	}
}

//-----------------------------------------------------------------------------------------

//Sort - sort files by size
func (files *TFiles) Sort(by string) {
	switch {
	case by == "name_asc":
		sort.Sort(nameSorter(*files))
	default:
		sort.Sort(sizeAndNameSorter(*files))
	}
}

//-----------------------------------------------------------------------------------------

//CalculateMaxLenFilename - calculating max length filename in files
func (files *TFiles) CalculateMaxLenFilename() int {
	var maxlen = 0
	var c = 0
	for _, f := range *files {
		if f.Depth <= InputArgs.Depth {
			c++
			maxlen = int(math.Max(float64(maxlen), float64(len(f.RelativePath)+1+len(f.Name))))
			//break if we up to defined limit
			if isExceedLimit(c + 1) {
				break
			}
		}
	}
	return maxlen
}

//-----------------------------------------------------------------------------------------

//PrintFilesSizes - print out sizes of files
func (files *TFiles) PrintFilesSizes() {
	maxlen := files.CalculateMaxLenFilename()

	//print results
	var strfmt = "%3d.| %-7s %-" + fmt.Sprintf("%d", maxlen+2) + "s | SIZE: %8.2f %-4s | DEPTH: %d %s"
	var c = 0
	for _, f := range *files {
		if f.Depth <= InputArgs.Depth && !f.IsNotAccessible {
			c++
			dirorfile := "PATH:"
			fmt.Printf(strfmt, c, dirorfile, f.RelativePath+f.Name, f.AdaptedSize, f.AdaptedUnit, f.Depth, es())

			//break if we up to defined limit in case limit > 0
			if isExceedLimit(c + 1) {
				break
			}
		}
	}
}

//return True if limit has exceeded
func isExceedLimit(x int) bool {
	return x > InputArgs.Limit && InputArgs.Limit != 0
}

//PrintOverallInfo -
func (info *OverallInfo) PrintOverallInfo() {
	fmt.Printf("\nOverall info:\n")
	fmt.Printf("   Total time: %s\n", info.totalTime)
	fmt.Printf("   Total dirs: %d\n", info.totalDirs)
	fmt.Printf("   Total files: %d\n", info.totalFiles)
	fmt.Printf("   Total links: %d\n", info.totalLinks)
	fmt.Printf("   Total size: %.2f %s\n", info.totalAdaptedSize, info.totalAdaptedUnit)
	fmt.Printf("   Total size (bytes): %d\n", info.totalSize)
	fmt.Printf("   Unaccessible dirs & files: %d\n", info.totalNotAccessibleFiles)
}

//GetOverallInfo -
func (files *TFiles) GetOverallInfo(totalTime time.Duration) *OverallInfo {

	r := &OverallInfo{}

	r.totalTime = totalTime

	for _, file := range *files {
		if file.Depth == 1 {
			r.totalSize += file.Size
		}
		if file.IsNotAccessible {
			r.totalNotAccessibleFiles++
		}
		if file.IsDir {
			r.totalDirs++
		} else {
			r.totalFiles++
		}

		if file.IsLink {
			r.totalLinks++
		}
	}

	r.totalAdaptedSize, r.totalAdaptedUnit = GetAdaptedSize(r.totalSize, "")

	return r
}
