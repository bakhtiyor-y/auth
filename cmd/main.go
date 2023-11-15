package main

import (
	"context"
	"flag"
	"log"

	"github.com/Shemistan/uzum_auth/internal/app"
)

func main() {
	flag.Parse()
	ctx := context.Background()

	a, err := app.NewApp(ctx)
	if err != nil {
		log.Fatal("failed to create app")
	}

	err = a.Run()
	if err != nil {
		log.Fatal("failed to run app")
	}
}
