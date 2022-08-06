package function

import (
	"fmt"
	"net/http"
)

func Start(h http.HandlerFunc) {
	start(h)
}

func start(h http.HandlerFunc) {
	http.HandleFunc("/", h)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic("Fail starting server")
	}
	fmt.Println("Server started")
}
