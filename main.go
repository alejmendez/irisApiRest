package main

import (
	"github.com/alejmendez/goApiRest/core"
)

func main() {
	app := core.GetServerInstance()
	app.Start()

	defer app.Close()
}
