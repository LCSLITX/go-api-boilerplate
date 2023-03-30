package main

import (
	"github.com/lucassauro/go-api-boilerplate/utils"
	"github.com/lucassauro/go-api-boilerplate/server"
)
func main() {
	err := server.Server()
	utils.Check(err)
}