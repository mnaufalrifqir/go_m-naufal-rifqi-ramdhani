package main

import (
	"mvc_alterra/route"
)

func main() {
	route := route.StartRoute()
	route.Start(":8000")
}
