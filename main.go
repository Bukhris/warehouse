package main

import	(
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
	"io"
		)

var (
	finalserve string
	forumid string
	name string
	parentid string
		)

func main() {

	db, err := sql.Open("sqlite3", "./sample.db")
	log.Println("website.db opened.")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	rows, err := db.Query(`select * from FORUM where parentid = "none"; ` )
		if err != nil {
			log.Fatal(err)
		}
	defer rows.Close()

	
	for rows.Next() {
		err := rows.Scan(&forumid, &name,&parentid)
		if err != nil {
			log.Fatal(err)
		}
	}
finalserve=`<!DOCTYPE html><html>
    <head>
        <title>WAREHOUSE 11</title>
    </head>
    <body>`+name+`</body></html>`
	
	log.Println(finalserve)
http.HandleFunc("/login",login)
http.HandleFunc("/",index)
http.HandleFunc("/register",register)
http.HandleFunc("/serve",serveforum)

http.ListenAndServeTLS(":9000","cert.pem","key.unencrypted.pem",nil)

}


func index(res http.ResponseWriter, req *http.Request){
	io.WriteString(res,finalserve)
	}
