package main

import (
  "fmt"
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
)

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

  insert, err := db.Query("INSERT INTO sample1 (name) VALUES ('Test')")

  if err != nil {
    panic(err.Error())
  } else if insert.Err() != nil {
    fmt.Println(insert.Err().Error())
  } else {
    fmt.Println(" Record inserted")
  }

  defer db.Close()
}
