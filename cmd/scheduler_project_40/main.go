package main

import (
	"fmt"
	"net/http"
	"time"
)

type Worker struct {
	lastJobId int
	queue     []time.Duration
}

func NewWorker() *Worker {
	w := new(Worker)
	w.lastJobId = 0
	w.queue = make([]time.Duration, 0)
	return w
}

func (w Worker) Add(duration time.Duration) {
	w.lastJobId++
	w.queue = append(w.queue, duration)
}

func main() {
	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
	})
	http.ListenAndServe(":80", nil)
}
