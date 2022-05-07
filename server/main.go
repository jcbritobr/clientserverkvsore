package main

import (
	"github.com/jcbritobr/cstodo/server/router"
)

func main() {
	e := router.Boostrap()
	e.Run(":8080")
}
