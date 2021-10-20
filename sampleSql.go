package main

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	// connect to a database
	conn, err := sql.Open("pgx", "host=localhost port=5432 dbname=test_connect user=ms password=mzx89/" )
	if err != nil {
		log.Fatal(fmt.Sprintf("unable to connect: %v\n", err))
	}
	defer conn.Close()

	log.Println("connected to database")

	// test connection
	err = conn.Ping()
	if err != nil {
		log.Fatal("connection test failed")
	}
	log.Println("Ping successful!")

	// get rows from table
	err = getAllRows(conn)
	if err != nil {
		log.Fatal(err)
	}

	// insert a row
	insert := `insert into users (first_name, last_name) values ($1, $2)`
	_, err = conn.Exec(insert, "Ed", "Smith")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("inserted a row")

	// get rows again
	err = getAllRows(conn)
	if err != nil {
		log.Fatal(err)
	}

	// update a row
	update := `update users set first_name = $1 where first_name = $2`
	_, err = conn.Exec(update, "Eddie", "Ed")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("updated one or more rows")

	// get rows again
	err = getAllRows(conn)
	if err != nil {
		log.Fatal(err)
	}

	// get one row by id
	oneRow := `select id, first_name, last_name from users where id = $1`
	var first_name, last_name string
	var id int
	row := conn.QueryRow(oneRow, 1)
	err = row.Scan(&id, &first_name, &last_name)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("oneRow returns", id, first_name,last_name)

	// delete a row
	delete := `delete from users where first_name = $1`
	_, err = conn.Exec(delete, "Ed")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("deleted one or more rows")

	// get rows again
	err = getAllRows(conn)
	if err != nil {
		log.Fatal(err)
	}
}

func getAllRows(conn *sql.DB) error {
	rows, err := conn.Query("select id, first_name, last_name from users")
	if err != nil {
		log.Println(err)
		return err
	} 
	defer rows.Close()

	var first_name, last_name string
	var id int

	for rows.Next() {
		err := rows.Scan(&id, &first_name, &last_name)
		if err != nil {
			log.Println(err)
			return err
		}
		fmt.Println("Record is", id, first_name, last_name)
	}
	// good practice to check for scanning error after ranging through select results
	if err = rows.Err(); err != nil {
		log.Fatal("error scanning rows", err)

	}
	fmt.Println("---------------------------")
	
	return nil
}

