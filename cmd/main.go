package main

import "go-core-study/internal/di"

func main() {
	app := di.NewAppContainer()
	app.Run()
}
