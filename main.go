/// 2>/dev/null; exec go run "$0" "$@"

package main

import (
	"fmt"
	"os"
)

func parseJSON(data []byte, pos *int) bool {

	if len(data) > 1 {

		skipWhiteSpace(data, pos)

		if !expect(data, pos, '{') {
			return false
		}

		skipWhiteSpace(data, pos)

		if !expect(data, pos, '}') {
			for {
				skipWhiteSpace(data, pos)

				if !expect(data, pos, '"') {
					return false
				}

				if !skipChars(data, pos) {
					return false
				}

				if !expect(data, pos, '"') {
					return false
				}

				skipWhiteSpace(data, pos)

				if !expect(data, pos, ':') {
					return false
				}

				if !parseValue(data, pos) {
					return false
				}

				skipWhiteSpace(data, pos)

				if *pos >= len(data) {
					return false
				}

				if data[*pos] != byte(',') {
					break
				} else {
					*pos ++
				}
			}

			skipWhiteSpace(data, pos)

			if !expect(data, pos, '}') {
				return false
			}
		}

		skipWhiteSpace(data, pos)

		// return *pos == len(data)
		return true
	} else {
		return false
	}
}

func skipWhiteSpace(input []byte, pos *int) {
	for _, value :=range input[*pos:] {
		if value == byte('\t') || value == byte('\n') || value == byte('\r') || value == byte(' '){
			*pos ++
		} else {
			break
		}
	}
}

func isValidEscape(b byte) bool {
	switch b {
	case '"', '\\', '/', 'b', 'f', 'n', 'r', 't', 'u':
		return true
	}
	return false
}

func isHexDigit(b byte) bool {
	return (b >= '0' && b <= '9') || (b >= 'A' && b <= 'F') || (b >= 'a' && b <= 'f')
}

func skipChars(input []byte, pos *int) bool {
	for *pos < len(input) {
		if input[*pos] == '"' {
			break
		} else if input[*pos] == 0x9 || input[*pos] == 0xa{
			return false
		} else if input[*pos] == '\\' {
			*pos ++
			if *pos >= len(input) {return false}
			if !isValidEscape(input[*pos]) {return false}

			if input[*pos] == 'u' {
				if *pos + 4 >= len(input) {return false}
				for i := 1; i <= 4; i++ {
					if !isHexDigit(input[*pos+i]) {return false}
				}
				*pos += 4
			}
			*pos ++
		} else {
			*pos ++
		}
	}

	return true
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
		} else if input[*pos] == byte(',') || input[*pos] == byte('}') || input[*pos] == byte(']') {
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
			break
		}
	}

	return true
}

func parseArray(input []byte, pos *int) bool {
	// if !skipChars(input, pos) {
	// 	return false
	// }

	if !expect(input, pos, '[') {
		return false
	}

	skipWhiteSpace(input, pos)

	if expect(input, pos, ']') {
		return true
	}

	for {
		if !parseValue(input, pos) {
			return false
		}

		skipWhiteSpace(input, pos)

		if expect(input, pos, ',') {
			continue
		}
		break
	}

	return expect(input, pos, ']')
}

func parseValue(data []byte, pos *int) bool {
	skipWhiteSpace(data, pos)

	if *pos >= len(data) {
		return false
	}

	switch data[*pos] {
	case '"':
		return parseString(data, pos)
	case 't', 'f', 'n':
		return keywordTypeMatch(data, pos, data[*pos])
	case '{':
		return parseJSON(data, pos)
	case '[':
		return parseArray(data, pos)
	default:
		return parseNum(data, pos, data[*pos])
	}
}

func validJSONPos(data []byte) (bool, int) {
	pos := 0

	skipWhiteSpace(data, &pos)

	if pos >= len(data) {
		return false, pos
	}

	switch data[pos] {
	case '{':
		return parseJSON(data, &pos) && pos == len(data), pos
	case '[':
		return parseArray(data, &pos) && pos == len(data), pos
	case '"':
		return parseString(data, &pos) && pos == len(data), pos
	case 't', 'f', 'n':
		return keywordTypeMatch(data, &pos, data[pos]) && pos == len(data), pos
	default:
		return parseNum(data, &pos, data[pos]) && pos == len(data), pos
	}
}

func validJSON(data []byte) bool {
	valid, _ := validJSONPos(data)
	return valid
}

func parseString(data []byte, pos *int) bool {
	if !expect(data, pos, '"') {
		return false
	}

	if !skipChars(data, pos) {
		return false
	}

	return expect(data, pos, '"')
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

	if validJSON(data) {
		fmt.Println("Valid JSON")
		os.Exit(0)
	} else {
		fmt.Println("Invalid JSON")
		os.Exit(1)
	}
}