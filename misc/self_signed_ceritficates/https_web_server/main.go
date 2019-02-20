package main

// Code to demostrate how to upload multipart forms to golang and save the files.

/*
 * In this exercise we will start a web server. The web server will take the
 * username phone number and email id of users.  We will give methods to create
 * a new user. List all the users. Find a user, delete a user. All of this will
 * be done based on the username which will be case sensitive.
 *
 * The users will be saved in a array of users. This array will be later saved
 * in a redis database which will be persistent.
 *
 * Going forward we will add user profile and authentication to this
 * application. The application will be oauth based.
 *
 * Users with profiles will be able to see the contacts saved by them only.
 *
 * The frontend will be written in react. The backend will be written in golang
 * and it will be deployed on kubernetes.
 */

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

const (
	PORT_NUMBER = "9999"
)

type user_struct struct {
	Name   string `json: name`
	Phone1 string `json: phone1`
	Email  string `json: email`
}

var userDatabase = make([]user_struct, 0)

// formatRequest generates ascii representation of a request
func formatRequest(r *http.Request) string {
	// https://medium.com/doing-things-right/pretty-printing-http-requests-in-golang-a918d5aaa000
	// Create return string
	var request []string
	// Add the request string
	url := fmt.Sprintf("%v %v %v", r.Method, r.URL, r.Proto)
	request = append(request, url)
	// Add the host
	request = append(request, fmt.Sprintf("Host: %v", r.Host))
	// Loop through headers
	for name, headers := range r.Header {
		name = strings.ToLower(name)
		for _, h := range headers {
			request = append(request, fmt.Sprintf("%v: %v", name, h))
		}
	}

	// If this is a POST, add post data
	if r.Method == "POST" {
		r.ParseForm()
		request = append(request, "\n")
		request = append(request, r.Form.Encode())
	}
	// Return the request as a string
	return strings.Join(request, "\n")
}

// printRequestBody will print the body of the request.
func printRequestBody(req *http.Request) {
	var bodyBytes []byte
	var err error
	if bodyBytes, err = ioutil.ReadAll(req.Body); err != nil {
		fmt.Printf("Error")
	}

	fmt.Printf("\n\nBody is\n%s\n\n", string(bodyBytes))
}

// PostUsersHandlerUrlEncodedForm will take urlencoded form and create a dictionary.
func PostUsersHandlerUrlEncodedForm(w http.ResponseWriter, req *http.Request) {

	req.ParseForm()
	for k, v := range req.Form {
		fmt.Println("keyx:", k)
		fmt.Println("valx:", strings.Join(v, ""))
	}
}

func PostUsersHandlerMultiPartFormWithImage(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile("uploadfile")

	fmt.Println("here 1")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Println("here 2")
	fmt.Fprintf(w, "%v", handler.Header)
	f, err := os.OpenFile("image_"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("here 3")
	defer f.Close()
	io.Copy(f, file)
}

func PostUsersHandlerMultiPartFormWithVideo() {
}

func PostUsersHandlerMultiPartFormWithResume() {
}

func PostUsersHandlerMultiPartFormWithTextFile() {
}

// PostUsersHandlerMultiPartForm will take a multi part form and save the user in the database.
func PostUsersHandlerMultiPartForm(w http.ResponseWriter, req *http.Request) {

	req.ParseForm()
	req.ParseMultipartForm(10000)
	fmt.Println("FORM DATA", req.FormValue("name"))
	fmt.Println("Printing req data ----------XX ", req.PostForm, "--------XXXXXX--------------")
	fmt.Println("Printing form data\n", strings.Repeat("#", 80))
	for k, v := range req.Form {
		fmt.Println("keyx:", k)
		fmt.Println("valx:", strings.Join(v, ""))
	}
}

// PostUsersHandlerREST takes a JSON formatted input, parses it and adds to the database.
func PostUsersHandlerREST(w http.ResponseWriter, req *http.Request) {
	var err error
	fmt.Println("Printing formmated request", formatRequest(req))
	fmt.Println("Printing request", req)
	req.ParseForm()
	fmt.Println("Printing req data", req.Form)
	fmt.Println("path", req.URL.Path)
	fmt.Println("scheme", req.URL.Scheme)
	req.ParseForm()
	fmt.Println("Printing the form sent", req.Form)

	var bodyBytes []byte
	if bodyBytes, err = ioutil.ReadAll(req.Body); err != nil {

		fmt.Printf("Error")
	}

	fmt.Printf("\n\nBody Is\n%s\n\n", string(bodyBytes))

	fmt.Println("\n\nRequest is\n", req)
	var t user_struct
	err = json.Unmarshal(bodyBytes, &t)
	if err != nil {
		fmt.Println("Error is ", err)
		panic(err)
	}
	fmt.Println("Structure is ", t)
	userDatabase = append(userDatabase, t)
	fmt.Println("User Data base is ", userDatabase)
}

func PostRootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "GetRootHandlerCalled %s %s", PORT_NUMBER, r.URL.Path[1:])
}

func GetRootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "GetRootHandlerCalled %s %s", PORT_NUMBER, r.URL.Path[1:])
}

// RootHandler will just return something to the user.
func RootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "GetRootHandlerCalled %s %s", PORT_NUMBER, r.URL.Path[1:])
}

// GetUsersHandler convert the users database to a JSON string and send as reply.
func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	/*
	 * fmt.Fprintf(w, "Port: %s GET Hi !!!!!!!! there, I was sent %s! ", PORT_NUMBER, r.URL.Path[1:])
	 */

	userDatabaseJson, err := json.Marshal(userDatabase)

	if err != nil {
		fmt.Println("Error in marhshalling")
	} else {
		fmt.Println("Marshalled data is ", userDatabaseJson)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	fmt.Fprintf(w, "%s", string(userDatabaseJson))
}

// PostUsersHandler handles the requests which are of the application/json content type.
func PostUsersHandler(w http.ResponseWriter, req *http.Request) {
	var err error
	fmt.Println("JSON JSONPrinting formmated request", formatRequest(req))
	fmt.Println("Printing request", req)
	req.ParseForm()
	fmt.Println("Printing req data", req.Form)

	fmt.Println("\n\nResponse is\n", req)
	decoder := json.NewDecoder(req.Body)
	var t user_struct
	err = decoder.Decode(&t)
	if err != nil {
		fmt.Println("Error is ", err)
		panic(err)
	}
	fmt.Println("Structure is ", t)
	userDatabase = append(userDatabase, t)
	fmt.Println("User Data base is ", userDatabase)
}

func main() {
	log.SetOutput(os.Stderr)
	router := mux.NewRouter()

	router.HandleFunc("/", RootHandler)
	router.HandleFunc("/", GetRootHandler).Methods("GET")
	router.HandleFunc("/", PostRootHandler).Methods("POST")
	router.HandleFunc("/users", GetUsersHandlerPost).Methods("GET", "OPTIONS")
	router.HandleFunc("/users", PostUsersHandler).Methods("POST")
	router.HandleFunc("/users", GetUsersHandler).Methods("GET")
	router.HandleFunc("/usersUEF", PostUsersHandlerUrlEncodedForm).Methods("POST")
	router.HandleFunc("/usersMPF", PostUsersHandlerMultiPartForm).Methods("POST")
	router.HandleFunc("/usersMPFWI", PostUsersHandlerMultiPartFormWithImage).Methods("POST")
	fmt.Printf("Running on port: %s", PORT_NUMBER)
	log.Fatal(http.ListenAndServeTLS(":"+PORT_NUMBER, "server.crt", "server.key", router))
}

// GetUsersHandlerPost
func GetUsersHandlerPost(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request is ", r)
	var err error
	var bodyBytes []byte
	if bodyBytes, err = ioutil.ReadAll(r.Body); err != nil {

		fmt.Printf("Error")
	}

	fmt.Printf("\n\nBody Is\n%s\n\n", string(bodyBytes))

	(w).Header().Set("Content-Type", "application/json")
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	json.NewEncoder(w).Encode("OKOK")
}
