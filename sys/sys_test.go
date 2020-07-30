package sys

import (
	"io/ioutil"
	"os"
	"reflect"
	"testing"
	"time"
)

func TestDu(t *testing.T) {
	type args struct {
		dir string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "ok",
			args: args{
				dir: "../sys",
			},
		}, {
			name: "err",
			args: args{
				dir: "../",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Du(tt.args.dir); got <= 0 {
				t.Errorf("Du() = %v, want > 0", got)
			}
		})
	}
}
func TestGetATime(t *testing.T) {
	type args struct {
		path        string
		defaultTime time.Time
	}
	now := time.Now()
	f, _ := ioutil.TempFile("", "")
	fInfo, _ := f.Stat()
	defer os.Remove(f.Name())
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{
			name: "err",
			args: args{
				path:        "",
				defaultTime: now,
			},
			want: now,
		}, {
			name: "ok",
			args: args{
				path:        f.Name(),
				defaultTime: now,
			},
			want: fInfo.ModTime(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetATime(tt.args.path, tt.args.defaultTime); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetATime() = %v, want %v", got, tt.want)
			}
		})
	}
}
