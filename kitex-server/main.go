package main

import (
	example "github.com/tim-pipi/orbitaltesting/kitex-server/kitex_gen/hello/example/helloservice"
	"log"
)

func main() {
	svr := example.NewServer(new(HelloServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
