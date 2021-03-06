package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
