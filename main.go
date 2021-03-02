package main

import	(
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
	"fmt"
		)

func main() {

	
http.HandleFunc("/login",login)
http.HandleFunc("/",index)
http.HandleFunc("/register",register)
//http.HandleFunc("/serve",serveforum)

log.Fatal(http.ListenAndServeTLS(":9000","cert.pem","key.unencrypted.pem",nil))

}

func index(res http.ResponseWriter, req *http.Request){
	var (
		forumid string
		name string
		parentid string
		)


	db, err := sql.Open("sqlite3", "./sample.db")
	log.Println("website.db opened.")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	rows, err := db.Query(`select * from FORUM ; ` )
		if err != nil {
			log.Fatal(err)
		}
	defer rows.Close()

	names :="" 
	for rows.Next() {
		err := rows.Scan(&forumid, &name,&parentid)
		log.Println(forumid,"===",name,"===",parentid)
		names=names+name+"<br>"
		if err != nil {log.Fatal(err)}
	}

	
finalserve:=`<!DOCTYPE html><html>
    <head>
        <title>WAREHOUSE 11</title>
    </head>
    <body>`+names+`</body></html>`
	fmt.Fprintf(res,finalserve)


	}
