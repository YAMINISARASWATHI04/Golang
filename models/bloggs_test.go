package models

import (
	// "context"
	"RestApiProject/util"
	"database/sql"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func CreateBlog(t *testing.T) Blog {
	setupTestDB()
	arg := Blog{
		ID:        uuid.New().String(),
		Author:    util.RandomAuthor(),
		Title:     util.RandomTitle(),
		Content:   util.RandomContent(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	resblog, err := SavetheBlogs(arg)
	require.NoError(t, err)
	require.NotEmpty(t, arg)

	// var resblog Blog

	query := `select id,author,title,content from blogs where id=$1`
	err = DB.QueryRow(query, arg.ID).Scan(&resblog.ID, &resblog.Author, &resblog.Title, &resblog.Content)

	require.NoError(t, err)

	require.Equal(t, arg.Author, resblog.Author)
	require.Equal(t, arg.Title, resblog.Title)
	require.Equal(t, arg.Content, resblog.Content)

	require.NotZero(t, resblog.ID)
	require.NotZero(t, resblog.CreatedAt)
	require.NotZero(t, resblog.UpdatedAt)
	return resblog
}

func TestSavetheBlogs(t *testing.T) { // we will use T object to manage the test case
	CreateBlog(t)

}

func TestGetBlogs(t *testing.T) {
	blog1 := CreateBlog(t)
	blogs, err := GetBlogs()
	require.NoError(t, err)
	require.NotEmpty(t, blogs)
	var blog2 Blog
	for _, b := range blogs {
		if b.ID == blog1.ID {
			blog2 = b
			break
		}
	}
	require.Equal(t, blog1.ID, blog2.ID)
	require.Equal(t, blog1.Author, blog2.Author)
	require.Equal(t, blog1.Title, blog2.Title)
	require.Equal(t, blog1.Content, blog2.Content)
	require.NotZero(t, blog2.ID)
	require.NotZero(t, blog2.CreatedAt)
	//require.Equal(t, blog1.CreatedAt, blog2.CreatedAt)
	require.WithinDuration(t, blog1.CreatedAt, blog2.CreatedAt, time.Second)
	require.WithinDuration(t, blog1.UpdatedAt, blog2.UpdatedAt, time.Second)

}

func TestGetSingleBlogByID(t *testing.T) {
	blog1 := CreateBlog(t)
	blog2, err := GetSingleBlogByID(blog1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, blog1)
	require.Equal(t, blog1.ID, blog2.ID)
	require.Equal(t, blog1.Author, blog2.Author)
	require.Equal(t, blog1.Title, blog2.Title)
	require.Equal(t, blog1.Content, blog2.Content)
	//require.Equal(t,blog1.CreatedAt,blog2.CreatedAt)
	require.WithinDuration(t, blog1.UpdatedAt, blog2.UpdatedAt, time.Second)
	require.NotZero(t, blog2.ID)

}

func TestUpdateTheBlogByID(t *testing.T) {
	blog1 := CreateBlog(t)

	arg := Blog{
		ID:        blog1.ID,
		Author:    blog1.Author,
		Title:     blog1.Title,
		Content:   blog1.Content,
		CreatedAt: blog1.CreatedAt,
		UpdatedAt: time.Now().UTC(),
	}

	err := UpdateTheBlogByID(arg.ID, arg)

	require.NoError(t, err)
	require.NotEmpty(t, blog1)
	require.NotEmpty(t, arg)
	require.Equal(t, blog1.ID, arg.ID)
	require.Equal(t, blog1.Author, arg.Author)
	require.Equal(t, blog1.Title, arg.Title)
	require.Equal(t, blog1.Content, arg.Content)
	require.NotZero(t, blog1.ID)
	require.NotZero(t, arg.ID)
	//require.Equal(t, blog1.UpdatedAt, arg.UpdatedAt)
	require.WithinDuration(t, blog1.CreatedAt, arg.CreatedAt, time.Second)
	require.WithinDuration(t, blog1.UpdatedAt, arg.UpdatedAt, time.Second)

}

func TestDeleteBlog(t *testing.T) {
	blog1 := CreateBlog(t)
	err := DeleteBlog(blog1.ID)
	require.NoError(t, err)

	blog2, err := GetSingleBlogByID(blog1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, blog2)
}
