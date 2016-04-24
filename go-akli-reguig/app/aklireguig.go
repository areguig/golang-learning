package main

import "net/http"

func init() {
	http.HandleFunc("/", index)

}

func index(w http.ResponseWriter, r *http.Request) {
	//http.Redirect(w, r, "/cv", http.StatusMovedPermanently)
}

//func static(w http.ResponseWriter, r *http.Request) {
//http.ServeFile(w, r, "static/"+r.URL.Path)
//}
