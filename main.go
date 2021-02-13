package main

import (
	"fmt"
	"post_crud_golang/database"
	"post_crud_golang/models"
	_ "time/tzdata"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello golang from docker!")
	database.Init()

	router := gin.Default()

	router.GET("/posts", func(context *gin.Context) {
		// これだと複数モデルが出てきた時に大変になるので、メソッド化するなどの工夫が必要だろう。
		posts := models.GetAll()
		context.JSON(200, gin.H{
			"posts": posts,
		})
	})
	router.POST("/posts", func(context *gin.Context) {
		title := context.PostForm("ベタ書きタイトル")
		content := context.PostForm("ベタ書きコンテンツ")
		models.Insert(title, content)
	})
	router.PUT("/posts")
	router.DELETE("/post/:id")
	router.PATCH("/post/:id")

	_ = router.Run(":8080")
}
