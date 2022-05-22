package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"

	"gonum.org/v1/gonum/stat/distuv"
)

const (
	n = 400
)

type Point struct {
	X, Y float64
}

type GraphData struct {
	Line []Point
	PDF  []Point
	CDF  []Point
}

type Distribution interface {
	Prob(float64) float64
	CDF(float64) float64
}

type DistRequest struct {
	Name string
	Mode string
	Args []float64
}

func (d *DistRequest) getData() GraphData {
	var data GraphData

	switch d.Name {
	case "Normal":
		dist := distuv.Normal{
			Mu:    d.Args[0],
			Sigma: d.Args[1],
		}
		begin, end := dist.Mu-4.0*dist.Sigma, dist.Mu+4.0*dist.Sigma
		data.Line = curvePoints(dist, begin, end)

		switch d.Mode {
		case "Distribution":
			data.CDF = curvePoints(dist, begin, end)
			data.PDF = []Point{}
		case "PDF":
			data.PDF = []Point{{d.Args[2], dist.Prob(d.Args[2])}}
			data.CDF = []Point{}
		case "CDF":
			begin, end := dist.Mu-4.0*dist.Sigma, math.Min(d.Args[2], 4.0*dist.Sigma)
			data.CDF = curvePoints(dist, begin, end)
			data.PDF = []Point{}
		}

	case "Student's T":
		dist := distuv.StudentsT{
			Mu:    d.Args[0],
			Sigma: d.Args[1],
			Nu:    d.Args[2],
		}
		begin, end := dist.Mu-4.0*dist.Sigma, dist.Mu+4.0*dist.Sigma
		data.Line = curvePoints(dist, begin, end)

		switch d.Mode {
		case "Distribution":
			data.CDF = curvePoints(dist, begin, end)
			data.PDF = []Point{}
		case "PDF":
			data.PDF = []Point{{d.Args[3], dist.Prob(d.Args[3])}}
			data.CDF = []Point{}
		case "CDF":
			begin, end := dist.Mu-4.0*dist.Sigma, math.Min(d.Args[3], 4.0*dist.Sigma)
			data.CDF = curvePoints(dist, begin, end)
			data.PDF = []Point{}
		}

	}

	return data
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
