package main

import "github.com/iniyusril/template/app"

func main() {
	app.InitializedEnvirontment()
	router := InitializedServer()
	router.Logger.Fatal(router.Start(":9001"))
}
