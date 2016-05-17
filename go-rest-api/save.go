package main

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/kurrik/oauth1a"
	"github.com/kurrik/twittergo"

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
	client, err = getCredentials()
	query := url.Values{}
	query.Set("q", "twitterapi")
	url := fmt.Sprintf("/1.1/search/tweets.json?%v", query.Encode())
	req, err = http.NewRequest("GET", url, nil)
	if err != nil {
		log.Debugf(ctx, "Could not parse request: %v\n", err)
	}
	resp, err = client.SendRequest(req)
	if err != nil {
		log.Debugf(ctx, "Could not send request: %v\n", err)
	}
	results = &twittergo.SearchResults{}
	err = resp.Parse(results)
	if err != nil {
		log.Debugf(ctx, "Problem parsing response: %v\n", err)
	}
	for i, tweet := range results.Statuses() {
		user := tweet.User()
		log.Debugf(ctx, "%v.) %v\n", i+1, tweet.Text())
		log.Debugf(ctx, "From %v (@%v) ", user.Name(), user.ScreenName())
		log.Debugf(ctx, "at %v\n\n", tweet.CreatedAt().Format(time.RFC1123))
	}

}

func getCredentials() (client *twittergo.Client, err error) {

	config := &oauth1a.ClientConfig{
		ConsumerKey:    "cW0kdWCjgnE8vpJGOvUxe4epL",
		ConsumerSecret: "GEcenuc4kLzZLAfYddfC3PovRVdAu3CL3n9sc61zikH4wK2eDw",
	}
	client = twittergo.NewClient(config, nil)
	return
}
