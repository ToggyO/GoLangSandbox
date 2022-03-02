package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Name string
	Age  int
}

var users = []User{
	{Name: "Oleg", Age: 30},
}

type UserRepository struct{}

func (u UserRepository) GetAll(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: GetAll")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
