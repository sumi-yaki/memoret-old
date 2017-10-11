package mydb

import (
//    "fmt"
//    "time"
    _ "github.com/go-sql-driver/mysql"
//    "github.com/gocraft/dbr"
//    "github.com/gocraft/dbr/dialect"
//    "log"
)

// tableの設計

func GetArticle(idstart, idend int) []Articles {
//func GetArticle() {
    //SELECT
    var atcls []Articles
//    var atcls []Test

    var id0,id1 int
    var order bool
    if idstart <= idend{id0 = idstart;id1 = idend;order = true}
    if idstart > idend{id0 = idend;id1 = idstart;order = false}

    sess.Select("*").
        From("articles").
        Where("id >= ? AND id <= ?", id0, id1).
        OrderDir("id",order).
        Load(&atcls)

//    fmt.Println(atcls)
    return atcls
}

// https://eurie.co.jp/blog/engineering/2015/12/go-lang-ormapper-dbr
// http://kaworu.jpn.org/kaworu/2008-03-21-2.php
func LastId() int {
    var i []int
//    sess.SelectBySql("SELECT 'articles'; LAST_INSERT_ID()").Load(&i)
////    sess.SelectBySql("SELECT id FROM articles ORDER BY id DESC LIMIT 1").Load(&i)
    sess.SelectBySql("SELECT id FROM articles ORDER BY id DESC LIMIT 1").Load(&i)
    return i[0]
}
