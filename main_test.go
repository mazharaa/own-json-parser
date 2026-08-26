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
		{"step3 invalid2", "testdata/step3/invalid2.json", false},
		{"step3 invalid3", "testdata/step3/invalid3.json", false},
		{"step3 invalid4", "testdata/step3/invalid4.json", false},
		{"step3 invalid5", "testdata/step3/invalid5.json", false},
		{"step3 invalid6", "testdata/step3/invalid6.json", false},
		{"step3 invalid7", "testdata/step3/invalid7.json", false},
		{"step3 invalid8", "testdata/step3/invalid8.json", false},
		{"step3 invalid9", "testdata/step3/invalid9.json", false},
		{"step3 invalid10", "testdata/step3/invalid10.json", false},
		{"step3 valid", "testdata/step3/valid.json", true},
		{"step3 valid2", "testdata/step3/valid2.json", true},
		{"step3 valid3", "testdata/step3/valid3.json", true},
		{"step3 valid4", "testdata/step3/valid4.json", true},
		{"step3 valid5", "testdata/step3/valid5.json", true},
		{"step3 valid6", "testdata/step3/valid6.json", true},
		{"step3 valid7", "testdata/step3/valid7.json", true},
		{"step3 valid8", "testdata/step3/valid8.json", true},
		{"step3 valid9", "testdata/step3/valid9.json", true},
		{"step4 invalid", "testdata/step4/invalid.json", false},
		{"step4 invalid2", "testdata/step4/invalid2.json", false},
		{"step4 invalid3", "testdata/step4/invalid3.json", false},
		{"step4 invalid4", "testdata/step4/invalid4.json", false},
		{"step4 invalid5", "testdata/step4/invalid5.json", false},
		{"step4 invalid6", "testdata/step4/invalid6.json", false},
		{"step4 invalid7", "testdata/step4/invalid7.json", false},
		{"step4 invalid8", "testdata/step4/invalid8.json", false},
		{"step4 valid", "testdata/step4/valid.json", true},
		{"step4 valid2", "testdata/step4/valid2.json", true},
		{"step4 valid3", "testdata/step4/valid3.json", true},
		{"step4 valid4", "testdata/step4/valid4.json", true},
		{"step4 valid5", "testdata/step4/valid5.json", true},
		{"step4 valid6", "testdata/step4/valid6.json", true},
		{"step4 valid7", "testdata/step4/valid7.json", true},
		{"step4 valid8", "testdata/step4/valid8.json", true},
		{"step4 valid9", "testdata/step4/valid9.json", true},
		{"edge valid-empty-array", "testdata/edge/valid-empty-array.json", true},
		{"edge valid-top-level-array", "testdata/edge/valid-top-level-array.json", true},
		{"edge valid-top-level-string", "testdata/edge/valid-top-level-string.json", true},
		{"edge valid-top-level-number", "testdata/edge/valid-top-level-number.json", true},
		{"edge valid-top-level-true", "testdata/edge/valid-top-level-true.json", true},
		{"edge valid-top-level-null", "testdata/edge/valid-top-level-null.json", true},
		{"edge valid-whitespace-before-comma", "testdata/edge/valid-whitespace-before-comma.json", true},
		{"edge invalid-missing-comma-array", "testdata/edge/invalid-missing-comma-array.json", false},
		{"edge invalid-trailing-value", "testdata/edge/invalid-trailing-value.json", false},
		{"edge invalid-garbage-after-number", "testdata/edge/invalid-garbage-after-number.json", false},
		{"edge pass5", "testdata/edge/pass5.json", true},
		{"edge fail2", "testdata/edge/fail2.json", false},
		{"edge fail3", "testdata/edge/fail3.json", false},
		{"edge fail4", "testdata/edge/fail4.json", false},
		{"edge fail5", "testdata/edge/fail5.json", false},
		{"edge fail6", "testdata/edge/fail6.json", false},
		{"edge fail7", "testdata/edge/fail7.json", false},
		{"edge fail8", "testdata/edge/fail8.json", false},
		{"edge fail9", "testdata/edge/fail9.json", false},
		{"edge fail10", "testdata/edge/fail10.json", false},
		{"edge fail11", "testdata/edge/fail11.json", false},
		{"edge fail12", "testdata/edge/fail12.json", false},
		{"edge fail13", "testdata/edge/fail13.json", false},
		{"edge fail14", "testdata/edge/fail14.json", false},
		{"edge fail15", "testdata/edge/fail15.json", false},
		{"edge fail16", "testdata/edge/fail16.json", false},
		{"edge fail17", "testdata/edge/fail17.json", false},
		{"edge pass4", "testdata/edge/pass4.json", true},
		{"edge fail19", "testdata/edge/fail19.json", false},
		{"edge fail20", "testdata/edge/fail20.json", false},
		{"edge fail21", "testdata/edge/fail21.json", false},
		{"edge fail22", "testdata/edge/fail22.json", false},
		{"edge fail23", "testdata/edge/fail23.json", false},
		{"edge fail24", "testdata/edge/fail24.json", false},
		{"edge fail25", "testdata/edge/fail25.json", false},
		{"edge fail26", "testdata/edge/fail26.json", false},
		{"edge fail27", "testdata/edge/fail27.json", false},
		{"edge fail28", "testdata/edge/fail28.json", false},
		{"edge fail29", "testdata/edge/fail29.json", false},
		{"edge fail30", "testdata/edge/fail30.json", false},
		{"edge fail31", "testdata/edge/fail31.json", false},
		{"edge fail32", "testdata/edge/fail32.json", false},
		{"edge fail33", "testdata/edge/fail33.json", false},
		{"edge fail34", "testdata/edge/fail34.json", true},
		{"edge fail35", "testdata/edge/fail35.json", false},
		{"edge pass1", "testdata/edge/pass1.json", true},
		{"edge pass2", "testdata/edge/pass2.json", true},
		{"edge pass3", "testdata/edge/pass3.json", true},
		{"reg invalid-top-level-neg-leading-zero", "testdata/edge/invalid-top-level-neg-leading-zero.json", false},
		{"reg valid-top-level-neg-zero", "testdata/edge/valid-top-level-neg-zero.json", true},
		{"reg invalid-bare-minus", "testdata/edge/invalid-bare-minus.json", false},
		{"reg valid-neg-zero-in-object", "testdata/edge/valid-neg-zero-in-object.json", true},
		{"reg valid-neg-zero-fraction", "testdata/edge/valid-neg-zero-fraction.json", true},
		{"reg invalid-raw-control-0x0b", "testdata/edge/invalid-raw-control-0x0b.json", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, e := os.ReadFile(tt.file)

			if e != nil {
				t.Fatalf("failed to read %s: %v", tt.file, e)
			}

			valid, pos := validJSONPos(data)

			if valid != tt.want {
				if pos <= len(data) {
					t.Logf("pos=%d len=%d remaining=%q", pos, len(data), string(data[pos:]))
				} else {
					t.Logf("pos=%d len=%d (pos beyond end)", pos, len(data))
				}
				t.Errorf("validJSON(%s) = %v, want %v", tt.file, valid, tt.want)
			}
		})
	}
}