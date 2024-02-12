package main

import "receipt/internal/router"

func main() {
	r := router.Init()
	r.Run(":8080")
}
