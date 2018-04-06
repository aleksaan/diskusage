package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"sort"
	"strings"
	"time"
)

type file struct {
	path  string
	size  int64
	isdir bool
}

type files []file

type sizeSorter []file

func (a sizeSorter) Len() int           { return len(a) }
func (a sizeSorter) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a sizeSorter) Less(i, j int) bool { return a[i].size < a[j].size }

type sizeDescSorter []file

func (a sizeDescSorter) Len() int           { return len(a) }
func (a sizeDescSorter) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a sizeDescSorter) Less(i, j int) bool { return a[i].size > a[j].size }

//main function
func main() {

	//start timer for calculating total execution time
	start := time.Now()

	fmt.Println("Start scanning")

	//read command line arguments

	//paths
	argpaths := flag.String("path", "", "Folders paths separated by semicolon (required)")

	//limit
	arglimit := flag.Int("limit", 10, "Limit of printed folders (10 by default)")

	//parse argument
	flag.Parse()

	//check for required argument is exists
	if *argpaths == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	//split paths from inline format to array
	var paths = strings.Split(*argpaths, ";")
	fmt.Println("Checking folders:", paths)

	//create
	var files files

	//loop throught folders and get info about size into &files object
	for _, p := range paths {
		getFileInfo(&files, strings.Trim(p, " "))
	}

	fmt.Printf("List of %d max size folders:\n", *arglimit)

	//sort list of folders by descending size
	files.sort("size", "desc")

	//print results
	for i, f := range files {
		//convert unit of size to a comfort for a reporting (b -> Kb, Mb, ...)
		unit, size := f.getComfortSize()
		fmt.Printf("%3d.| DIR: %-100s | SIZE: %.2f %s\n", i+1, f.path, size, unit)

		//break if we up to defined limit
		if i >= *arglimit-1 {
			break
		}
	}

	//print total time
	fmt.Println("Finish scanning")
	elapsed := time.Since(start)
	fmt.Printf("Total time: %s", elapsed)
}

//sorting for files arrays
func (f files) sort(by string, order string) {
	switch {
	case by == "size" && order == "desc":
		sort.Sort(sizeDescSorter(f))
	case by == "size" && order == "":
		sort.Sort(sizeSorter(f))
	default:
	}
}

var unitsOfSize = []string{"b", "Kb", "Mb", "Gb", "Tb", "Pb"}

//get unit correlated to a power of 1024
func getUnitByPowerOf1024(power float64) string {
	p := math.Min(power, float64(len(unitsOfSize)))
	return unitsOfSize[int(p-1)]
}

//get unit correlated to a power of 1024
/* func getSizeByUnit(size float64, unit string) float64 {
	return unitsOfSize[int(p-1)]
} */

//convert size to ident dimension
func (f file) getComfortSize() (string, float64) {
	var power float64 = 1    //=1 by default
	var maxpower float64 = 4 //Mb
	var size = float64(f.size)

	//calculate more comfortable scale
	for power = 2; power <= maxpower; power++ {
		if size < math.Pow(1024, power-1) {
			break //breaks on power which is one more bigger than needed
		}
	}

	//return unit & size in units
	return getUnitByPowerOf1024(power - 1), size / math.Pow(1024, power-2)
}

//calcs dir size
func getFileInfo(files *files, path string) (*file, error) {
	size := int64(0)

	//if file or folder is not accessible then return nil
	dir, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	//if file then return it size
	if !dir.IsDir() {
		size = dir.Size()
		return &file{path: path, size: dir.Size(), isdir: false}, nil
	}

	//read content of folder
	fs, _ := ioutil.ReadDir(path)

	//calc total size throught folder content
	for _, f := range fs {
		cf, err := getFileInfo(files, path+"/"+f.Name())
		if err == nil {
			size = size + cf.size
		}
	}

	//generate object
	f := file{path: path, size: size, isdir: true}

	//append to a global list
	*files = append(*files, f)

	//and return back
	return &f, nil
}
