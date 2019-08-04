package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

var myClient = &http.Client{Timeout: 10 * time.Second}

//empty struct that will eventually hold data from the API that we hit on a route.

type Message struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func splashPage(w http.ResponseWriter, r *http.Request) {
	message := "hello world!"
	fmt.Fprintf(w, message)
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Path
	message = strings.TrimPrefix(message, "/")
	message = "Hello " + message
	fmt.Fprintf(w, message)
}

func getJsonArray(w http.ResponseWriter, r *http.Request) {

	res, err := myClient.Get("https://jsonplaceholder.typicode.com/todos")

	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	var todos []Message

	json.Unmarshal([]byte(body), &todos)

	//proper json array solution found here: https://stackoverflow.com/questions/28411394/golang-and-json-with-array-of-struct

	jsonData, _ := json.Marshal(todos)

	w.Write(jsonData)

}

//w and r are the (req, res) callbacks that javascript uses to handle server requests

func getSingleJson(w http.ResponseWriter, r *http.Request) {
	res, err := myClient.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}
	m := Message{}

	json.Unmarshal(body, &m)

	newJson, err := json.Marshal(m)

	if err != nil {
		panic(err)
	}

	w.Write(newJson)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", splashPage)
	router.HandleFunc("/world", sayHello)
	router.HandleFunc("/todos", getJsonArray)
	if err := http.ListenAndServe(":8080", router); err != nil {
		panic(err)
	}
}
