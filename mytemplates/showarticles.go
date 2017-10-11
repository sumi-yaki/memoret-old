package mytemplates

import (
    "../mydb/"
    "../article"
    "fmt"
    "time"
)

func ShowArticles(title string, atcls []mydb.Articles) string {
    s := ""
    head1 := 
`<html>
   <head>
       <meta charset="UTF-8">
       <title>`

    head2 :=
`</title>
   </head>
        <body>`
    
    tail :=
`        </body>
</html>`

    for _,v := range atcls {
        idstring := fmt.Sprintf("%d",v.Id)
        s += "ID: "+idstring+"<br>\n" 

        layout2 := "2006-01-02 15:04:05"
        t,_ := time.Parse(layout2,v.LastModified)
        locjst := time.FixedZone("JST", 9*60*60)
        t = t.In(locjst)
        s += "LastModified: "+t.String()+"<br>\n" 

        s += article.Scan(v.InputComment)+"<br>\n"
        s += "<a href=\"/m/e/"+ idstring + "\">Edit</a>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href=\"/m/new\">New</a><hr>\n"
    }

    return head1+title+head2+s+tail
}
