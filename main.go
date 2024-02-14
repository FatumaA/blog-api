package main

import (
	"blog-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	blog := &routes.Blog{}

	router := gin.Default()
	router.GET("/blogs", blog.GetBlogs)
	router.GET("/blogs/:id", blog.GetBlog)
	router.POST("/blogs", blog.CreateBlog) // Convert blog.CreateBlog to gin.HandlerFunc
	router.PUT("/blogs/:id", blog.UpdateBlog)
	router.DELETE("/blogs/:id", blog.DeleteBlog)

	router.Run("localhost:8000")
}
