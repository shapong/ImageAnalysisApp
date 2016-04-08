package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
    "encoding/json"
    "html/template"
)

type Page struct {
	APIKey string
	Curl string
}

var apiKey interface{}
var apiStr string
var ok bool
var curl_cmd string

func handler(w http.ResponseWriter, req *http.Request) {
    
    vs := os.Getenv("VCAP_SERVICES")
    
	var jsonblob = []byte(vs)
    var f interface{}
	err := json.Unmarshal(jsonblob, &f)
    m := f.(map[string]interface{})
	if err != nil {
	} else {

       for k, v := range m {
        _ = k
        switch vv := v.(type) {
            case []interface{}:
                
                for i, val := range vv {
                    _ = i
                    switch concreteVal := val.(type) {
                    case map[string]interface{}:
                        for key, val := range val.(map[string]interface{}) {
                            _ = key
                            switch concreteVal := val.(type) {
                            case map[string]interface{}:
                                
                                   for key, val := range val.(map[string]interface{}) {
                                       _ = key
                                        switch concreteVal := val.(type) {
                                        case map[string]interface{}:

                                        case []interface{}:

                                        default:
                                            apiKey = concreteVal
                                            apiStr, ok = apiKey.(string)
                                            if !ok {
                                                apiStr = ""
                                            }
                                            
                                            curl_cmd = "curl \"https://api.havenondemand.com/1/api/sync/ocrdocument/v1?url=https%3A%2F%2Fwww.havenondemand.com%2Fsample-content%2Fimages%2Fbowers.jpg&apikey=" + apiStr + "\""
                                            
                                            //fmt.Fprintln(w, "curl \"https://api.havenondemand.com/1/api/sync/ocrdocument/v1?url=https%3A%2F%2Fwww.havenondemand.com%2Fsample-content%2Fimages%2Fbowers.jpg&apikey=" + apiStr + "\"" )
                                            
                                        }
                                    }
                                
                                
                            case []interface{}:

                            default:
                                _ = concreteVal
                            }
                        }

                    default:
                        _ = concreteVal

                    }
                }
                    

            default:
          }
       }
    }
    
    curl_cmd = "curl \"https://api.havenondemand.com/1/api/sync/ocrdocument/v1?url=https%3A%2F%2Fwww.havenondemand.com%2Fsample-content%2Fimages%2Fbowers.jpg&apikey=" + apiStr + "\""
    p := &Page{APIKey: apiStr, Curl: curl_cmd}
    renderTemplate(w, "main", p)
    
	if req.URL.Path == "/crash" {
		os.Exit(1)
	}
}

var templates = template.Must(template.ParseFiles("main.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", handler)
	addr := ":" + os.Getenv("PORT")
	fmt.Printf("Listening on %v\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
