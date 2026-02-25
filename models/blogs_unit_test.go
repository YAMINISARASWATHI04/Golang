package models

import (
	"testing"
	"time"

	"github.com/google/uuid"

	"RestApiProject/util"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	// "time"
	
)

func TestSaveTheBlogs(t *testing.T){
	db,mock,err :=sqlmock.New()
	require.NoError(t,err)
	defer db.Close()

	blog := Blog{
		ID:        uuid.New().String(),
		Author:    util.RandomAuthor(),
		Title:     util.RandomTitle(),
		Content:   util.RandomContent(),
		// CreatedAt: time.Now().UTC(),
		// UpdatedAt: time.Now().UTC(),
	}

	mock.ExpectExec("INSERT INTO blogs").
	WithArgs(blog.ID,blog.Author,blog.Title ,blog.Content,sqlmock.AnyArg(),sqlmock.AnyArg(),).
		WillReturnResult(sqlmock.NewResult(1,1))

	arg,err :=SavetheBlogs(db,blog)
	require.NoError(t,err)
	require.NotZero(t,arg.ID)
	require.NotZero(t,blog.ID)
	require.Equal(t,arg.ID,blog.ID)
	require.Equal(t,arg.Author,blog.Author)
	require.Equal(t,arg.Title,blog.Title)
	require.Equal(t,arg.Content,blog.Content)
	require.NotZero(t,arg.CreatedAt)
	require.NotZero(t,arg.UpdatedAt)

	require.NoError(t,mock.ExpectationsWereMet())

}

func TestGetBlogs(t *testing.T){
	db,mock,err :=sqlmock.New()
	require.NoError(t,err)
	defer db.Close()
	id:=  	   uuid.New().String()
	author:=  util.RandomAuthor()
	title:=   util.RandomTitle()
	content:=  util.RandomContent()
	created_at :=time.Now()
	updated_at :=time.Now()

	rows:=sqlmock.NewRows([] string{"id","author","title","content","created_at","updated_at"}).
	AddRow(id,author,title,content,created_at,updated_at)

	mock.ExpectQuery("SELECT \\* FROM blogs").WillReturnRows(rows)

	blog,err:=GetBlogs(db)

	require.NoError(t,err)
	require.Len(t,blog,1)
	require.Equal(t,id,blog[0].ID)
	require.Equal(t,author,blog[0].Author)
	require.Equal(t,title,blog[0].Title)
	require.Equal(t,content,blog[0].Content)
	require.WithinDuration(t,created_at,blog[0].CreatedAt,time.Second)
	require.WithinDuration(t,updated_at,blog[0].UpdatedAt,time.Second)
	require.NotZero(t,blog[0].ID)
	require.NotZero(t,blog[0].CreatedAt)
	require.NotZero(t,blog[0].UpdatedAt)

}

func TestGetSingleBlogByID(t *testing.T){
	db,mock,err :=sqlmock.New()
	require.NoError(t,err)
	defer db.Close()
	

	id:=  	   uuid.New().String()
	author:=  util.RandomAuthor()
	title:=   util.RandomTitle()
	content:=  util.RandomContent()
	created_at :=time.Now()
	updated_at :=time.Now()

	rows:=sqlmock.NewRows([] string{"id","author","title","content","created_at","updated_at"}).
	AddRow(id,author,title,content,created_at,updated_at)

	mock.ExpectQuery("SELECT \\* FROM blogs").WillReturnRows(rows)

	blog,err:=GetSingleBlogByID(db,id)

	require.NoError(t,err)
	// require.Len(t,blog,1)
	require.Equal(t,id,blog.ID)
	require.Equal(t,author,blog.Author)
	require.Equal(t,title,blog.Title)
	require.Equal(t,content,blog.Content)
	require.WithinDuration(t,created_at,blog.CreatedAt,time.Second)
	require.WithinDuration(t,updated_at,blog.UpdatedAt,time.Second)
	require.NotZero(t,blog.ID)
	require.NotZero(t,blog.CreatedAt)
	require.NotZero(t,blog.UpdatedAt)
}

func TestUpdateTheBlogByID(t *testing.T){
	db,mock,err :=sqlmock.New()
	require.NoError(t,err)
	defer db.Close()
	id:=uuid.New().String()

	blog :=Blog{
		Author:  util.RandomAuthor(),
		Title:   util.RandomTitle(),
		Content: util.RandomContent(),
	}
	mock.ExpectExec("UPDATE blogs SET").
	WithArgs(
		id,
		blog.Author,
		blog.Title,
		blog.Content,
		sqlmock.AnyArg(),
	).
	WillReturnResult(sqlmock.NewResult(1,1))

	err = UpdateTheBlogByID(db,id,blog)
	require.NoError(t, mock.ExpectationsWereMet())
	require.NoError(t,err)
}

func TestDeleteBlog(t *testing.T){
	db,mock,err :=sqlmock.New()
	require.NoError(t,err)
	defer db.Close()

	id:=uuid.New().String()

	mock.ExpectExec("DELETE FROM blogs").WithArgs(id).WillReturnResult(sqlmock.NewResult(1,1))

	err = DeleteBlog(db,id)
	require.NoError(t,err)
	require.NoError(t, mock.ExpectationsWereMet())

}
