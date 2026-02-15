package models

// To store event data
import (
	// "encoding/json"
	// "errors"
	"fmt"
	"RestApiProject/db"
	// "github.com/google/uuid"
	// "os"
	// "time"
)



func SavetheBlogs(blog Blog) error  {
	query := `INSERT INTO blogs (id, author, title, content, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)` // Question marks are placeholders and protect us from SQL injection attacks by ensuring that user input is treated as data rather than executable code. The database driver will safely escape the values before executing the query.
	stmt,err := db.DB.Prepare(query) // To prepare the SQL statement for execution. It returns a prepared statement that can be executed multiple times with different parameters.
	// when we use prepare it stores in memory and we can execute it multiple times with different parameters which will improve the performance of our application by reducing the time taken to parse and compile the SQL statement for each execution.
	if err!=nil{
		return err
	}
	_, err =stmt.Exec(blog.ID,blog.Author,blog.Title,blog.Content,blog.CreatedAt,blog.UpdatedAt) // To execute the prepared statement with the provided parameters. It will insert a new row into the blogs table with the values from the blog struct.
	if err!=nil{
		return err
	}
	// To generate a unique identifier for the blog post using the UUID package. This ensures that each blog post has a unique ID, which can be used to retrieve, update, or delete the post later.

	defer stmt.Close()
	return nil
}
func GetBlogs() ([]Blog, error) {
	query := `SELECT * FROM blogs` // To retrieve all blog posts from the database. It selects the relevant columns from the blogs table.
	rows, err := db.DB.Query(query) // To execute the SQL query and return the result set as rows. It returns a sql.Rows object that can be iterated over to access each row of data.
	if err != nil {
		fmt.Println("Error executing query:", err)
		return nil, err
	}
	defer rows.Close() 

	var blogs []Blog
	for rows.Next() { // To iterate over the rows returned by the query. It returns true if there is another row to process and false when there are no more rows.
		var blog Blog
		err := rows.Scan(&blog.ID, &blog.Author, &blog.Title, &blog.Content, &blog.CreatedAt, &blog.UpdatedAt) // To read the values from the current row into the fields of the blog struct. The Scan method takes pointers to the fields of the struct and populates them with the corresponding values from the row.
		if err != nil {
			fmt.Println("Error scanning row:", err)
			return nil, err
		}
		blogs = append(blogs, blog) // To add the blog struct to the slice of blogs. This allows us to collect all the blog posts retrieved from the database into a single slice that can be returned to the caller.
	}
	
	
	return blogs,nil
}
