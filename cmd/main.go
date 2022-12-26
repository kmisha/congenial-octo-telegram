package main

import (
	"auth/server"
	"auth/store"
	"fmt"
	"log"
	"net/http"
)

const PORT = 5000

func main() {
	str := store.NewInMemoryUsersStore()
	srv := server.NewAuthServer(str)

	fmt.Printf("Start server at %d PORT", PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", PORT), srv))
}
