package app

import (
	"log"
	"net"
	"net/http"
	"sync"
)

func (a *App) Run() error {
	wg := sync.WaitGroup{}
	wg.Add(3)

	go func() {
		defer wg.Done()

		log.Fatal(a.runAuthGRPC())
	}()

	go func() {
		defer wg.Done()

		log.Fatal(a.runLoginGRPC())
	}()

	go func() {
		defer wg.Done()
		log.Fatal(a.runHTTP())
	}()

	go func() {
		defer wg.Done()

		log.Fatal(a.runDocumentation())
	}()

	wg.Wait()
	return nil
}

func (a *App) runAuthGRPC() error {
	listener, err := net.Listen("tcp", a.appConfig.App.PortAuthGRPC)
	if err != nil {
		return err
	}

	log.Println("Auth GRPC server running on port:", a.appConfig.App.PortAuthGRPC)

	return a.grpcAuthServer.Serve(listener)
}

func (a *App) runLoginGRPC() error {
	listener, err := net.Listen("tcp", a.appConfig.App.PortLoginGRPC)
	if err != nil {
		return err
	}

	log.Println("Login GRPC server running on port:", a.appConfig.App.PortLoginGRPC)

	return a.grpcLoginServer.Serve(listener)
}

func (a *App) runHTTP() error {
	log.Println("HTTP server is running on port:", a.appConfig.App.PortAuthHTTP)

	return http.ListenAndServe(a.appConfig.App.PortAuthHTTP, a.muxAuth)
}

func (a *App) runDocumentation() error {
	log.Println("Swagger documentation running on port:", a.appConfig.App.PortDocs)

	return http.ListenAndServe(a.appConfig.App.PortDocs, a.reDoc.Handler())
}
