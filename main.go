package main

import "github.com/alejmendez/goApiRest/bootstrap"

func main() {
	app := bootstrap.NewApplication()

	defer app.Close()
}
