package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type Post struct {
	Id   int
	Name string
	Text string
}

func main() {

	//open a connection
	db, e := sql.Open("mysql", "<username>:<password>@/<db-name>")
	ErrorCheck(e)

	//closing db connection after all work done
	defer db.Close()

	PingDB(db)

	//insert into DB
	//prepare
	stmt1, e := db.Prepare("insert into post(Name,Text)values(?,?)")
	ErrorCheck(e)

	//execute
	res1, e := stmt1.Exec("Nature", "content about nature")
	ErrorCheck(e)

	id1, e := res1.LastInsertId()
	ErrorCheck(e)

	fmt.Println("inserted id", id1)

	//insert
	//prepare
	stmt2, e := db.Prepare("insert into post(Name,Text)values(?,?)")
	ErrorCheck(e)

	//execute
	res2, e := stmt2.Exec("Animal", "content about animal")
	ErrorCheck(e)

	id2, e := res2.LastInsertId()
	ErrorCheck(e)

	fmt.Println("inserted id", id2)

	//insert
	//prepare
	stmt3, e := db.Prepare("insert into post(Name,Text)values(?,?)")
	ErrorCheck(e)

	//execute
	res3, e := stmt3.Exec("Birds", "content about birds")
	ErrorCheck(e)

	id3, e := res3.LastInsertId()
	ErrorCheck(e)

	fmt.Println("inserted id", id3)

	//insert
	//prepare
	stmt4, e := db.Prepare("insert into post(Name,Text)values(?,?)")
	ErrorCheck(e)

	//execute
	res4, e := stmt4.Exec("Trees", "content about trees")
	ErrorCheck(e)

	id4, e := res4.LastInsertId()
	ErrorCheck(e)

	fmt.Println("inserted id", id4)

	//insert
	//prepare
	stmt5, e := db.Prepare("insert into post(Name,Text)values(?,?)")
	ErrorCheck(e)

	//execute
	res5, e := stmt5.Exec("Water", "content about water")
	ErrorCheck(e)

	id5, e := res5.LastInsertId()
	ErrorCheck(e)

	fmt.Println("inserted id", id5)

	//query all data
	rows, e := db.Query("select * from post")
	ErrorCheck(e)

	var posts = Post{}

	for rows.Next() {
		e = rows.Scan(&posts.Id, &posts.Name, &posts.Text)
		ErrorCheck(e)
		fmt.Println(posts)
	}

	//delete data
	stmt, e := db.Prepare("delete from post where id=?")
	ErrorCheck(e)

	res, e := stmt.Exec("5")
	ErrorCheck(e)

	a, e := res.RowsAffected()
	ErrorCheck(e)

	fmt.Println(a)

}

func PingDB(db *sql.DB) {
	err := db.Ping()
	ErrorCheck(err)
}

func ErrorCheck(err error) {
	if err != nil {
		panic(err.Error())
	}
}
