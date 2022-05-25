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
	Points []Point
	PDF    []Point
	CDF    []Point
}

type Distribution interface {
	Prob(float64) float64
	CDF(float64) float64
}

func curvePoints(dist Distribution, begin, end float64) []Point {
	if begin >= end {
		return []Point{}
	}

	points := make([]Point, n)
	step := (end - begin) / float64(n)

	for i := 0; i < n; i++ {
		x := begin + step*float64(i)
		y := dist.Prob(x)

		points[i] = Point{x, y}
	}

	return points
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

		fmt.Println("dist: ", distReq.Name)
		fmt.Println("mode: ", distReq.Mode)
		fmt.Println("inputs: ", distReq.Args)

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

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
}
