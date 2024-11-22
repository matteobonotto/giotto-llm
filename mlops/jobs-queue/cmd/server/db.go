package main

import (
	"database/sql"
	"os"

	"jobs-daemon/internal/types"

	_ "github.com/mattn/go-sqlite3"
)

var jobs_db_path = os.Getenv("JOBS_DB_PATH")

func ListArcJobs() ([]types.ArcJob, error) {
	sqliteDatabase, _ := sql.Open("sqlite3", jobs_db_path)
	defer sqliteDatabase.Close()
	row, err := sqliteDatabase.Query("SELECT * FROM arc_jobs ORDER BY jobId")
	if err != nil {
		return []types.ArcJob{}, err
	}
	defer row.Close()

	jobs := make([]types.ArcJob, 0, 10)
	for row.Next() {
		var jobId int
		var jobName string
		var jobCmd string
		var jobConfig string
		var started int
		var done int
		var failed int
		row.Scan(&jobId, &jobName, &jobCmd, &jobConfig, &started, &done, &failed)
		newJob := types.ArcJob{
			JobId:     jobId,
			JobName:   jobName,
			JobCmd:    jobCmd,
			JobConfig: jobConfig,
			Started:   started,
			Done:      done,
			Failed:    failed,
		}
		jobs = append(jobs, newJob)
	}
	return jobs, nil
}

func InsertArcJob(j types.ArcJob) error {
	sqliteDatabase, _ := sql.Open("sqlite3", jobs_db_path)
	defer sqliteDatabase.Close()

	insertJobSQL := `INSERT INTO arc_jobs(JobName, JobCmd, JobConfig, Started, Done, Failed) VALUES (?, ?, ?, ?, ?, ?)`

	statement, err := sqliteDatabase.Prepare(insertJobSQL)
	if err != nil {
		return err
	}

	_, err = statement.Exec(j.JobName, j.JobCmd, j.JobConfig, j.Started, j.Done, j.Failed)
	if err != nil {
		return err
	}

	return nil
}

func UpdateArcJob(j types.ArcJob) error {
	sqliteDatabase, _ := sql.Open("sqlite3", jobs_db_path)
	defer sqliteDatabase.Close()

	updateJobSQL := `UPDATE arc_jobs SET JobName=?, JobCmd=?, JobConfig=?, Started=?, Done=?, Failed=? WHERE JobId=?`

	statement, err := sqliteDatabase.Prepare(updateJobSQL)
	if err != nil {
		return err
	}

	_, err = statement.Exec(j.JobName, j.JobCmd, j.JobConfig, j.Started, j.Done, j.Failed, j.JobId)
	if err != nil {
		return err
	}

	return nil
}
