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

	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		fmt.Printf("Error opening database: %v\n", err)
		return
	}
	defer db.Close()

	sqlStmt := `create table foo (id integer not null primary key, name text); delete from foo;`

	if _, err := db.Exec(sqlStmt); err != nil {
		fmt.Printf("Error creating table: %v\n", err)
		return
	}

	if err := insertWithTransaction(db); err != nil {
		fmt.Printf("Error inserting data: %v\n", err)
		return
	}

	queryAll(db)

	stmt, err := db.Prepare("select id, name from foo where id = ?")
	if err != nil {
		fmt.Printf("Error preparing statement: %v\n", err)
		return
	}
	defer stmt.Close()
	
	var user user
	if err := stmt.QueryRow("3").Scan(&user.id, &user.name); err != nil {
		fmt.Printf("Error scanning row: %v\n", err)
		return
	}
	fmt.Println(user)
}

func insertWithTransaction(db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	
	stmt, err := tx.Prepare("insert into foo(id, name) values(?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	
	for i := 1; i < 10; i++ {
		if _, err := stmt.Exec(i, fmt.Sprintf("name-%03d", i)); err != nil {
			return err
		}
	}
	return tx.Commit()
}

func queryAll(db *sql.DB) {
	rows, err := db.Query("select id, name from foo")
	if err != nil {
		fmt.Printf("Error querying: %v\n", err)
		return
	}
	defer rows.Close()
	
	for rows.Next() {
		var id int
		var name string
		if err := rows.Scan(&id, &name); err != nil {
			fmt.Printf("Error scanning row: %v\n", err)
			continue
		}
		fmt.Println("id: ", id, " name: ", name)
	}
}
