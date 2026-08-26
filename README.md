# Own JSON Parser

A from-scratch JSON parser built as part of the [Build Your Own JSON Parser](https://codingchallenges.fyi/challenges/challenge-json-parser) coding challenge. No external JSON libraries — the lexer, parser, and validator are all hand-written.

## Features

- Full JSON grammar support per [RFC 8259](https://datatracker.ietf.org/doc/html/rfc8259):
  - Objects `{}`, arrays `[]`, strings `""`, numbers, booleans, and `null`
  - Nested objects and arrays to any depth
  - Any valid JSON value as a top-level document (not just objects)
  - Whitespace (` `, `\t`, `\n`, `\r`) anywhere between tokens
- String escape sequences: `\"`, `\\`, `\b`, `\f`, `\n`, `\r`, `\t`, and `\uXXXX` (with hex digit validation)
- Number grammar: optional `-`, no leading zeros, optional fraction, optional exponent (`e`/`E` with `+`/`-`)
- Rejects malformed input: unquoted keys, trailing commas, missing commas, invalid escapes, garbage after values, mismatched brackets

## Usage

```bash
go build -o jsonparser .

./jsonparser path/to/file.json
echo $?   # 0 = valid JSON, 1 = invalid JSON
```

The program reads a single file path as its only argument, prints `Valid JSON` or `Invalid JSON` to stdout, and exits with code 0 (valid) or 1 (invalid) per the challenge spec.

```bash
./jsonparser testdata/step1/valid.json && echo "valid"   # prints Valid JSON
./jsonparser testdata/step1/invalid.json; echo $?        # prints Invalid JSON, exit 1
```

## Architecture

The parser uses a positional recursive-descent design — a single `pos` cursor advances through the input as each parsing function consumes what it expects.

```
main
 └─ validJSON          top-level entry: skip ws, dispatch on first char, require full consumption
     └─ parseJSON      object  { key: value, ... }
     └─ parseArray     array   [ value, ... ]
     └─ parseString    string  "..."
     └─ keywordTypeMatch  true / false / null
     └─ parseNum       number  -?\d+(\.\d+)?([eE][+-]?\d+)?
         └─ parseValue dispatcher: routes each value by its first character
```

### Functions

| Function | Role |
|----------|------|
| `parseJSON` | Parses one object: `{` then key/value pairs separated by `,`, then `}` |
| `parseArray` | Parses one array: `[` then values separated by `,`, then `]` (empty array allowed) |
| `parseValue` | Dispatcher: inspects the current character and routes to the right sub-parser |
| `parseString` | Parses `"..."` via `expect('"')` + `skipChars` + `expect('"')` |
| `skipChars` | Advances through string content; validates escape sequences and `\uXXXX` hex digits; rejects unescaped control chars |
| `keywordTypeMatch` | Matches `true`, `false`, `null` character by character |
| `parseNum` | Parses a number; enforces no leading zeros, fraction rules, exponent rules |
| `skipWhiteSpace` | Advances past ` `, `\t`, `\n`, `\r` |
| `expect` | Checks the current char matches, advances if so, else fails |
| `isValidEscape` | Validates a single escape character after `\` |
| `isHexDigit` | Validates a hex digit for `\uXXXX` |
| `validJSONPos` | Entry point that returns validity plus the final position (used by tests) |
| `validJSON` | Public entry point: returns just validity |

### Key design decisions

- **Positional cursor** — all parsers share one `pos` pointer; nothing is copied or tokenized up front
- **`parseValue` as dispatcher** — keeps objects, arrays, and scalars decoupled; new value types slot in as one `switch` case
- **Recursion for nesting** — `parseJSON`/`parseArray` call `parseValue`, which can call back into them
- **Top-level dispatch** — `validJSON` handles any value type as a document root, per RFC 8259
- **Strict termination** — parsers `break` at the first non-consuming character and let callers enforce separators; garbage is rejected downstream

## Testing

```bash
go test ./...
```

111 test cases in `main_test.go` covering all six challenge steps plus regression edges.

### Test data layout

```
testdata/
├── step1/   step-by-step challenge fixtures (valid{N}.json / invalid{N}.json)
├── step2/
├── step3/
├── step4/
└── edge/    JSON.org test suite (pass1–5, fail2–33) + hand-written regression cases
```

The `edge/` fixtures come from the official [JSON.org test suite](http://www.json.org/JSON_checker/test.zip): `pass1`–`pass5` must be accepted, `fail2`–`fail33` must be rejected. The hand-written `valid-*` / `invalid-*` files capture bugs found during development (top-level scalars, missing commas, garbage after numbers, whitespace before commas).

## Challenge Progress

| Step | Goal | Status |
|------|------|--------|
| 0 | Environment setup | Done |
| 1 | Parse `{}` vs invalid, exit codes | Done |
| 2 | String keys and values | Done |
| 3 | Numbers, booleans, null | Done |
| 4 | Nested objects and arrays | Done |
| 5 | Own tests + JSON.org suite | Done — all 5 pass files accepted, all 32 fail files rejected |
