package common

import (
	"testing"
)

func TestIsZero(t *testing.T) {
	type args struct {
		x interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "int_非零值",
			args: args{1},
			want: false,
		}, {
			name: "int_零值",
			args: args{0},
			want: true,
		}, {
			name: "string_非零值",
			args: args{"s"},
			want: false,
		}, {
			name: "string_零值",
			args: args{""},
			want: true,
		}, {
			name: "slice_非零值_1",
			args: args{[]int{1}},
			want: false,
		}, {
			name: "slice_非零值_2",
			args: args{[]int{}},
			want: false,
		}, {
			name: "slice_零值",
			args: args{nil},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsZero(tt.args.x); got != tt.want {
				t.Errorf("IsZero() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAllF(t *testing.T) {
	type args struct {
		f func(interface{}) bool
		x []interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "空参数",
			args: args{
				f: func(interface{}) bool { return true },
			},
			want: true,
		}, {
			name: "全真",
			args: args{
				f: func(interface{}) bool { return true },
				x: []interface{}{1},
			},
			want: true,
		}, {
			name: "全假",
			args: args{
				f: func(interface{}) bool { return false },
				x: []interface{}{1},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AllF(tt.args.f, tt.args.x...); got != tt.want {
				t.Errorf("AllF() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAllZero(t *testing.T) {
	type args struct {
		x []interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "空参数",
			args: args{},
			want: true,
		}, {
			name: "全不是零值",
			args: args{[]interface{}{1, "s"}},
			want: false,
		}, {
			name: "不全是零值",
			args: args{[]interface{}{1, ""}},
			want: false,
		}, {
			name: "全是零值",
			args: args{[]interface{}{0, ""}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AllZero(tt.args.x...); got != tt.want {
				t.Errorf("All() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAnyF(t *testing.T) {
	type args struct {
		f func(interface{}) bool
		x []interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "空参数",
			args: args{
				f: func(interface{}) bool { return true },
			},
			want: false,
		}, {
			name: "全真",
			args: args{
				f: func(interface{}) bool { return true },
				x: []interface{}{1},
			},
			want: true,
		}, {
			name: "全假",
			args: args{
				f: func(interface{}) bool { return false },
				x: []interface{}{1},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AnyF(tt.args.f, tt.args.x...); got != tt.want {
				t.Errorf("AnyF() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAnyZero(t *testing.T) {
	type args struct {
		x []interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "空参数",
			args: args{},
			want: false,
		}, {
			name: "全不是零值",
			args: args{[]interface{}{1, "s"}},
			want: false,
		}, {
			name: "不全是零值",
			args: args{[]interface{}{1, ""}},
			want: true,
		}, {
			name: "全是零值",
			args: args{[]interface{}{0, ""}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AnyZero(tt.args.x...); got != tt.want {
				t.Errorf("AnyZero() = %v, want %v", got, tt.want)
			}
		})
	}
}
