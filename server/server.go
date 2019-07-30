package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

var myClient = &http.Client{Timeout: 10 * time.Second}

type Message struct {
	UserId    int  `json:"userId"`
	Id        int  `json:"id"`
	Title     int  `json:"title"`
	Completed bool `json:"completed"`
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
	http.HandleFunc("/", splashPage)
	http.HandleFunc("/world", sayHello)
	http.HandleFunc("/todos", getSingleJson)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
