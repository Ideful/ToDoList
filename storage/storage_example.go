package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	db,err:= sql.Open("postgres","port=5432 sslmode=disable dbname = DB" )	// ya hz kuda podklyachilsya
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}


	// get max val
	var id int
	toadd,_:=db.Query(`select max(id) from person`)
	for toadd.Next() {
		toadd.Scan(&id)
	}
	id++
//  post
	sqlStatement := `
	INSERT INTO person (id,name, age, gender, address)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id
`

	err=db.QueryRow(sqlStatement, id,"Olejka", 123, "male", "Ekb").Scan(&id)
	if err != nil {
		log.Fatal(err)
	}

	// get
	rows,err:=db.Query(`SELECT * from person`)
	for rows.Next() {
		var id int
		var name string
		var age int
		var gender string
		var address string
		err := rows.Scan(&id,&name,&age,&gender,&address) 
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d, Name: %s, Age: %d, Gender: %s, Address: %s\n", id, name, age, gender, address)
	}
}

 // у типа в видосе в структуре энкриптед
// а в бд тоже энкриптед



