package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	 
)


func main() {

	resp, err := http.Get("http://azalead.com/")
	if err != nil {
		fmt.Println("eRROr ",err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("eRROr ",err)
	}
	fmt.Println(string(body))

}

