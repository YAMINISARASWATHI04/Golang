package routes

import (	
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/blog", getBlog) // This is an endpoint for GET request
	server.POST("/blog", createBlog)
	server.GET("/blog/:id", getBlogbyId)
	server.DELETE("/blog/:id", deleteBlog)
	server.GET("/blog/count", NumbeofBlogs)
	server.PUT("/blog/:id", updateBlog)
}