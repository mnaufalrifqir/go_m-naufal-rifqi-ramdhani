package main

import "MVC/route"

func main() {
	route := route.StartRoute()
	route.Start(":8000")
}
