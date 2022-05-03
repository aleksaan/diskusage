package analyzer

import (
	"io/ioutil"
	"os"
	"time"

	"github.com/aleksaan/diskusage/config"
	"github.com/aleksaan/diskusage/files"
	"github.com/aleksaan/diskusage/models"
)

const (
	AppTitle   = "https://github.com/aleksaan/diskusage"
	AppVersion = "2.8.0"
	AppAuthor  = "Alexander Anufriev"
	AppYear    = "2022"
)

var Result = &models.TResult{}

//FinalFiles -
var FinalFiles = &models.TFiles{}
var cfg *config.Config
var basePath string

//var c = make(chan int)
var countFiles int

//Run -
func Run() {
	Result.About = models.TAbout{
		AppTitle:   AppTitle,
		AppVersion: AppVersion,
		AppAuthor:  AppAuthor,
		AppYear:    AppYear,
	}

	startProcess()
	cfg = config.Cfg
	basePath = files.AddPathSeparator(cfg.Path)
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
		if cfg.SizeCalculatingMethod == "cumulative" || (cfg.SizeCalculatingMethod == "plain" && !file.IsDir) {
			dirsize += file.Size
		}

		//*Files = append(*Files, *file)
		addToOverallInfo(file)

		//increase count of processed files and folders and out it to console
		countFiles++
		//c <- countFiles

		if depth <= cfg.Depth {
			*FinalFiles = append(*FinalFiles, *file)
		}
	}

	return dirsize
}

func setAdaptedFileSize(file *models.TFile) {
	file.AdaptedSize, file.AdaptedUnit = files.GetAdaptedSize(file.Size, &cfg.Units)
}

//-----------------------------------------------------------------------------------------

//ScanFile - scan dir/file parameters
func scanFile(path string, name string, depth int) *models.TFile {
	f := &models.TFile{}
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
	//go console.PrintCountFilesToConsole(c)
}

func endProcess() {
	t := time.Now()
	Result.Overall.TotalTime = t.Sub(startTime)
	//close(c)
}
