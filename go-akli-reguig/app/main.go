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
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	urlshortener "google.golang.org/api/urlshortener/v1"

)

func init() {
	http.HandleFunc("/twitter/search", handleTwitterSearch)
	http.HandleFunc("/shorten", handleShorten)
	http.HandleFunc("/", handleGoToNewURL)
}


func handleGoToNewURL(w http.ResponseWriter, r *http.Request) {
  http.Redirect(w, r, "https://areguig.github.io/", http.StatusMovedPermanently)
}

func handleTwitterSearch(w http.ResponseWriter, r *http.Request) {
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
	w.Header().Set("Access-Control-Allow-Origin","https://areguig.github.io")
	w.Write(js)
}

func handleShorten(w http.ResponseWriter, r *http.Request) {
	longUrl :=  r.FormValue("url")
	//longUrl := r.URL.Query().Get("url")
	ctx := appengine.NewContext(r)
	shortenedRes, _:=shortenURL(ctx,longUrl)
	w.Header().Set("Access-Control-Allow-Origin","https://areguig.github.io")
	w.Write([]byte(shortenedRes))
}


func shortenURL(ctx context.Context, url string) (string, error) {
        transport := &oauth2.Transport{
                Source: google.AppEngineTokenSource(ctx, urlshortener.UrlshortenerScope),
                Base:   &urlfetch.Transport{Context: ctx},
        }
        client := &http.Client{Transport: transport}
        svc, err := urlshortener.New(client)
        if err != nil {
                return "", err
        }
        resp, err := svc.Url.Insert(&urlshortener.Url{LongUrl: url}).Do()
        if err != nil {
                return "", err
        }
        return resp.Id, nil
}
