package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/css-import-loop/", func(w http.ResponseWriter, r *http.Request) {
		numT := r.URL.Path[len("/css-import-loop/"):]
		num, err := strconv.ParseInt(numT, 10, 64)
		if err != nil {
			http.Error(w, "Invalid number", http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "text/css")
		fmt.Fprintf(w, `
@import url("/css-import-loop/%d");
`, num+1)
	})

	http.HandleFunc("/{$}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintf(w, `<html>
    <head>
    <link rel="stylesheet" href="/css-import-loop/1" type="text/css">
    <title>Loading...</title>
    </head>
</html>`)
	})

	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}
