package mytemplates

const Postq = 1

func Formnew(title, htmlArticle, hiddenId, inputText string) string {
    head1 := 
`<html>
   <head>
       <meta charset="UTF-8">
       <title>`

    head2 :=
`</title>
       <style type="text/css">
           textarea {
               width: 90%;
               height: 70%;
           }
       </style>
   </head>
        <body>`
    
    form1 :=
`              <form action="/m/newpost" method="post">
                  <input type='submit' name='text' value='text'/>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
                  <input type='hidden' name='hiddenid' value='`
   form15 :=
`'>
                  <input type='reset' name='clear' value='clear' />&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
                  <input type='submit' name='ok' value='ok' /><br>
                  <textarea name='inputComment'>`

    form2 := 
`</textarea><br>
                  <input type='submit' name='text' value='text'/>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
                  <input type='reset' name='clear' value='clear' />&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
                  <input type='submit' name='ok' value='ok' />
              </form>
        </body>
</html>`

    return head1+title+head2+htmlArticle+form1+hiddenId+form15+inputText+form2
}
