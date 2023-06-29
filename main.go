package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Run()
}
