package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	operation "rpn_calculator/calculate"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		parameters := strings.Fields(line)
		result, err := operation.Operate(parameters)

		if err != nil {
			fmt.Println("ERROR")
		} else {
			fmt.Println(result)
		}
	}
}
