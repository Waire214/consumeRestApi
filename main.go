package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type user struct {
	UserId int    `json:"userId"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var users []user
	json.Unmarshal(body, &users)
	log.Println(users)
}
