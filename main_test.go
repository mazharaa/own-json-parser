package main

import (
	"os"
	"testing"
)

func TestValidJSON(t *testing.T) {
	tests := []struct {
		name		string
		file		string
		want		bool
	} {
		{"step1 invalid", "testdata/step1/invalid.json", false},
		{"step1 valid", "testdata/step1/valid.json", true},
		// {"step2 invalid", "testdata/step2/invalid.json", false},
		// {"step2 invalid2", "testdata/step2/invalid2.json", false},
		// {"step2 valid", "testdata/step2/valid.json", true},
		// {"step2 valid2", "testdata/step2/valid2.json", true},
		// {"step3 invalid", "testdata/step3/invalid.json", false},
		// {"step3 valid", "testdata/step3/valid.json", true},
		// {"step4 invalid", "testdata/step4/invalid.json", false},
		// {"step4 valid", "testdata/step4/valid.json", true},
		// {"step4 valid2", "testdata/step4/valid2.json", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, e := os.ReadFile(tt.file)

			if e != nil {
				t.Fatalf("failed to read %s: %v", tt.file, e)
			}

			got := parseJSON(data)
			if got != tt.want {
				t.Errorf("parseJSON(%s) = %v, want %v", tt.file, got, tt.want)
			}
		})
	}
}