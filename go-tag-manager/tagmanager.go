package main

import (
    "fmt"
    "net/http"
    "html/template"
    "encoding/base64"

    "appengine"
    "appengine/user"
)

func init() {
    http.HandleFunc("/tag",getTag)
    http.HandleFunc("/",index)
}



func getTag(w http.ResponseWriter, r *http.Request) {
	 az := r.URL.Query().Get("az")
   decoded, _ := base64.URLEncoding.DecodeString(az)
	 fmt.Println("az: ",az,"\n decoded :",string(decoded))
	 fmt.Fprint(w, string(decoded))

}

func index(w http.ResponseWriter, r *http.Request) {
  c := appengine.NewContext(r)
  u := user.Current(c)
  if u == nil {
    url, err := user.LoginURL(c, r.URL.String())
    if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }
    w.Header().Set("Location", url)
    w.WriteHeader(http.StatusFound)
    return
  }
  t, _ := template.ParseFiles("templates/index.html")
  t.Execute(w,u)
}
