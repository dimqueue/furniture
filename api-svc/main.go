package main

import (
	"github.com/dmytroserhiienko02/furniture/api-svc/internal/cli"
)

// @title Furniture API
// @version 1.0
// @description Furniture API

// @host localhost:8080
// @BasePath /
func main() {
	if err := cli.Run(); err != nil {
		panic(err)
	}
}
