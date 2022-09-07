package filter

import (
	"encoding/json"
	"fmt"

	"github.com/aleksaan/diskusage/pkg/printer"
)

// func WriteJSONToConsole() {
// 	//r, err := json.MarshalIndent(result, "", "  ")
// 	r, err := json.Marshal(results)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return
// 	}
// 	fmt.Println("JSON")
// 	fmt.Println(string(r))
// }

// func WriteTextToConsole() {
// 	printConfig(cfg)
// 	printFilesHR(&results.Files)
// 	printOverall(&results.Overall)
// }

func WriteJSONToConsole() {
	r, err := json.Marshal(fresults)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("%s", string(r))
}

func WriteHumanReadableToConsole() {
	printConfig()
	printer.PrintFiles(&fresults.Files, "", 50)
}

func printConfig() {
	fmt.Println("-------------------\nArguments:")
	fmt.Printf("   %-10s %s\n", "filter:", cfg.Filter)
	fmt.Printf("   %-10s %d\n", "depth:", cfg.Depth)
	fmt.Printf("   %-10s %d\n", "top:", cfg.Top)
	fmt.Printf("   %-10s %v\n", "hr:", cfg.Hr)
	fmt.Printf("   %-10s %s\n", "size:", cfg.Size)
}
