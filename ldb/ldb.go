package ldb

import (
    "fmt"
    "sync"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var once sync.Once

type database struct {
	instance *sql.DB
	cnt int
}

type Option func(db *database)

var db *database

func GetInstance() *sql.DB {
	db = new(database)
    if db.instance == nil {
        once.Do(
            func() {
                fmt.Println("Creating ldb instance now.")
				idb, err := sql.Open("sqlite3", "./example.db")
				db.instance = idb
				if err != nil {
					panic(err)
				}
                
            })
    } else {
        fmt.Println("ldb instance already created.")
    }

    return db.instance
}
 
func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}