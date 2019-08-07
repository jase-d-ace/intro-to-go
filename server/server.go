package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

var myClient = &http.Client{Timeout: 10 * time.Second}

const (
	host   = "localhost"
	port   = 5432
	user   = "jase"
	dbname = "golang_test_db"
)

//empty struct that will eventually hold data from the API that we hit on a route.
type Message struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type Person struct {
	Id       int    `json:"id"`
	Nickname string `json:"nickname"`
	Age      int    `json:"age"`
}

//tutorial on properly adding a psql database found here: https://www.sohamkamani.com/blog/2017/10/18/golang-adding-database-to-web-application/

//make the database methods available to the entire application to DRY code
//these methods will then be applied to a "controller" struct that we'll use to call the methods
type Store interface {
	CreatePerson(person *Person) error
	GetPersons() ([]*Person, error)
	GetOnePerson(id int) (*Person, error)
}

type dbStore struct {
	db *sql.DB
}

//defining the methods that we set up in our interface
//interfaces hold methods that can be applied to structs, effectively making an OOP-style of calling methods.
//function signature: keyword func + struct (optional) + function name + parameters + return type
func (store *dbStore) CreatePerson(person *Person) error {
	_, err := store.db.Query(`INSERT INTO persons(nickname, age) VALUES($1, $2) RETURNING id`, person.Nickname, person.Age)
	return err
}

func (store *dbStore) GetOnePerson(id int) (*Person, error) {
	personRes := &Person{}
	res := store.db.QueryRow(`SELECT * FROM persons WHERE id=$1`, id)

	//so res.Scan is the way that populates the empty person struct with sql data
	if err := res.Scan(&personRes.Id, &personRes.Nickname, &personRes.Age); err != nil {
		fmt.Println("Ooops, messed up", err)
		return nil, err
	}
	return personRes, nil

}

func (store *dbStore) GetPersons() ([]*Person, error) {
	persons := []*Person{}
	res, err := store.db.Query(`SELECT * FROM persons`)

	if err != nil {
		panic(err)
	}

	defer res.Close()

	for res.Next() {
		person := &Person{}
		if err := res.Scan(&person.Id, &person.Nickname, &person.Age); err != nil {
			fmt.Println("Ooops, messed up", err)
			return nil, err
		}
		persons = append(persons, person)
	}
	return persons, nil
}

//make the store available application-wide. Defining it as a struct before isn't enough; now we instantiate using this variable and the following function
var store Store

func InitStore(s Store) {
	store = s
}

//empty arrays that we use in the other functions to hold data from our API calls

var todos []Message

var Persons []Person

var extraTodos []Message

func getPersonHandler(w http.ResponseWriter, r *http.Request) {
	persons, err := store.GetPersons()

	jsonData, err := json.Marshal(persons)
	if err != nil {
		panic(err)
	}
	w.Write(jsonData)
}

func getOnePersonHandler(w http.ResponseWriter, r *http.Request) {
	personId := mux.Vars(r)["id"]
	id, err := strconv.Atoi(personId)

	if err != nil {
		panic(err)
	}
	person, err := store.GetOnePerson(id)
	fmt.Println("looking for this dude: ", id)
	if err != nil {
		panic(err)
	}
	jsonData, err := json.Marshal(person)
	w.Write(jsonData)
}

func splashPage(w http.ResponseWriter, r *http.Request) {
	message := "hello world!"
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

func createPersonHandler(w http.ResponseWriter, r *http.Request) {

	newPerson := Person{}

	reqBody, err := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &newPerson)

	store.CreatePerson(&newPerson)

	jsonPerson, _ := json.Marshal(newPerson)

	if err != nil {
		panic(err)
	}
	w.Write(jsonPerson)
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"dbname=%s sslmode=disable", host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()

	if err != nil {
		panic(err)
	}
	fmt.Println("db connection live")

	InitStore(&dbStore{db: db})

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", splashPage)
	router.HandleFunc("/todos", getJsonArray).Methods("GET")
	router.HandleFunc("/people", getPersonHandler).Methods("GET")
	router.HandleFunc("/todos", createPersonHandler).Methods("POST")
	router.HandleFunc("/people/{id}", getOnePersonHandler).Methods("GET")
	if err := http.ListenAndServe(":8080", router); err != nil {
		panic(err)
	}
}
