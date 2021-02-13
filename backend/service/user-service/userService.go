package userService

import (
	"strconv"
	"encoding/json"
	"net/http"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	. "schedule-app/utils"
	userModel "schedule-app/models/user-model"
)

func Main(w http.ResponseWriter, r *http.Request){
	log.Println(r.Method, r.URL.Path )

	if r.Method == "POST" && r.URL.Path == "/user"{
		AddUser(w, r)
		return
	}

	if r.Method == "GET" && r.URL.Path == "/user"{
		getUser(w, r)
		return
	}

	if r.Method == "POST" && r.URL.Path == "/user/login"{
		login(w, r)
		return
	}

}

func getUser(w http.ResponseWriter, r *http.Request){
	// Example for Normal Find query
	var page int64 = 1
	var limit int64 = 10
	if keys, ok := r.URL.Query()["page"]; ok {
		i , err :=  strconv.ParseInt(keys[0], 10, 64)
		if err != nil {
			log.Println(err)
		}
		page = i
	}

	if keys, ok := r.URL.Query()["per_page"]; ok {
		i , err :=  strconv.ParseInt(keys[0], 10, 64)
		if err != nil {
			log.Println(err)
		}
		limit = i
	}

	data, pagination := userModel.Get(page, limit, bson.M{})

	b, _ := json.Marshal(map[string]interface{}{
		"data" : data,
		"pagination" : pagination,
	})
	
	w.Write([]byte(b))
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

	b, err := json.Marshal(map[string]interface{}{
		"data" : user,
	})
	if err != nil {
		log.Println(err)
		return
	}

	w.Write([]byte(string(b)))
}

func login(w http.ResponseWriter, r *http.Request){
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println(err)
	}

	log.Println(user)

	data, _ := userModel.Get(1, 1, bson.M{
		"name" : user.Name,
	})

	if len(data) >= 1 {
		b, err := json.Marshal(map[string]interface{}{
			"data" : data[0],
		})
		if err != nil {
			http.Error(w, "Error parse data", http.StatusNotAcceptable)
			w.Write([]byte("Error Parse Data"))
		}
		w.Write([]byte(b))
	} else {
		http.Error(w, "Login Error", http.StatusNotAcceptable)
	}
}