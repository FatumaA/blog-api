// Package routes Blog API.
//
//	Schemes: http
//	BasePath: /
//	Version: 1.0.0
//	Host: localhost:8080
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
// swagger:meta
package routes

// swagger:model
type NewBlog struct {
	// The title of the blog
	//
	// example: My first blog
	Title string `json:"title"`

	// The description of the blog
	//
	// example: This is my first blog
	Description string `json:"description"`

	// The body of the blog
	//
	// example: This is the body of my first blog
	Body string `json:"body"`

	// The author of the blog
	//
	// example: John Doe
	Author string `json:"author"`

	// The publication status of the blog
	//
	// example: true
	IsPublished bool `json:"isPublished"`
}

// swagger:route GET /blogs blogs getBlogs
//
// GetBlogs returns all blogs.
//
// Responses:
//
//	200: successResponse
func GetBlogs() {}

// swagger:parameters getBlog
type BlogIDParam struct {
	// The ID of the blog
	//
	// in: path
	// required: true
	ID int `json:"id"`
}

// swagger:route GET /blogs/{id} blogs getBlog
//
// GetBlog returns a blog by its ID.
//
// Responses:
//
//	200: successResponse
//	404: errorResponse
func GetBlog() {}

// swagger:parameters createBlog
type NewBlogBodyParam struct {
	// The blog data
	//
	// in: body
	// required: true
	Blog NewBlog `json:"blog"`
}

// swagger:route POST /blogs blogs createBlog
//
// CreateBlog creates a new blog and returns it.
//
// Responses:
//
//	201: successResponse
//	400: errorResponse
func CreateBlog() {}

// swagger:parameters updateBlog
type UpdateBlogParams struct {
	// The ID of the blog
	//
	// in: path
	// required: true
	// default: 1
	ID int `json:"id"`

	// The blog data
	//
	// in: body
	// required: true
	Blog NewBlog `json:"blog"`
}

// swagger:route PUT /blogs/{id} blogs updateBlog
//
// UpdateBlog updates a blog by its ID.
//
// Responses:
//
//	200: successResponse
//	400: errorResponse
//	404: errorResponse
func UpdateBlog() {}

// swagger:parameters deleteBlog
type DeleteBlogParams struct {
	// The ID of the blog
	//
	// in: path
	// required: true
	ID int `json:"id"`
}

// swagger:route DELETE /blogs/{id} blogs deleteBlog
//
// DeleteBlog deletes a blog by its ID.
//
// Responses:
//
//	200: successResponse
func DeleteBlog() {}

// A BlogResponse model
//
// swagger:response BlogResponse
type BlogResponse struct {
	// in: body
	Blog Blog `json:"blog"`
}

// A ErrorResponse model
//
// swagger:response errorResponse
type ErrorResponse struct {
	// The error message
	//
	// in: body
	// required: true
	Message string `json:"message"`
}

// A SuccessResponse model
//
// swagger:response successResponse
type SuccessResponse struct {
	// The success message
	//
	// in: body
	// required: true
	Message string `json:"message"`
}
