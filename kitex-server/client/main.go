package main

import (
    "context"
    "log"

    "github.com/tim-pipi/orbitaltesting/kitex-server/kitex_gen/hello/example/helloservice"
    "github.com/tim-pipi/orbitaltesting/kitex-server/kitex_gen/hello/example"
    "github.com/cloudwego/kitex/client"
)

func main() {
    c, err := helloservice.NewClient("kitex-server", client.WithHostPorts("0.0.0.0:8888"))
    if err != nil {
        log.Fatal(err)
    }

    req := &example.HelloReq{Name: "Timothy"}
    resp, err := c.HelloMethod(context.Background(), req)
    if err != nil {
        log.Fatal(err)
    }
    log.Println(resp)
}
