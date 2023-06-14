package main

import (
	"context"
	example "github.com/tim-pipi/orbitaltesting/kitex-server/kitex_gen/hello/example"
)

// HelloServiceImpl implements the last service interface defined in the IDL.
type HelloServiceImpl struct{}

// HelloMethod implements the HelloServiceImpl interface.
func (s *HelloServiceImpl) HelloMethod(ctx context.Context, request *example.HelloReq) (resp *example.HelloResp, err error) {
	// TODO: Your code here...
    return &example.HelloResp{RespBody: "Hello, " + request.Name}, nil
}
