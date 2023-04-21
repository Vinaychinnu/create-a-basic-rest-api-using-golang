package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Name string `json:"name"`
}

func homePage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "This method is not allowed", http.StatusMethodNotAllowed)
		return
	}
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println("new user created: ", user)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "hello " + user.Name})
	fmt.Println("%s", user.Name)

}
func main() {
	http.HandleFunc("/hello", homePage)
	fmt.Println("server is running on port:8564...")
	err := http.ListenAndServe(":8564", nil)
	if err != nil {
		panic(err)
	}

}
