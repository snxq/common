package common

import (
	"reflect"
	"testing"
)

type ConvertA struct {
	Nothing string `json:"a"`
}

type ConvertB struct {
	Something string `json:"a"`
}

type ConvertC struct {
	A string
}

type ConvertD struct {
	Nothing string `json:"b"`
}

func TestConvertByJSON(t *testing.T) {
	type args struct {
		x interface{}
		y interface{}
	}
	tests := []struct {
		name       string
		args       args
		wantErr    bool
		wantResult interface{}
	}{
		{
			name:       "ok_same_tag",
			args:       args{x: &ConvertA{"test"}, y: &ConvertB{}},
			wantResult: &ConvertB{"test"},
		}, {
			name:       "ok_has_no_tag",
			args:       args{x: &ConvertA{"test"}, y: &ConvertC{}},
			wantResult: &ConvertC{"test"},
		}, {
			name:       "faild_diff_tag",
			args:       args{x: &ConvertA{"test"}, y: &ConvertD{}},
			wantResult: &ConvertD{},
		}, {
			name:       "faild_x_nil",
			args:       args{x: nil, y: &ConvertB{}},
			wantResult: &ConvertB{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ConvertByJSON(tt.args.x, tt.args.y); (err != nil) != tt.wantErr {
				t.Errorf("ConvertByJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(tt.args.y, tt.wantResult) {
				t.Errorf("ConvertByJSON() got = %v, wantResult = %v", tt.args.y, tt.wantResult)
			}
		})
	}
}
