package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	var jobs_db_path = os.Getenv("JOBS_DB_PATH")
	log.Printf("Creating database as %s\n", jobs_db_path)

	file, err := os.Create(jobs_db_path)
	if err != nil {
		log.Fatal(err.Error())
	}
	file.Close()

	log.Println("Creating table arc_jobs")
	sqliteDatabase, _ := sql.Open("sqlite3", jobs_db_path)
	defer sqliteDatabase.Close()

	createTableSQL := `CREATE TABLE arc_jobs (
		"JobID" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"JobName" TEXT,
		"JobCmd" TEXT,
		"JobConfig" TEXT,
		"Started" TEXT,
		"Done" INTEGER,
		"Failed" INTEGER
		);`

	statement, err := sqliteDatabase.Prepare(createTableSQL)
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec()

	log.Println("Done")
}
