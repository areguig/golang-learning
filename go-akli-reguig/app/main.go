package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"net/url"

	"github.com/ChimeraCoder/anaconda"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/urlfetch"
)

func init() {
	http.HandleFunc("/twitter/search", twitterSearch)
}

func twitterSearch(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	query := r.URL.Query().Get("q")
	hash := r.URL.Query().Get("h")
	log.Debugf(ctx, "Twitter search", string(query))
	anaconda.SetConsumerKey("cW0kdWCjgnE8vpJGOvUxe4epL")
	anaconda.SetConsumerSecret("GEcenuc4kLzZLAfYddfC3PovRVdAu3CL3n9sc61zikH4wK2eDw")
	api := anaconda.NewTwitterApi("", "")
	api.HttpClient.Transport = &urlfetch.Transport{Context: ctx}
	v := url.Values{
		"result_type":      {"mixed"},
		"count":            {"1000"},
		"include_entities": {"false"},
	}
	if hash == "true" {
		query = "#" + query
	} else {
		query = "@" + query
	}
	searchResult, _ := api.GetSearch(url.QueryEscape(string(query)), v)
	js, err := json.Marshal(searchResult.Statuses[rand.Intn(len(searchResult.Statuses))])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
