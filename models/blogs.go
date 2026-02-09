package models

// To store event data
import "time"

type Blog struct {
	ID        int       `json:"id"`
	Author    string    `json:"author"` // struct tags we need to mention for the one which is important
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

var blogs = []Blog{}

func (e Blog) Save() {
	// store in json file later
	blogs = append(blogs, e)
}

// Its just a function not a method
func GetAllBlogs() []Blog {
	return blogs

}
