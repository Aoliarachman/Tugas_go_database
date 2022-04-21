package belajar_db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func TestOpenConnection(t *testing.T) {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/belajar_db?parseTime=True")
	if err != nil {
		panic(err)
	}
	defer db.Close()
}
