package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/articles/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from /articles/")
	})

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from /users")
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Server", "GO Server")
		fmt.Fprintf(w, `<html>
			<body>
				Hello, Gopher!
			</body>
			</html>
			`)
	})
	http.ListenAndServe(":3000", nil)
}
