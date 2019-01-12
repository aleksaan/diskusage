package tests

import (
	"testing"

	"github.com/aleksaan/diskusage/diskusage"
)

func TestCleanPathRelativeToAbs(t *testing.T) {
	argpath := "../data/dir1/"
	waitedPath := "D:\\_appl\\go\\src\\github.com\\aleksaan\\diskusage\\data\\dir1\\"
	cleanPath := diskusage.CleanPath(&argpath, true)
	if cleanPath != waitedPath {
		t.Errorf("\n\nCleanPath does not work correct\n Wait: %s\n Got:%s\n", waitedPath, cleanPath)
	}
}

func TestCleanPathAbsToAbs(t *testing.T) {

	argpath := "D:\\_appl\\go/src\\github.com\\aleksaan\\diskusage\\data\\dir1\\"
	waitedPath := "D:\\_appl\\go\\src\\github.com\\aleksaan\\diskusage\\data\\dir1\\"
	cleanPath := diskusage.CleanPath(&argpath, false)
	if cleanPath != waitedPath {
		t.Errorf("\n\nCleanPath does not work correct\n Wait: %s\n Got:%s\n", waitedPath, cleanPath)
	}
}

func TestAddPathSeparatorToDisk(t *testing.T) {

	argpath := "C:\\"
	waitedPath := argpath
	cleanPath := diskusage.CleanPath(&argpath, false)
	if cleanPath != waitedPath {
		t.Errorf("\n\nCleanPath does not work correct\nWait: %s\nGot: %s\n", waitedPath, cleanPath)
	}
}

func TestScanFile(t *testing.T) {
	//files := &diskusage.TFiles{}

	diskusage.InputArgs.Depth = 1
	diskusage.InputArgs.Limit = 1
	argpath := "../data/dir1/"
	diskusage.InputArgs.SetPath(&argpath)

	name := "file1.txt"

	file := diskusage.ScanFile(diskusage.InputArgs.Path, name, 1)

	if file.IsNotAccessible {
		t.Errorf("\nFile %s/%s is not accessible\n", diskusage.InputArgs.Path, name)
		//t.Errorf("%s", file.RelativePath)
		t.Errorf("%s", file.IsNotAccessibleMessage)
	}

	if file.Size != 18 || file.IsDir {
		t.Errorf("File %s cannot be read correct (function diskusage.ScanFile is not work)", diskusage.InputArgs.Path)
	}
}

func TestScanDir(t *testing.T) {
	files := &diskusage.TFiles{}

	diskusage.InputArgs.Depth = 1
	diskusage.InputArgs.Limit = 1
	argpath := "../data/dir1"
	diskusage.InputArgs.SetPath(&argpath)

	diskusage.ScanDir(files, diskusage.InputArgs.Path, 1)

	if len(*files) != 2 {
		t.Errorf("Dir %s cannot be read correct (function diskusage.ScanDir is not work)", diskusage.InputArgs.Path)
	}
}

func TestScanWrongSymlinkFile(t *testing.T) {
	//files := &diskusage.TFiles{}

	diskusage.InputArgs.Depth = 1
	diskusage.InputArgs.Limit = 1
	argpath := "../data/dir2_symlink/"
	diskusage.InputArgs.SetPath(&argpath)
	name := "symlink_to_deleted_file.txt.lnk"

	file := diskusage.ScanFile(diskusage.InputArgs.Path, name, 1)

	if file.IsNotAccessible == true {
		t.Errorf("Wrong symlink file %s cannot be read correct (function diskusage.ScanFile is not work)", diskusage.InputArgs.Path)
	}
}

func TestScanAllUsers(t *testing.T) {
	//files := &diskusage.TFiles{}

	diskusage.InputArgs.Depth = 1
	diskusage.InputArgs.Limit = 1
	argpath := "C:/Users/"
	diskusage.InputArgs.SetPath(&argpath)
	name := "All Users"

	file := diskusage.ScanFile(diskusage.InputArgs.Path, name, 1)

	if file.IsNotAccessible == true {
		t.Errorf("Wrong symlink file %s cannot be read correct (function diskusage.ScanFile is not work)", diskusage.InputArgs.Path)
	}
}
