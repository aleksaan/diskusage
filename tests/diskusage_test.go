package tests

import (
	"path/filepath"
	"testing"

	"github.com/aleksaan/diskusage/diskusage"
)

func TestScanFile(t *testing.T) {
	//files := &diskusage.TFiles{}

	diskusage.InputArgs.Depth = 1
	diskusage.InputArgs.Limit = 1

	path, _ := filepath.Abs("../data/dir1/file1.txt")

	file := diskusage.ScanFile(path, 1)

	if file.Path != path || file.Size != 18 || file.IsDir {
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

	path, _ := filepath.Abs("../data/dir2_symlink/symlink_to_deleted_file.txt.lnk")

	file := diskusage.ScanFile(path, 1)

	if file.IsNotAccessible == true {
		t.Errorf("Wrong symlink file %s cannot be read correct (function diskusage.ScanFile is not work)", path)
	}
}
