// Code generated by protoc-gen-goclay. DO NOT EDIT.
// source: testService.proto

/*
Package foopkg is a self-registering gRPC and JSON+Swagger service definition.

It conforms to the github.com/utrack/clay/v2/transport Service interface.
*/
package foopkg

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/go-chi/chi"
	"github.com/go-openapi/spec"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/utilities"
	"github.com/pkg/errors"
	"github.com/utrack/clay/v2/transport"
	"github.com/utrack/clay/v2/transport/httpclient"
	"github.com/utrack/clay/v2/transport/httpruntime"
	"github.com/utrack/clay/v2/transport/httpruntime/httpmw"
	"github.com/utrack/clay/v2/transport/httptransport"
	"github.com/utrack/clay/v2/transport/swagger"
	"google.golang.org/grpc"
)

// Update your shared lib or downgrade generator to v1 if there's an error
var _ = transport.IsVersion2

var _ = ioutil.Discard
var _ chi.Router
var _ runtime.Marshaler
var _ bytes.Buffer
var _ context.Context
var _ fmt.Formatter
var _ strings.Reader
var _ errors.Frame
var _ httpruntime.Marshaler
var _ http.Handler
var _ url.Values
var _ base64.Encoding
var _ httptransport.MarshalerError
var _ utilities.DoubleArray

// SampleServiceDesc is a descriptor/registrator for the SampleServiceServer.
type SampleServiceDesc struct {
	svc  SampleServiceServer
	opts httptransport.DescOptions
}

// NewSampleServiceServiceDesc creates new registrator for the SampleServiceServer.
// It implements httptransport.ConfigurableServiceDesc as well.
func NewSampleServiceServiceDesc(svc SampleServiceServer) *SampleServiceDesc {
	return &SampleServiceDesc{
		svc: svc,
	}
}

// RegisterGRPC implements service registrator interface.
func (d *SampleServiceDesc) RegisterGRPC(s *grpc.Server) {
	RegisterSampleServiceServer(s, d.svc)
}

// Apply applies passed options.
func (d *SampleServiceDesc) Apply(oo ...transport.DescOption) {
	for _, o := range oo {
		o.Apply(&d.opts)
	}
}

// SwaggerDef returns this file's Swagger definition.
func (d *SampleServiceDesc) SwaggerDef(options ...swagger.Option) (result []byte) {
	if len(options) > 0 || len(d.opts.SwaggerDefaultOpts) > 0 {
		var err error
		var s = &spec.Swagger{}
		if err = s.UnmarshalJSON(_swaggerDef_testService_proto); err != nil {
			panic("Bad swagger definition: " + err.Error())
		}

		for _, o := range d.opts.SwaggerDefaultOpts {
			o(s)
		}
		for _, o := range options {
			o(s)
		}
		if result, err = s.MarshalJSON(); err != nil {
			panic("Failed marshal spec.Swagger definition: " + err.Error())
		}
	} else {
		result = _swaggerDef_testService_proto
	}
	return result
}

// RegisterHTTP registers this service's HTTP handlers/bindings.
func (d *SampleServiceDesc) RegisterHTTP(mux transport.Router) {
	chiMux, isChi := mux.(chi.Router)

	{
		// Handler for Foo, binding: GET /v1/sample/foo/{a}
		var h http.HandlerFunc
		h = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer r.Body.Close()

			unmFunc := unmarshaler_goclay_SampleService_Foo_0(r)
			rsp, err := _SampleService_Foo_Handler(d.svc, r.Context(), unmFunc, d.opts.UnaryInterceptor)

			if err != nil {
				if err, ok := err.(httptransport.MarshalerError); ok {
					httpruntime.SetError(r.Context(), r, w, errors.Wrap(err.Err, "couldn't parse request"))
					return
				}
				httpruntime.SetError(r.Context(), r, w, err)
				return
			}

			if ctxErr := r.Context().Err(); ctxErr != nil && ctxErr == context.Canceled {
				w.WriteHeader(499) // Client Closed Request
				return
			}

			_, outbound := httpruntime.MarshalerForRequest(r)
			w.Header().Set("Content-Type", outbound.ContentType())
			err = outbound.Marshal(w, rsp)
			if err != nil {
				httpruntime.SetError(r.Context(), r, w, errors.Wrap(err, "couldn't write response"))
				return
			}
		})

		h = httpmw.DefaultChain(h)

		if isChi {
			chiMux.Method("GET", pattern_goclay_SampleService_Foo_0, h)
		} else {
			panic("query URI params supported only for chi.Router")
		}
	}

	{
		// Handler for Bar, binding: POST /v1/sample/bar
		var h http.HandlerFunc
		h = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer r.Body.Close()

			unmFunc := unmarshaler_goclay_SampleService_Bar_0(r)
			rsp, err := _SampleService_Bar_Handler(d.svc, r.Context(), unmFunc, d.opts.UnaryInterceptor)

			if err != nil {
				if err, ok := err.(httptransport.MarshalerError); ok {
					httpruntime.SetError(r.Context(), r, w, errors.Wrap(err.Err, "couldn't parse request"))
					return
				}
				httpruntime.SetError(r.Context(), r, w, err)
				return
			}

			if ctxErr := r.Context().Err(); ctxErr != nil && ctxErr == context.Canceled {
				w.WriteHeader(499) // Client Closed Request
				return
			}

			_, outbound := httpruntime.MarshalerForRequest(r)
			w.Header().Set("Content-Type", outbound.ContentType())
			err = outbound.Marshal(w, rsp)
			if err != nil {
				httpruntime.SetError(r.Context(), r, w, errors.Wrap(err, "couldn't write response"))
				return
			}
		})

		h = httpmw.DefaultChain(h)

		if isChi {
			chiMux.Method("POST", pattern_goclay_SampleService_Bar_0, h)
		} else {
			mux.Handle(pattern_goclay_SampleService_Bar_0, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.Method != "POST" {
					w.WriteHeader(http.StatusMethodNotAllowed)
					return
				}
				h(w, r)
			}))
		}
	}

}

type SampleService_httpClient struct {
	c    *http.Client
	host string
}

// NewSampleServiceHTTPClient creates new HTTP client for SampleServiceServer.
// Pass addr in format "http://host[:port]".
func NewSampleServiceHTTPClient(c *http.Client, addr string) *SampleService_httpClient {
	if strings.HasSuffix(addr, "/") {
		addr = addr[:len(addr)-1]
	}
	return &SampleService_httpClient{c: c, host: addr}
}

func (c *SampleService_httpClient) Foo(ctx context.Context, in *FooRequest, opts ...grpc.CallOption) (*FooResponse, error) {
	mw, err := httpclient.NewMiddlewareGRPC(opts)
	if err != nil {
		return nil, err
	}

	path := pattern_goclay_SampleService_Foo_0_builder(in)

	buf := bytes.NewBuffer(nil)

	m := httpruntime.DefaultMarshaler(nil)

	req, err := http.NewRequest("GET", c.host+path, buf)
	if err != nil {
		return nil, errors.Wrap(err, "can't initiate HTTP request")
	}
	req = req.WithContext(ctx)

	req.Header.Add("Accept", m.ContentType())

	req, err = mw.ProcessRequest(req)
	if err != nil {
		return nil, err
	}
	rsp, err := c.c.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "error from client")
	}
	defer rsp.Body.Close()

	rsp, err = mw.ProcessResponse(rsp)
	if err != nil {
		return nil, err
	}

	if rsp.StatusCode >= 400 {
		b, _ := ioutil.ReadAll(rsp.Body)
		return nil, errors.Errorf("%v %v: server returned HTTP %v: '%v'", req.Method, req.URL.String(), rsp.StatusCode, string(b))
	}

	ret := FooResponse{}

	err = m.Unmarshal(rsp.Body, &ret)

	return &ret, errors.Wrap(err, "can't unmarshal response")
}

func (c *SampleService_httpClient) Bar(ctx context.Context, in *BarRequest, opts ...grpc.CallOption) (*BarResponse, error) {
	mw, err := httpclient.NewMiddlewareGRPC(opts)
	if err != nil {
		return nil, err
	}

	path := pattern_goclay_SampleService_Bar_0_builder(in)

	buf := bytes.NewBuffer(nil)

	m := httpruntime.DefaultMarshaler(nil)

	if err = m.Marshal(buf, in); err != nil {
		return nil, errors.Wrap(err, "can't marshal request")
	}

	req, err := http.NewRequest("POST", c.host+path, buf)
	if err != nil {
		return nil, errors.Wrap(err, "can't initiate HTTP request")
	}
	req = req.WithContext(ctx)

	req.Header.Add("Accept", m.ContentType())

	req, err = mw.ProcessRequest(req)
	if err != nil {
		return nil, err
	}
	rsp, err := c.c.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "error from client")
	}
	defer rsp.Body.Close()

	rsp, err = mw.ProcessResponse(rsp)
	if err != nil {
		return nil, err
	}

	if rsp.StatusCode >= 400 {
		b, _ := ioutil.ReadAll(rsp.Body)
		return nil, errors.Errorf("%v %v: server returned HTTP %v: '%v'", req.Method, req.URL.String(), rsp.StatusCode, string(b))
	}

	ret := BarResponse{}

	err = m.Unmarshal(rsp.Body, &ret)

	return &ret, errors.Wrap(err, "can't unmarshal response")
}

// patterns for SampleService
var (
	pattern_goclay_SampleService_Foo_0 = "/v1/sample/foo/{a}"

	pattern_goclay_SampleService_Foo_0_builder = func(in *FooRequest) string {
		values := url.Values{}

		u := url.URL{
			Path:     fmt.Sprintf("/v1/sample/foo/%v", in.A),
			RawQuery: values.Encode(),
		}
		return u.String()
	}

	unmarshaler_goclay_SampleService_Foo_0_boundParams = &utilities.DoubleArray{Encoding: map[string]int{"a": 0}, Base: []int{1, 1, 0}, Check: []int{0, 1, 2}}

	pattern_goclay_SampleService_Bar_0 = "/v1/sample/bar"

	pattern_goclay_SampleService_Bar_0_builder = func(in *BarRequest) string {
		values := url.Values{}

		u := url.URL{
			Path:     fmt.Sprintf("/v1/sample/bar"),
			RawQuery: values.Encode(),
		}
		return u.String()
	}

	unmarshaler_goclay_SampleService_Bar_0_boundParams = &utilities.DoubleArray{Encoding: map[string]int{"": 0}, Base: []int{1, 1, 0}, Check: []int{0, 1, 2}}
)

// marshalers for SampleService
var (
	unmarshaler_goclay_SampleService_Foo_0 = func(r *http.Request) func(interface{}) error {
		return func(rif interface{}) error {
			req := rif.(*FooRequest)

			if err := errors.Wrap(runtime.PopulateQueryParameters(req, r.URL.Query(), unmarshaler_goclay_SampleService_Foo_0_boundParams), "couldn't populate query parameters"); err != nil {
				return httpruntime.TransformUnmarshalerError(err)
			}

			rctx := chi.RouteContext(r.Context())
			if rctx == nil {
				panic("Only chi router is supported for GETs atm")
			}
			for pos, k := range rctx.URLParams.Keys {
				if err := errors.Wrapf(runtime.PopulateFieldFromPath(req, k, rctx.URLParams.Values[pos]), "can't read '%v' from path", k); err != nil {
					return httptransport.NewMarshalerError(httpruntime.TransformUnmarshalerError(err))
				}
			}

			return nil
		}
	}

	unmarshaler_goclay_SampleService_Bar_0 = func(r *http.Request) func(interface{}) error {
		return func(rif interface{}) error {
			req := rif.(*BarRequest)

			if err := errors.Wrap(runtime.PopulateQueryParameters(req, r.URL.Query(), unmarshaler_goclay_SampleService_Bar_0_boundParams), "couldn't populate query parameters"); err != nil {
				return httpruntime.TransformUnmarshalerError(err)
			}

			inbound, _ := httpruntime.MarshalerForRequest(r)
			if err := errors.Wrap(inbound.Unmarshal(r.Body, &req), "couldn't read request JSON"); err != nil {
				return httptransport.NewMarshalerError(httpruntime.TransformUnmarshalerError(err))
			}
			return nil
		}
	}
)

var _swaggerDef_testService_proto = []byte(`{
  "swagger": "2.0",
  "info": {
    "title": "testService.proto",
    "version": "version not set"
  },
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "paths": {
    "/v1/sample/bar": {
      "post": {
        "operationId": "SampleService_Bar",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/foopkgBarResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/foopkgBarRequest"
            }
          }
        ],
        "tags": [
          "SampleService"
        ]
      }
    },
    "/v1/sample/foo/{a}": {
      "get": {
        "operationId": "SampleService_Foo",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/foopkgFooResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "a",
            "in": "path",
            "required": true,
            "type": "string",
            "format": "int64"
          }
        ],
        "tags": [
          "SampleService"
        ]
      }
    }
  },
  "definitions": {
    "foopkgBarRequest": {
      "type": "object",
      "properties": {
        "a": {
          "type": "string",
          "format": "int64"
        }
      }
    },
    "foopkgBarResponse": {
      "type": "object",
      "properties": {
        "sum": {
          "type": "string",
          "format": "int64"
        },
        "error": {
          "type": "string"
        }
      }
    },
    "foopkgFooResponse": {
      "type": "object",
      "properties": {
        "sum": {
          "type": "string",
          "format": "int64"
        },
        "error": {
          "type": "string"
        }
      }
    }
  }
}

`)
