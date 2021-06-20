package main

import (
	"github.com/dotslashbin/pabeadle_services/srvsapp"
)

func main() {
	service := srvsapp.Gin()
	service.Run()
}
