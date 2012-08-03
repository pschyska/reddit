package main

import (
  "log"
  "fmt"
  "github.com/pschyska/reddit"
)

func main() {
	items, err := reddit.Get("golang")
	if err != nil {
		log.Fatal(err)
	}
	for _, child := range items {
		fmt.Println(child)
	}
}

