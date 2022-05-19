package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Point struct {
	X, Y int
}

type Points []Point

func api(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		data := Points{
			{1, 2},
			{2, 3},
			{5, 5},
			{10, 2},
			{11, 1},
		}

		jsonData, err := json.Marshal(data)
		if err != nil {
			panic(err)
		}

		fmt.Println(string(jsonData))

		w.Header().Set("Content-Type", "application/json")
		// w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Write(jsonData)
	}
}

func main() {
	fs := http.FileServer(http.Dir("./public"))

	http.Handle("/", fs)
	http.HandleFunc("/api", api)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
