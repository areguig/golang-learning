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


// >>>>>>>>> Web stuff. 
// Renders the html template tmpl filled with p content 
// in the writer w
func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	t, _ := template.ParseFiles(tmpl+".html")
	t.Execute(w,p)
}

// Handles /view/* endpoints
func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, err := loadPage (title)
	if err !=nil{
		http.Redirect(w,r,"/edit/"+title,http.StatusFound)
	}
	renderTemplate(w,"view",p)
}

// Handles /edit/* endpoints
func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title:title}
	}
	renderTemplate(w,"edit",p)
}

// Handle /save/* endpoints
func saveHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/save/"):]
	body := r.FormValue("body")
	p := &Page{Title:title,Body:[]byte(body)}
	p.save()
	http.Redirect(w,r,"/view"+title,http.StatusFound)
}

// Main 
func main() {
	http.HandleFunc("/view/",viewHandler)
	http.HandleFunc("/edit/",editHandler)
	http.HandleFunc("/save/",saveHandler)
	fmt.Println("Starting the server...")
	http.ListenAndServe(":8080",nil)
}

