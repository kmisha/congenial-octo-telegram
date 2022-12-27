package main

import (
	"fmt"
	"log"
	"net/http"

	"auth/server"
	"auth/store"

	"github.com/go-chi/jwtauth/v5"
)

const PORT = 5000

func main() {
	str := store.NewInMemoryUsersStore()
	token := jwtauth.New("HS256", []byte("secret"), nil)
	srv := server.NewAuthServer(str, *token)

	log.Printf("Start server at %d PORT", PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", PORT), srv.Router))
}
