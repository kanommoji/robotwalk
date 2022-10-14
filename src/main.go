package main

import (
	"net/http"
	robotwalk "robotwalk/src/service"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		walks := r.URL.Query().Get("walks")
		w.Header().Set("Content-Type", "application/text")
		walkig, err := robotwalk.RobotWalk(walks)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		w.Write([]byte(walkig))
	})
	http.ListenAndServe(":8080", nil)
}
