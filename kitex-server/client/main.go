package client

import (
    "context"
    "log"
    "time"

    "github.com/tim-pipi/orbitaltesting/kitex-server/kitex_gen/hello/example"
    "github.com/tim-pipi/orbitaltesting/kitex-server/kitex_gen/hello"
    "github.com/cloudwego/kitex/client"
)

func main() {
    c, err := example.NewClient("kitex-server", client.WithHostPorts("0.0.0.0:8888"))
    if err != nil {
        log.Fatal(err)
    }

    req := &hello.HelloRequest{Name: "Timothy"}
    resp, err := c.Echo(context.Background(), req, callopt.WithRPCTimeout(3*time.Second))
    if err != nil {
        log.Fatal(err)
    }
    log.Println(resp)
}
