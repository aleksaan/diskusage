package console

import (
	"fmt"
	"os"
	"time"
)

//PrintCountFilesToConsole - print count of processed files and folder to console
func PrintCountFilesToConsole(c chan int) {
	var t time.Time
	for i := range c {
		t = time.Now()
		fmt.Fprintf(os.Stdout, "\r%s: Files and folders processed: %10d ", t.Format("15:04:05"), i)
	}
}
