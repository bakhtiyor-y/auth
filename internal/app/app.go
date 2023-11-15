package app

import (
	"context"
	"fmt"
	"github.com/bakhtiyor-y/auth/internal/service/login_v1"
	"log"
	"runtime"

	gateway_runtime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/jmoiron/sqlx"
	"github.com/kelseyhightower/envconfig"
	_ "github.com/lib/pq"
	"github.com/mvrilo/go-redoc"
	"google.golang.org/grpc"

	"github.com/bakhtiyor-y/auth/dev"
	"github.com/bakhtiyor-y/auth/internal/models"
	auth_system "github.com/bakhtiyor-y/auth/internal/service/auth_v1"
	login_system "github.com/bakhtiyor-y/auth/internal/service/login_v1"
	"github.com/bakhtiyor-y/auth/internal/storage/postgresql"
)

type App struct {
	appConfig *models.Config
	muxAuth   *gateway_runtime.ServeMux

	grpcAuthServer     *grpc.Server
	grpcLoginServer    *grpc.Server
	authSystemService  auth_system.IAuthSystemService
	loginSystemService login_v1.ILoginService
	reDoc              redoc.Redoc

	db *sqlx.DB
}

func NewApp(ctx context.Context) (*App, error) {
	runtime.GOMAXPROCS(runtime.NumCPU())

	a := &App{}
	a.setConfig()
	a.initDB()
	a.initReDoc()
	a.initGRPCServer()
	if err := a.initHTTPServer(ctx); err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) setConfig() {
	if dev.DEBUG {
		err := dev.SetConfig()
		if err != nil {
			log.Fatal("failed to get config:", err.Error())
		}

	}
	conf := models.Config{}

	envconfig.MustProcess("", &conf)

	a.appConfig = &conf
}

func (a *App) getAuthSystemService() auth_system.IAuthSystemService {
	storage := postgresql.NewStorage(a.db)

	if a.authSystemService == nil {
		a.authSystemService = auth_system.NewAuthSystemService(storage, a.appConfig.App.KeyForHashingPassword)
	}

	return a.authSystemService
}

func (a *App) getLoginSystemService() login_system.ILoginService {
	storage := postgresql.NewStorage(a.db)

	if a.loginSystemService == nil {
		a.loginSystemService = login_system.NewLoginSystemService(a.appConfig.App.AccessSecret, storage)
	}

	return a.loginSystemService
}
func (a *App) getSqlConnectionString() string {
	sqlConnectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%v",
		a.appConfig.DB.User,
		a.appConfig.DB.Password,
		a.appConfig.DB.Host,
		a.appConfig.DB.Port,
		a.appConfig.DB.Database,
		a.appConfig.DB.SSLMode,
	)

	return sqlConnectionString
}
