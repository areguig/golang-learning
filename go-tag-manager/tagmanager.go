package main

import (
    "fmt"
    "net/http"
    "encoding/base64"
)

func init() {
    http.HandleFunc("/tag",getTag)
}



func getTag(w http.ResponseWriter, r *http.Request) {
	az := r.URL.Path[len("/tag?az="):]
	 decoded, _ := base64.URLEncoding.DecodeString(az)
	 fmt.Println("az: ",az,"\n decoded :",string(decoded))
	 fmt.Fprint(w, string(decoded))
	
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello, world!")
}