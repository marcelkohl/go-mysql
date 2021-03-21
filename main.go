package main

import (
  "fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Tag struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}

func main() {
  fmt.Println("Go MySQL")

  db, err := sql.Open("mysql", "admin:admin@tcp(127.0.0.1:8306)/docker_db")

  if err != nil {
    panic(err.Error())
  } else if db.Ping() != nil {
    fmt.Println(db.Ping().Error())
  } else {
    fmt.Println(" Connected!")
  }

  // Insert record
  insert, err := db.Query("INSERT INTO sample1 (name) VALUES ('Test')")

  if err != nil {
    panic(err.Error())
  } else if insert.Err() != nil {
    fmt.Println(insert.Err().Error())
  } else {
    fmt.Println(" Record inserted")
  }

  // Select and print records
  // results, err := db.Query("SELECT id, name FROM sample1")
  //
  // if err != nil {
  //   panic(err.Error())
  // } else if results.Err() != nil {
  //   fmt.Println(results.Err().Error())
  // }
  //
  // for results.Next() {
  //   var tag Tag
  //
  //   err = results.Scan(&tag.ID, &tag.Name)
  //
  //   if err != nil {
  //     panic(err.Error())
  //   }
  //
  //   fmt.Println(tag.Name)
  // }

  // Select just one row
  // var tag Tag
  //
  // err = db.QueryRow("SELECT id, name FROM sample1 WHERE id=?", 2).Scan(&tag.ID, &tag.Name)
  //
  // if err != nil {
  //   panic(err.Error())
  // }
  //
  // fmt.Println(fmt.Sprintf("%v - %v", tag.ID, tag.Name))

  defer db.Close()
}
