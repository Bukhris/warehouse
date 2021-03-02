package main

import	(
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	//"os"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"fmt"
		)



func register(res http.ResponseWriter, req *http.Request){
	if req.Method=="GET"{fmt.Fprint(res,`<!DOCTYPE html><html><head><title>WAREHOUSE 11</title></head>
	<body><form method="post"><br>
	  <input type="text" id="id" name="id" name=maxlength="10"><br>
	  <input type="password" id="pwd" name="pwd">
	  <input type="submit" value="Register"></form></body></html>`)
	}else if req.Method=="POST"{
		log.Println("OH SHIT A POST REQUEST")

		err := req.ParseForm()
		if err != nil {log.Fatal(err)}

		id:=req.FormValue("id")
		pwd:=req.FormValue("pwd")

		log.Println(id,pwd)
		regacc(id,pwd)
	}
}



func regacc(id string, pwd string) {

var(authid =id
	authpwd []byte =[]byte(pwd)
)

/*	if _, err := os.Stat("authlog"); err == nil{
		_, err := os.Open("authlog") 
		if err != nil {log.Fatal(err)}
		log.Println("Log opened.")
	}else{
		_, err := os.Create("authlog")
		if err != nil {log.Fatal(err)}
		log.Println("No log file found, Log Created.")
	}*/

	db, err := sql.Open("sqlite3", "./sample.db")
	log.Println("website DB opened.")
	if err != nil {log.Fatal(err)}
	defer db.Close()

/*	hashrow:= db.QueryRow(`select pwd from USER where userid  ="`+authid+`";`)
	hashrow.Scan(&hash)
		if err != nil {log.Fatal(err)}*/
	authpwdhashed, err := bcrypt.GenerateFromPassword(authpwd, bcrypt.DefaultCost)
	if err != nil {log.Fatal(err)}
	/*log.Println("REGACC///UNHASHED:",string(authpwd),"HASHED:",string(authpwdhashed))
	log.Println("REGACC///UNHASHED:",string(authpwd)==string(authpwdhashed))
	log.Println("cmp",bcrypt.CompareHashAndPassword(authpwdhashed,authpwd))*/
	
 _, err = db.Exec(`insert into user values ("`+authid+`","`+string(authpwdhashed)+`","`+authid+`",1,1852000,0,"pics/johnny.png",1852000,0,1852000,0,-1,"");`)
		if err != nil {log.Fatal(err)}

	
}
