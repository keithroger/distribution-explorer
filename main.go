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
	Distribution string
	Mode         string
	Args         []float64
}

func (d *DistRequest) getData() []Point {
	data := make([]Point, n)

	switch d.Distribution {
	case "Normal":
		switch d.Mode {
		case "Distribution":
			dist := distuv.Normal{
				Mu:    d.Args[0],
				Sigma: d.Args[1],
			}

			for i := 0; i < 200; i++ {
				data[i] = Point{float64(i), dist.Prob(float64(i))}
			}
		}

	case "Student's T":
		// something
	}

	return data
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
