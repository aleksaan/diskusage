package diskusage

import (
	"fmt"
	"math"
	"sort"
)

//TFile - struct for file object
type TFile struct {
	Path                   string
	Size                   int64
	IsDir                  bool
	Depth                  int
	IsNotAccessible        bool
	IsNotAccessibleMessage string
	AdaptedSize            float64
	AdaptedUnit            string
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
	return a[i].Size > a[j].Size || (a[i].Size == a[j].Size && a[i].Path < a[j].Path)
}

//-----------------------------------------------------------------------------------------

//Sort - sort files by size
func (files *TFiles) Sort(by string, order string) {
	switch {
	case by == "size" && order == "desc":
		sort.Sort(sizeDescSorter(*files))
	case by == "size_name" && order == "desc":
		sort.Sort(sizeAndNameSorter(*files))
	case by == "size" && order == "":
		sort.Sort(sizeSorter(*files))
	default:
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
			maxlen = int(math.Max(float64(maxlen), float64(len(f.Path))))
			//break if we up to defined limit
			if c+1 >= InputArgs.Limit-1 {
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
	var strfmt = "%3d.| %-5s %-" + fmt.Sprintf("%d", maxlen+2) + "s | SIZE: %6.2f %-4s | DEPTH: %d \n"
	var isnotaccessiblemessage = "%3d.| %s: %-" + fmt.Sprintf("%d", maxlen+2) + "s | SIZE: %6.2f %-4s | DEPTH: %d \n"
	var c = 0
	for _, f := range *files {
		if f.Depth <= InputArgs.Depth {
			c++
			if f.IsNotAccessible {
				fmt.Printf(isnotaccessiblemessage, c, "UNKNOWN", f.Path, f.AdaptedSize, f.AdaptedUnit, f.Depth)
			} else {
				dirorfile := "DIR:"
				if !f.IsDir {
					dirorfile = "FILE:"
				}
				fmt.Printf(strfmt, c, dirorfile, f.Path, f.AdaptedSize, f.AdaptedUnit, f.Depth)
			}

			//break if we up to defined limit
			if c+1 >= InputArgs.Limit-1 {
				break
			}
		}
	}
}
