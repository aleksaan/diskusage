package console

import (
	"bufio"
	"fmt"
	"os"
)

//WaitExit -
func WaitExit(isNeedWaiting bool) {
	if isNeedWaiting {
		runReadRune("\nPress [Y] for exit: ")
	}
}

func runReadRune(prompt string) {

	for {
		fmt.Print(prompt)

		reader := bufio.NewReader(os.Stdin)
		c, _, err := reader.ReadRune()
		if err != nil {
			fmt.Println(err)
			return
		}

		//fmt.Printf("%q\n", c)

		if c == 'y' || c == 'Y' {
			break
		}
	}
}
