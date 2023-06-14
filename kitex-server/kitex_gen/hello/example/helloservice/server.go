// Code generated by Kitex v0.6.0. DO NOT EDIT.
package helloservice

import (
	server "github.com/cloudwego/kitex/server"
	example "github.com/tim-pipi/orbitaltesting/kitex-server/kitex_gen/hello/example"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler example.HelloService, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}
