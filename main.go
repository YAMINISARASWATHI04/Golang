package main

import (
	"RestApiProject/models"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func main() {
	fmt.Println("Hello, World!")
	server := gin.Default() // It will setup an engine for us and return it. The default engine is a Logger and Recovery middleware already attached.
	// Handler for GET request to the root path
	server.GET("/blog", getBlog) // This is an endpoint for GET request
	server.POST("/blog", createBlog)
	server.GET("/blog/:id", getBlogbyId)
	server.DELETE("/blog/:id", deleteBlog)
	server.GET("/blog/count", NumbeofBlogs)
	server.PUT("/blog/:id", updateBlog)
	// server.GET("/blog", getBlog)
	server.Run(":3000") // listen and serve on
}

// In this gin.Context already has request and response objects
func getBlog(context *gin.Context) { //context parameter will be set by gin and will be pointer to gin
	// Reach out to a database
	// Write in a file
	// Make an API call to another service
	blogs := models.GetAllBlogs()
	context.JSON(http.StatusOK, blogs) // To send the response
	// gin.H is a shortcut for map[string]interface{} which is a map with string keys and values
}

func createBlog(context *gin.Context) {
	var blog models.Blog

	err := context.ShouldBindJSON(&blog) // parse the JSON data from an HTTP request body and bind it to a Go struct or object, validating the data based on struct tags in the process
	// we need to pass a pointer to the blog to modify data in the blog variable

	if (blog.Author==" " || blog.Title==" " || blog.Content==" "){
		context.JSON(http.StatusBadRequest,gin.H{"message":"Please fill all the fields"})
	}
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "couldnt parse request",
			"error":   err.Error(),
		})
		return
	}
	if (blog.Author=="" || blog.Title=="" || blog.Content==""){
		context.JSON(http.StatusBadRequest,gin.H{"message":"Please fill all the fields"})
		return
	}
	// blogs := models.GetAllBlogs()
	blog.ID = uuid.New().String()

	blog.CreatedAt = time.Now()
	blog.UpdatedAt = time.Now()

	if err := models.SaveBlogs(blog); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to save blog",
			"error":   err.Error(),
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Blog created successfully", "blog": blog})

}
func NumbeofBlogs(context *gin.Context) {
	blogs := models.GetAllBlogs()
	context.JSON(http.StatusOK, gin.H{"number_of_blogs": len(blogs)})

}

func getBlogbyId(context *gin.Context) {
	id := context.Param("id")

	blog, err := models.GetBlogByID(id)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"message": "blog not found",
			"error":   err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, blog)
}

func deleteBlog(context *gin.Context) {
	id := context.Param("id")

	err := models.DeleteBlogByID(id)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"message": "blog not found",
			"error":   err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Blog deleted successfully"})
}

func updateBlog(context *gin.Context) {
	id := context.Param("id")

	var updatedBlog models.Blog
	err := context.ShouldBindJSON(&updatedBlog)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "couldnt parse request",
			"error":   err.Error(),
		})
		return
	}

	err = models.UpdateBlogByID(id, updatedBlog)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"message": "blog not found",
			"error":   err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Blog updated successfully", "blog": updatedBlog})
}
