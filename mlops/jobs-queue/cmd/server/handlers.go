package main

import (
	"encoding/json"
	"fmt"
	"jobs-daemon/internal/types"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func listAPI(rw http.ResponseWriter, req *http.Request) {
	jobs, err := ListArcJobs()
	if err != nil {
		log.Fatalln("listAPI(): Failed read list of jobs from SQLite database:", err)
	}
	log.Printf("Listed %d ARC jobs\n", len(jobs))

	jsonData, err := json.Marshal(jobs)
	if err != nil {
		log.Fatalln("listAPI(): Failed to marshal into JSON data:", err)
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)

	_, err = rw.Write(jsonData)
	if err != nil {
		log.Fatalln("listAPI(): Failed to write response:", err)
	}
}

func insertAPI(rw http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var newArcJob types.ArcJob
	err := decoder.Decode(&newArcJob)
	if err != nil {
		log.Fatalln("insertAPI(): failed to decode req.Body:", err)
	}

	err = InsertArcJob(newArcJob)
	if err != nil {
		log.Fatalln("insertAPI(): failed to insert new job in SQLite database:", err)
	}
	log.Printf("Inserted new ARC job %v\n", newArcJob)

	rw.Header().Set("Content-Type", "text/plain; charset=utf-8")
	msg := fmt.Sprintf("%s - POST inserted %v\n", TimeStamp(), newArcJob)
	rw.Write([]byte(msg))
}

func updateAPI(rw http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var newArcJob types.ArcJob
	err := decoder.Decode(&newArcJob)
	if err != nil {
		log.Fatalln("updateAPI(): failed decode JSON from req.Body:", err)
	}

	err = UpdateArcJob(newArcJob)
	if err != nil {
		log.Fatalln("updateAPI(): failed to update job in SQLite database:", err)
	}
	log.Printf("Updated ARC job %v\n", newArcJob)

	rw.Header().Set("Content-Type", "text/plain; charset=utf-8")
	msg := fmt.Sprintf("%s - POST updated %v\n", TimeStamp(), newArcJob)
	rw.Write([]byte(msg))
}
