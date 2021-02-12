package scheduleService

import (
	"time"
	"strconv"
	"encoding/json"
	"net/http"
	"log"
	scheduleModel "schedule-app/models/schedule-model"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	. "schedule-app/utils"
)

func Main(w http.ResponseWriter, r *http.Request){
	log.Println(r.Method, r.URL.Path )

	if r.Method == "POST" && r.URL.Path == "/schedule"{
		addSchedule(w, r)
		return
	}

	if r.Method == "GET" && r.URL.Path == "/schedule"{
		getSchedule(w, r)
		return
	}

	if r.Method == "PUT"{
		updateSchedule(w, r)
		return
	}

}

func getClashSchedule(schedule Schedule) bool{
	fstart_date := time.Date(schedule.StartDate.Time().Year(), schedule.StartDate.Time().Month(), schedule.StartDate.Time().Day(), 0, 0, 0, 0, schedule.StartDate.Time().Location())
	fend_date := time.Date(schedule.EndDate.Time().Year(), schedule.EndDate.Time().Month(), schedule.EndDate.Time().Day(), 23, 59, 59, 0, schedule.EndDate.Time().Location())
	filter := bson.M{
		"start_date" : bson.M{
			"$gt": fstart_date,
			"$lt": fend_date,
		},
		"end_date" : bson.M{
			"$gt" : schedule.StartDate.Time(),
		},
		"_id" : bson.M {
			"$ne" : schedule.ID,
		},
	}
	if !schedule.ID.IsZero() {
		filter["_id"] = bson.M {
			"$ne" : schedule.ID,
		}
	}
	schedules , _ := scheduleModel.Get(1, 1, filter)
	log.Println("not Equal", schedule.ID, schedule)

	if len(schedules) >= 1 {
		return true
	} else {
		return false
	}
}

func addSchedule(w http.ResponseWriter, r *http.Request){
	var schedule Schedule
	err := json.NewDecoder(r.Body).Decode(&schedule)
	if err != nil {
		log.Println(err)
	}


	if schedule.AssignUserId.IsZero(){
		http.Error(w, "invalid assign user id", http.StatusNotAcceptable)
		return
	}

	if schedule.EndDate == 0 || schedule.EndDate == 0 {
		http.Error(w, "invalid start_date, end_date", http.StatusNotAcceptable)
		return
	}

	if schedule.StartDate.Time().After(schedule.EndDate.Time()) {
		http.Error(w, "start_date need to be lower than end_date", http.StatusNotAcceptable)
		return
	}

	if getClashSchedule(schedule) {
		http.Error(w, "create clashed schedule", http.StatusNotAcceptable)
		return
	}

	schedule.Status = "created"

	schedule = scheduleModel.Add(schedule)

	b, err := json.Marshal(schedule)
	if err != nil {
		log.Println(err)
		return
	}

	w.Write([]byte(string(b)))
}


func getSchedule(w http.ResponseWriter, r *http.Request){
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

	data, pagination := scheduleModel.Get(page, limit, bson.M{})

	b, _ := json.Marshal(map[string]interface{}{
		"data" : data,
		"pagination" : pagination,
	})
	
	w.Write([]byte(b))
}

func updateSchedule(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	var schedule Schedule
	err := json.NewDecoder(r.Body).Decode(&schedule)
	if err != nil {
		log.Println(err)
	}

	if schedule.ID, err = primitive.ObjectIDFromHex(vars["id"]); err != nil {
		http.Error(w, "invalid update id", http.StatusNotAcceptable)
		return
	}

	if schedules , _ := scheduleModel.Get(1,1,bson.M{"_id" : schedule.ID}); len(schedules) != 1 {
		http.Error(w, "schedule not found", http.StatusNotAcceptable)
		return
	} else {
		schedule.CreatedAt = schedules[0].CreatedAt
	}

	if schedule.AssignUserId.IsZero(){
		http.Error(w, "invalid assign user id", http.StatusNotAcceptable)
		return
	}

	if schedule.EndDate == 0 || schedule.EndDate == 0 {
		http.Error(w, "invalid start_date, end_date", http.StatusNotAcceptable)
		return
	}

	if schedule.StartDate.Time().After(schedule.EndDate.Time()) {
		http.Error(w, "start_date need to be lower than end_date", http.StatusNotAcceptable)
		return
	}

	if getClashSchedule(schedule) {
		http.Error(w, "create clashed schedule", http.StatusNotAcceptable)
		return
	}

	schedule.Status = "created"

	if result, err := scheduleModel.Update(schedule.ID, schedule); err != nil || result.MatchedCount == 0 {
		http.Error(w, "fail update data", http.StatusNotAcceptable)
		return 
	}

	b, err := json.Marshal(schedule)
	if err != nil {
		log.Println(err)
		return
	}

	w.Write([]byte(string(b)))
}