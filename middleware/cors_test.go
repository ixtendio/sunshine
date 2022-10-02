package middleware

import (
	"net/url"
	"testing"
)

func Test_isSameOrigin(t *testing.T) {
	type args struct {
		reqUrl string
		origin string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "http => true",
			args: args{
				reqUrl: "http://domain.com:80",
				origin: "http://domain.com:80",
			},
			want: true,
		},
		{
			name: "http with port => true",
			args: args{
				reqUrl: "http://domain.com",
				origin: "http://domain.com:80",
			},
			want: true,
		},
		{
			name: "https with port => true",
			args: args{
				reqUrl: "https://domain.com",
				origin: "https://domain.com:443",
			},
			want: true,
		},
		{
			name: "http without scheme => false",
			args: args{
				reqUrl: "domain.com:80",
				origin: "http://domain.com:80",
			},
			want: false,
		},
		{
			name: "case insensitive => false",
			args: args{
				reqUrl: "http://domain.com:80",
				origin: "http://DOMAIN.com:80",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reqUrl, err := url.Parse(tt.args.reqUrl)
			if err != nil {
				t.Errorf("Failed parsing the url: %s, err:%v", tt.args.reqUrl, err)
			}
			if got := isSameOrigin(reqUrl, tt.args.origin); got != tt.want {
				t.Errorf("isSameOrigin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isValidOrigin(t *testing.T) {
	type args struct {
		origin string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "empty origin => false",
			args: args{origin: ""},
			want: false,
		},
		{
			name: "origin with % => false",
			args: args{origin: "domain%20.com"},
			want: false,
		},
		{
			name: "null origin => true",
			args: args{origin: "null"},
			want: true,
		},
		{
			name: "file:// origin => true",
			args: args{origin: "file://domain.com"},
			want: true,
		},
		{
			name: "valid origin => true",
			args: args{origin: "https://domain.com"},
			want: true,
		},
		{
			name: "invalid origin => false",
			args: args{origin: "domain.com"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidOrigin(tt.args.origin); got != tt.want {
				t.Errorf("isValidOrigin() = %v, want %v", got, tt.want)
			}
		})
	}
}