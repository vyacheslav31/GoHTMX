package main

import (
	"os"

	"vyacheslav31/gohtmx/internal/app"
)

func main() {
	env := os.Getenv("ENV")
	app.NewApp(env).Start()
}
