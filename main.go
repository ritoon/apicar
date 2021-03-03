package main

import (
	"apicar/service"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	service.Init(r)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
