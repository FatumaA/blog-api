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

	intID, err := strconv.Atoi(id)
	if err != nil || intID < 0 || intID > len(blogs) {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Blog not found"})
		return
	}

	var blogToBeUpdated Blog
	if err := c.BindJSON(&blogToBeUpdated); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Invalid JSON format"})
		return
	}

	fmt.Println("updating blog with id: ", id)

	// blogs[intID] = blogToBeUpdated
	blogs[intID].Title = blogToBeUpdated.Title
	blogs[intID].Description = blogToBeUpdated.Description
	blogs[intID].Body = blogToBeUpdated.Body
	blogs[intID].Author = blogToBeUpdated.Author
	blogs[intID].IsPublished = blogToBeUpdated.IsPublished
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": blogToBeUpdated})
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
