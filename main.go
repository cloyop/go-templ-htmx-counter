package main

import (
	"fmt"
	"net/http"

	"github.com/cloyop/go-templ-htmx-counter/views"
	"github.com/cloyop/go-templ-htmx-counter/views/components"
)

func main() {

	count := 0
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		views.Index(fmt.Sprint(count)).Render(r.Context(), w)
	})
	http.HandleFunc("/increase", func(w http.ResponseWriter, r *http.Request) {
		count++
		components.Counter(fmt.Sprint(count)).Render(r.Context(), w)
	})
	http.HandleFunc("/decrease", func(w http.ResponseWriter, r *http.Request) {
		count--
		components.Counter(fmt.Sprint(count)).Render(r.Context(), w)
	})
	http.ListenAndServe("localhost:3000", nil)
}
