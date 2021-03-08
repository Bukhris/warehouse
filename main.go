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
log.Fatal(http.ListenAndServeTLS(":9000","cert.pem","key.unencrypted.pem",nil))
}

func index(res http.ResponseWriter, req *http.Request){
//load DB
	db, err := sql.Open("sqlite3", "./main.db")
	var sqlcolumns [11]string
	if err != nil {log.Fatal(err)}
	defer db.Close()

	if req.URL.Path=="/"{	
//serve main page
		rows, err := db.Query(`select * from FORUM where parentid = "none";`)
			if err != nil {log.Fatal(err)}
		defer rows.Close()

		result := "" 
		for rows.Next() {
			err := rows.Scan(&sqlcolumns[0], &sqlcolumns[1],&sqlcolumns[2])
			result=result+"<a href=/"+sqlcolumns[0]+">"+sqlcolumns[1]+"</a><br>" 
	 		if err != nil {log.Fatal(err)}}
	finalserve:=`<!DOCTYPE html><html>
	    <head>
	        <title>Sample</title></head>
	    <body>`+result+`</body></html>`
		fmt.Fprintf(res,finalserve)
	}else{
//serve forum and posts if any
			rows, err := db.Query(`select * from FORUM where parentid = "`+req.URL.Path[1:]+`";`)
			if err != nil {log.Fatal(err)}
		defer rows.Close()

		result := "" 
		for rows.Next() {
			err := rows.Scan(&sqlcolumns[0], &sqlcolumns[1],&sqlcolumns[2])
			result=result+"<a href=/"+sqlcolumns[0]+">"+sqlcolumns[1]+"</a><br>" 
			log.Println(result)
	 		if err != nil {log.Fatal(err)}}
//load posts
			log.Println(result)
			rows, err = db.Query(`select * from POST where forumid="`+req.URL.Path[1:]+`";`)
			if err != nil {log.Fatal(err)}
			defer rows.Close()
				 		 
		for rows.Next() { 
			err := rows.Scan(&sqlcolumns[0], &sqlcolumns[1],&sqlcolumns[2],&sqlcolumns[3],&sqlcolumns[4],&sqlcolumns[5],&sqlcolumns[6],&sqlcolumns[7],&sqlcolumns[8],&sqlcolumns[9],&sqlcolumns[10])
			result=result+"<a href=/"+sqlcolumns[0]+">POST:"+sqlcolumns[2]+"</a><br>" 
			if err != nil {log.Fatal(err)}}
				 			 		
	 			finalserve:=`<!DOCTYPE html><html>
	    <head>
	        <title>Sample</title></head>
	    <body>`+result+`</body></html>`
		fmt.Fprintf(res,finalserve)

	}

}
