package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/aleksaan/diskusage/diskusage"
)

//-----------------------------------------------------------------------------------------
//main function
func main() {

	//start timer
	start := time.Now()

	//gets command line arguments
	parseCLIArguments()
	diskusage.InputArgs.PrintArguments()

	//start scanning files
	fmt.Printf("Start scanning\n")

	//get files
	files := &diskusage.TFiles{}
	diskusage.ScanDir(files, diskusage.InputArgs.Paths, 1)

	//sort files by size
	files.Sort("size_name", "desc")

	//print files results to console
	files.PrintFilesSizes()

	//finish work and calculate elapsed time
	fmt.Printf("Finish scanning\n")
	elapsed := time.Since(start)
	fmt.Printf("Total time: %s", elapsed)
}

//-----------------------------------------------------------------------------------------

//ParseCLIArguments - parses input arguments
func parseCLIArguments() {

	var argpaths = flag.String("path", "", "Folders paths separated by semicolon (required)")
	var arglimit = flag.Int("limit", diskusage.LimitDefault, fmt.Sprintf("Number of folders printed in a results (%d by default)", diskusage.LimitDefault))
	var argfixunit = flag.String("fixunit", "", "Fixed size unit for a results represetation. Should be one of {b, Kb, Mb, Gb, Tb, Pb}.")
	var argdepth = flag.Int("depth", diskusage.DepthDefault, "Depth of subfolders.")

	//parse argument
	flag.Parse()

	t := &diskusage.InputArgs

	//processing arguments
	t.SetPaths(argpaths)
	//checkError(err)

	t.SetLimit(arglimit)
	//checkError(err)

	t.SetFixUnit(argfixunit)
	//checkError(err)

	t.SetDepth(argdepth)
	//checkError(err)
}
