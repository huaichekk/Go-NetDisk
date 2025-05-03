package main

import (
	"github.com/gin-gonic/gin"
	"go-netdisk/initialize"
	"log"
)

func main() {
	r := gin.Default()
	initialize.InitEngine(r)
	log.Fatalln(r.Run(":8080"))
}
