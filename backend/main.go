// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	scheduleService "schedule-app/service/schedule-service"
	userService "schedule-app/service/user-service"
	utilsService "schedule-app/utils"
	"flag"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	utilsService.CreateConnection()
	flag.Parse()
	router := mux.NewRouter()


	schedule := router.PathPrefix("/schedule").Subrouter()
		schedule.HandleFunc("", scheduleService.Main)
		schedule.HandleFunc("/{id}", scheduleService.Main)
	user := router.PathPrefix("/user").Subrouter()
		user.HandleFunc("", userService.Main)

	log.Fatal(http.ListenAndServe(":8081", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(router)))
}
