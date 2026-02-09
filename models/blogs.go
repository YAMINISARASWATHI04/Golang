package models

// To store event data
import "time"

type Blog struct {
	ID        int
	Author    string
	Title     string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

var blogs =[] Blog {}
func (e Blog) Save(){
	// store in json file later
	blogs = append(blogs,e)
}
// Its just a function not a method
func GetAllBlogs() [] Blog{
	return blogs

}