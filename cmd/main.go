package main

import (
	"golang-travel-api/internal/router"
)

func main() {
	r := router.Init()

	if err := r.Run(":9090"); err != nil {
		return
	}
}
