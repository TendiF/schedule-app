package shiftService

import (
	"time"
	"strconv"
	"encoding/json"
	"net/http"
	"log"
	shiftModel "schedule-app/models/shift-model"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	. "schedule-app/utils"
)

func Main(w http.ResponseWriter, r *http.Request){
	log.Println(r.Method, r.URL.Path )

	if r.Method == "POST" && r.URL.Path == "/shift"{
		addShift(w, r)
		return
	}

	if r.Method == "GET" && r.URL.Path == "/shift"{
		getShift(w, r)
		return
	}

	if r.Method == "PUT"{
		updateShift(w, r)
		return
	}

	if r.Method == "DELETE"{
		deleteShift(w, r)
		return
	}

}

func getClashShift(shift Shift) bool{
	fstart_date := time.Date(shift.StartDate.Time().Year(), shift.StartDate.Time().Month(), shift.StartDate.Time().Day(), 0, 0, 0, 0, shift.StartDate.Time().Location())
	fend_date := time.Date(shift.EndDate.Time().Year(), shift.EndDate.Time().Month(), shift.EndDate.Time().Day(), 23, 59, 59, 0, shift.EndDate.Time().Location())
	filter := bson.M{
		"start_date" : bson.M{
			"$gt": fstart_date,
			"$lt": fend_date,
		},
		"end_date" : bson.M{
			"$gt" : shift.StartDate.Time(),
		},
		"_id" : bson.M {
			"$ne" : shift.ID,
		},
	}
	if !shift.ID.IsZero() {
		filter["_id"] = bson.M {
			"$ne" : shift.ID,
		}
	}
	shifts , _ := shiftModel.Get(1, 1, filter)

	if len(shifts) >= 1 {
		return true
	} else {
		return false
	}
}

func addShift(w http.ResponseWriter, r *http.Request){
	var shift Shift
	err := json.NewDecoder(r.Body).Decode(&shift)
	if err != nil {
		log.Println(err)
	}

	if shift.AssignUserId.IsZero(){
		http.Error(w, "invalid assign user id", http.StatusNotAcceptable)
		return
	}

	if shift.EndDate == 0 || shift.EndDate == 0 {
		http.Error(w, "invalid start_date, end_date", http.StatusNotAcceptable)
		return
	}

	if shift.StartDate.Time().After(shift.EndDate.Time()) {
		http.Error(w, "start_date need to be lower than end_date", http.StatusNotAcceptable)
		return
	}

	if getClashShift(shift) {
		http.Error(w, "create clashed shift", http.StatusNotAcceptable)
		return
	}

	shift.Status = "created"

	shift = shiftModel.Add(shift)

	b, err := json.Marshal(shift)
	if err != nil {
		log.Println(err)
		return
	}

	w.Write([]byte(string(b)))
}


func getShift(w http.ResponseWriter, r *http.Request){
	// Example for Normal Find query
	var page int64 = 1
	var limit int64 = 10
	var filter = bson.M{}
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

	if keys, ok := r.URL.Query()["id_user"]; ok {
		if ID, err := primitive.ObjectIDFromHex(keys[0]); err == nil {
			filter["assign_user_id"] = ID
		} else {
			http.Error(w, "invalid id filter", http.StatusNotAcceptable)
			return
		}
	}
	

		if keys, ok := r.URL.Query()["id"]; ok {
			if ID, err := primitive.ObjectIDFromHex(keys[0]); err == nil {
				filter["_id"] = ID
			} else {
				http.Error(w, "invalid id filter", http.StatusNotAcceptable)
				return
			}
		}

	data, pagination := shiftModel.Get(page, limit, filter)

	b, _ := json.Marshal(map[string]interface{}{
		"data" : data,
		"pagination" : pagination,
	})
	
	w.Write([]byte(b))
}

func updateShift(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	var shift Shift
	err := json.NewDecoder(r.Body).Decode(&shift)
	if err != nil {
		log.Println(err)
	}

	log.Println(shift.Status)


	if shift.ID, err = primitive.ObjectIDFromHex(vars["id"]); err != nil {
		http.Error(w, "invalid update id", http.StatusNotAcceptable)
		return
	}

	if shifts , _ := shiftModel.Get(1,1,bson.M{"_id" : shift.ID}); len(shifts) != 1 {
		http.Error(w, "shift not found", http.StatusNotAcceptable)
		return
	} else {
		if shifts[0].Status == "published" {
			http.Error(w, "shift status already publish", http.StatusNotAcceptable)
			return
		}
		shift.CreatedAt = shifts[0].CreatedAt
	}

	if shift.AssignUserId.IsZero(){
		http.Error(w, "invalid assign user id", http.StatusNotAcceptable)
		return
	}

	if shift.EndDate == 0 || shift.EndDate == 0 {
		http.Error(w, "invalid start_date, end_date", http.StatusNotAcceptable)
		return
	}

	if shift.StartDate.Time().After(shift.EndDate.Time()) {
		http.Error(w, "start_date need to be lower than end_date", http.StatusNotAcceptable)
		return
	}

	if getClashShift(shift) {
		http.Error(w, "create clashed shift", http.StatusNotAcceptable)
		return
	}

	if result, err := shiftModel.Update(shift.ID, shift); err != nil || result.MatchedCount == 0 {
		http.Error(w, "fail update data", http.StatusNotAcceptable)
		return 
	}

	b, err := json.Marshal(shift)
	if err != nil {
		log.Println(err)
		return
	}

	w.Write([]byte(string(b)))
}

func deleteShift(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	var shift Shift
	var err error

	if shift.ID, err = primitive.ObjectIDFromHex(vars["id"]); err != nil {
		http.Error(w, "invalid update id", http.StatusNotAcceptable)
		return
	}
	
	if shifts , _ := shiftModel.Get(1,1,bson.M{"_id" : shift.ID}); len(shifts) != 1 {
		http.Error(w, "shift not found", http.StatusNotAcceptable)
		return
	} else {
		if shifts[0].Status == "published" {
			http.Error(w, "shift status already publish", http.StatusNotAcceptable)
			return
		}
		shift.CreatedAt = shifts[0].CreatedAt
	}

	result, err := shiftModel.Delete(shift.ID)

	if err != nil{
		http.Error(w, "fail delete data", http.StatusNotAcceptable)
		return
	}

	b, err := json.Marshal(result)
	if err != nil {
		log.Println(err)
		return
	}

	w.Write([]byte(string(b)))
}