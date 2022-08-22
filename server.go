package main

import (
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

var lastMessage = 0

func main() {
	r := NewRouter()
	http.ListenAndServe(":8080", r)
}

// NewRouter returns an HTTP handler that implements the routes for the API
func NewRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		print("Somebody wants GET")
		w.Write([]byte(strconv.Itoa(lastMessage)))
		lastMessage = 0
	})
	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
		println("Somebody wants POST")
		b, _ := ioutil.ReadAll(r.Body)
		lastMessage, _ = strconv.Atoi(string(b))
		println("lastMessage is %d", lastMessage)
		//lastMessage = int(binary.BigEndian.Uint64(b))

	})
	return r
}
