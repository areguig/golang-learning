package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"html/template"
)

type Page struct {
	Title string
	Body []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	fmt.Println("Saving the file ",filename)
	return ioutil.WriteFile(filename,p.Body,0600)
}

func loadPage(title string) (*Page,error) {
	filename := title+".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil{
		return nil,err
	}
	return &Page{Title:title,Body:body}, nil
}


// Web stuff. 

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, _ := loadPage (title)
	t, _ := template.ParseFiles("view.html")
	t.Execute(w,p)
	
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title:title}
	}
	t, _ := template.ParseFiles("edit.html")
	t.Execute(w,p)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	//title := r.URL.Path[len("/edit/"):]
	
}


func main() {
	http.HandleFunc("/view/",viewHandler)
	http.HandleFunc("/edit/",editHandler)
	fmt.Println("Starting the server...")
	http.ListenAndServe(":8080",nil)
}

