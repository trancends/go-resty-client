package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/go-resty/resty/v2"
)

type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	client := resty.New()
	resp, err := client.R().Get("https://jsonplaceholder.typicode.com/todos")
	if err != nil {
		log.Println("Error: ", err.Error())
	}

	var todos []Todo
	err = json.Unmarshal(resp.Body(), &todos)
	if err != nil {
		log.Println("error: ", err.Error())
	}
	fmt.Println(len(todos))
	fmt.Println(todos[3].Title)
	// rumus pagination
	page := 1
	size := 10
	startIndex := (page - 1) * size
	endIndex := startIndex + size

	if endIndex > len(todos) {
		endIndex = len(todos)
		startIndex = endIndex - size
	}

	fmt.Println("Page:", page)
	fmt.Println("Size:", size)
	fmt.Println("startIndex:", startIndex)
	fmt.Println("endIndex:", endIndex)

	fmt.Println(len(todos[startIndex:endIndex]))
	fmt.Println(todos[startIndex:endIndex])
}
