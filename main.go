package main

import (
	"gotest_superindo/pkg"
	"gotest_superindo/route"
)

func main() {
	pkg.InitDB()

	router := route.Router()

	router.Run(":8181")
}
