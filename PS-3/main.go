package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		currentTime := time.Now()
		fmt.Fprintf(w, `
			<html>
			<body>
				<h1>Current Date and Time</h1>
				<p>%s</p>
			</body>
			</html>
		`, currentTime.Format("2006-01-02 15:04:05"))
	})


	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
