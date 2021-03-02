package main

import	(
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"fmt"
	"math/rand"
		)



func login(res http.ResponseWriter, req *http.Request){
	if req.Method=="GET"{fmt.Fprint(res,`<!DOCTYPE html><html><head><title>WAREHOUSE 11</title></head>
	<body><form method="post"><br>
	  <input type="text" id="id" name="id"><br>
	  <input type="password" id="pwd" name="pwd">
	  <input type="submit" value="Login"></form></body></html>`)
	}else if req.Method=="POST"{
		log.Println("OH SHIT A POST REQUEST")

		err := req.ParseForm()
		if err != nil {log.Fatal(err)}

		id:=req.FormValue("id")
		pwd:=req.FormValue("pwd")

		log.Println(id,pwd)
		if checkpwd(id,pwd)==0{
			session:= http.Cookie{
				Name: "logon",
				Value: fmt.Sprintf("The integer is: %d",rand.Int()),
				Secure: true,
				MaxAge: 0,
			}
			
			http.SetCookie(res, &session)
				
		}
	}
}





func checkpwd(id string, pwd string)int {

var(hash string
	authid =id
	authpwd  []byte =[]byte(pwd)
)


	if _, err := os.Stat("authlog"); err == nil{
		_, err := os.Open("authlog") 
		if err != nil {log.Fatal(err)}
		log.Println("Log opened.")
	}else{
		_, err := os.Create("authlog")
		if err != nil {log.Fatal(err)}
		log.Println("No log file found, Log Created.")
	}

	db, err := sql.Open("sqlite3", "./sample.db")
	log.Println("website DB opened.")
	if err != nil {log.Fatal(err)}
	defer db.Close()

	hashrow:= db.QueryRow(`select pwd from USER where userid  ="`+authid+`";`)
	
	hashrow.Scan(&hash)
		if err != nil {log.Fatal(err)}

	err = bcrypt.CompareHashAndPassword([]byte(hash),authpwd)	
	if err==nil{return 0}else{return 1}
}
