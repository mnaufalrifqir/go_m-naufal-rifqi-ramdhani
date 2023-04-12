package main

import (
	"docker/route"
)

func main() {
	route := route.StartRoute()
	route.Start(":8000")
}
