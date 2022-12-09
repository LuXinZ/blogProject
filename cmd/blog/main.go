package main

import "github.com/LuXinZ/blogProject/internal/blogServer"

func main() {
	application := blogServer.Application{
		Address: "8081",
	}
	application.Run()
}
