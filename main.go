package main

import (
	"github.com/asim/go-micro/plugins/registry/nacos/v4"
	_ "github.com/asim/go-micro/plugins/registry/nacos/v4"
	httpServer "github.com/asim/go-micro/plugins/server/http/v4"
	"github.com/gin-gonic/gin"
	"go-micro.dev/v4"
	log "go-micro.dev/v4/logger"
	"go-micro.dev/v4/server"
)

var (
	service = "helloworld"
	version = "latest"
)

func main() {
	/*	// Create service
		srv := micro.NewService(
			micro.Name(service),
			micro.Version(version),
		)
		srv.Init()

		// Register handler
		pb.RegisterHelloworldHandler(srv.Server(), new(handler.Helloworld))
		pb.RegisterGreeterHandler(srv.Server(),new(handler.Greeter))
		// Run service
		if err := srv.Run(); err != nil {
			log.Fatal(err)
		}*/
	test()
}

func test() {

	srv := httpServer.NewServer(
		server.Name("helloworld"),
		server.Address(":80"),
	)

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Recovery())

	// register router
	/*demo := newDemo()
	demo.InitRouter(router)
	*/

	group := router.Group("/a")

	group.POST("/m1", func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": "a m1 ok!"})
	})

	group.GET("/m2", func(c *gin.Context) {

		c.JSON(200, gin.H{"msg": "a m2 ok!"})
	})

	groupb := router.Group("/b")

	groupb.POST("/m1", func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": "b m1 ok!"})
	})

	groupb.Any("/m2", func(c *gin.Context) {
		log.Info("xx")
		c.JSON(200, gin.H{"msg": "b m2 ok!"})
	})

	hd := srv.NewHandler(router)
	if err := srv.Handle(hd); err != nil {
		log.Error(err)
	}

	//var s=[]string{"192.168.144.17:8888"}
	service := micro.NewService(
		micro.Server(srv),
		micro.Registry(nacos.NewRegistry()),
		//micro.Registry(nacos.NewRegistry(nacos.WithAddress([] string{"192.168.144.17:8888"}))),
	)
	service.Init()
	//pb.RegisterHelloworldHandler(service.Server(), new(handler.Helloworld))
	//pb.RegisterGreeterHandler(service.Server(), new(handler.Greeter))
	service.Run()
}

/*type demo struct{}

func newDemo() *demo {
	return &demo{}
}

func (a *demo) InitRouter(router *gin.Engine) {
	router.POST("/demo", a.demo)
}

func (a *demo) demo(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "call go-micro v3 http server success OK HTTP"})
}*/
