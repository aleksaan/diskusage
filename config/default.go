package config

import "os"

var (
	defaultDepth                 = 5
	defaultLimit                 = 20
	defaultSizeCalculatingMethod = "cumulative"
	defaultUnits                 = ""
	defaultFilterByObject        = "folders&files"
	defaultToTextFile            = ""
	defaultPath                  = ""
	defaultSort                  = "desc"
)

func initDefaultsValues() {
	defaultPath, _ = os.Getwd()

}
