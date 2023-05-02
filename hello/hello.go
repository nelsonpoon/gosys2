package main

import (
    "fmt"
    "golang.org/x/example/stringutil"
    "example.com/ui"
    //"example/stringutil"
    "example.com/ldb"
    "strconv"
    "time"
)

const create string = `
  CREATE TABLE IF NOT EXISTS userinfo (
    username VARCHAR(255) NULL,
    department VARCHAR(255) NULL,
    created DATETIME NULL
  );`

func main() {
    mldb := ldb.GetInstance()
    

    _, err := mldb.Exec(create)
    checkErr(err)

    stmt, err := mldb.Prepare("INSERT INTO userinfo(username, department, created) values(?,?,?)")
    checkErr(err)

    res, err := stmt.Exec("astaxie", "研發部門", "2012-12-09")
    checkErr(err)
    
    id, err := res.LastInsertId()
    checkErr(err)
    fmt.Println("Add record"+ strconv.FormatInt(id,10))
    
    

    time.Sleep(1 * 100 * time.Millisecond)

    /*rows, err := mldb2.Query("SELECT * FROM userinfo")
    checkErr(err)
    
    for rows.Next() {
        var un string
        var dp string
        var credt time.Time
        err = rows.Scan(&un,&dp,&credt)
        checkErr(err)
        fmt.Println(un)
        fmt.Println(dp)
        fmt.Println(credt)
    }*/

    
    fmt.Println(stringutil.ToUpper("Hello"))
    window := ui.createApp()
    ui.WelcomeWin(window)
    ui.LoginWin(window)
}

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}