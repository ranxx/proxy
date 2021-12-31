package components

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/ranxx/proxy/errors"
)

var encode = json.Marshal
var contentType = "application/json; charset=utf-8"

// EncodeHTTPGenericResponse ....
func EncodeHTTPGenericResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if response == nil {
		return responseToServer(ctx, errors.Response{Msg: "success"}, w)
	}
	return responseToServer(ctx, errors.Response{Msg: "success", Data: response}, w)
}

// ServerErrorEncoder ...
func ServerErrorEncoder(ctx context.Context, err error, w http.ResponseWriter) {
	resp := errors.Convert2Response(err)
	_ = responseToServer(ctx, resp, w)
}

func responseToServer(ctx context.Context, resp errors.Response, w http.ResponseWriter) error {
	data, err := encode(resp)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Length", fmt.Sprintf("%d", len(data)))
	w.Header().Set("Content-Type", contentType)
	_, err = w.Write(data)
	return err
}

// StripPrefix  ...
func StripPrefix(prefix string, h http.Handler) http.Handler {
	if prefix == "" {
		return h
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		p := strings.TrimPrefix(r.URL.Path, prefix)
		rp := strings.TrimPrefix(r.URL.RawPath, prefix)
		if index := strings.Index(r.URL.Path, prefix); index >= 0 {
			p = r.URL.Path[index+len(prefix):]
		}
		if index := strings.Index(r.URL.RawPath, prefix); index >= 0 {
			rp = r.URL.RawPath[index+len(prefix):]
		}
		if len(p) < len(r.URL.Path) && (r.URL.RawPath == "" || len(rp) < len(r.URL.RawPath)) {
			r2 := new(http.Request)
			*r2 = *r
			r2.URL = new(url.URL)
			*r2.URL = *r.URL
			r2.URL.Path = p
			r2.URL.RawPath = rp
			h.ServeHTTP(w, r2)
		} else {
			http.NotFound(w, r)
		}
	})
}
