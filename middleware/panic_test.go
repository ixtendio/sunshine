package middleware

import (
	"context"
	"github.com/ixtendio/gow/errors"
	"github.com/ixtendio/gow/request"
	"github.com/ixtendio/gow/response"
	"github.com/ixtendio/gow/router"
	"reflect"
	"testing"
)

func TestPanicMiddleware(t *testing.T) {
	type want struct {
		resp   response.HttpResponse
		errMsg string
	}
	tests := []struct {
		name    string
		handler router.Handler
		want
	}{
		{
			name: "error are returns",
			handler: func(ctx context.Context, r *request.HttpRequest) (response.HttpResponse, error) {
				return nil, errors.ErrWrongCredentials
			},
			want: want{
				resp:   nil,
				errMsg: errors.ErrWrongCredentials.Error(),
			},
		},
		{
			name: "panic is handled and error returned",
			handler: func(ctx context.Context, r *request.HttpRequest) (response.HttpResponse, error) {
				panic("a panic message")
			},
			want: want{
				resp:   nil,
				errMsg: "recover from panic, err: a panic message",
			},
		},
		{
			name: "error is nil",
			handler: func(ctx context.Context, r *request.HttpRequest) (response.HttpResponse, error) {
				return response.HtmlHttpResponseOK("ok"), nil
			},
			want: want{
				resp:   response.HtmlHttpResponseOK("ok"),
				errMsg: "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Panic()
			resp, err := m(tt.handler)(context.Background(), nil)
			if err != nil {
				if err.Error() != tt.want.errMsg {
					t.Errorf("Panic() = %v, want %v", err.Error(), tt.want.errMsg)
				}
			} else if !reflect.DeepEqual(tt.want.resp, resp) {
				t.Errorf("Panic() = %v, want %v", resp, tt.want.resp)
			}
		})
	}
}