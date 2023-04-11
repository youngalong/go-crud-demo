package main

import "go-crud-demo/router"

func main() {
	router.SetRoute().Run(":8080")
}
