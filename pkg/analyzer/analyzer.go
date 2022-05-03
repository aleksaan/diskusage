package analyzer

import (
	"io/ioutil"
	"os"
	"time"

	"github.com/aleksaan/diskusage/pkg/files"
	"github.com/aleksaan/diskusage/pkg/models"
)

const (
	AppTitle   = "https://github.com/aleksaan/diskusage"
	AppVersion = "2.8.0"
	AppAuthor  = "Alexander Anufriev"
	AppYear    = "2022"
)

var Result = &models.TAnalyserResult{}

//FinalFiles -
var FinalFiles = &models.TFiles{}

//var c = make(chan int)
var countFiles int

//Start -
func Start() {
	Result.About = models.TAbout{
		AppTitle:   AppTitle,
		AppVersion: AppVersion,
		AppAuthor:  AppAuthor,
		AppYear:    AppYear,
	}

	Result.Config = *cfg

	startProcess()
	scanDir(files.AddPathSeparator(cfg.Path), 1)
	if cfg.Hr {
		WriteHumanReadableToConsole()
	} else {
		WriteJSONToConsole()
	}
	endProcess()
}

//-----------------------------------------------------------------------------------------

//ScanDir - scan directory and return its size
func scanDir(path string, depth uint) (int64, int64) {
	//read content of folder
	osfiles, _ := ioutil.ReadDir(path)

	var dirSize, dirSizeNoSF int64

	//calc total size throught folder content
	for _, osfile := range osfiles {

		file := scanFile(path, osfile.Name(), depth)
		if file.IsDir {
			newpath := files.AddPathSeparator(path + osfile.Name())
			file.SizeSubFoldersExcludes, file.SizeSubFoldersIncludes = scanDir(newpath, depth+1)
		}

		dirSize += file.SizeSubFoldersIncludes
		if !file.IsDir {
			dirSizeNoSF += file.SizeSubFoldersExcludes
		}

		addToOverallInfo(file)

		//increase count of processed files and folders and out it to console
		countFiles++

		if depth <= cfg.Depth || cfg.Depth <= 0 {
			Result.Files = append(Result.Files, *file)
		}
	}

	return dirSizeNoSF, dirSize
}

// func setAdaptedFileSize(file *models.TFile) {
// 	file.AdaptedSize, file.AdaptedUnit = files.GetAdaptedSize(file.Size, &cfg.Units)
// }

//-----------------------------------------------------------------------------------------

//ScanFile - scan dir/file parameters
func scanFile(path string, name string, depth uint) *models.TFile {
	f := &models.TFile{}
	//f.Name = name
	f.RelativePath = path[len(cfg.Path):] + name
	f.Depth = depth

	//if file or folder is not accessible then return nil
	pathName := files.CleanPath(&path, false) + name
	//f.FullPath = pathName

	//dirstat, _ := os.Stat(pathName)
	dir, err := os.Lstat(pathName)

	if err != nil {
		f.IsNotAccessible = true
		f.IsNotAccessibleMessage = err.Error()
		return f
	}

	if dir.Mode()&os.ModeSymlink != 0 {
		f.IsLink = true
		//f.LinkedDirPath = "Unknown"
	}

	f.IsDir = dir.IsDir()

	if !dir.IsDir() {
		f.SizeSubFoldersExcludes, f.SizeSubFoldersIncludes = dir.Size(), dir.Size()
	}

	return f
}

func startProcess() {
	startTime = time.Now()
	//go console.PrintCountFilesToConsole(c)
}

func endProcess() {
	t := time.Now()
	Result.Overall.TotalTime = t.Sub(startTime)
	//close(c)
}
