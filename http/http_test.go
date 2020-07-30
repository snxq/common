package http

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestDo(t *testing.T) {
	type args struct {
		req *http.Request
		res interface{}
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{"code": 999, "msg": "success"}`)
	}))
	defer ts.Close()
	type Response struct {
		Code int32  `json:"code" yaml:"code"`
		Msg  string `json:"msg" yaml:"msg"`
	}
	getReq := func(url string) *http.Request {
		req, _ := http.NewRequest(http.MethodGet, url, nil)
		return req
	}
	tests := []struct {
		name    string
		args    args
		want    *Response
		wantErr bool
	}{
		{
			name: "ok_json",
			args: args{
				req: getReq(ts.URL),
				res: &Response{},
			},
			want:    &Response{999, "success"},
			wantErr: false,
		}, {
			name: "err_do",
			args: args{
				req: getReq("no_exist_url"),
				res: &Response{},
			},
			want:    &Response{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Do(tt.args.req, tt.args.res); (err != nil) != tt.wantErr {
				t.Errorf("Do() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(tt.args.res, tt.want) {
				t.Errorf("Do() res = %v, want %v", tt.args.res, tt.want)
			}
		})
	}
}
