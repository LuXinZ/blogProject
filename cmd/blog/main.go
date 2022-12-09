package main

import "github.com/LuXinZ/blogProject/internal/blogApp"

func main() {
	application := blogApp.Application{
		Address: "8081",
	}
	application.Run()
}
