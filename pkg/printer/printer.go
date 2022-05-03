package printer

import (
	"fmt"
	"math"
	"unicode/utf8"

	"github.com/aleksaan/diskusage/pkg/files"
	"github.com/aleksaan/diskusage/pkg/models"
)

func PrintFiles(ff *models.TFiles, units string) {
	fmt.Printf("-------------------\nResults:\n")
	maxlen := calculateMaxLenFilename(ff)
	var strfmt = "   %3d.| %-7s %-" + fmt.Sprintf("%d", maxlen+2) + "s | FULL SIZE: %8.2f %-4s | CLEAN SIZE: %8.2f %-4s | DEPTH: %d\n"
	var dirorfile = "PATH:"
	for i, fi := range *ff {
		fsizeIncl, funitIncl := files.GetAdaptedSize(fi.SizeSubFoldersIncludes, &units)
		fsizeExl, funitExl := files.GetAdaptedSize(fi.SizeSubFoldersExcludes, &units)
		fmt.Printf(strfmt, i+1, dirorfile, fi.RelativePath, fsizeIncl, funitIncl, fsizeExl, funitExl, fi.Depth)
		if i > 48 {
			fmt.Printf("%3s\n", "...")
			fi = (*ff)[len(*ff)-1]
			fsizeIncl, funitIncl := files.GetAdaptedSize(fi.SizeSubFoldersIncludes, &units)
			fsizeExl, funitExl := files.GetAdaptedSize(fi.SizeSubFoldersExcludes, &units)
			fmt.Printf(strfmt, len(*ff), dirorfile, fi.RelativePath, fsizeIncl, funitIncl, fsizeExl, funitExl, fi.Depth)
			break
		}
	}
}

//calculateMaxLenFilename -
func calculateMaxLenFilename(ff *models.TFiles) int {
	var maxlen = 0
	for i, f := range *ff {
		if i == len(*ff)-1 {
			_ = maxlen
		}
		if (i < 50) || (i == len(*ff)-1) {
			strlen := utf8.RuneCountInString(f.RelativePath)
			//+ 1 + utf8.RuneCountInString(f.Name)
			maxlen = int(math.Max(float64(maxlen), float64(strlen)))
		}
	}
	return maxlen
}
