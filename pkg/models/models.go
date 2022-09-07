package models

import (
	"sort"
	"time"
)

type TAnalyserResult struct {
	About   TAbout          `json:"about"`
	Config  TAnalyserConfig `json:"config"`
	Files   TFiles          `json:"results"`
	Overall TOverallInfo    `json:"totals"`
}

type TFilterResult struct {
	About  TAbout        `json:"about"`
	Config TFilterConfig `json:"config"`
	Files  TFiles        `json:"results"`
}

type TAbout struct {
	AppTitle   string `json:"title"`
	AppVersion string `json:"version"`
	AppAuthor  string `json:"author"`
	AppYear    string `json:"year"`
}

//TFile - struct for file object
type TFile struct {
	Number                 int    `json:"number"`
	RelativePath           string `json:"path"`
	SizeSubFoldersIncludes int64  `json:"sizeSubFoldesIncludes"`
	SizeSubFoldersExcludes int64  `json:"sizeSubFoldesExcludes"`
	IsDir                  bool   `json:"isDir"`
	IsLink                 bool   `json:"isLink"`
	Depth                  uint   `json:"depth"`
	IsNotAccessible        bool   `json:"isNotAccessible"`
	IsNotAccessibleMessage string `json:"isNotAccessibleMessage"`
}

//Config - utility configuration
type TAnalyserConfig struct {
	Path   string `json:"path"`
	Depth  uint   `json:"depth"`
	Hr     bool   `json:"hr"`
	HrRows uint   `json:"hrrows"`
}

//Config - utility configuration
type TFilterConfig struct {
	Hr     bool   `json:"hr"`
	Depth  uint   `json:"depth"`
	Top    uint   `json:"limit"`
	Filter string `json:"filter"`
	Size   string `json:"size"`
	Path   string `json:"path"`
}

//TOverallInfo -
type TOverallInfo struct {
	TotalTime               time.Duration `json:"totalTime"`
	TotalDirs               int64         `json:"totalDirs"`
	TotalFiles              int64         `json:"totalFiles"`
	TotalLinks              int64         `json:"totalLinks"`
	TotalSize               int64         `json:"totalSize"`
	TotalAdaptedSize        float64       `json:"totalAdaptedSize"`
	TotalAdaptedUnit        string        `json:"totalAdaptedUnit"`
	TotalNotAccessibleFiles int64         `json:"totalNotAccessibleFiles"`
}

//TFiles - struct for files array object
type TFiles []TFile

type CumulativeSizeSorter []TFile

func (a CumulativeSizeSorter) Len() int      { return len(a) }
func (a CumulativeSizeSorter) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a CumulativeSizeSorter) Less(i, j int) bool {
	return a[i].SizeSubFoldersIncludes > a[j].SizeSubFoldersIncludes ||
		(a[i].SizeSubFoldersIncludes == a[j].SizeSubFoldersIncludes && (a[i].RelativePath < a[j].RelativePath))
}

type PlainSizeSorter []TFile

func (a PlainSizeSorter) Len() int      { return len(a) }
func (a PlainSizeSorter) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a PlainSizeSorter) Less(i, j int) bool {
	return a[i].SizeSubFoldersExcludes > a[j].SizeSubFoldersExcludes ||
		(a[i].SizeSubFoldersExcludes == a[j].SizeSubFoldersExcludes && (a[i].RelativePath < a[j].RelativePath))
}

// //-----------------------------------------------------------------------------------------

//Sort - sort files by size
func (files *TFiles) Sort(t string) {
	switch {
	case t == "f":
		sort.Sort(CumulativeSizeSorter(*files))
	case t == "c":
		sort.Sort(PlainSizeSorter(*files))
	}
}
