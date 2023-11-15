package auth_v1

import (
	"context"
	"log"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (a *Auth) Healthz(_ context.Context, _ *emptypb.Empty) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

// nolint
func recoveryFunction(args ...interface{}) {
	if recoveryMessage := recover(); recoveryMessage != nil {
		log.Println(recoveryMessage)
		log.Println("!!! !!! Произошла ПАНИКА !!! !!! request:", args)
	}
}
