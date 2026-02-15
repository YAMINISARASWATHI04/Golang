package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3" // we will not use this package directly rather we use it through sql package
	// underscore indicates that we are importing the package for its side effects (initialization) and not directly using it in our code
	// that will expose functionality of sql pa
)

var DB *sql.DB

// sql package doesnot contain database drivers we need to use thirdpparty libraries

func InitDB(){ // To initialize the database connection and create the necessary tables if they do not exist
	var err error
	DB,err =sql.Open("sqlite3","api.db") // this will create a file named api.db in the current directory if it does not exist

	if err!=nil{
	panic("couldn't connect to the database") // if there is an error while opening the database connection, we panic and print the error message
	}
// gin helps us to recover from panics and return a 500 Internal Server Error response to the client instead of crashing the server
	DB.SetMaxOpenConns(10) // To confgure connection pooling for better performance and resource management
	// Connection pooling is a performance optimization technique that maintains a cache of active database connections, allowing them to be reused for multiple requests rather than creating a new connection each time
	// when we open this it controls how many connections can be open at a time to the database. By default, it is set to 0 which means unlimited connections. We can set it to a specific number to limit the number of connections that can be open at a time.

	// we dont need to open and close the connection for each request, we can reuse the same connection for multiple requests which will improve the performance of our application
	DB.SetMaxIdleConns(5) // So that if connections are closed atleast 5 connections will be kept open and ready to use for the next request. This will improve the performance of our application by reducing the time taken to establish a new connection for each request.

	if err = DB.Ping(); err != nil {
		panic("database not reachable") // FOrces the application to attempt to connect to the database and verify that it is reachable. If the connection cannot be established, it will return an error which we can handle appropriately (in this case, we panic and print the error message)
	}


	CREATE_TABLES()
}

func CREATE_TABLES(){ 
	// To create the necessary tables in the database if they do not exist
	createblogtable := `CREATE TABLE IF NOT EXISTS blogs (
	id TEXT PRIMARY KEY,
	author TEXT NOT NULL,
	title TEXT NOT NULL,
	content TEXT NOT NULL,
	created_at DATETIME NOT NULL,
	updated_at DATETIME	 NOT NULL
);`
	if DB == nil {
		panic("DB is nil before Exec")
	}

	_,err := DB.Exec(createblogtable)

	if err!=nil{
		panic(err)
	}	
}
