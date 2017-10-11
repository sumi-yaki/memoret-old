package mydb

import (
//    "fmt"
    "time"
    _ "github.com/go-sql-driver/mysql"
    "github.com/gocraft/dbr"
//    "github.com/gocraft/dbr/dialect"
    "log"
)

// tableの設計
type Articles struct {
//type Test struct {
    Id              int64      `db:"id"`
//    LastModified  time.Time  `db:"last_modified"`
    LastModified    string     `db:"last_modified"`
    InputComment    string     `db:"input_comment"`
}

var conn *dbr.Connection
var sess *dbr.Session

func Init() {
    conn, _ = dbr.Open("mysql", "<IDENTITYTOMYSQL>", nil)
    sess = conn.NewSession(nil)
}

func NewArticle(text string) {

    // INSERT
        _, err := sess.InsertInto("articles").
        Columns("last_modified", "input_comment").
        Values(time.Now().UTC(), text).
        Exec()
 
    if err != nil {
        log.Fatal(err)
    } 
}

func UpdateArticle(id int, text string) {
        _, err := sess.Update("articles").
        Set("input_comment",text).
        Set("last_modified",time.Now().UTC()).
        Where("id = ?", id).
        Exec()

    if err != nil {
        log.Fatal(err)
    }
}
