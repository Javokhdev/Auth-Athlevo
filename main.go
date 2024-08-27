package main

import (
	"auth-athlevo/config"
	"auth-athlevo/internal/app"
)

func main() {
	cfg := config.Load()
	app.Run(&cfg)
}
