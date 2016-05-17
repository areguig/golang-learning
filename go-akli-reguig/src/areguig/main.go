package main

import (
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

func init() {
	http.HandleFunc("/twitter/search", twitterSearch)
}

func twitterSearch(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	q := r.URL.Query().Get("q")
	log.Debugf(ctx, "Twitter search ", string(q))

}
