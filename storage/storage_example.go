package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	db,err:= sql.Open("postgres","user=postgres port=5432 sslmode=disable dbname=postgres password=qwerty" )	
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	var id int
	sqlStatement := `INSERT INTO users (username,password) VALUES ($1, $2) RETURNING id`

	err=db.QueryRow(sqlStatement,"qww2e","z3xc").Scan(&id)
	if err != nil {
		log.Fatal(err)
	}

	// get
	rows,err:=db.Query(`SELECT * from users`)
	for rows.Next() {
		var id int
		var name string
		var pwd string
		err := rows.Scan(&id,&name,&pwd) 
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("id: %d, name: %s, pwd: %s\n",id, name,pwd)
	}
}

 // у типа в видосе в структуре энкриптед
// а в бд тоже энкриптед



