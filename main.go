package main
import "fmt"
// import "net/http"
import "github.com/gin-gonic/gin"
import "RestApiProject/models"
func main() {
	fmt.Println("Hello, World!")
	server :=gin.Default() // It will setup an engine for us and return it. The default engine is a Logger and Recovery middleware already attached.
	// Handler for GET request to the root path
	server.GET("/blog",getBlog) // This is an endpoint for GET request 
	server.Run(":8080") // listen and serve on
}

// In this gin.Context already has request and response objects
func getBlog(context *gin.Context){ //context parameter will be set by gin and will be pointer to gin
	// Reach out to a database
	// Write in a file
	// Make an API call to another service
	context.JSON(200,gin.H{
			"ID":        "245",
			"Author":    "Yamini",
			"Title":     "This is a post request",
			"Content":   "This is the content of the post request",
			"CreatedAt": "2024-06-01T12:00:00Z",
			"UpdatedAt": "2024-06-01T12:00:00Z",
	}) // To send the response
// gin.H is a shortcut for map[string]interface{} which is a map with string keys and values
}