package main

import (
	"fmt"
	"post_crud_golang/database"
	_ "time/tzdata"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello golang from docker!")
	database.Init()

	router := gin.Default()

	router.GET("/posts")
	router.POST("/posts")
	router.PUT("/posts")
	router.DELETE("/post/:id")
	router.PATCH("/post/:id")

	_ = router.Run(":8080")
}
