package models

// To store event data
import (
	"encoding/json"
	"errors"
	"fmt"
	
	"os"
	"time"
)

type Blog struct {
	ID        string    `json:"id"`
	Author    string    `json:"author"` // struct tags we need to mention for the one which is important
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

const blogfile = "blogs.json"

func SaveBlogs(blog Blog) error {

	blogs, err := readBlogsFromFile()
	if err != nil {
		return fmt.Errorf("Error reading blogs from file: %v", err)
	}
	blogs = append(blogs, blog) // Append the new blog to the existing slice of blogs

	// Marshal means converts the struct into json format and then we can write it to the file

	blogdata, err := json.MarshalIndent(blogs, "", " ") // Marshal the slice of structs into JSON formatted bytes
	if err != nil {
		return fmt.Errorf("Error marshaling data: %v", err)
	}
	return os.WriteFile(blogfile, blogdata, 0644)
	// Write the JSON data to the file, creating it if it doesn't exist (0644 is standard permissions)

}

// Its just a function not a method
func GetAllBlogs() []Blog {

	blogs, err := readBlogsFromFile()
	if err != nil {
		fmt.Println("Error reading blogs from file:", err)
		return []Blog{}
	}
	return blogs

}

func readBlogsFromFile() ([]Blog, error) {
	blogs, err := os.ReadFile(blogfile)
	if err != nil {
		fmt.Println("Error reading file:", err)

		if errors.Is(err, os.ErrNotExist) {
			return []Blog{}, nil // Return an empty slice if the file doesn't exist
		}
		return nil, fmt.Errorf("Error reading file: %v", err)
	}
	var blogList []Blog
	// converts the json data into the struct format and then we can use it in our code
	err = json.Unmarshal(blogs, &blogList)
	if err != nil {
		return nil, fmt.Errorf("Error unmarshaling data: %v", err)
	}
	return blogList, nil

}

func GetBlogByID(id string) (*Blog, error) {
	blogs, err := readBlogsFromFile()
	if err != nil {
		return nil, fmt.Errorf("Error reading blogs from file: %v", err)
	}
	for i := range blogs {
		if blogs[i].ID == id {
			return &blogs[i], nil
		}
	}
	return nil, errors.New("blog not found")
}

func DeleteBlogByID(id string) error {
	blogs, err := readBlogsFromFile()
	if err != nil {
		return fmt.Errorf("Error reading blogs from file: %w", err)
	}
	
	exist := false

	for i := range blogs {
		if blogs[i].ID == id {
			blogs = append(blogs[:i], blogs[i+1:]...) // Remove the blog with the specified ID
			exist = true
			break

		}
	}
	if !exist {
		return errors.New("blog not found")
	}	

	
	blogdata, err := json.MarshalIndent(blogs, "", " ") // Marshal the updated slice of blogs into JSON formatted bytes
	if err != nil {
		return fmt.Errorf("Error marshaling data: %v", err)
	}
	return os.WriteFile(blogfile, blogdata, 0644) // Write the updated JSON data back to the file

}

func UpdateBlogByID(id string, updatedBlog Blog) error {
	blogs, err := readBlogsFromFile()
	if err != nil {
		return fmt.Errorf("Error reading blogs from file: %v", err)
	}
	exist := false
	for i := range blogs {
		if blogs[i].ID == id {
			updatedBlog.ID = id
			updatedBlog.CreatedAt = blogs[i].CreatedAt
			updatedBlog.UpdatedAt = time.Now()
			blogs[i] = updatedBlog // Update the blog with the new data
			exist = true
			break
		}
	}
	if !exist {
		return errors.New("blog not found")
	}

	blogdata, err := json.MarshalIndent(blogs, "", " ") // Marshal the updated slice of blogs into JSON formatted bytes
	if err != nil {
		return fmt.Errorf("Error marshaling data: %v", err)
	}
	return os.WriteFile(blogfile, blogdata, 0644) // Write the updated JSON data back to the file

}