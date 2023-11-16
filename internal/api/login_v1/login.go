package login_v1

import (
	"context"
	"github.com/bakhtiyor-y/auth/internal/models"
	"github.com/bakhtiyor-y/auth/internal/service/login_v1"
	pb "github.com/bakhtiyor-y/auth/pkg/login_v1"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/metadata"
	"log"
)

type Login struct {
	pb.UnimplementedLoginV1Server
	Service login_v1.ILoginService
}

func (s *Login) Login(ctx context.Context, req *pb.Login_Request) (*pb.Login_Response, error) {

	res, err := s.Service.Login(ctx, &models.AuthUser{
		Login:    req.GetLogin(),
		Password: req.GetPassword(),
	})

	if err != nil {
		return nil, err
	}

	return &pb.Login_Response{
		RefreshToken: res.Refresh,
		AccessToken:  res.Access,
	}, nil
}

func (s *Login) Check(ctx context.Context, _ *empty.Empty) (*pb.Check_Response, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Println("No metadata found in context")
	}

	tokens := md.Get("authorization")
	if len(tokens) == 0 {
		log.Println("No authorization token found in metadata")
	}

	resp, err := s.Service.Check(ctx)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
