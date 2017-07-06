package main

import (
	"Waxt/server/env"
	"Waxt/server/serv"
)

func main() {
	en := env.InitEnvironment()

	serv.StartServer(en)
}
