package app

import (
	"context"
	"log"

	gateway_runtime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/bakhtiyor-y/auth/docs"
	auth_system_v1 "github.com/bakhtiyor-y/auth/internal/api/auth_v1"
	login_system_v1 "github.com/bakhtiyor-y/auth/internal/api/login_v1"

	pb_auth "github.com/bakhtiyor-y/auth/pkg/auth_v1"
	pb_login "github.com/bakhtiyor-y/auth/pkg/login_v1"
)

func (a *App) initDB() {
	sqlConnectionString := a.getSqlConnectionString()
	var err error
	a.db, err = sqlx.Open("postgres", sqlConnectionString)
	if err != nil {
		log.Fatal("failed to opening connection to db: ", err.Error())
	}

	// Проверка соединения с базой данных
	if err = a.db.Ping(); err != nil {
		log.Fatal("failed to connect to the database: ", err.Error())
	}
}

func (a *App) initGRPCServer() {
	a.grpcAuthServer = grpc.NewServer()
	pb_auth.RegisterAuthV1Server(
		a.grpcAuthServer,
		&auth_system_v1.Auth{
			AuthService: a.getAuthSystemService(),
		},
	)

	a.grpcLoginServer = grpc.NewServer()
	pb_login.RegisterLoginV1Server(
		a.grpcLoginServer,
		&login_system_v1.Login{
			Service: a.getLoginSystemService(),
		})
}

func (a *App) initHTTPServer(ctx context.Context) error {
	a.muxAuth = gateway_runtime.NewServeMux()
	optsAuth := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	err := pb_auth.RegisterAuthV1HandlerFromEndpoint(ctx, a.muxAuth, a.appConfig.App.PortAuthGRPC, optsAuth)
	if err != nil {
		return err
	}
	return nil
}

func (a *App) initReDoc() {
	a.reDoc = docs.Initialize()
}
