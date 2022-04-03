package store

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Records struct {
	db *sql.DB
}

type Issue struct {
	TimeStamp      string
	Id             string
	Subject        string
	StatusId       string
	AssignedToName string
	EstimatedHours string
}

func (rs *Records) Open() error {
	dbfile := "./test.db"

	// os.Remove(dbfile)

	db, err := sql.Open("sqlite3", dbfile)
	if err != nil {
		log.Fatal(err)
		return err
	}

	rs.db = db

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS "issues" ("timestamp" TEXT, "id" TEXT, "subject" TEXT, "status_id" TEXT, "assigned_to_name" TEXT, "estimated_hours" TEXT )`)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

func (rs *Records) Add(col Issue) error {
	_, err := rs.db.Exec(`
      INSERT INTO "issues" ( "timestamp", "id", "subject", "status_id", "assigned_to_name", "estimated_hours" ) VALUES( ?, ?, ?, ?, ?, ? )`,
		col.TimeStamp, col.Id, col.Subject, col.StatusId, col.AssignedToName, col.EstimatedHours)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

func (rs *Records) Close() {
	rs.db.Close()
}
