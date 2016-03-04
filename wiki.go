// wiki / plaintext-snippets
// A simple wiki based on the tutorial "Writing Web Applications"
// See also https://golang.org/doc/articles/wiki/
package main

import (
	"html/template"
	"io/ioutil"
	"net/http"
)

//A wiki page has a title and a body (the page content)
type Page struct {
    Title string
    Body  []byte
}

//Save a wiki page by storing it's content in a textfile named by it's title
func (p *Page) save() error {
    filename := p.Title + ".txt"
    return ioutil.WriteFile(filename, p.Body, 0600)
}

//Load a wiki page by restoring it's content from a textfile named by it's title
func loadPage(title string) (*Page, error) {
    filename := title + ".txt"
    body, err := ioutil.ReadFile(filename)
    if err != nil {
        return nil, err
    }
    return &Page{Title: title, Body: body}, nil
}

//Render a html template
func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
    t, _ := template.ParseFiles(tmpl + ".html")
    t.Execute(w, p)
}

//Handler: View a wiki page
func viewHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/view/"):]
    p, _ := loadPage(title)
    renderTemplate(w, "view", p)
}

//Handler: Edit a wiki page
func editHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/edit/"):]
    p, err := loadPage(title)
    if err != nil {
        p = &Page{Title: title}
    }
    renderTemplate(w, "edit", p)
}

func main() {
    http.HandleFunc("/view/", viewHandler)
    http.HandleFunc("/edit/", editHandler)
    //http.HandleFunc("/save/", saveHandler)
    http.ListenAndServe(":8080", nil)
}
