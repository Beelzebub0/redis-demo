package main

import (
	"log"
	"net/http"

	"github.com/Beelzebub0/redis-demo/rest"
)

func main() {
	port := ":8080"
	http.HandleFunc("/product", rest.Product)
	log.Fatal(http.ListenAndServe(port, nil))
}
