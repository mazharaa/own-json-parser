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
					if data[pos] == byte('t') || data[pos] == byte('f') || data[pos] == byte('n') {
						if !keywordTypeMatch(data, &pos, data[pos]) {
							return false
						}
					} else {
						if !parseNum(data, &pos, data[pos]) {
							return false
						}
					}
				} else {
					skipChars(data, &pos)

					if !expect(data, &pos, '"') {
						return false
					}
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

func keywordTypeMatch(input []byte, pos *int, firstChar byte) bool {
	keywordMatch := []byte("")
	*pos ++

	if firstChar == 't' && (*pos + 2) < len(input) {
		keywordMatch = []byte("rue")
		
		for _, val := range keywordMatch {
			if input[*pos] != byte(val) {
				return false
			}
			*pos ++
		}
	} else if firstChar == 'f' && (*pos + 3) < len(input) {
		keywordMatch = []byte("alse")
		
		for _, val := range keywordMatch {
			if input[*pos] != byte(val) {
				return false
			}
			*pos ++
		}
	} else if firstChar == 'n' && (*pos + 2) < len(input) {
		keywordMatch = []byte("ull")
		
		for _, val := range keywordMatch {
			if input[*pos] != byte(val) {
				return false
			}
			*pos ++
		}
	} else {
		return false
	}

	return true
}

func parseNum(input []byte, pos *int, firstElement byte) bool {
	countDecimal := 0
	firstElementChecked := false
	for *pos < len(input) {
		if !firstElementChecked {
			if (input[*pos] == byte('-')) || (input[*pos] >= byte('0') && input[*pos] <= byte('9')) {
				*pos ++
				firstElementChecked = true
			} else {
				return false
			}
		} else if input[*pos] >= byte('0') && input[*pos] <= byte('9') {
			if firstElement == byte('0') {
				if countDecimal != 1 {
					return false
				}
			}
			*pos ++
		} else if input[*pos] == byte('.') {
			if countDecimal > 0 || input[*pos - 1] == byte('-'){
				return false
			}
			*pos ++
			countDecimal ++
		} else if (input[*pos] == byte('+') || input[*pos] == byte('-')) {
			if input[*pos - 1] != byte('e') && input[*pos - 1] != byte('E') {
				return false
			}

			*pos++
		} else if input[*pos] == byte(',') || input[*pos] == byte('}') || input[*pos] == byte('\n') {
			if input[*pos - 1] == byte('.') || input[*pos - 1] == byte('e') || input[*pos - 1] == byte('E') || input[*pos - 1] == byte('-') || input[*pos - 1] == byte('+') {
				return false
			}
			break
		} else if input[*pos] == byte('E') || input[*pos] == byte('e'){
			if input[*pos - 1] == byte('-') {
				return false
			}
			*pos ++
		} else  {
			return false
		}
	}

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