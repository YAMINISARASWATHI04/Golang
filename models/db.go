package models

import (
	"database/sql"
	_ "github.com/jackc/pgx/v5/stdlib" // we will not use this package directly rather we use it through sql package
	"github.com/jackc/pgx/v5"
	// underscore indicates that we are importing the package for its side effects (initialization) and not directly using it in our code
	// that will expose functionality of sql pa
	"context"
    // "fmt"
    // "log"
	// "github.com/alextanhongpin/dbtx"
    _ "github.com/lib/pq"
)

var DB *sql.DB


// sql package doesnot contain database drivers we need to use thirdpparty libraries

func InitDB(){ // To initialize the database connection and create the necessary tables if they do not exist
	var err error
	DB,err =sql.Open("pgx","postgres://postgres:Yadamma%402004@localhost:5432/mydb?sslmode=disable")

	if err!=nil{
		panic("couldn't connect to the database") 
	}
	DB.SetMaxOpenConns(10)
	// we dont need to open and close the connection for each request, we can reuse the same connection for multiple requests which will improve the performance of our application
	DB.SetMaxIdleConns(5) // So that if connections are closed atleast 5 connections will be kept open and ready to use for the next request. This will improve the performance of our application by reducing the time taken to establish a new connection for each request.

	if err = DB.Ping(); err != nil {
		panic("database not reachable") // FOrces the application to attempt to connect to the database and verify that it is reachable. If the connection cannot be established, it will return an error which we can handle appropriately (in this case, we panic and print the error message)
	}

	CREATE_TABLES()
}

func Connection() (*pgx.Conn, error) {
    conn, err := pgx.Connect(context.Background(), "postgres://postgres:Yadamma%402004@localhost:5432/mydb?sslmode=disable")
    if err != nil {
        return nil, err
    }
    
    
	return conn, nil


}



func CREATE_TABLES(){ 
	// To create the necessary tables in the database if they do not exist
	createblogtable := `CREATE TABLE IF NOT EXISTS blogs (
	id TEXT PRIMARY KEY,
	author TEXT NOT NULL,
	title TEXT NOT NULL,
	content TEXT NOT NULL,
	created_at TIMESTAMP DEFAULT now(),
	updated_at TIMESTAMP DEFAULT now()
);`
	if DB == nil {
		panic("DB is nil before Exec")
	}

	_,err := DB.Exec(createblogtable)

	if err!=nil{
		panic(err)
	}	
}
