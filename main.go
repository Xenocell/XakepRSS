package main

func main() {

	xakepParse := NewXakepParse()

	httpRouter := NewHttpRouter()
	httpRouter.InitXakepRoute(xakepParse)

	httpServer := NewHttpServer(httpRouter, ":3000")
	httpServer.Start()

}
