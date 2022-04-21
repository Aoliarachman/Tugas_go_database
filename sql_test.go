package belajar_db

import (
	"context"
	"fmt"
	"strconv"
	"testing"
)

func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "INSERT INTO comments(email,comment) VALUES('?','?')"
	_, err := db.ExecContext(ctx, script)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert comment")
}

func TestQuerySql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "SELECT email, comment FROM comments"
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var email, comment string
		err = rows.Scan(&email, &comment)
		if err != nil {
			panic(err)
		}
		fmt.Println("email:", email)
		fmt.Println("comment", comment)
	}
}

func TestQuerySqlComplex(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "SELECT id, email, comment" +
		" created_at FROM comments"
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
	}
	defer rows.Close()

	for rows.Next() {
		var id string
		var email string
		var comment string

		err = rows.Scan(&id, &email, &comment)
		if err != nil {
			panic(err)
		}
		fmt.Println("Id", id)
		fmt.Println("Email", email)
		fmt.Println("comment", comment)
	}
}

func TestSqlInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	email := "admin@gmail.com'; #"
	comment := "salah"

	script := "SELECT email FROM comments WHERE email = '" + email +
		"' AND comment = '" + comment + "'LIMIT 1"

	fmt.Println(script)
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		var email string
		err := rows.Scan(&email)
		if err != nil {
			panic(err)
		}
		fmt.Println("Sukses Login", email)
	} else {
		fmt.Println("Gagal Login")
	}
}

func TestSqlInjectionSafe(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	email := "admin'; #"
	comment := "salah"

	script := "SELECT email FROM comments WHERE email = ? AND comment = ? LIMIT 1"
	fmt.Println(script)
	rows, err := db.QueryContext(ctx, script, email, comment)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	if rows.Next() {
		var email string
		err := rows.Scan(&email)
		if err != nil {
			panic(err)
		}
		fmt.Println("Sukses Login", email)
	} else {
		fmt.Println("Gagal Login")
	}
}

func TestExecSqlParameter(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	email := "Eko'; DROP TABLE user; #"
	comment := "Eko"

	script := "INSERT INTO comments(email, comment) VALUES(?,?)"
	_, err := db.ExecContext(ctx, script, email, comment)
	if err != nil {
		panic(err)
	}
	fmt.Println("Success insert new user")
}

func TestAutoIncrement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	email := "eko@gmail.com"
	comment := "Test komen"

	script := "INSERT INTO comments(email, comment) VALUES(?,?)"
	result, err := db.ExecContext(ctx, script, email, comment)
	if err != nil {
		panic(err)
	}
	insertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	fmt.Println("Success insert new comment with id", insertId)
}

func TestPrepareStatment(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	script := "INSERT INTO comments(email,comment) VALUES(?,?)"
	statement, err := db.PrepareContext(ctx, script)
	if err != nil {
		panic(err)
	}

	defer statement.Close()

	for i := 0; i < 10; i++ {
		email := "eko" + strconv.Itoa(i) + "@gmail.com"
		comment := "komentar ke " + strconv.Itoa(i)

		result, err := statement.ExecContext(ctx, email, comment)
		if err != nil {
			panic(err)
		}

		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}

		fmt.Println("Comment Id", id)
	}
}

func TestTransaction(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	script := "INSERT INTO comments(email, comment) VALUES(?,?)"
	// DO TRANSACTION

	for i := 0; i < 10; i++ {
		email := "eko" + strconv.Itoa(i) + "@gmail.com"
		comment := "komentar ke " + strconv.Itoa(i)

		result, err := tx.ExecContext(ctx, script, email, comment)
		if err != nil {
			panic(err)
		}
		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}

		fmt.Println("Comment Id", id)
	}

	err = tx.Rollback()
	if err != nil {
		panic(err)
	}
}
