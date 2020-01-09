package main

import "music-pc-server/application"

func main() {
	core := application.NewApplication()
	core.Start()
}
