package controllers

import (
	"GoConn/db_client"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type Post struct {
	//if you like to accept a null entry just put * before data type
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

func GetPosts(c *gin.Context) {
	var posts []Post

	db_client.DBClient.Select(&posts, "SELECT id, title, conent, created_at FROM posts;") //rows (query)

	c.JSON(http.StatusOK, posts)
}

func GetPost(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	var post Post

	db_client.DBClient.Get(&post, "SELECT id, title, content, created_at FROM posts WHERE id = ?;", id) // single row (queryRow)

	c.JSON(http.StatusOK, post)
}
