package main

import	(
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"log"
	)
func initfile(){
	os.Remove("./main.db")
	log.Println("main.db deleted.")
	db, err := sql.Open("sqlite3", "./main.db")
	log.Println("website.db created and opened.")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
        _, err = db.Exec(`CREATE TABLE FORUM(forumid PRIMARY KEY, name, parentid REFERENCES FORUM(forumid));
CREATE TABLE USER(userid PRIMARY KEY, confirmation, pwd, username, type, signupdate, status, pic, lastloginattemptdate, failedlogincount, lastlogindate, lastposteid REFERENCES POST(postid), messagecount, signature);
CREATE TABLE MESSAGE(messageid PRIMARY KEY, idsend REFERENCES USER(userid), idrecv REFERENCES USER(userid), messagefile, attachmentcount);
CREATE TABLE POST(postid PRIMARY KEY, forumid REFERENCES FORUM(forumid), title, posterid REFERENCES USER(userid), postfile, date, issticky, isarchived, commentcount, lastcommentdate, attachmentcount);
CREATE TABLE COMMENT(commentid PRIMARY KEY, posterid REFERENCES USER(userid), postid REFERENCES POST(postid), commentfile, date, attachmentcount);
CREATE TABLE ATTACHMENT(attachmentid PRIMARY KEY, articleid, attachmentfile, FOREIGN KEY (articleid) REFERENCES POST(postid), FOREIGN KEY (articleid) REFERENCES COMMENT(commentid) );`)
        if err != nil {
                log.Fatal("%q:\n", err)
                return
        }


}

func main() {
initfile();
log.Println("Success!")
}
