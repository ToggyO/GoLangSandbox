package api

import (
	"github.com/gorilla/mux"
	"hello/packages/api/handlers"
	"log"
	"net/http"
)

func StartApi() {
	r := mux.NewRouter().StrictSlash(true)

	var userRepository = new(api.UserRepository)
	r.HandleFunc("/all", userRepository.GetAll)

	log.Println(http.ListenAndServe(":10000", r))
}
