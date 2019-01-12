package tests

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/aleksaan/diskusage/diskusage"
)

func TestScanFile(t *testing.T) {
	//files := &diskusage.TFiles{}

	diskusage.InputArgs.Depth = 1
	diskusage.InputArgs.Limit = 1

	path, _ := filepath.Abs("../data/dir1/")
	name := "file1.txt"

	file := diskusage.ScanFile(path, name, 1)

	if file.RelativePath != path || file.Size != 18 || file.IsDir {
		t.Errorf("File %s cannot be read correct (function diskusage.ScanFile is not work)", path)
	}
}

func TestScanDir(t *testing.T) {
	files := &diskusage.TFiles{}

	diskusage.InputArgs.Depth = 1
	diskusage.InputArgs.Limit = 1

	path, _ := filepath.Abs("../data/dir1")

	diskusage.ScanDir(files, path, 1)

	if len(*files) != 2 {
		t.Errorf("Dir %s cannot be read correct (function diskusage.ScanDir is not work)", path)
	}
}

func TestScanWrongSymlinkFile(t *testing.T) {
	//files := &diskusage.TFiles{}

	diskusage.InputArgs.Depth = 1
	diskusage.InputArgs.Limit = 1

	path, _ := filepath.Abs("../data/dir2_symlink/")
	name := "symlink_to_deleted_file.txt.lnk"

	file := diskusage.ScanFile(path, name, 1)

	if file.IsNotAccessible == true {
		t.Errorf("Wrong symlink file %s cannot be read correct (function diskusage.ScanFile is not work)", path)
	}
}

func TestScanAllUsers(t *testing.T) {
	//files := &diskusage.TFiles{}

	diskusage.InputArgs.Depth = 1
	diskusage.InputArgs.Limit = 1
	path := filepath.Clean("C:/Users/")
	name := "All Users"

	file := diskusage.ScanFile(path+string(os.PathSeparator), name, 1)

	if file.IsNotAccessible == true {
		t.Errorf("Wrong symlink file %s cannot be read correct (function diskusage.ScanFile is not work)", path)
	}
}
