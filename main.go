package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Point struct {
	X, Y float64
}

type GraphData struct {
	Curve []Point
	PDF   []Point
	CDF   []Point
}

type DistuvDist interface {
	Prob(float64) float64
	CDF(float64) float64
}

func api(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		decoder := json.NewDecoder(r.Body)
		var distReq DistRequest
		err := decoder.Decode(&distReq)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}

		data := distReq.getData()

		jsonData, err := json.Marshal(data)
		if err != nil {
			panic(err)
		}

		w.Header().Set("Content-Type", "application/json")
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
