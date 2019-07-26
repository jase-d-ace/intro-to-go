package httpreqs

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var myClient = &http.Client{Timeout: 10 * time.Second}

//structs need capital letters. go figure.

type Message struct {
	UserId    float64
	Id        float64
	Title     string
	Completed bool
}

func getJson(target interface{}) error {
	res, err := myClient.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return json.NewDecoder(res.Body).Decode(target)
}

func MakeReq() {
	m := Message{}
	getJson(&m)
	fmt.Println(m)
}
