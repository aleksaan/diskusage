package diskusage

import (
	"io/ioutil"
	"math"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

//pairs of key and scale power x, when 1024^x is scale of size
var sizeUnits = map[string]float64{
	"b":  1,
	"Kb": 2,
	"Mb": 3,
	"Gb": 4,
	"Tb": 5,
	"Pb": 6,
}

var sortValues = map[string]float64{
	"name_asc":  1,
	"size_desc": 2,
}

var sortedKeysSizeUnits = []string{"b", "Kb", "Mb", "Gb", "Tb", "Pb"}

//-----------------------------------------------------------------------------------------

//ScanDir - scan directory and return its size
func ScanDir(files *TFiles, path string, depth int) int64 {
	//read content of folder
	osfiles, _ := ioutil.ReadDir(path)

	var dirsize int64

	//calc total size throught folder content
	for _, osfile := range osfiles {

		file := ScanFile(path, osfile.Name(), depth)
		if file.IsDir {
			file.Size = ScanDir(files, AddPathSeparator(path+osfile.Name()), depth+1)
		}

		file.SetAdaptedSizeOfFile(&InputArgs)
		dirsize += file.Size
		*files = append(*files, *file)
	}

	return dirsize
}

//SetAdaptedSizeOfFile - set file properties: AdaptedSize & AdaptedUnit
func (file *TFile) SetAdaptedSizeOfFile(inputArgs *TInputArgs) {
	file.AdaptedSize, file.AdaptedUnit = GetAdaptedSize(file.Size, inputArgs.FixUnit)
}

//-----------------------------------------------------------------------------------------

//ScanFile - scan dir/file parameters
func ScanFile(path string, name string, depth int) *TFile {
	f := &TFile{}
	f.Name = name
	f.RelativePath = path[len(InputArgs.Path):]
	f.Depth = depth

	//if file or folder is not accessible then return nil
	pathName := CleanPath(&path, false) + name

	//dirstat, _ := os.Stat(pathName)
	dir, err := os.Lstat(pathName)

	if err != nil {
		f.IsNotAccessible = true
		f.IsNotAccessibleMessage = err.Error()
		return f
	}

	linkdir, linkerr := filepath.EvalSymlinks(pathName)
	if linkerr != nil {
		f.IsNotAccessible = true
		f.IsNotAccessibleMessage = linkerr.Error()
		return f
	}

	if strings.ToLower(linkdir) != strings.ToLower(pathName) {
		f.IsLink = true
		f.LinkedDirPath = linkdir
	}

	f.IsDir = dir.IsDir()

	if !dir.IsDir() {
		f.Size = dir.Size()
	}

	return f
}

//GetAdaptedSize - get file size adapted to InputArgs.FixUnit units or to a flexible useful units
func GetAdaptedSize(sizeB int64, fixunit string) (float64, string) {

	var size = float64(sizeB)
	var unit string
	var power float64

	if len(fixunit) > 0 {
		unit = fixunit
		power = sizeUnits[fixunit]
	} else {
		for _, unit = range sortedKeysSizeUnits {
			power = sizeUnits[unit]
			if size < math.Pow(1024, power) {
				break
			}
		}
	}
	return (size / math.Pow(1024, power-1)), unit
}

//CleanPath - get absolute path like C:\temp\
func CleanPath(path *string, isrelativeclean bool) string {
	if isrelativeclean {
		abspath, _ := filepath.Abs(*path)
		return AddPathSeparator(filepath.Clean(abspath))
	}

	return AddPathSeparator(filepath.Clean(*path))
}

//AddPathSeparator - add os path separator to string
func AddPathSeparator(path string) string {
	cleanPath := filepath.Clean(path)
	lastSymbol := cleanPath[(len(cleanPath) - 1):]
	if lastSymbol == string(os.PathSeparator) {
		return filepath.Clean(path)
	}
	return filepath.Clean(path) + string(os.PathSeparator)
}

//end line symbol
func es() string {
	switch runtime.GOOS {
	case "windows":
		return "\r\n"
	default:
		return "\n"
	}
}
