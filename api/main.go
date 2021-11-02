package main

import (
	"fmt"
	"log"
	"net/http"
	"oath-go/src/config"
	"oath-go/src/router"

	"github.com/rs/cors"
)

func main() {
	config.LoadEnv()
	r := router.Generate()

	c := cors.AllowAll()

	fmt.Println("Server is running...")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.APIPort), c.Handler(r)))
}
