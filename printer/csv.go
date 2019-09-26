package printer

import (
	"fmt"
	"log"
	"time"
)

// //SaveToCsv - saving results to csv file
// func SaveToCsv(files *files.TFiles, cfg *config.Config) {
// 	if cfg.Printer.ToFile != CsvDefault {
// 		file, err := os.Create(InputArgs.CsvFileName)
// 		fmt.Printf("Saving results to %s...", InputArgs.CsvFileName)
// 		checkError("Cannot create file", err)
// 		defer file.Close()

// 		writer := csv.NewWriter(file)
// 		defer writer.Flush()

// 		writer.Comma = ';'

// 		//write header
// 		value := []string{"rownum", "file name", "size (bytes)", "adapted size", "units", "depth"}
// 		err = writer.Write(value)
// 		var i = 0

// 		//write values
// 		for _, f := range *files {
// 			if f.Depth <= InputArgs.Depth && !f.IsNotAccessible {
// 				i++

// 				value = []string{strconv.Itoa(i), f.RelativePath + f.Name, fmt.Sprintf("%d", f.Size), fmt.Sprintf("%f8.2", f.AdaptedSize), f.AdaptedUnit, strconv.Itoa(f.Depth)}
// 				err := writer.Write(value)
// 				checkError("Cannot write to file", err)

// 				//break if we up to defined limit in case limit > 0
// 				if isExceedLimit(i + 1) {
// 					break
// 				}
// 			}
// 		}
// 		fmt.Printf("OK\n")

// 	}
// }

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}

func getDefaultCsvFileName() string {
	tnow := time.Now()
	tnowstr := tnow.Format("20060102_150405")
	return fmt.Sprintf("./results/result_%s.csv", tnowstr)
}
