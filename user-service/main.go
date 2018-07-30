package main

import (
	"log"
	"net"
	"google.golang.org/grpc"
	pb "blog-micro/user-service/proto"
)

const (
	PORT = ":50051"
)


func main()  {
	listener, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("listen on: %s\n", PORT)

	server := grpc.NewServer()
	repo := UserRepository{}
	token := TokenService{}

	// 向 rRPC 服务器注册微服务
	// 此时会把我们自己实现的微服务 service 与协议中的 ShippingServiceServer 绑定
	pb.RegisterUserServiceServer(server, &handler{repo: repo, tokenService: token})

	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
