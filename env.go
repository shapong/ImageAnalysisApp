package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type (
	// User represents the structure of our resource
	User struct {
		ID     bson.ObjectId `json:"id" bson:"_id"`
		Name   string        `json:"name" bson:"name"`
		Gender string        `json:"gender" bson:"gender"`
		Age    int           `json:"age" bson:"age"`
	}
)

func NewUserController(s *mgo.Session) *UserController {
	return &UserController{s}
}

type UserController struct {
	session *mgo.Session
}

type Page struct {
	DB       string
	Host     string
	Hostname string
	Name     string
	Password string
	Port     string
	URI      string
	Username string
	NewUser  string
}

type ServiceDetails struct {
	Instances []InstanceDetails `json:"sharikamongo"`
}

type InstanceDetails struct {
	Creds Credentials `json:"credentials"`
}

type Credentials struct {
	DB       string `json:"db"`
	Host     string `json:"host"`
	Hostname string `json:"hostname"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Port     string `json:"port"`
	URI      string `json:"uri"`
	Username string `json:"username"`
}

var apiKey interface{}
var apiStr string
var ok bool
var curl_cmd string

func handler(w http.ResponseWriter, req *http.Request) {

	vs := os.Getenv("VCAP_SERVICES")
	//fmt.Fprintln(w, "hey111....................: "+vs)
	//fmt.Fprintln(w, strings.Join(os.Environ(), "\n"))

	var jsonblob = []byte(vs)
	//fmt.Fprintln(w, "hey....................: "+vs)

	var serviceInfo ServiceDetails
	err := json.Unmarshal(jsonblob, &serviceInfo)
	//fmt.Fprintln(w, "MY STRUCTURE..........................:"+serviceInfo.Instances[0].Creds.DB)
	p := &Page{
		Username: "<Insert here>",
		DB:       "<Insert here>",
		Port:     "<Insert here>",
		Password: "<Insert here>",
		Name:     "<Insert here>",
		Host:     "<Insert here>",
		Hostname: "<Insert here>",
		URI:      "<Insert here>",
		NewUser:  "<Insert here>",
	}
	if err != nil {
		renderTemplate(w, "main", p)
	} else {
		for _, instance := range serviceInfo.Instances {
			db := instance.Creds.DB
			host := instance.Creds.Host
			hostname := instance.Creds.Hostname
			name := instance.Creds.Name
			pwd := instance.Creds.Password
			port := instance.Creds.Port
			url := instance.Creds.URI
			user := instance.Creds.Username

			//fmt.Fprintf(w, "Heyyyy hii heloo ... and hamsika ..... ")
			uc := NewUserController(getSession(url))
			newUser, err := uc.CreateUser(w, "sharika", "female", 24, db)
			if err != nil {
				p = &Page{
					Username: user,
					Port:     port,
					Password: pwd,
					URI:      url,
					Name:     name,
					Hostname: hostname,
					Host:     host,
					DB:       db,
					NewUser:  "New user could not be added",
				}
			} else {
				p = &Page{
					Username: user,
					Port:     port,
					Password: pwd,
					URI:      url,
					Name:     name,
					Hostname: hostname,
					Host:     host,
					DB:       db,
					NewUser:  string(newUser),
				}
			}

		}
		renderTemplate(w, "main", p)
	}

	if req.URL.Path == "/crash" {
		os.Exit(1)
	}
	//renderTemplate(w, "main", p)
}

var templates = template.Must(template.ParseFiles("main.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func getSession(mongoURI string) *mgo.Session {
	// Connect to our local mongo
	s, err := mgo.Dial(mongoURI)

	// Check if connection error, is mongo running?
	if err != nil {
		panic(err)
	}
	return s
}

// CreateUser creates a new user resource
func (uc UserController) CreateUser(w http.ResponseWriter, name string, gender string, age int, db string) ([]byte, error) {
	// Stub an user to be populated from the body
	u := User{
		Name:   name,
		Gender: gender,
		Age:    age,
	}

	// Populate the user data
	//json.NewDecoder(r.Body).Decode(&u)

	// Add an Id
	u.ID = bson.NewObjectId()
	//fmt.Fprintf(w, "Heyyyy hii heloo ... and hamsika ..... %s", u.ID)

	// Write the user to mongo
	err := uc.session.DB(db).C("users").Insert(u)
	if err != nil {
		return nil, err
		//fmt.Fprintf(w, "Heyyyy hii heloo ... and hamsika ..... %s", err)
	}

	// Marshal provided interface into JSON structure
	uj, err := json.Marshal(u)
	if err != nil {
		return nil, err
		//fmt.Fprintf(w, "Heyyyy hii heloo ... and hamsika ..... %s", err)
	}
	// Write content-type, statuscode, payload
	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(201)
	//fmt.Fprintf(w, "Heyyyy hii heloo ... and hamsika ..... %s", uj)
	return uj, nil
}

// GetUser retrieves an individual user resource
func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, id string, db string) {
	// Grab id
	//id := p.ByName("id")

	// Verify id is ObjectId, otherwise bail
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(404)
		return
	}

	// Grab id
	oid := bson.ObjectIdHex(id)

	// Stub user
	u := User{}

	// Fetch user
	if err := uc.session.DB(db).C("users").FindId(oid).One(&u); err != nil {
		w.WriteHeader(404)
		return
	}

	// Marshal provided interface into JSON structure
	uj, _ := json.Marshal(u)

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", uj)
}

func main() {
	http.HandleFunc("/", handler)
	addr := ":" + os.Getenv("PORT")
	fmt.Printf("Listening on %v\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
