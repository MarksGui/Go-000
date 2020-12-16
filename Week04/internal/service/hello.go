package service

import (
	pb "Go-000/Week04/api/hello"
	"Go-000/Week04/internal/biz"
	"Go-000/Week04/internal/pkg/errcode"
	"context"
	"errors"

	"github.com/jinzhu/gorm"
)

type HelloService struct {
	biz *biz.HelloBIZ
}

func NewHelloService(biz *biz.HelloBIZ) *HelloService {
	return &HelloService{biz: biz}
}

func (s *HelloService) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	hello, err := s.biz.GetHello(in.Id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errcode.ErrNotFound
		}
	}
	return &pb.HelloReply{Message: hello.Name}, nil
}
