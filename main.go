package main

import (
	"RestApiProject/db"
	"RestApiProject/routes"
	// "RestApiProject/models"
	"fmt"
	"os"
	// "net/http"
	// "time"
	"context"


	"github.com/gin-gonic/gin"
	// "github.com/google/uuid"
)

func main() {
	fmt.Println("Hello, World!")
	conn,err :=db.Connection()

	if (err!=nil){
		fmt.Fprintf(os.Stderr,"Unable to connect to databse",err)
	}

	defer conn.Close(context.Background()) 

	db.InitDB()
	
    fmt.Println("Connection established succesfully")
	server := gin.Default() // It will setup an engine for us and return it. The default engine is a Logger and Recovery middleware already attached.
	
	routes.RegisterRoutes(server) 
	server.Run(":3000") 
}

// In this gin.Context already has request and response objects
