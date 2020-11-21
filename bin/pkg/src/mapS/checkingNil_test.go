package maps

import "testing"

func TestIsNil(t *testing.T) {
	type args struct {
		m map[string]interface{}
	}
	caseNotNil := map[string]interface{}{
		"1": "1",
		"2": [2]string{"1", "2"},
		"3": 3, // print out "Unknown Type"
	}
	caseNil := map[string]interface{}{
		"1": []string{},
		"2": []interface{}{},
		"3": nil,
		"4": "",
	}
	caseEmptymap := make(map[string]interface{})
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"t1", args{m: caseNotNil}, false},
		{"t2", args{m: caseNil}, true},
		{"t2", args{m: caseEmptymap}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsNil(tt.args.m); got != tt.want {
				t.Errorf("IsNil() = %v", got)
			}
		})
	}

}
