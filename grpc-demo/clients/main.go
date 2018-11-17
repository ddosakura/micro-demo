package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	pb "github.com/ddosakura/go-micro-demo/grpc-demo/services/proto/talk" // 导入 protoc 自动生成的包
	"google.golang.org/grpc"
)

const (
	ADDRESS = "localhost:50001"
	data    = `{
	"description": "ABC",
	"containers": [
		{
			"data": "dbdbdbdfb"
		},
		{
			"data": "avfdkvdfv"
		}
	]
}`
)

// 解析
func parse(data string) (*pb.Topic, error) {
	var topic *pb.Topic
	err := json.Unmarshal([]byte(data), &topic)
	return topic, err
}

func main() {
	// 连接到 gRPC 服务器
	conn, err := grpc.Dial(ADDRESS, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("connect error: %v", err)
	}
	defer conn.Close()

	// 初始化 gRPC 客户端
	client := pb.NewTaklkingServiceClient(conn)

	// 解析
	topic, err := parse(data)
	if err != nil {
		fmt.Println(topic, err)
	}

	// 调用 RPC
	// 存储到仓库
	resp, err := client.CreateTalk(context.Background(), topic)
	if err != nil {
		log.Fatalf("create topic error: %v", err)
	}

	// 是否成功
	log.Printf("created: %t", resp.Created)

	// 列出所有
	resp, err = client.GetTalk(context.Background(), &pb.GetRequest{})
	if err != nil {
		log.Fatalf("failed to list topics: %v", err)
	}
	for _, c := range resp.Topics {
		log.Printf("%+v", c)
	}
}
