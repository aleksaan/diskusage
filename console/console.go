package console

import (
	"fmt"
	"log"
	"os"

	ui "github.com/gizak/termui/v3"
)

const (
	ColumnsCount = 80
	RowsCount    = 40
	AppTitle     = "diskusage"
	AppVersion   = "2.0.1"
)

const (
	BellCharacter = "\a"
)

type AsciiFont string

const (
	AsciiFont2D AsciiFont = "2d"
	AsciiFont3D AsciiFont = "3d"
)

func Init() {

	fmt.Printf("\033]0;%s\007", AppTitle)

	if err := ui.Init(); err != nil {
		log.Fatalf("Failed to initialize ui: %v", err)
	}
}

func Close() {
	ui.Close()
}

func Exit(message string) {
	if len(message) > 0 {
		println(message)
	}
	os.Exit(0)
}
