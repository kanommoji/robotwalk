package main

import (
	"net/http"
	"robotwalk"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		walk := r.URL.Query().Get("walk")
		robotwalk.RobotWalk(walk)
		w.Header().Set("Content-Type", "application/text")
		w.Write([]byte(robotwalk.ReadResult()))
	})
	http.ListenAndServe(":8080", nil)
}
