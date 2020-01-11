package main

import (
	"music-pc-server/internal/app/Application"
)

func main() {
	app := Application.NewApplication()
	app.Init()
}
