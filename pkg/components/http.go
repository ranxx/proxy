package components

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	ggJsonPB "github.com/gogo/protobuf/jsonpb"
	ggProto "github.com/gogo/protobuf/proto"
	"github.com/ranxx/proxy/errors"
)

var encode = json.Marshal
var contentType = "application/json; charset=utf-8"

// EncodeHTTPGenericResponse ....
func EncodeHTTPGenericResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if response == nil {
		return responseToServer(ctx, errors.Response{Msg: "success", Data: struct{}{}}, w)
	}
	return responseToServer(ctx, errors.Response{Msg: "success", Data: response}, w)
}

// ServerErrorEncoder ...
func ServerErrorEncoder(ctx context.Context, err error, w http.ResponseWriter) {
	resp := errors.Convert2Response(err)
	_ = responseToServer(ctx, resp, w)
}

func responseToServer(ctx context.Context, resp errors.Response, w http.ResponseWriter) error {
	data, err := []byte{}, (error)(nil)
	if pdata, ok1 := resp.Data.(ggProto.Message); ok1 {
		marshaller := ggJsonPB.Marshaler{
			EmitDefaults: false,
			OrigName:     true,
		}
		writer := bytes.NewBuffer(make([]byte, 0, 128))
		err = marshaller.Marshal(writer, pdata)
		if err != nil {
			return err
		}
		data = writer.Bytes()
	}

	if resp.Data == nil {
		resp.Data = struct{}{}
	}

	if len(data) > 0 {
		data = []byte(fmt.Sprintf(`{"code":%d,"msg":"%s","data":%s}`, resp.Code, resp.Msg, string(data)))
	} else {
		data, err = encode(resp)
	}

	if err != nil {
		return err
	}

	w.Header().Set("Content-Length", fmt.Sprintf("%d", len(data)))
	w.Header().Set("Content-Type", contentType)
	if resp.HTTPStatus > 0 {
		w.WriteHeader(resp.HTTPStatus)
	}
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
