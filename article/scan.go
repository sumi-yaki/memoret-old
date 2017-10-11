package article

import (
    "strings"
    "regexp"
)

func Scan (s string) string {
    r1 := regexp.MustCompile(`&`)
    r2 := regexp.MustCompile(` `)
    r3 := regexp.MustCompile(`<`)
    r4 := regexp.MustCompile(`>`)
    r5 := regexp.MustCompile(`"`)

    rhttp := regexp.MustCompile(`[hH]ttp://.*/.*$`); rhttp.Longest()
    rhttps := regexp.MustCompile(`[hH]ttps://.*/.*$`); rhttps.Longest()
    rftp := regexp.MustCompile(`[fF]tp://.*/.*$`); rftp.Longest()

    rets := ""
    arr := strings.Split(s,"\n")
    for i,v := range arr {

         if rhttp.MatchString(v) {
              http := rhttp.FindString(v)
              wohttp := rhttp.ReplaceAllString(v,"")
              http = "<a href=\"" + http +"\">" + http + "</a>"
              arr[i] = Scan(wohttp)+http
              continue
         }
         if rhttps.MatchString(v) {
              https := rhttps.FindString(v)
              wohttps := rhttps.ReplaceAllString(v,"")
              https = "<a href=\"" + https +"\">" + https + "</a>"
              arr[i] = Scan(wohttps)+https
              continue
         }
         if rftp.MatchString(v) {
              ftp := rftp.FindString(v)
              woftp := rftp.ReplaceAllString(v,"")
              ftp = "<a href=\"" + ftp +"\">" + ftp + "</a>"
              arr[i] = Scan(woftp)+ftp
              continue
         }

         v = r1.ReplaceAllString(v,"&amp;")
         v = r2.ReplaceAllString(v,"&nbsp;")
         v = r3.ReplaceAllString(v,"&lt;")
         v = r4.ReplaceAllString(v,"&gt;")
         v = r5.ReplaceAllString(v,"&quot;")

         arr[i] = v
    }
    for i,v := range(arr) {
         if i < len(arr) - 1 {
             v += "<br>\n"
         }

         rets += v
    }
    return rets
}
