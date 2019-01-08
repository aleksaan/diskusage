package diskusage

import (
	"io/ioutil"
	"math"
	"os"
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

var sortedKeysSizeUnits = []string{"b", "Kb", "Mb", "Gb", "Tb", "Pb"}

//-----------------------------------------------------------------------------------------

//ScanDir - scan directory and return its size
func ScanDir(files *TFiles, path string, depth int) int64 {
	//read content of folder
	osfiles, _ := ioutil.ReadDir(path)

	var dirsize int64

	//calc total size throught folder content
	for _, osfile := range osfiles {
		file := ScanFile(path+"/"+osfile.Name(), depth)
		if file.IsDir {
			file.Size = ScanDir(files, file.Path, depth+1)
		}
		file.SetAdaptedSizeOfFile(&InputArgs)
		dirsize += file.Size
		*files = append(*files, *file)
	}

	return dirsize
}

//-----------------------------------------------------------------------------------------

//ScanFile - scan dir/file parameters
func ScanFile(path string, depth int) *TFile {
	f := &TFile{}
	f.Path = path
	f.Depth = depth

	//if file or folder is not accessible then return nil
	dir, err := os.Stat(path)
	if err != nil {
		f.IsNotAccessible = true
		f.IsNotAccessibleMessage = err.Error()
		return f
	}

	f.IsDir = dir.IsDir()

	if !dir.IsDir() {
		f.Size = dir.Size()
	}

	return f
}

//SetAdaptedSizeOfFile - set file size adapted to InputArgs.FixUnit units or to a flexible useful units
func (f *TFile) SetAdaptedSizeOfFile(inputargs *TInputArgs) {

	var size = float64(f.Size)
	var unit string
	var power float64

	if len(inputargs.FixUnit) > 0 {
		unit = inputargs.FixUnit
		power = sizeUnits[inputargs.FixUnit]
	} else {
		for _, unit = range sortedKeysSizeUnits {
			power = sizeUnits[unit]
			if size < math.Pow(1024, power) {
				break
			}
		}
	}
	f.AdaptedUnit = unit
	f.AdaptedSize = size / math.Pow(1024, power-1)
}
