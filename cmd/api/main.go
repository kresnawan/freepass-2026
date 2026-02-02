package main

import (
	"canteen/internal/env"
	"canteen/internal/routes"
)

func main() {
	env.InitEnv()
	routes.InitRoute()
}
