package main

import "github.com/dihanto/gosnap/internal/app/router"

func main() {
	e := router.NewRouter()
	e.Start(":8000")
}
