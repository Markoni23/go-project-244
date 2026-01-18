package code_test

import (
	"code"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenDiff(t *testing.T) {
	tests := []struct {
		name       string
		firstFile  string
		secondFile string
		want       string
		wantErr    bool
	}{
		{
			"plain json",
			"testdata/fixture/test1.json",
			"testdata/fixture/test2.json",
			"{\n  + added_field: 4\n  - changed_field: 2\n  + changed_field: 2.4\n    normal_field: 1\n  - removed_field: 3\n}", false},
		/*
			test1.json
			{
				"normal_field": "1",
				"changed_field": 2.0,
				"removed_field": 3
			}

			test2.json
			{
				"normal_field": "1",
				"changed_field": 2.4,
				"added_field": 4
			}

			want
			{
			  + added_field: 4
			  - changed_field: 2
			  + changed_field: 2.4
			    normal_field: 1
			  - removed_field: 3
			}
		*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := code.GenDiff(tt.firstFile, tt.secondFile)
			if tt.wantErr {
				assert.NotNil(t, gotErr, "expected error")
			}

			assert.Equal(t, got, tt.want)
		})
	}
}
