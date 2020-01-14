package main

import (
	"music-pc-server/internal/app/application"
)

func main() {
	app := application.NewApplication()
	app.Init()
}
