package main

import (
	"fmt"
	"net/http"

	"github.com/vladimirok5959/golang-server-sessions/session"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Init session
		sess := session.New(w, r, "./tmp")
		defer sess.Close()

		if r.URL.Path == "/" {
			var counter int

			// Get value or set default
			if sess.IsSetInt("counter") {
				counter = sess.GetInt("counter", 0)
			} else {
				counter = 0
			}

			// Increment value
			counter++

			// Update
			sess.SetInt("counter", counter)

			w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
			w.Header().Set("Content-Type", "text/html")
			if _, err := w.Write([]byte(`
				<div>Hello World!</div>
				<div>Counter: ` + fmt.Sprintf("%d", counter) + `</div>
			`)); err != nil {
				fmt.Printf("%s\n", err.Error())
			}
		} else {
			w.WriteHeader(http.StatusNotFound)
			if _, err := w.Write([]byte(`<div>Error 404!</div>`)); err != nil {
				fmt.Printf("%s\n", err.Error())
			}
		}
	})

	// Delete expired session files
	if err := session.Clean("./tmp"); err != nil {
		fmt.Printf("%s\n", err.Error())
	}

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("%s\n", err.Error())
	}
}
