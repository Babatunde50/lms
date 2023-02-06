package models

import (
	"database/sql"
	"log"
	"strings"
	"time"

	setting "github.com/Babatunde50/lms/pkg/settings"
	_ "github.com/jackc/pgx/v5/stdlib"
)

var db *sql.DB

const maxOpenDbConn = 10
const maxIdleDbConn = 5
const maxDbLifetime = 5 * time.Minute

func Setup() {

	var dsn strings.Builder

	dsn.WriteString("host=" + setting.DatabaseSetting.Host)
	dsn.WriteString(" port=" + setting.DatabaseSetting.Port)
	dsn.WriteString(" dbname=" + setting.DatabaseSetting.Name)
	dsn.WriteString(" user=" + setting.DatabaseSetting.User)
	dsn.WriteString(" password=" + setting.DatabaseSetting.Password)

	log.Println("Connecting to Database...")

	d, err := newDatabase(dsn.String())

	if err != nil {
		log.Fatal(err)
	}

	d.SetMaxOpenConns(maxOpenDbConn)
	d.SetMaxIdleConns(maxIdleDbConn)
	d.SetConnMaxIdleTime(maxDbLifetime)

	err = testDB(d)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to DB successfully!")

	db = d

}

func testDB(d *sql.DB) error {
	err := d.Ping()
	if err != nil {
		return err
	}
	return nil
}

func newDatabase(dsn string) (*sql.DB, error) {
	db, err := sql.Open(setting.DatabaseSetting.Type, dsn)

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
