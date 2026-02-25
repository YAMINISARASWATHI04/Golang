package models

// To store event data
import (
	// "encoding/json"
	// "errors"
	"fmt"
	// "RestApiProject/db"
	// "context"
	// "github.com/google/uuid"
	// "os"
	"database/sql"
	"time"
)


func SavetheBlogs(db *sql.DB, blog Blog) (Blog, error) {
	blog.CreatedAt = time.Now().UTC()
	blog.UpdatedAt = time.Now().UTC()

	query := `INSERT INTO blogs (id, author, title, content, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)` // Question marks are placeholders and protect us from SQL injection attacks by ensuring that user input is treated as data rather than executable code. The database driver will safely escape the values before executing the query.
	// stmt, err := DB.Prepare(query)
	// if err != nil {
	// 	return err
	// }
	// defer stmt.Close()

	_, err := db.Exec(query, blog.ID, blog.Author, blog.Title, blog.Content, blog.CreatedAt, blog.UpdatedAt) // To execute the prepared statement with the provided parameters. It will insert a new row into the blogs table with the values from the blog struct.
	if err != nil {
		return Blog{}, err
	}

	return blog, nil
}

func GetBlogs(db *sql.DB) ([]Blog, error) {
	query := `SELECT * FROM blogs`
	rows, err := db.Query(query) // To execute the SQL query and return the result set as rows. It returns a sql.Rows object that can be iterated over to access each row of data.
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

	return blogs, nil
}

func GetSingleBlogByID(db *sql.DB, id string) (Blog, error) {
	query := `SELECT * FROM blogs WHERE id = $1`
	row := db.QueryRow(query, id)
	var blog Blog
	err := row.Scan(&blog.ID, &blog.Author, &blog.Title, &blog.Content, &blog.CreatedAt, &blog.UpdatedAt) // To read the values from the retrieved row into the fields of the blog struct. The Scan method takes pointers to the fields of the struct and populates them with the corresponding values from the row.
	if err != nil {
		fmt.Println("Error scanning row:", err)
		return Blog{}, err
	}
	return blog, nil
}

func UpdateTheBlogByID(db *sql.DB, id string, updatedBlog Blog) error {
	query := `UPDATE blogs SET author = $2, title = $3, content = $4, updated_at = $5 WHERE id = $1`
	// stmt, err := db.Prepare(query) // To prepare the SQL statement for execution. It returns a prepared statement that can be executed multiple times with different parameters.
	// if err != nil {
	// 	return err
	// }
	result, err := db.Exec(query,id, updatedBlog.Author, updatedBlog.Title, updatedBlog.Content, updatedBlog.UpdatedAt) // To execute the prepared statement with the provided parameters. It will update the corresponding row in the blogs table with the new values from the updatedBlog struct where the ID matches.
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no blog found with id %s", id)
	}

	// defer db.Close()
	return err
}

func DeleteBlog(db *sql.DB,id string) error {
	query := "DELETE FROM blogs WHERE ID=$1"

	// stmt, err := db.Prepare(query)

	// if err != nil {
	// 	return err
	// }

	// defer stmt.Close()

	result, err := db.Exec(query,id)

	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return fmt.Errorf("Blog not found")
	}

	return nil

}
