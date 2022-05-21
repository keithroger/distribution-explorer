package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"gonum.org/v1/gonum/stat/distuv"
)

const (
	n = 200
)

type Point struct {
	X, Y float64
}

type DistRequest struct {
	Name string
	Mode string
	Args []float64
}

type Distribution interface {
	Prob(float64) float64
	CDF(float64) float64
}

func (d *DistRequest) getData() []Point {
	var data []Point

	switch d.Name {
	case "Normal":
		switch d.Mode {
		case "Distribution":
			dist := distuv.Normal{
				Mu:    d.Args[0],
				Sigma: d.Args[1],
			}

			begin, end := -4.0*dist.Sigma, 4.0*dist.Sigma
			data = createPoints(dist, begin, end)

		}

	case "Student's T":
		// something
	}

	return data
}

func createPoints(dist Distribution, begin, end float64) []Point {
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

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
