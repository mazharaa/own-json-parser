/// 2>/dev/null; exec go run "$0" "$@"

package main

import (
	"fmt"
	"os"
)

func parseJSON(data []byte) bool {
	foundFirstCurl := false
	foundPair := false

	for i:=range data {
		if string(data[i]) == " " {
			continue
		} else if string(data[i]) == "{" {
			if !foundFirstCurl {
				foundFirstCurl = true
				continue
			} else if foundPair{
				foundPair = false
				continue
			}
		} else if string(data[i]) == "}" {
			if foundPair{
				return false
			} else if foundFirstCurl {
				foundPair = true
				continue
			}
		} else {
			return false
		}
	}
	return foundPair
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Error: No Argument")
		os.Exit(1)
	}
	
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