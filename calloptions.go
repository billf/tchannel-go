// Copyright (c) 2015 Uber Technologies, Inc.

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package tchannel

// Format is the arg scheme used for a specific call.
type Format string

// The list of formats supported by tchannel.
const (
	HTTP   Format = "http"
	JSON   Format = "json"
	Raw    Format = "raw"
	Thrift Format = "thrift"
)

func (f Format) String() string {
	return string(f)
}

// CallOptions are options for a specific call.
type CallOptions struct {
	// Format is arg scheme used for this call, sent in the "as" header.
	// This header is only set if the Format is set.
	Format Format

	// ShardKey determines where this call request belongs, used with ringpop applications.
	ShardKey string

	// RequestState stores request state across retry attempts.
	RequestState *RequestState

	// RoutingDelegate identifies a service capable of routing a request to its
	// intended recipient.
	RoutingDelegate string

	// callerName can only be used when forwarding a request. It can only be set internally,
	// e.g. by calling (*InboundCall).CallOptions() when forwarding a request
	callerName string
}

var defaultCallOptions = &CallOptions{}

func (c *CallOptions) setHeaders(headers transportHeaders) {
	headers[ArgScheme] = Raw.String()
	c.overrideHeaders(headers)
}

// overrideHeaders sets headers if the call options contains non-default values.
func (c *CallOptions) overrideHeaders(headers transportHeaders) {
	if c.Format != "" {
		headers[ArgScheme] = c.Format.String()
	}
	if c.ShardKey != "" {
		headers[ShardKey] = c.ShardKey
	}
	if c.RoutingDelegate != "" {
		headers[RoutingDelegate] = c.RoutingDelegate
	}
	if c.callerName != "" {
		headers[CallerName] = c.callerName
	}
}

// setResponseHeaders copies some headers from the incoming call request to the response.
func setResponseHeaders(reqHeaders, respHeaders transportHeaders) {
	respHeaders[ArgScheme] = reqHeaders[ArgScheme]
}
