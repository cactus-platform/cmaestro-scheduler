package main

import (
	"context"
	"fmt"
	"log"

	"github.com/cactus-platform/scheduler/internal/bootstrap"
)

func main() {
	ctx := context.Background()
	app, err := bootstrap.NewFromEnv(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("App:", app)
}
