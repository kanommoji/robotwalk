package main

import (
	"fmt"
	"net/http"
	"robotwalk"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		walk := r.URL.Query().Get("walk")
		robotwalk.RobotWalk(walk)
		fmt.Println("id =>", walk)
		fmt.Print(robotwalk.ReadTable())
	})
	http.ListenAndServe(":8080", nil)
}
