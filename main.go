package main

import "music-pc-server/application"

func main() {
	core := new(application.Application)
	core.Start()
}
