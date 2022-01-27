package main

import (
	"context"
	"fmt"
	httpClient "github.com/asim/go-micro/plugins/client/http/v4"
	"github.com/asim/go-micro/plugins/registry/nacos/v4"
	"go-micro.dev/v4/client"
	"go-micro.dev/v4/selector"
	"log"
	"time"
)

func main() {
	/*	srv := micro.NewService(
			micro.Name(service),
			micro.Version(version),
		)
		srv.Init()*/
	//client := pb.NewGreeterService("greeter", srv.Client())
	/*var r= &pb.Request{
		Name: "yujikuan",
	}*/
	//client.Hello()

	CallHttpServer()
}

func CallHttpServer() {
	//r:=registry.NewRegistry()
	r := nacos.NewRegistry(nacos.WithAddress([]string{"localhost:8888"}))
	//r := consul.NewRegistry(registry.Addrs("localhost:8500"))
	//r := registry.NewRegistry(registry.Addrs("localhost:8500"))
	s := selector.NewSelector(selector.Registry(r))
	// new client
	c := httpClient.NewClient(client.Selector(s))
	// create request/response
	funcName(c)
	funcName2(c)

}

func funcName(c client.Client) {
	request := c.NewRequest("helloworld", "/a/m2", "", client.WithContentType("application/json"))

	response := new(map[string]interface{})
	// call service
	err := c.Call(context.Background(), request, response)
	log.Printf("err:%v response:%#v\n", err, response)
}

func funcName2(c client.Client) {
	request := c.NewRequest("helloworld", "/b/m1", "", client.WithContentType("application/json"))
	method := request.Method()
	fmt.Println("method=" + method)
	response := new(map[string]interface{})
	// call service
	err := c.Call(context.Background(), request, response, client.WithRequestTimeout(time.Second*30), client.WithDialTimeout(time.Second*30))
	log.Printf("err:%v response:%#v\n", err, response)
}
