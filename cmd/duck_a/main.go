package main

import (
	"github.com/aleksaan/diskusage/pkg/analyzer"
)

func main() {
	analyzer.CfgInit()
	analyzer.Start()
}
