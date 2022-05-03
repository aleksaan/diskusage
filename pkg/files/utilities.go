package files

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
	"runtime"
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

//GetAdaptedSize - get file size adapted to InputArgs.FixUnit units or to a flexible useful units
func GetAdaptedSize(sizeB int64, units *string) (float64, string) {

	var size = float64(sizeB)
	var unit string
	var power float64

	if len(*units) > 0 {
		unit = *units
		power = sizeUnits[*units]
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

//CreateFile - creates output file with <filename>
func CreateFile(filename *string) *os.File {
	// open output file
	f, err := os.Create(*filename)
	if err != nil {
		panic(err)
	}
	return f
}

//CheckFileIsExist - check file is exist on the disk
func CheckFileIsExist(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}

//PrintEndOfLine - printing end of line symbol
func PrintEndOfLine(f *os.File) {
	fmt.Fprintf(f, "%s", es())
}

//es - produces EOF symbol
func es() string {
	switch runtime.GOOS {
	case "windows":
		return "\r\n"
	default:
		return "\n"
	}
}
