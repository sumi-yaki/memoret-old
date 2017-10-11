package main

import (
//    "fmt"
    "./mytemplates"
    "./article"
    "./mydb"
    "net/http"
    "gopkg.in/labstack/echo.v1"
    mw "gopkg.in/labstack/echo.v1/middleware"
    "strconv"
)

const TLSkeyDir = "<TLSKEYDIR>"

func main(){

    mydb.Init()

    e := echo.New()
    e.Use(mw.Logger())
    e.Use(mw.Recover())
    e.Use(mw.BasicAuth(func(usr, pwd string) bool {
        if usr == "<YOURUSERNAME>" && pwd == "<PASSWORD>" {
                return true
        }
                return false
    }))

    m := e.Group("/m")

    m.Get("/new", func(c *echo.Context) error {
                     return c.HTML(http.StatusOK, mytemplates.Formnew("Input new article","","0",""))
    })
    m.Post("/newpost", func(c *echo.Context) error {
                          res := c.Form("inputComment")
                          hiddenid := c.Form("hiddenid")
                          if c.Form("text") == "text" {
                              return c.HTML(http.StatusOK,mytemplates.Formnew("New article",article.Scan(res),hiddenid,res))
                          } else { 
                              if hiddenid == "0" {
                                    mydb.NewArticle(res)
                                    return c.String(http.StatusOK, "OK submitted!")
                              } else {
                                    ihiddenid,_ := strconv.Atoi(hiddenid)
                                    mydb.UpdateArticle(ihiddenid,res)
//                                    return c.Redirect(http.StatusTemporaryRedirect,"/m/l/"+hiddenid+"/"+hiddenid)
                                    return c.Redirect(http.StatusMovedPermanently,"/m/l/"+hiddenid+"/"+hiddenid)
                              }
                         }
    })
    m.Get("/l/:start/:end", func(c *echo.Context) error {
                          istart, _ := strconv.Atoi(c.Param("start"))
                          iend, _ := strconv.Atoi(c.Param("end"))
                          articles := mydb.GetArticle(istart,iend)
                          artcl := mytemplates.ShowArticles("Memoret articles",articles)
                          return c.HTML(http.StatusOK, artcl)
    })

    m.Get("/sh/:last/:num", func(c *echo.Context) error {
                          ilast,_ := strconv.Atoi(c.Param("last"))
                          inum,_ := strconv.Atoi(c.Param("num"))
                          articles := mydb.GetArticle(mydb.LastId()-ilast+1,mydb.LastId()-inum+1)
                          artcl := mytemplates.ShowArticles("Memoret articles",articles)
                          return c.HTML(http.StatusOK, artcl)
    })

    m.Get("/e/:id", func(c *echo.Context) error {
                     hiddenid := c.Param("id")
                     ihiddenid,_ := strconv.Atoi(hiddenid)
                     atcls := mydb.GetArticle(ihiddenid,ihiddenid)
                     aic := atcls[0].InputComment
         
                     return c.HTML(http.StatusOK, mytemplates.Formnew("Edit an article",article.Scan(aic),hiddenid,aic))
    })

    m.Get("/a/:id", func(c *echo.Context) error {
                         return c.Redirect(http.StatusMovedPermanently, "/m/l/"+c.Param("id")+"/"+c.Param("id"))
    })
    e.RunTLS(":<PORTNUMBER>",TLSkeyDir+"/public_key", TLSkeyDir+"/private_key")
}
