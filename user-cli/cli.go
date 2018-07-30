package main

import (
	pb "blog-micro/user-service/proto"
	"io/ioutil"
	"encoding/json"
	"errors"
	"google.golang.org/grpc"
	"log"
	"os"
	"context"
)

const (
	ADDRESS           = "localhost:8080"
	DEFAULT_INFO_FILE = "user.json"
)

// 读取 user.json 中记录的用户信息
func parseFile(fileName string) (*pb.User, error) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	var user pb.User
	err = json.Unmarshal(data, &user)
	if err != nil {
		return nil, errors.New("consignment.json file content error")
	}
	return &user, nil
}

func main() {
	// 连接到 gRPC 服务器
	conn, err := grpc.Dial(ADDRESS, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("connect error: %v", err)
	}
	defer conn.Close()

	// 初始化 gRPC 客户端
	client := pb.NewUserServiceClient(conn)

	// 在命令行中指定新的货物信息 json 文件
	infoFile := DEFAULT_INFO_FILE
	if len(os.Args) > 1 {
		infoFile = os.Args[1]
	}

	// 解析货物信息
	user, err := parseFile(infoFile)
	if err != nil {
		log.Fatalf("parse info file error: %v", err)
		return
	}

	// 调用 RPC
	// 将货物存储到我们自己的仓库里
	resp, err := client.Create(context.Background(), user)
	if err != nil {
		log.Fatalf("create user error: %v", err)
		return
	}

	// 新货物是否托运成功
	log.Printf("created: %t", resp.User)
}