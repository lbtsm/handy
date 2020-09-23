package router

import (
	"net/http"
	"testing"
)

func Hello2World(resp http.ResponseWriter, req *http.Request) {
	resp.WriteHeader(http.StatusOK)
	_, _ = resp.Write([]byte("ok"))
}

func Ping(resp http.ResponseWriter, req *http.Request) {
	resp.WriteHeader(http.StatusOK)
	_, _ = resp.Write([]byte("pong"))
}

func TestGenerateUrlPathByFunc(t *testing.T) {
	SetProjectName("fireLoad")
	type args struct {
		f      interface{}
		prefix []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "FuncNameHaveNum",
			args: args{f: Hello2World},
			want: "fireLoad/v1.0/hello-2-world",
		},
		{
			name: "standard-ping",
			args: args{f: Ping},
			want: "fireLoad/v1.0/ping",
		},
		{
			name: "appointPrefix",
			args: args{f: Ping, prefix: []string{"v1.1"}},
			want: "fireLoad/v1.1/ping",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateUrlPathByFunc(tt.args.f, tt.args.prefix...); got != tt.want {
				t.Errorf("GenerateUrlPathByFunc() = %v, want %v", got, tt.want)
			}
		})
	}
}
