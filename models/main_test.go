

package models

import (
	"database/sql"
	// "testing"
	// "time"

	_ "github.com/lib/pq"
)

func setupTestDB() {
	connStr := "postgres://postgres:Yadamma%402004@localhost:5432/mydb?sslmode=disable"

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
}


// func TestDB(m *testing.M){
// 	dbDriver :="postgres"
// 	dbSource :="postgres://postgres:Yadamma%402004@localhost:5432/mydb?sslmode=disable"
// 	conn,err :=sql.Open(dbDriver,dbSource)
// 	if err!=nil{
// 		panic(err)
// 	}

// 	os.Exit(m.Run()) 
// 	// Run runs the tests. It returns an exit code to pass to os.Exit. The exit code is zero when all tests pass, and non-zero for any kind of failure.

// }
