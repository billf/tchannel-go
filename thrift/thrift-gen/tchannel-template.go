package main

var tchannelTmpl = `
// @generated Code generated by thrift-gen. Do not modify.

// Package {{ .Package }} is generated code used to make or handle TChannel calls using Thrift.
package {{ .Package }}

import (
"fmt"

athrift "{{ .Imports.Thrift }}"
"{{ .Imports.TChannel }}"

{{ range .Includes }}
	"{{ .Import }}"
{{ end }}
)


{{ range .Includes }}
	var _ = {{ .Package }}.GoUnusedProtection__
{{ end }}


// Interfaces for the service and client for the services defined in the IDL.

{{ range .Services }}
// {{ .Interface }} is the interface that defines the server handler and client interface.
type {{ .Interface }} interface {
	{{ if .HasExtends }}
		{{ .ExtendsServicePrefix }}{{ .ExtendsService.Interface }}

	{{ end }}
	{{ range .Methods }}
		{{ .Name }}({{ .ArgList }}) {{ .RetType }}
	{{ end }}
}
{{ end }}

// Implementation of a client and service handler.

{{/* Generate client and service implementations for the above interfaces. */}}
{{ range $svc := .Services }}
type {{ .ClientStruct }} struct {
	{{ if .HasExtends }}
		{{ .ExtendsServicePrefix }}{{ .ExtendsService.Interface }}

	{{ end }}
	thriftService string
	client        thrift.TChanClient
}


func {{ .InheritedClientConstructor }}(thriftService string, client thrift.TChanClient) *{{ .ClientStruct }} {
	return &{{ .ClientStruct }}{
		{{ if .HasExtends }}
			{{ .ExtendsServicePrefix }}{{ .ExtendsService.InheritedClientConstructor }}(thriftService, client),
		{{ end }}
		thriftService,
		client,
	}
}

// {{ .ClientConstructor }} creates a client that can be used to make remote calls.
func {{ .ClientConstructor }}(client thrift.TChanClient) {{ .Interface }} {
	return {{ .InheritedClientConstructor }}("{{ .ThriftName }}", client)
}

{{ range .Methods }}
	func (c *{{ $svc.ClientStruct }}) {{ .Name }}({{ .ArgList }}) {{ .RetType }} {
		var resp {{ .ResultType }}
		args := {{ .ArgsType }}{
			{{ range .Arguments }}
				{{ .ArgStructName }}: {{ .Name }},
			{{ end }}
		}
		success, err := c.client.Call(ctx, c.thriftService, "{{ .ThriftName }}", &args, &resp)
		if err == nil && !success {
			{{ range .Exceptions }}
				if e := resp.{{ .ArgStructName }}; e != nil {
					err = e
				}
			{{ end }}
		}

		{{ if .HasReturn }}
			return resp.GetSuccess(), err
		{{ else }}
			return err
		{{ end }}
	}
{{ end }}

type {{ .ServerStruct }} struct {
	{{ if .HasExtends }}
		thrift.TChanServer

	{{ end }}
	handler {{ .Interface }}

	interceptorRunner thrift.InterceptorRunner
}

// {{ .ServerConstructor }} wraps a handler for {{ .Interface }} so it can be
// registered with a thrift.Server.
func {{ .ServerConstructor }}(handler {{ .Interface }}) thrift.TChanServer {
	return &{{ .ServerStruct }}{
		{{ if .HasExtends }}
			TChanServer: {{ .ExtendsServicePrefix }}{{ .ExtendsService.ServerConstructor }}(handler),
		{{ end }}
		handler: handler,
	}
}

func (s *{{ .ServerStruct }}) Service() string {
	return "{{ .ThriftName }}"
}

func (s *{{ .ServerStruct }}) Methods() []string {
	return []string{
		{{ range .Methods }}
			"{{ .ThriftName }}", 
		{{ end }}
		{{ range .InheritedMethods }}
			"{{ . }}",
		{{ end }}
	}
}

// RegisterInterceptors registers the provided interceptors with the server.
func (s *{{ .ServerStruct }}) RegisterInterceptorRunner(runner thrift.InterceptorRunner) {
	s.interceptorRunner = runner
}

func (s *{{ .ServerStruct }}) Handle(ctx {{ contextType }}, methodName string, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	switch methodName {
		{{ range .Methods }}
			case "{{ .ThriftName }}":
				return s.{{ .HandleFunc }}(ctx, protocol)
		{{ end }}
		{{ range .InheritedMethods }}
			case "{{ . }}":
				return s.TChanServer.Handle(ctx, methodName, protocol)
		{{ end }}
		default:
			return false, nil, fmt.Errorf("method %v not found in service %v", methodName, s.Service())
	}
}

{{ range .Methods }}
	func (s *{{ $svc.ServerStruct }}) {{ .HandleFunc }}(ctx {{ contextType }}, protocol athrift.TProtocol) (handled bool, resp athrift.TStruct, err error) {
		var req {{ .ArgsType }}
		var res {{ .ResultType }}
		serviceMethod := "{{ .ThriftName }}::{{ .Name }}"

		if readErr := req.Read(protocol); readErr != nil {
			return false, nil, readErr
		}

		postRun, err := s.interceptorRunner.RunPre(ctx, serviceMethod, &req)

		defer func () {
			err = postRun(resp, err)
			if err != nil {
				resp = nil
				{{ if .HasExceptions }}
				switch v := err.(type) {
					{{ range .Exceptions }}
					case {{ .ArgType }}:
						if v == nil {
							err = fmt.Errorf("Handler for {{ .Name }} returned non-nil error type {{ .ArgType }} but nil value")
						} else {
							res.{{ .ArgStructName }} = v
							err = nil
							resp = &res
						}
					{{ end }}
				}
				{{ end }}
			}
		}()

		if err != nil {
			return false, nil, err
		}

		{{ if .HasReturn }}
		r, err :=
		{{ else }}
		err =
		{{ end }}
			s.handler.{{ .Name }}({{ .CallList "req" }})

		{{ if .HasReturn }}
		if err == nil {
			res.Success = {{ .WrapResult "r" }}
		}
		{{ end }}

		return err == nil, &res, err
	}
{{ end }}

{{ end }}
`
