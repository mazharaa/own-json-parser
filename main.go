/// 2>/dev/null; exec go run "$0" "$@"

package main

import (
	"fmt"
	"os"
)

func parseJSON(data []byte) bool {
	pos := 0

	if len(data) > 1 {

		skipWhiteSpace(data[pos:], &pos)

		if !expect(data, &pos, '{') {
			return false
		}

		skipWhiteSpace(data[pos:], &pos)

		if !expect(data, &pos, '}') {
			for {
				skipWhiteSpace(data[pos:], &pos)

				if !expect(data, &pos, '"') {
					return false
				}

				skipChars(data, &pos)

				if !expect(data, &pos, '"') {
					return false
				}

				skipWhiteSpace(data[pos:], &pos)

				if !expect(data, &pos, ':') {
					return false
				}

				skipWhiteSpace(data[pos:], &pos)

				if !expect(data, &pos, '"') {
					return false
				}

				skipChars(data, &pos)

				if !expect(data, &pos, '"') {
					return false
				}

				skipWhiteSpace(data[pos:], &pos)

				if pos >= len(data) {
					return false
				}

				if data[pos] != byte(',') {
					break
				} else {
					pos ++
				}
			}

			skipWhiteSpace(data[pos:], &pos)

			if !expect(data, &pos, '}') {
				return false
			}
		}

		skipWhiteSpace(data[pos:], &pos)

		return pos == len(data)
	} else {
		return false
	}
}

func skipWhiteSpace(input []byte, pos *int) {
	for _, value :=range input {
		if value == byte('\t') || value == byte('\n') || value == byte('\r') || value == byte(' '){
			*pos ++
		} else {
			break
		}
	}
}

func skipChars(input []byte, pos *int) {
	for *pos < len(input) {
		if input[*pos] == '"' {
			break
		} else if input[*pos] == '\\' {
			*pos ++
			if *pos >= len(input) {
				break
			}
			if input[*pos] == 'u' {
				*pos += 4
			}
			*pos ++
		} else {
			*pos ++
		}
	}
}

func expect(input []byte, pos *int, char byte) bool {
	if *pos >= len(input) || input[*pos] != char {
		return false
	}
	*pos ++
	return true
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