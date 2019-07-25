package httpreqs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func MakeReq() {
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")

	if err != nil {
		log.Fatalln(err)
	}

	defer res.Body.Close()

	resp, err := ioutil.ReadAll(res.Body)

	type Message struct {
		userId    float64
		id        float64
		title     string
		completed bool
	}

	var m Message

	json.Unmarshal(resp, &m)

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(m)

}
