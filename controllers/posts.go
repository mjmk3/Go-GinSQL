package controllers

import (
	"GoConn/db_client"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Post struct {
	ID        int64     `json: "id"`
	Title     string    `json: "title"`
	Content   string    `json: "content"`
	CreatedAt time.Time `json: "created_at"`
}

func CreatePost(c *gin.Context) {
	var reqBody Post
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error":   true,
			"message": "invalid request body",
		})
		return
	}

	res, _ := db_client.DBClient.Exec("INSERT INTO posts (title, content) VALUES (?, ?)",
		reqBody.Title,
		reqBody.Content,
	)

	id, _ := res.LastInsertId()
	c.JSON(http.StatusCreated, gin.H{
		"error": false,
		"id":    id,
	})
}
