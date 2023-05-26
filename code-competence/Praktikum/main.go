package main

import (
	"code_competence/route"
)

func main() {
	route := route.StartRoute()
	route.Start(":8000")
}
