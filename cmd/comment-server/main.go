package main

import "github.com/LuXinZ/blogProject/internal/commentServer"

func main() {
	app := commentServer.Application{}
	app.Run()
}
