package main

import (
	"context"
	"log"
	"net"

	pb "github.com/ddosakura/go-micro-demo/grpc-demo/services/proto/talk" // 导入 protoc 自动生成的包
	"google.golang.org/grpc"
)

const (
	PORT = ":50001"
)

type IRepository interface {
	Create(topic *pb.Topic) (*pb.Topic, error)
	GetAll() []*pb.Topic
}

type Repository struct {
	topics []*pb.Topic
}

func (repo *Repository) Create(topic *pb.Topic) (*pb.Topic, error) {
	repo.topics = append(repo.topics, topic)
	return topic, nil
}

func (repo *Repository) GetAll() []*pb.Topic {
	return repo.topics
}

// 定义微服务，实现 topic.pb.go 中的 xxxServiceServer 接口
type service struct {
	repo Repository
}

func (s *service) CreateTalk(ctx context.Context, req *pb.Topic) (*pb.Response, error) {
	topic, err := s.repo.Create(req)
	if err != nil {
		return nil, err
	}
	resp := &pb.Response{Created: true, Topic: topic}
	return resp, nil
}

func (s *service) GetTalk(context.Context, *pb.GetRequest) (*pb.Response, error) {
	a := s.repo.GetAll()
	resp := &pb.Response{Topics: a}
	return resp, nil
}

func main() {
	listener, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("listen on: %s\n", PORT)

	server := grpc.NewServer()
	repo := Repository{}

	// 向 rRPC 服务器注册微服务
	pb.RegisterTaklkingServiceServer(
		server,
		&service{
			repo,
		},
	)

	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
