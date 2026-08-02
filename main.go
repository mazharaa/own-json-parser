/// 2>/dev/null; exec go run "$0" "$@"

package main

import (
	"fmt"
	"os"
)

func parseJSON(data []byte) bool {
	if len(data) > 1 {
		if string(data[0]) == "{" && string(data[len(data)-1]) == "}" {
			return true
		}
	}

	return false
}

func main() {
	filePath := os.Args[1]
	data, e := os.ReadFile(filePath)

	if e != nil {
		fmt.Println("Error:", e)
		os.Exit(1)
	}

	if parseJSON(data) {
		fmt.Println("Valid JSON")
		os.Exit(0)
	} else {
		fmt.Println("Invalid JSON")
		os.Exit(1)
	}
}