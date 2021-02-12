package userService

import (
	"encoding/json"
	"net/http"
	"log"
	. "schedule-app/utils"
	userModel "schedule-app/models/user-model"
)

func Main(w http.ResponseWriter, r *http.Request){
	log.Println(r.Method, r.URL.Path )

	if r.Method == "POST" && r.URL.Path == "/user"{
		AddUser(w, r)
		return
	}

}

func AddUser(w http.ResponseWriter, r *http.Request){
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println(err)
	}

	if user.Name == "" {
		http.Error(w, "required Name", http.StatusNotAcceptable)
		return
	}

	user = userModel.Add(user)

	b, err := json.Marshal(user)
	if err != nil {
		log.Println(err)
		return
	}

	w.Write([]byte(string(b)))
}