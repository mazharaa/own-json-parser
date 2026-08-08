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
		{"step1 invalid2", "testdata/step1/invalid2.json", false},
		{"step1 invalid3", "testdata/step1/invalid3.json", false},
		{"step1 invalid4", "testdata/step1/invalid4.json", false},
		{"step1 invalid5", "testdata/step1/invalid5.json", false},
		{"step1 invalid6", "testdata/step1/invalid6.json", false},
		{"step1 invalid7", "testdata/step1/invalid7.json", false},
		{"step1 valid", "testdata/step1/valid.json", true},
		{"step1 valid2", "testdata/step1/valid2.json", true},
		{"step1 valid3", "testdata/step1/valid3.json", true},
		{"step1 valid4", "testdata/step1/valid4.json", true},
		{"step1 valid5", "testdata/step1/valid5.json", true},
		{"step2 invalid", "testdata/step2/invalid.json", false},
		{"step2 invalid2", "testdata/step2/invalid2.json", false},
		{"step2 invalid3", "testdata/step2/invalid3.json", false},
		{"step2 invalid4", "testdata/step2/invalid4.json", false},
		{"step2 invalid5", "testdata/step2/invalid5.json", false},
		{"step2 invalid6", "testdata/step2/invalid6.json", false},
		{"step2 invalid7", "testdata/step2/invalid7.json", false},
		{"step2 invalid8", "testdata/step2/invalid8.json", false},
		{"step2 invalid9", "testdata/step2/invalid9.json", false},
		{"step2 valid", "testdata/step2/valid.json", true},
		{"step2 valid2", "testdata/step2/valid2.json", true},
		{"step2 valid3", "testdata/step2/valid3.json", true},
		{"step2 valid4", "testdata/step2/valid4.json", true},
		{"step2 valid5", "testdata/step2/valid5.json", true},
		{"step2 valid6", "testdata/step2/valid6.json", true},
		{"step2 valid7", "testdata/step2/valid7.json", true},
		{"step2 valid8", "testdata/step2/valid8.json", true},
		{"step3 invalid", "testdata/step3/invalid.json", false},
		{"step3 invalid2", "testdata/step3/invalid.json2", false},
		{"step3 invalid3", "testdata/step3/invalid.json3", false},
		{"step3 invalid4", "testdata/step3/invalid.json4", false},
		{"step3 invalid5", "testdata/step3/invalid.json5", false},
		{"step3 invalid6", "testdata/step3/invalid.json6", false},
		{"step3 invalid7", "testdata/step3/invalid.json7", false},
		{"step3 invalid8", "testdata/step3/invalid.json8", false},
		{"step3 invalid9", "testdata/step3/invalid.json9", false},
		{"step3 invalid10", "testdata/step3/invalid.json10", false},
		{"step3 valid", "testdata/step3/valid.json", true},
		{"step3 valid2", "testdata/step3/valid.json2", true},
		{"step3 valid3", "testdata/step3/valid.json3", true},
		{"step3 valid4", "testdata/step3/valid.json4", true},
		{"step3 valid5", "testdata/step3/valid.json5", true},
		{"step3 valid6", "testdata/step3/valid.json6", true},
		{"step3 valid7", "testdata/step3/valid.json7", true},
		{"step3 valid8", "testdata/step3/valid.json8", true},
		{"step3 valid9", "testdata/step3/valid.json9", true},
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