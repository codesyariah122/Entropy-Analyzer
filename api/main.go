// main.go
package main

import (
	"entropi-api/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		AllowCredentials: true,
	}))

	r.POST("/entropi/text", controllers.HitungEntropiText)
	r.POST("/entropi/image", controllers.HitungEntropiImage)
	r.POST("/entropi/csv", controllers.HitungEntropiCSV)
	r.GET("/entropi/pdf", controllers.GeneratePDF)

	r.Run(":8080")
}
