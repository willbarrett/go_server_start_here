package main

import (
  "database/sql"
  "github.com/gorilla/mux"
  "net/http"
  "os"
  "log"
  
  _ "github.com/mattn/go-sqlite3"
)

var database *sql.DB

func HomeHandler(w http.ResponseWriter, request *http.Request) {
  w.Write([]byte("Hello World"))
}

func OpenDatabaseConnection() {
  var err error
  database, err = sql.Open("sqlite3", "file:development.db")
  if err != nil {
    log.Fatal(err)
    os.Exit(1)
  }
}

func execStatementOrCrash(sql string) {
  statement, err := database.Prepare("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, name TEXT)")
  if err != nil {
    log.Fatal(err)
    os.Exit(1)
  }
  statement.Exec()
}

func SetupDatabase() {
  execStatementOrCrash("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, name TEXT)")
}

func main() {
  OpenDatabaseConnection()
  SetupDatabase()

  r := mux.NewRouter()
  r.HandleFunc("/", HomeHandler)
  http.Handle("/", r)
  http.ListenAndServe(":8080", nil)
}
