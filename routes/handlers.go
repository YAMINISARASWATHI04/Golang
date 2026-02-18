package routes
import (	
	"net/http"
	"RestApiProject/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"time"

)

func getBlog(context *gin.Context) { 
	blogs ,err := models.GetBlogs()
	if err != nil {		
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to get blogs",
			"error":   err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, blogs) 
	// gin.H is a shortcut for map[string]interface{} which is a map with string keys and values
}

func createBlog(context *gin.Context) {
	var blog models.Blog

	err := context.ShouldBindJSON(&blog) 

	if blog.Author == " " || blog.Title == " " || blog.Content == " " {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Please fill all the fields"})
	}
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "couldnt parse request",
			"error":   err.Error(),
		})
		return
	}
	if blog.Author == "" || blog.Title == "" || blog.Content == "" {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Please fill all the fields"})
		return
	}
	// blogs := models.GetAllBlogs()
	// To generate a unique identifier for the blog post using the UUID package. This ensures that each blog post has a unique ID, which can be used to retrieve, update, or delete the post later.
	blog.ID = uuid.New().String()

	blog.CreatedAt = time.Now()
	blog.UpdatedAt = time.Now()

	if err := models.SavetheBlogs(blog); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to save blog",
			"error":   err.Error(),
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{ "blog": blog})

}
func NumbeofBlogs(context *gin.Context) {
	blogs := models.GetAllBlogs()
	context.JSON(http.StatusOK, gin.H{"number_of_blogs": len(blogs)})

}

func getBlogbyId(context *gin.Context) {
	id := context.Param("id") // To extract the value of the id parameter from the URL path. The Param method takes the name of the parameter as an argument and returns its value as a string.

	blog, err := models.GetSingleBlogByID(id)
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

	err := models.DeleteBlog(id)
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

	if updatedBlog.Author == "" || updatedBlog.Title == "" || updatedBlog.Content == "" {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Please fill all the fields"})
		return
	}
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "couldnt parse request",
			"error":   err.Error(),
		})
		return
	}
	updatedBlog.UpdatedAt = time.Now()

	err = models.UpdateTheBlogByID(id, updatedBlog)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"message": "blog not found",
			"error":   err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{ "blog": updatedBlog})
}
