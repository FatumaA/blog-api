package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Blog struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Body        string `json:"body"`
	Author      string `json:"author"`
	IsPublished bool   `json:"isPublished"`
}

var blogs = []Blog{
	{
		ID:          1,
		Title:       "My first blog",
		Description: "This is my first blog",
		Body:        "This is the body of my first blog",
		Author:      "John Doe",
		IsPublished: true,
	},
	{
		ID:          2,
		Title:       "My second blog",
		Description: "This is my second blog",
		Body:        "This is the body of my second blog",
		Author:      "Jane Doe",
		IsPublished: false,
	},
	{
		ID:          3,
		Title:       "My third blog",
		Description: "This is my third blog",
		Body:        "This is the body of my third blog",
		Author:      "John Doe",
		IsPublished: true,
	},
	{
		ID:          4,
		Title:       "My fourth blog",
		Description: "This is my fourth blog",
		Body:        "This is the body of my fourth blog",
		Author:      "Jane Doe",
		IsPublished: true,
	},
}

// GetBlogs returns all blogs
func (b *Blog) GetBlogs(c *gin.Context) {
	fmt.Printf("getting blogs")
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": blogs})
}

// GetBlog returns a single blog
func (b *Blog) GetBlog(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("getting blog with id: ", id)

	intID, err := strconv.Atoi(id)
	if err != nil || intID < 0 || intID >= len(blogs) {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Something went wrong" + err.Error()})
	}

	//find blog whose id matched the param id
	for _, blog := range blogs {
		if blog.ID == intID {
			c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": blog})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Blog not found"})
}

// CreateBlog creates a new blog
func (b *Blog) CreateBlog(c *gin.Context) {
	var incomingBlog Blog

	incomingBlog.ID = len(blogs) + 1
	c.BindJSON(&incomingBlog)
	blogs = append(blogs, incomingBlog)

	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "data": incomingBlog})
}

// UpdateBlog updates a blog
func (b *Blog) UpdateBlog(c *gin.Context) {
	id := c.Param("id")

	// Convert the ID to an integer
	intID, _ := strconv.Atoi(id)
	// if blogID == -1 {
	// 	c.JSON(400, gin.H{"error": "Invalid ID"})
	// 	return
	// }

	// Find the blog with the matching ID
	// index := findBlogIndex(intID)

	for index, blog := range blogs {
		if blog.ID == intID {
			// Parse the request body to get the updated blog data
			var updatedBlog Blog
			if err := c.BindJSON(&updatedBlog); err != nil {
				c.JSON(400, gin.H{"error": "Invalid JSON"})
				return
			}

			// Update the blog with the new data
			updatedBlog.ID = intID
			blogs[index] = updatedBlog

			// Respond with the updated blog
			c.JSON(200, updatedBlog)
			c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": updatedBlog})
			return
		}
	}
	// if index == -1 {
	// 	c.JSON(404, gin.H{"error": "Blog not found"})
	// 	return
	// }
}

// DeleteBlog deletes a blog
func (b *Blog) DeleteBlog(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("deleting blog with id: ", id)

	intID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Something went wrong" + err.Error()})
	}
	for index, blog := range blogs {
		if blog.ID == intID {
			blogs = append(blogs[:index], blogs[index+1:]...)
			c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Blog deleted successfully"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Blog could not be deleted. Blog not found"})
}
