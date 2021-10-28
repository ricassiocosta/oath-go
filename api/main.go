package main

import (
	"fmt"
	"log"
	"net/http"
	"oath-go/src/config"
	"oath-go/src/router"
)

func main() {
	config.LoadEnv()
	r := router.Generate()

	fmt.Println("Server is running...")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.APIPort), r))
}
