package main

import (
	"context"
	"fmt"
	"github.com/asim/go-micro/plugins/registry/nacos/v4"
	"go-micro.dev/v4"
	pbcache "helloworld/cache/proto"
	"time"
)

func main() {
	srv := micro.NewService(
		micro.Address(":7777"),

		micro.Registry(nacos.NewRegistry(nacos.WithAddress([]string{"192.168.144.17:8888"}))),
	)

	srv.Init()

	clientx := srv.Client()
	service := pbcache.NewCacheService("helloworld", clientx)
	request := &pbcache.GetRequest{}
	request.Key = "hello"
	putRequest := &pbcache.PutRequest{}
	putRequest.Key = "hello"
	putRequest.Value = "hello value"
	putRequest.Duration = time.Hour.String()
	put, err2 := service.Put(context.Background(), putRequest)
	if err2 != nil {
		panic(err2)
	}
	fmt.Println(put.String())
	response, err := service.Get(context.Background(), request)
	if err != nil {
		panic(err)
	}
	fmt.Println(response.Value + "," + response.Expiration)
}
