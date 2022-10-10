package main

import (
	"github.com/alejmendez/goApiRest/core"
)

func main() {
	app, _ := core.NewServer()

	defer app.Close()
}
