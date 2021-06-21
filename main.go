package main

import (
	"github.com/dotslashbin/pabeadle_services/app"
)

func main() {
	service := app.Gin()
	service.Run()
}
