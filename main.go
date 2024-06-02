package main

import "github.com/burp-backend/app"

func main() {
	app := app.NewApp()
	app.Init()
	app.Run()
}
