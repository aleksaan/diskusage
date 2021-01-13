package files

import (
	"sort"
	"strings"
)

//TFile - struct for file object
type TFile struct {
	Number                 int     `yaml:"number"`
	RelativePath           string  `yaml:"-"`
	FullPath               string  `yaml:"path"`
	Name                   string  `yaml:"-"`
	Size                   int64   `yaml:"sizeInBytes"`
	IsDir                  bool    `yaml:"isDir"`
	IsLink                 bool    `yaml:"isLink"`
	LinkedDirPath          string  `yaml:"-"`
	Depth                  int     `yaml:"depth"`
	IsNotAccessible        bool    `yaml:"-"`
	IsNotAccessibleMessage string  `yaml:"-"`
	AdaptedSize            float64 `yaml:"adaptedSize"`
	AdaptedUnit            string  `yaml:"adaptedUnit"`
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
