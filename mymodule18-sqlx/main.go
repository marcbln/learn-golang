package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

// multi-statement Exec behavior varies between database drivers;  pq will exec them all, sqlite3,mariadb  won't
var schemaPerson = `
CREATE OR REPLACE TABLE person(
    first_name VARCHAR(255),
    last_name VARCHAR(255),
    email VARCHAR(255)
);
`

var schemaPlace = `
CREATE OR REPLACE TABLE place (
    country VARCHAR(255),
    city VARCHAR(255) NULL,
    telcode INTEGER
);
`

type Person struct {
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Email     string
}

type Place struct {
	Country string
	City    sql.NullString
	TelCode int
}

func main() {
	db, err := sqlx.Open("mysql", "root:11111@tcp(127.0.0.1:3306)/hello")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// force a connection and test that it worked
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// ---- exec the schema or fail
	db.MustExec(schemaPerson)
	db.MustExec(schemaPlace)

	// ---- insert some data
	tx := db.MustBegin()
	tx.MustExec("INSERT INTO person (first_name, last_name, email) VALUES (?, ?, ?)", "Jason", "Moiron", "jmoiron@jmoiron.net")
	tx.MustExec("INSERT INTO person (first_name, last_name, email) VALUES (?, ?, ?)", "John", "Doe", "johndoeDNE@gmail.net")
	tx.MustExec("INSERT INTO person (first_name, last_name, email) VALUES (?, ?, ?)", "Marc", "Doe", "marc@example")
	tx.MustExec("INSERT INTO place (country, city, telcode) VALUES (?, ?, ?)", "United States", "New York", "1")
	tx.MustExec("INSERT INTO place (country, telcode) VALUES (?, ?)", "Hong Kong", "852")
	tx.MustExec("INSERT INTO place (country, telcode) VALUES (?, ?)", "Singapore", "65")
	// Named queries can use structs, so if you have an existing struct (i.e. person := &Person{}) that you have populated, you can pass it in as &person
	tx.NamedExec("INSERT INTO person (first_name, last_name, email) VALUES (:first_name, :last_name, :email)", &Person{"Jane", "Citizen", "jane.citzen@example.com"})
	tx.Commit()

	// ---- Query the database, storing results in a []Person (wrapped in []interface{})
	people := []Person{}
	err = db.Select(&people, "SELECT * FROM person ORDER BY first_name ASC")
	if err != nil {
		log.Fatal(err)
	}
	jason, john := people[0], people[1]
	fmt.Printf("%#v\n", jason)
	fmt.Printf("%#v\n", john)

	// ---- get a single result, a-la database/sql's QueryRow
	marc := Person{}
	err = db.Get(&marc, "SELECT * FROM person WHERE first_name=?", "Marc")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", marc)

	// ---- Loop through rows using only one struct
	fmt.Println(" ---- Loop through rows using only one struct")
	place := Place{}
	rows, err := db.Queryx("SELECT * FROM place")
	for rows.Next() {
		err := rows.StructScan(&place)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("%#v\n", place)
	}

	// ---- named queries using map/struct as parameter
	// Selects Mr. Smith from the database
	rows, err = db.NamedQuery(`SELECT * FROM person WHERE first_name=:fn`, map[string]interface{}{"fn": "Bin"})
	// Named queries can also use structs.  Their bind names follow the same rules
	// as the name -> db mapping, so struct fields are lowercased and the `db` tag
	// is taken into consideration.
	rows, err = db.NamedQuery(`SELECT * FROM person WHERE first_name=:first_name`, jason)

	// ---- batch insert with structs
	personStructs := []Person{
		{FirstName: "Ardie", LastName: "Savea", Email: "asavea@ab.co.nz"},
		{FirstName: "Sonny Bill", LastName: "Williams", Email: "sbw@ab.co.nz"},
		{FirstName: "Ngani", LastName: "Laumape", Email: "nlaumape@ab.co.nz"},
	}
	_, err = db.NamedExec(`INSERT INTO person (first_name, last_name, email) 
									VALUES (:first_name, :last_name, :email)`, personStructs)

	// ---- batch insert with maps
	personMaps := []map[string]interface{}{
		{"first_name": "Ardie", "last_name": "Savea", "email": "asavea@ab.co.nz"},
		{"first_name": "Sonny Bill", "last_name": "Williams", "email": "sbw@ab.co.nz"},
		{"first_name": "Ngani", "last_name": "Laumape", "email": "nlaumape@ab.co.nz"},
	}
	_, err = db.NamedExec(`INSERT INTO person (first_name, last_name, email)
        											VALUES (:first_name, :last_name, :email)`, personMaps)

	// ---- Loop through rows using only one struct
	fmt.Println(" ---- Loop through rows using only one struct")
	person := Person{}
	rows, err = db.Queryx("SELECT * FROM person")
	for rows.Next() {
		err := rows.StructScan(&person)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("%#v\n", person)
	}

}
