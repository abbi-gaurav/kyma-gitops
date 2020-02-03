package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = fmt.Fprintf(writer, "Kyma GitOps with Flux and Go app!!!")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))

}
