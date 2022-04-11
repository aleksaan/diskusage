package models

import (
	"sort"
	"strings"
	"time"

	"github.com/aleksaan/diskusage/config"
)

type TResult struct {
	About   TAbout       `json:"about"`
	Files   TFiles       `json:"results"`
	Overall TOverallInfo `json:"totals"`
}

type TAbout struct {
	AppTitle   string `json:"title"`
	AppVersion string `json:"version"`
	AppAuthor  string `json:"author"`
	AppYear    string `json:"year"`
}

//TFile - struct for file object
type TFile struct {
	Number                 int     `json:"number"`
	RelativePath           string  `json:"-"`
	FullPath               string  `json:"path"`
	Name                   string  `json:"-"`
	Size                   int64   `json:"sizeInBytes"`
	IsDir                  bool    `json:"isDir"`
	IsLink                 bool    `json:"isLink"`
	LinkedDirPath          string  `json:"-"`
	Depth                  int     `json:"depth"`
	IsNotAccessible        bool    `json:"-"`
	IsNotAccessibleMessage string  `json:"-"`
	AdaptedSize            float64 `json:"adaptedSize"`
	AdaptedUnit            string  `json:"adaptedUnit"`
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
func (files *TFiles) Sort() {
	switch {
	case config.Cfg.Sort == "name_asc":
		sort.Sort(nameSorter(*files))
	default:
		sort.Sort(sizeAndNameSorter(*files))
	}
}
