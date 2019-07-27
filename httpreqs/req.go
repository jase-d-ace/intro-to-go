package httpreqs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var myClient = &http.Client{Timeout: 10 * time.Second}

//structs need capital letters. go figure.
//json solution found here: https://stackoverflow.com/questions/17156371/how-to-get-json-response-in-golang

type Message struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func getJson(target interface{}) error {
	res, err := myClient.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return err
	}
	return json.Unmarshal(body, target)
}

func MakeReq() {
	m := Message{}
	getJson(&m)
	newJson, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", newJson)
}
