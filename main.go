package main

import (
	"project/configs"
	"project/routes"
)

func main() {
	configs.InitDB()
	e := routes.New()
	e.Start(":8000")
}
