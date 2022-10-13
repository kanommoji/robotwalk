package main

import (
	"net/http"
	robotwalk "robotwalk/src/service"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		walk := r.URL.Query().Get("walk")
		walkig := robotwalk.RobotWalk(walk)
		w.Header().Set("Content-Type", "application/text")
		w.Write([]byte(walkig))
	})
	http.ListenAndServe(":8080", nil)
}
