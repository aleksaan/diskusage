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

//входные аргументы
var argpaths *string     //list of folders to analyze separated by ;
var arglimit *int        //limit folders in results
var arglimitdefault = 10 //default value for a arglimit
//var argmaxunit *string   //max possible size unit in a results presetation (b, Kb, Mb, ...). Has a lower priority than "argfixunit". Must be in sizeUnits.
var argfixunit *string //fixed size unit in a results presetation (b, Kb, Mb, ...). Has a upper priority than "argmaxunit". Must be in sizeUnits.

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

//main function
func main() {

	//start timer for calculating total execution time
	start := time.Now()

	//gets command line arguments
	var res = parseCLIArguments()

	//check for required argument is exists
	if !res {
		fmt.Println("Incorrect arguments! Program is finished.")
		os.Exit(1)
	}

	//split paths from inline format to array
	var paths = strings.Split(*argpaths, ";")
	var files files

	fmt.Println("Start scanning")

	//loop throught folders and get info about size into &files object
	for _, p := range paths {
		getFileInfo(&files, strings.Trim(p, " "))
	}

	//sort list of folders by descending size
	files.sort("size", "desc")

	var maxlen = 0
	for i, f := range files {
		maxlen = int(math.Max(float64(maxlen), float64(len(f.path))))
		//break if we up to defined limit
		if i >= *arglimit-1 {
			break
		}
	}
	//print results
	var strfmt = "%3d.| DIR: %-" + fmt.Sprintf("%d", maxlen+5) + "s | SIZE: %.2f %s\n"
	for i, f := range files {
		unit, size := f.getSize(*argfixunit)
		fmt.Printf(strfmt, i+1, f.path, size, unit)

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

//parses input arguments
func parseCLIArguments() bool {
	fmt.Println("Parsing input arguments")

	var res = true

	//path - folders paths separated by semicolon (required)
	argpaths = flag.String("path", "", "Folders paths separated by semicolon (required)")
	//limit - number of folders printed in a results
	arglimit = flag.Int("limit", arglimitdefault, fmt.Sprintf("Number of folders printed in a results (%d by default)", arglimitdefault))
	//maxunit - max unit for a comfort mode representation
	//argmaxunit = flag.String("maxunit", "Gb", "Max possible size unit for a results represetation. One from {b, Kb, Mb, Gb, Tb, Pb}.")
	//fixedunit - unit for a fixed mode representation. If argument is that means fixed mode representation is on, else is off
	argfixunit = flag.String("fixunit", "", "Fixed size unit for a results represetation. Should be one of {b, Kb, Mb, Gb, Tb, Pb}.")

	//parse argument
	flag.Parse()

	//processing arguments
	if len(*argpaths) == 0 {
		fmt.Println("Error! Argument 'path' could not be an empty string")
		res = false
	}

	if *arglimit < 1 {
		fmt.Printf("Argument 'limit' is negative (%d) and has been set to default value (%d)\n", *arglimit, arglimitdefault)
		*arglimit = arglimitdefault //set to default value
	}

	//if val, ok := sizeUnits[*argmaxunit]; ok && len(*argmaxunit) > 0 {
	//	fmt.Println("Argument 'maxunit' is not in allowable range {b, Kb, Mb, Gb, Tb, Pb}")
	//	res = false
	//}

	if _, ok := sizeUnits[*argfixunit]; !ok && len(*argfixunit) > 0 {
		fmt.Println("Error! Argument 'fixunit' is not in allowable range {b, Kb, Mb, Gb, Tb, Pb}")
		res = false
	}

	if res && len(*argfixunit) > 0 {
		fmt.Printf("Results will be represented with fixed units style in '%s'\n", *argfixunit)
	}

	//prints arguments
	if res {
		fmt.Println("Input arguments:")
		fmt.Printf("   path: %s\n", *argpaths)
		fmt.Printf("   limit: %d\n", *arglimit)
		fmt.Printf("   fixunit: %s\n", *argfixunit)
	}

	return res
}

//sorts files array
func (f files) sort(by string, order string) {
	switch {
	case by == "size" && order == "desc":
		sort.Sort(sizeDescSorter(f))
	case by == "size" && order == "":
		sort.Sort(sizeSorter(f))
	default:
	}
}

//converts size to ident dimension
func (f file) getSize(fixunit string) (string, float64) {

	var size = float64(f.size)
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
	return unit, size / math.Pow(1024, power-1)
}

//calculates dir size
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
