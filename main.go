package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg"
)

func main() {
	db := pg.Connect(&pg.Options{
		User:     "postgres",
		Password: "postgres",
		Database: "test",
	})

	r := gin.Default()
	ginConfig := cors.DefaultConfig()
	ginConfig.AllowOrigins = []string{"*"}
	ginConfig.AllowCredentials = true
	ginConfig.AllowHeaders = []string{
		"Access-Control-Allow-Origin",
		"Origin",
		"Content-Type",
		"Content-Length",
		"Accept-Encoding",
		"X-CSRF-Token",
		"Authorization",
		"Referer",
		"X-Size",
	}
	r.Use(cors.New(ginConfig))

	group := r.Group("api/todos")

	group.GET("", func(c *gin.Context) {
		tasks := make([]Task, 0)
		_ = db.Model(&tasks).Select()

		c.JSON(200, tasks)
	})

	group.POST("", func(c *gin.Context) {
		var task Task
		c.Bind(&task)

		task.CreatedAt = time.Now()

		_ = db.Insert(&task)

		c.JSON(200, task)
	})

	r.Run("0.0.0.0:8000")
}

type Task struct {
	ID        int       `json:"id"`
	Name      string    `json:"name" binding:"required"`
	CreatedAt time.Time `json:"createdAt"`
}
