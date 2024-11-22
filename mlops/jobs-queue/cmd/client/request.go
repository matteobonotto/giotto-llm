package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"jobs-daemon/internal/types"
	"net/http"
)

const HOST = "cluster-manager"
const PORT = 5050

func ListArcJobs() ([]types.ArcJob, error) {
	bodyReader := bytes.NewReader([]byte(""))
	requestURL := fmt.Sprintf("http://%s:%d/jobs/list", HOST, PORT)
	req, err := http.NewRequest(http.MethodGet, requestURL, bodyReader)

	if err != nil {
		return []types.ArcJob{}, err
	}

	req.Header.Set("Content-Type", "text/plain; charset=utf-8")

	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return []types.ArcJob{}, err
	}

	var allArcJobs []types.ArcJob
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&allArcJobs)
	if err != nil {
		return []types.ArcJob{}, err
	}

	return allArcJobs, nil
}

func InsertArcJob(newArcJob types.ArcJob) error {
	requestURL := fmt.Sprintf("http://%s:%d/jobs/insert", HOST, PORT)
	jsonData, err := json.Marshal(newArcJob)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodPost, requestURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

func UpdateArcJob(newArcJob types.ArcJob) error {
	requestURL := fmt.Sprintf("http://%s:%d/jobs/update", HOST, PORT)
	jsonData, err := json.Marshal(newArcJob)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodPost, requestURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
