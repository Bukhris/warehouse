package main

import	(
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
	"fmt"
		)

var (
	serve string
	forumid string
	name string
	parentid string
		)


func serveforum(res http.ResponseWriter, req *http.Request){

	fmt.Println(string(req.URL))
	
	db, err := sql.Open("sqlite3", "./sample.db")
	log.Println("SERVEFORUM: website.db opened.")
	if err != nil {log.Fatal(err)}
	defer db.Close()

	rows, err := db.Query(`select * from FORUM where parentid = "none"; ` )
	if err != nil {log.Fatal(err)}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&forumid, &name,&parentid)
		if err != nil {log.Fatal(err)}
	}

	finalserve=`<!DOCTYPE html><html>
    <head>
        <title>WAREHOUSE 11</title>
    </head>
    <body>`+name+`</body></html>`
	
	log.Println(serve)


	fmt.Fprintf(res,serve)

}

