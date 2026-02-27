package models

import (
	"database/sql"
	"testing"

	// "testing"
	// "time"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
)

func setupTestDB(t *testing.T) *sql.DB {
	connStr := "postgres://postgres:Yadamma%402004@localhost:5432/mydb?sslmode=disable"

	var err error
	DB, err = sql.Open("postgres", connStr)
	require.NoError(t,err)
	err=DB.Ping()

	t.Cleanup(func() {
        DB.Close()
    })
	
	return DB
}


