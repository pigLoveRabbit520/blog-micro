package main

import (
	"log"
	"net"
	"google.golang.org/grpc"
	pb "blog-micro/user-service/proto"
	"blog-micro/user-service/config"
	"fmt"
)

func main()  {
	err := config.IniConfig()
	if err != nil {
		panic(err.Error())
	}
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", config.Conf.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return
	}
	log.Printf("listen on: %d\n", config.Conf.Port)
	db, err := CreateConnection(config.Conf)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return
	}
	server := grpc.NewServer()
	repo := &UserRepository{db:db}
	token := &TokenService{repo:repo}

	// 向 rRPC 服务器注册微服务
	// 此时会把我们自己实现的微服务 service 与协议中的 ShippingServiceServer 绑定
	pb.RegisterUserServiceServer(server, &handler{repo: repo, tokenService: token})

	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
