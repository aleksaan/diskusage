package analyzer

import (
	"io/ioutil"
	"os"
	"time"

	"github.com/aleksaan/diskusage/config"
	"github.com/aleksaan/diskusage/console"
	"github.com/aleksaan/diskusage/files"
)

var sortValues = map[string]float64{
	"name_asc":  1,
	"size_desc": 2,
}

//Files -
var Files = &files.TFiles{}

//FinalFiles -
var FinalFiles = &files.TFiles{}
var cfg *config.Config
var basePath string
var c = make(chan int)
var countFiles int

//Run -
func Run() {
	startProcess()
	cfg = config.Cfg
	basePath = files.AddPathSeparator(*cfg.Analyzer.Path)
	scanDir(basePath, 1)
	calcAdaptedSizeInOverallInfo()
	endProcess()
}

//-----------------------------------------------------------------------------------------

//ScanDir - scan directory and return its size
func scanDir(path string, depth int) int64 {
	//read content of folder
	osfiles, _ := ioutil.ReadDir(path)

	var dirsize int64

	//calc total size throught folder content
	for _, osfile := range osfiles {

		file := scanFile(path, osfile.Name(), depth)
		if file.IsDir {
			newpath := files.AddPathSeparator(path + osfile.Name())
			file.Size = scanDir(newpath, depth+1)
		}

		setAdaptedFileSize(file)
		if *cfg.Analyzer.SizeCalculatingMethod == "cumulative" || (*cfg.Analyzer.SizeCalculatingMethod == "plain" && !file.IsDir) {
			dirsize += file.Size
		}

		//*Files = append(*Files, *file)
		addToOverallInfo(file)

		//increase count of processed files and folders and out it to console
		countFiles++
		c <- countFiles

		if depth <= *cfg.Filter.Depth {
			*FinalFiles = append(*FinalFiles, *file)
		}
	}

	return dirsize
}

func setAdaptedFileSize(file *files.TFile) {
	file.AdaptedSize, file.AdaptedUnit = files.GetAdaptedSize(file.Size, cfg.Printer.Units)
}

//-----------------------------------------------------------------------------------------

//ScanFile - scan dir/file parameters
func scanFile(path string, name string, depth int) *files.TFile {
	f := &files.TFile{}
	f.Name = name
	f.RelativePath = path[len(basePath):]
	f.Depth = depth

	//if file or folder is not accessible then return nil
	pathName := files.CleanPath(&path, false) + name

	//dirstat, _ := os.Stat(pathName)
	dir, err := os.Lstat(pathName)

	if err != nil {
		f.IsNotAccessible = true
		f.IsNotAccessibleMessage = err.Error()
		return f
	}

	if dir.Mode()&os.ModeSymlink != 0 {
		f.IsLink = true
		f.LinkedDirPath = "Unknown"
	}

	f.IsDir = dir.IsDir()

	if !dir.IsDir() {
		f.Size = dir.Size()
	}

	return f
}

func startProcess() {
	startTime = time.Now()
	go console.PrintCountFilesToConsole(c)
}

func endProcess() {
	t := time.Now()
	OverallInfo.TotalTime = t.Sub(startTime)
	close(c)
}
