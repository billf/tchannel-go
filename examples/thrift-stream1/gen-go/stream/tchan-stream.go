// @generated Code generated by thrift-gen. Do not modify.

// Package stream is generated code used to make or handle TChannel calls using Thrift.
package stream

import (
	"fmt"
	"io"

	athrift "github.com/apache/thrift/lib/go/thrift"
	"github.com/uber/tchannel-go"
	"github.com/uber/tchannel-go/thrift"
)

// Used to avoid unused warnings for non-streaming services.
var _ = tchannel.NewChannel
var _ = io.Reader(nil)

// Interfaces for the service and client for the services defined in the IDL.

// TChanTestStream is the interface that defines the server handler and client interface.
type TChanTestStream interface {
}

// TChanTestStreamServer is the interface that must be implemented by a handler.
type TChanTestStreamServer interface {
	BothStream(ctx thrift.Context, call TestStreamBothStreamInCall) error
	OutStream(ctx thrift.Context, prefix string, call TestStreamOutStreamInCall) error
}

// TChanTestStreamClient is the interface is used to make remote calls.
type TChanTestStreamClient interface {
	BothStream(ctx thrift.Context) (TestStreamBothStreamOutCall, error)
	OutStream(ctx thrift.Context, prefix string) (TestStreamOutStreamOutCall, error)
}

// TChanTestStream2 is the interface that defines the server handler and client interface.
type TChanTestStream2 interface {
	TChanTestStream
}

// TChanTestStream2Server is the interface that must be implemented by a handler.
type TChanTestStream2Server interface {
	TChanTestStreamServer

	OutStream2(ctx thrift.Context, prefix string, call TestStream2OutStream2InCall) error
}

// TChanTestStream2Client is the interface is used to make remote calls.
type TChanTestStream2Client interface {
	TChanTestStreamClient

	OutStream2(ctx thrift.Context, prefix string) (TestStream2OutStream2OutCall, error)
}

// Implementation of a client and service handler.

type tchanTestStreamClient struct {
	thriftService string
	client        thrift.TChanStreamingClient
}

func NewTChanTestStreamInheritedClient(thriftService string, client thrift.TChanStreamingClient) *tchanTestStreamClient {
	return &tchanTestStreamClient{
		thriftService,
		client,
	}
}

// NewTChanTestStreamClient creates a client that can be used to make remote calls.
func NewTChanTestStreamClient(client thrift.TChanStreamingClient) TChanTestStreamClient {
	return NewTChanTestStreamInheritedClient("TestStream", client)
}

func (c *tchanTestStreamClient) BothStream(ctx thrift.Context) (TestStreamBothStreamOutCall, error) {
	call, writer, err := c.client.StartCall(ctx, "TestStream::BothStream")
	if err != nil {
		return nil, err
	}

	outCall := &testStreamBothStreamOutCall{
		call: call,
	}

	outCall.writer = writer

	return outCall, nil
}

func (c *tchanTestStreamClient) OutStream(ctx thrift.Context, prefix string) (TestStreamOutStreamOutCall, error) {
	call, writer, err := c.client.StartCall(ctx, "TestStream::OutStream")
	if err != nil {
		return nil, err
	}

	outCall := &testStreamOutStreamOutCall{
		call: call,
	}

	args := TestStreamOutStreamArgs{
		Prefix: prefix,
	}
	if err := thrift.WriteStruct(writer, &args); err != nil {
		return nil, err
	}

	return outCall, nil
}

type tchanTestStreamServer struct {
	handler TChanTestStreamServer
}

// NewTChanTestStreamServer wraps a handler for TChanTestStreamServer so it can be
// registered with a thrift.Server.
func NewTChanTestStreamServer(handler TChanTestStreamServer) thrift.TChanStreamingServer {
	return &tchanTestStreamServer{
		handler,
	}
}

func (s *tchanTestStreamServer) Service() string {
	return "TestStream"
}

func (s *tchanTestStreamServer) Methods() []string {
	return []string{}
}

func (s *tchanTestStreamServer) Handle(ctx thrift.Context, methodName string, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	switch methodName {

	default:
		return false, nil, fmt.Errorf("method %v not found in service %v", methodName, s.Service())
	}
}

func (s *tchanTestStreamServer) StreamingMethods() []string {
	return []string{
		"BothStream",
		"OutStream",
	}
}

func (s *tchanTestStreamServer) HandleStreaming(ctx thrift.Context, call *tchannel.InboundCall) error {
	methodName := call.MethodString()
	arg3Reader, err := call.Arg3Reader()
	if err != nil {
		return err
	}
	switch methodName {
	case "TestStream::BothStream":
		return s.handleBothStream(ctx, call, arg3Reader)
	case "TestStream::OutStream":
		return s.handleOutStream(ctx, call, arg3Reader)
	}
	return fmt.Errorf("method %v not found in service %v", methodName, s.Service())
}

func (s *tchanTestStreamServer) handleBothStream(ctx thrift.Context, tcall *tchannel.InboundCall, arg3Reader io.ReadCloser) error {
	call := &testStreamBothStreamInCall{
		call: tcall,
		ctx:  ctx,
	}

	call.reader = arg3Reader

	err :=
		s.handler.BothStream(ctx, call)
	if err != nil {
		// TODO: encode any Thrift exceptions here.
		return err
	}

	if err := call.checkWriter(); err != nil {
		return err
	}

	// TODO: we may want to Close the writer if it's not already closed.

	return nil
}

func (s *tchanTestStreamServer) handleOutStream(ctx thrift.Context, tcall *tchannel.InboundCall, arg3Reader io.ReadCloser) error {
	call := &testStreamOutStreamInCall{
		call: tcall,
		ctx:  ctx,
	}

	var req TestStreamOutStreamArgs
	if err := thrift.ReadStruct(arg3Reader, &req); err != nil {
		return err
	}

	err :=
		s.handler.OutStream(ctx, req.Prefix, call)
	if err != nil {
		// TODO: encode any Thrift exceptions here.
		return err
	}

	if err := call.checkWriter(); err != nil {
		return err
	}

	// TODO: we may want to Close the writer if it's not already closed.

	return nil
}

// TestStreamBothStreamInCall is the object used to stream arguments and write
// response headers for incoming calls.
type TestStreamBothStreamInCall interface {
	// Read returns the next argument, if any is available. If there are no more arguments left,
	// it will return io.EOF.
	Read() (*SString, error)

	// SetResponseHeaders sets the response headers. This must be called before any
	// streaming responses are sent.
	SetResponseHeaders(headers map[string]string) error

	// Write writes a result to the response stream. The written items may not
	// be sent till Flush or Done is called.
	Write(arg *SString) error

	// Flush flushes headers (if they have not yet been sent) and any written results.
	Flush() error

	// Done closes the response stream and should be called after all results have been written.
	Done() error
}

// testStreamBothStreamInCall is the implementation for TestStreamBothStreamInCall.
type testStreamBothStreamInCall struct {
	call *tchannel.InboundCall
	ctx  thrift.Context

	reader io.ReadCloser

	writer tchannel.ArgWriter
}

func (c *testStreamBothStreamInCall) Read() (*SString, error) {
	var req SString
	if err := thrift.ReadStreamStruct(c.reader, func(protocol athrift.TProtocol) error {
		return req.Read(protocol)
	}); err != nil {
		return nil, err
	}

	return &req, nil
}

func (c *testStreamBothStreamInCall) SetResponseHeaders(headers map[string]string) error {
	if c.writer != nil {
		// arg3 is already being written, headers must be set first
		return fmt.Errorf("cannot set headers after writing streaming responses")
	}

	c.ctx.SetResponseHeaders(headers)
	return nil
}

func (c *testStreamBothStreamInCall) writeResponseHeaders() error {
	if c.writer != nil {
		// arg3 is already being written, headers must be set first
		return fmt.Errorf("cannot set headers after writing streaming responses")
	}

	// arg2 writer should be used to write headers
	arg2Writer, err := c.call.Response().Arg2Writer()
	if err != nil {
		return err
	}

	headers := c.ctx.ResponseHeaders()
	if err := thrift.WriteHeaders(arg2Writer, headers); err != nil {
		return err
	}

	return arg2Writer.Close()
}

// checkWriter creates the arg3 writer if it has not been created.
// Before the arg3 writer is created, response headers are sent.
func (c *testStreamBothStreamInCall) checkWriter() error {
	if c.writer == nil {
		if err := c.writeResponseHeaders(); err != nil {
			return err
		}

		writer, err := c.call.Response().Arg3Writer()
		if err != nil {
			return err
		}
		c.writer = writer
	}
	return nil
}

func (c *testStreamBothStreamInCall) Write(arg *SString) error {
	if err := c.checkWriter(); err != nil {
		return err
	}
	return thrift.WriteStreamStruct(c.writer, arg)
}

// Flush flushes headers (if they have not yet been sent) and any written results.
func (c *testStreamBothStreamInCall) Flush() error {
	if err := c.checkWriter(); err != nil {
		return err
	}
	return c.writer.Flush()
}

// Done closes the response stream and should be called after all results have been written.
func (c *testStreamBothStreamInCall) Done() error {
	if err := c.checkWriter(); err != nil {
		return err
	}
	return c.writer.Close()
}

// TestStreamBothStreamOutCall is the object used to stream arguments/results and
// read response headers for outgoing calls.
type TestStreamBothStreamOutCall interface {
	// Write writes an argument to the request stream. The written items may not
	// be sent till Flush or Done is called.
	Write(arg *SString) error

	// Flush flushes all written arguments.
	Flush() error

	// Done closes the request stream and should be called after all arguments have been written.
	Done() error

	// Read returns the next result, if any is available. If there are no more
	// results left, it will return io.EOF.
	Read() (*SString, error)

	// ResponseHeaders returns the response headers sent from the server. This will
	// block until server headers have been received.
	ResponseHeaders() (map[string]string, error)
}

// testStreamBothStreamOutCall is the implementation for TestStreamBothStreamOutCall.
type testStreamBothStreamOutCall struct {
	call            *tchannel.OutboundCall
	responseHeaders map[string]string
	reader          io.ReadCloser
	writer          tchannel.ArgWriter
}

// Write writes an argument to the request stream. The written items may not
// be sent till Flush or Done is called.
func (c *testStreamBothStreamOutCall) Write(arg *SString) error {
	return thrift.WriteStreamStruct(c.writer, arg)
}

// Flush flushes all written arguments.
func (c *testStreamBothStreamOutCall) Flush() error {
	return c.writer.Flush()
}

// Done closes the request stream and should be called after all arguments have been written.
func (c *testStreamBothStreamOutCall) Done() error {
	if err := c.writer.Close(); err != nil {
		return err
	}

	return nil
}

func (c *testStreamBothStreamOutCall) checkReader() error {
	if c.reader == nil {
		arg2Reader, err := c.call.Response().Arg2Reader()
		if err != nil {
			return err
		}

		c.responseHeaders, err = thrift.ReadHeaders(arg2Reader)
		if err != nil {
			return err
		}
		if err := arg2Reader.Close(); err != nil {
			return err
		}

		reader, err := c.call.Response().Arg3Reader()
		if err != nil {
			return err
		}

		c.reader = reader
	}
	return nil
}

func (c *testStreamBothStreamOutCall) Read() (*SString, error) {
	if err := c.checkReader(); err != nil {
		return nil, err
	}
	var res SString
	if err := thrift.ReadStreamStruct(c.reader, func(protocol athrift.TProtocol) error {
		return res.Read(protocol)
	}); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *testStreamBothStreamOutCall) ResponseHeaders() (map[string]string, error) {
	if err := c.checkReader(); err != nil {
		return nil, err
	}
	return c.responseHeaders, nil
}

// TestStreamOutStreamInCall is the object used to stream arguments and write
// response headers for incoming calls.
type TestStreamOutStreamInCall interface {

	// SetResponseHeaders sets the response headers. This must be called before any
	// streaming responses are sent.
	SetResponseHeaders(headers map[string]string) error

	// Write writes a result to the response stream. The written items may not
	// be sent till Flush or Done is called.
	Write(arg *SString) error

	// Flush flushes headers (if they have not yet been sent) and any written results.
	Flush() error

	// Done closes the response stream and should be called after all results have been written.
	Done() error
}

// testStreamOutStreamInCall is the implementation for TestStreamOutStreamInCall.
type testStreamOutStreamInCall struct {
	call *tchannel.InboundCall
	ctx  thrift.Context

	writer tchannel.ArgWriter
}

func (c *testStreamOutStreamInCall) SetResponseHeaders(headers map[string]string) error {
	if c.writer != nil {
		// arg3 is already being written, headers must be set first
		return fmt.Errorf("cannot set headers after writing streaming responses")
	}

	c.ctx.SetResponseHeaders(headers)
	return nil
}

func (c *testStreamOutStreamInCall) writeResponseHeaders() error {
	if c.writer != nil {
		// arg3 is already being written, headers must be set first
		return fmt.Errorf("cannot set headers after writing streaming responses")
	}

	// arg2 writer should be used to write headers
	arg2Writer, err := c.call.Response().Arg2Writer()
	if err != nil {
		return err
	}

	headers := c.ctx.ResponseHeaders()
	if err := thrift.WriteHeaders(arg2Writer, headers); err != nil {
		return err
	}

	return arg2Writer.Close()
}

// checkWriter creates the arg3 writer if it has not been created.
// Before the arg3 writer is created, response headers are sent.
func (c *testStreamOutStreamInCall) checkWriter() error {
	if c.writer == nil {
		if err := c.writeResponseHeaders(); err != nil {
			return err
		}

		writer, err := c.call.Response().Arg3Writer()
		if err != nil {
			return err
		}
		c.writer = writer
	}
	return nil
}

func (c *testStreamOutStreamInCall) Write(arg *SString) error {
	if err := c.checkWriter(); err != nil {
		return err
	}
	return thrift.WriteStreamStruct(c.writer, arg)
}

// Flush flushes headers (if they have not yet been sent) and any written results.
func (c *testStreamOutStreamInCall) Flush() error {
	if err := c.checkWriter(); err != nil {
		return err
	}
	return c.writer.Flush()
}

// Done closes the response stream and should be called after all results have been written.
func (c *testStreamOutStreamInCall) Done() error {
	if err := c.checkWriter(); err != nil {
		return err
	}
	return c.writer.Close()
}

// TestStreamOutStreamOutCall is the object used to stream arguments/results and
// read response headers for outgoing calls.
type TestStreamOutStreamOutCall interface {

	// Read returns the next result, if any is available. If there are no more
	// results left, it will return io.EOF.
	Read() (*SString, error)

	// ResponseHeaders returns the response headers sent from the server. This will
	// block until server headers have been received.
	ResponseHeaders() (map[string]string, error)
}

// testStreamOutStreamOutCall is the implementation for TestStreamOutStreamOutCall.
type testStreamOutStreamOutCall struct {
	call            *tchannel.OutboundCall
	responseHeaders map[string]string
	reader          io.ReadCloser
}

func (c *testStreamOutStreamOutCall) checkReader() error {
	if c.reader == nil {
		arg2Reader, err := c.call.Response().Arg2Reader()
		if err != nil {
			return err
		}

		c.responseHeaders, err = thrift.ReadHeaders(arg2Reader)
		if err != nil {
			return err
		}
		if err := arg2Reader.Close(); err != nil {
			return err
		}

		reader, err := c.call.Response().Arg3Reader()
		if err != nil {
			return err
		}

		c.reader = reader
	}
	return nil
}

func (c *testStreamOutStreamOutCall) Read() (*SString, error) {
	if err := c.checkReader(); err != nil {
		return nil, err
	}
	var res SString
	if err := thrift.ReadStreamStruct(c.reader, func(protocol athrift.TProtocol) error {
		return res.Read(protocol)
	}); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *testStreamOutStreamOutCall) ResponseHeaders() (map[string]string, error) {
	if err := c.checkReader(); err != nil {
		return nil, err
	}
	return c.responseHeaders, nil
}

type tchanTestStream2Client struct {
	TChanTestStreamClient

	thriftService string
	client        thrift.TChanStreamingClient
}

func NewTChanTestStream2InheritedClient(thriftService string, client thrift.TChanStreamingClient) *tchanTestStream2Client {
	return &tchanTestStream2Client{
		NewTChanTestStreamInheritedClient(thriftService, client),
		thriftService,
		client,
	}
}

// NewTChanTestStream2Client creates a client that can be used to make remote calls.
func NewTChanTestStream2Client(client thrift.TChanStreamingClient) TChanTestStream2Client {
	return NewTChanTestStream2InheritedClient("TestStream2", client)
}

func (c *tchanTestStream2Client) OutStream2(ctx thrift.Context, prefix string) (TestStream2OutStream2OutCall, error) {
	call, writer, err := c.client.StartCall(ctx, "TestStream2::OutStream2")
	if err != nil {
		return nil, err
	}

	outCall := &testStream2OutStream2OutCall{
		call: call,
	}

	args := TestStream2OutStream2Args{
		Prefix: prefix,
	}
	if err := thrift.WriteStruct(writer, &args); err != nil {
		return nil, err
	}

	return outCall, nil
}

type tchanTestStream2Server struct {
	thrift.TChanStreamingServer

	handler TChanTestStream2Server
}

// NewTChanTestStream2Server wraps a handler for TChanTestStream2Server so it can be
// registered with a thrift.Server.
func NewTChanTestStream2Server(handler TChanTestStream2Server) thrift.TChanStreamingServer {
	return &tchanTestStream2Server{
		NewTChanTestStreamServer(handler),
		handler,
	}
}

func (s *tchanTestStream2Server) Service() string {
	return "TestStream2"
}

func (s *tchanTestStream2Server) Methods() []string {
	return []string{

		"BothStream",
		"OutStream",
	}
}

func (s *tchanTestStream2Server) Handle(ctx thrift.Context, methodName string, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	switch methodName {

	case "BothStream":
		return s.TChanStreamingServer.Handle(ctx, methodName, protocol)
	case "OutStream":
		return s.TChanStreamingServer.Handle(ctx, methodName, protocol)
	default:
		return false, nil, fmt.Errorf("method %v not found in service %v", methodName, s.Service())
	}
}

func (s *tchanTestStream2Server) StreamingMethods() []string {
	return []string{
		"OutStream2",

		"BothStream",
		"OutStream",
	}
}

func (s *tchanTestStream2Server) HandleStreaming(ctx thrift.Context, call *tchannel.InboundCall) error {
	methodName := call.MethodString()
	arg3Reader, err := call.Arg3Reader()
	if err != nil {
		return err
	}
	switch methodName {
	case "TestStream2::OutStream2":
		return s.handleOutStream2(ctx, call, arg3Reader)
	}
	return fmt.Errorf("method %v not found in service %v", methodName, s.Service())
}

func (s *tchanTestStream2Server) handleOutStream2(ctx thrift.Context, tcall *tchannel.InboundCall, arg3Reader io.ReadCloser) error {
	call := &testStream2OutStream2InCall{
		call: tcall,
		ctx:  ctx,
	}

	var req TestStream2OutStream2Args
	if err := thrift.ReadStruct(arg3Reader, &req); err != nil {
		return err
	}

	err :=
		s.handler.OutStream2(ctx, req.Prefix, call)
	if err != nil {
		// TODO: encode any Thrift exceptions here.
		return err
	}

	if err := call.checkWriter(); err != nil {
		return err
	}

	// TODO: we may want to Close the writer if it's not already closed.

	return nil
}

// TestStream2OutStream2InCall is the object used to stream arguments and write
// response headers for incoming calls.
type TestStream2OutStream2InCall interface {

	// SetResponseHeaders sets the response headers. This must be called before any
	// streaming responses are sent.
	SetResponseHeaders(headers map[string]string) error

	// Write writes a result to the response stream. The written items may not
	// be sent till Flush or Done is called.
	Write(arg *SString) error

	// Flush flushes headers (if they have not yet been sent) and any written results.
	Flush() error

	// Done closes the response stream and should be called after all results have been written.
	Done() error
}

// testStream2OutStream2InCall is the implementation for TestStream2OutStream2InCall.
type testStream2OutStream2InCall struct {
	call *tchannel.InboundCall
	ctx  thrift.Context

	writer tchannel.ArgWriter
}

func (c *testStream2OutStream2InCall) SetResponseHeaders(headers map[string]string) error {
	if c.writer != nil {
		// arg3 is already being written, headers must be set first
		return fmt.Errorf("cannot set headers after writing streaming responses")
	}

	c.ctx.SetResponseHeaders(headers)
	return nil
}

func (c *testStream2OutStream2InCall) writeResponseHeaders() error {
	if c.writer != nil {
		// arg3 is already being written, headers must be set first
		return fmt.Errorf("cannot set headers after writing streaming responses")
	}

	// arg2 writer should be used to write headers
	arg2Writer, err := c.call.Response().Arg2Writer()
	if err != nil {
		return err
	}

	headers := c.ctx.ResponseHeaders()
	if err := thrift.WriteHeaders(arg2Writer, headers); err != nil {
		return err
	}

	return arg2Writer.Close()
}

// checkWriter creates the arg3 writer if it has not been created.
// Before the arg3 writer is created, response headers are sent.
func (c *testStream2OutStream2InCall) checkWriter() error {
	if c.writer == nil {
		if err := c.writeResponseHeaders(); err != nil {
			return err
		}

		writer, err := c.call.Response().Arg3Writer()
		if err != nil {
			return err
		}
		c.writer = writer
	}
	return nil
}

func (c *testStream2OutStream2InCall) Write(arg *SString) error {
	if err := c.checkWriter(); err != nil {
		return err
	}
	return thrift.WriteStreamStruct(c.writer, arg)
}

// Flush flushes headers (if they have not yet been sent) and any written results.
func (c *testStream2OutStream2InCall) Flush() error {
	if err := c.checkWriter(); err != nil {
		return err
	}
	return c.writer.Flush()
}

// Done closes the response stream and should be called after all results have been written.
func (c *testStream2OutStream2InCall) Done() error {
	if err := c.checkWriter(); err != nil {
		return err
	}
	return c.writer.Close()
}

// TestStream2OutStream2OutCall is the object used to stream arguments/results and
// read response headers for outgoing calls.
type TestStream2OutStream2OutCall interface {

	// Read returns the next result, if any is available. If there are no more
	// results left, it will return io.EOF.
	Read() (*SString, error)

	// ResponseHeaders returns the response headers sent from the server. This will
	// block until server headers have been received.
	ResponseHeaders() (map[string]string, error)
}

// testStream2OutStream2OutCall is the implementation for TestStream2OutStream2OutCall.
type testStream2OutStream2OutCall struct {
	call            *tchannel.OutboundCall
	responseHeaders map[string]string
	reader          io.ReadCloser
}

func (c *testStream2OutStream2OutCall) checkReader() error {
	if c.reader == nil {
		arg2Reader, err := c.call.Response().Arg2Reader()
		if err != nil {
			return err
		}

		c.responseHeaders, err = thrift.ReadHeaders(arg2Reader)
		if err != nil {
			return err
		}
		if err := arg2Reader.Close(); err != nil {
			return err
		}

		reader, err := c.call.Response().Arg3Reader()
		if err != nil {
			return err
		}

		c.reader = reader
	}
	return nil
}

func (c *testStream2OutStream2OutCall) Read() (*SString, error) {
	if err := c.checkReader(); err != nil {
		return nil, err
	}
	var res SString
	if err := thrift.ReadStreamStruct(c.reader, func(protocol athrift.TProtocol) error {
		return res.Read(protocol)
	}); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *testStream2OutStream2OutCall) ResponseHeaders() (map[string]string, error) {
	if err := c.checkReader(); err != nil {
		return nil, err
	}
	return c.responseHeaders, nil
}