package main

import (
	"database/sql"
	"errors"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {

	// sql.Open() does not establish any connections to the database!
	db, err := sql.Open("pgx", "postgres://postgres:test@localhost:5432/postgres")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// The first actual connection to the datastore will be established lazily, when it’s needed for the first time
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	//  If a function name includes Query, it is designed to ask a question of the database, and will return a set of rows
	rows, err := db.Query("select id, name from users where id = $1", 1)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var (
		id   int
		name string
	)
	for rows.Next() {
		if err = rows.Scan(&id, &name); err != nil {
			log.Fatal(err)
		}
		log.Println(id, name)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	// It is good to prepare queries to be used multiple times
	selectUserById, err := db.Prepare("select id, name from users where id = $1")
	if err != nil {
		log.Fatal(err)
	}
	defer selectUserById.Close()
	rows, err = selectUserById.Query(1)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		if err = rows.Scan(&id, &name); err != nil {
			log.Fatal(err)
		}
		log.Println(id, name)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	// If a query returns at most one row, I can use a shortcut
	err = db.QueryRow("select name from users where id = $1", 100).Scan(&name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Println("no rows")
		} else {
			log.Fatal(err)
		}
	}
	log.Printf("name: %s", name)

	// Use Exec() to accomplish an INSERT, UPDATE, DELETE, or another statement that doesn’t return rows
	insertIntoUsers, err := db.Prepare("insert into users(name) values ($1)")
	if err != nil {
		log.Fatal(err)
	}
	res, err := insertIntoUsers.Exec("Margarita")
	if err != nil {
		log.Fatal(err)
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("affected count = %d\n", rowCnt)

	// It is ok to use like that
	_, err = insertIntoUsers.Exec("Maria")
	if err != nil {
		log.Fatal(err)
	}
}
