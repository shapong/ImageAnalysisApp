package main

import (
	"html/template"
    "encoding/json"
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

var apiKey interface{}
func loadPage() (*HavenAPI, error) {
    
    vs := os.Getenv("VCAP_SERVICES")
    
	var jsonblob = []byte(vs)
	// var animals Blah
    var f interface{}
	err := json.Unmarshal(jsonblob, &f)
    m := f.(map[string]interface{})

	if err != nil {
		//fmt.Fprintln(w, "error:", err)
	} else {

       for k, v := range m {
        switch vv := v.(type) {
            case []interface{}:
                //parseArray(vv)
                _ = k
                
                for i, val := range vv {
                    _ = i
                    switch concreteVal := val.(type) {
                    case map[string]interface{}:
                        //fmt.Fprintln(w, i)
                        for key, val := range val.(map[string]interface{}) {
                            _ = key
                            switch concreteVal := val.(type) {
                            case map[string]interface{}:
                                //fmt.Fprintln(w, key)
                                
                                   for key, val := range val.(map[string]interface{}) {
                                       _ = key
                                       switch concreteVal := val.(type) {
                                       case map[string]interface{}:
                                            //fmt.Fprintln(w, key)
                                            //parseMap(val.(map[string]interface{}))
                                       case []interface{}:
                                            //fmt.Fprintln(w, key)
                                            //parseArray(val.([]interface{}))
                                       default:
                                            apiKey = concreteVal
                                            //fmt.Fprintln(w, key, "hamsika:", concreteVal)
                                       }
                                    }
                                
                                
                            case []interface{}:
                                //fmt.Fprintln(w,"hamsi", key)
                                //parseArray(val.([]interface{}))
                            default:
                                _ = concreteVal
                                //fmt.Fprintln(w, key, ":", concreteVal)
                            }
                        }

                    default:
                        _ = concreteVal
                        //fmt.Fprintln(w, "Index", i, ":", concreteVal)

                    }
                }
                    

            default:
                //fmt.Fprintln(w, k, "is of a type I don't know how to handle")
          }
       }
    }
    
    //apiKey := os.Getenv("VCAP_SERVICES")
    apiStr, ok := apiKey.(string)
    if !ok {
        return &HavenAPI{Title: ""}, errors.New("Missing API Key")
    }
    
	return &HavenAPI{Title: apiStr}, nil
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	p, err := loadPage()
	if err != nil {
		p = &HavenAPI{Title: ""}
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
