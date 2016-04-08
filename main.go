// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"html/template"
	// "io/ioutil"
	"net/http"
    "errors"
    "os"
	// "regexp"
)

//HavenAPI struct
type HavenAPI struct {
	Title string
}

func loadPage() (*HavenAPI, error) {
    apiKey := os.Getenv("VCAP_SERVICES")
    if apiKey == "" {
        return &HavenAPI{Title: ""}, errors.New("Missing API Key")
    }
	return &HavenAPI{Title: apiKey}, nil
}

// func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
// 	p, err := loadPage(title)
// 	if err != nil {
// 		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
// 		return
// 	}
// 	renderTemplate(w, "view", p)
// }

func editHandler(w http.ResponseWriter, r *http.Request) {
	p, err := loadPage()
	if err != nil {
		p = &HavenAPI{Title: "468f4fa3-530e-4898-8301-e98e32c43591"}
	}
	renderTemplate(w, "edit", p)
}

var templates = template.Must(template.ParseFiles("edit.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, p *HavenAPI) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	// http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/", editHandler)
	// http.HandleFunc("/save/", makeHandler(saveHandler))

	http.ListenAndServe(":8080", nil)
}
