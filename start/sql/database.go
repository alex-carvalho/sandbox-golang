package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type user struct {
	id   int
	name string
}

func main() {
	os.Remove("./foo.db")

	db, _ := sql.Open("sqlite3", "./foo.db")
	defer db.Close()

	sqlStmt := `create table foo (id integer not null primary key, name text); delete from foo;`

	db.Exec(sqlStmt)

	insertWithTransaction(db)

	queryAll(db)

	stmt, _ := db.Prepare("select id, name from foo where id = ?")
	defer stmt.Close()
	var user user
	stmt.QueryRow("3").Scan(&user.id, &user.name)
	fmt.Println(user)
}

func insertWithTransaction(db *sql.DB) {
	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("insert into foo(id, name) values(?, ?)")
	defer stmt.Close()
	for i := 1; i < 10; i++ {
		stmt.Exec(i, fmt.Sprintf("name-%03d", i))
	}
	tx.Commit()
}

func queryAll(db *sql.DB) {
	rows, _ := db.Query("select id, name from foo")
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		rows.Scan(&id, &name)
		fmt.Println("id: ", id, " name: ", name)
	}
}
