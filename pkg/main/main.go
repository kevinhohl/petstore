package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kevinhohl/petstore/pkg/api"
)

var commit string

func main() {
	fmt.Println("Started PetStore API")
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", api.GetRouter(commit)))
}
