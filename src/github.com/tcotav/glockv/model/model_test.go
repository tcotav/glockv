package model

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"testing"
)

var db *sql.DB
var lkey LKey

func TestCreateKV(t *testing.T) {
	_, err := CreateKV(lkey.Lat, lkey.Long, lkey.Url, lkey.Id)
	if err != nil {
		t.Error(err)
	}
}
func TestGetKV(t *testing.T) {
	keyList, err := GetKV(lkey.Lat, lkey.Long)
	if err != nil {
		t.Error(err)
	}

	lenKeyList := len(keyList)
	if len(keyList) != 1 {
		t.Errorf("Incorrect list length -- expected 1, got: %d", lenKeyList)
	}
	/*for i := 0; i < len(keyList); i++ {
		fmt.Println(keyList[i])
	}*/
}

const createTable = "CREATE TABLE LocValMap ( id INTEGER PRIMARY KEY AUTOINCREMENT, loc varchar(256) not null, lat double not null, lng double not null, extrakey varchar(129), url varchar(256) not null, createdat date);"

const dropTable = "DROP TABLE LocValMap;"

const test_database = "./geoloc.db"

func setup() int {
	// create a sample database
	db, err := sql.Open("sqlite3", test_database)
	if err != nil {
		fmt.Printf("Error getting db connection %s", err)
		return 2
	}

	stmt, err := db.Prepare(createTable)
	if err != nil {
		fmt.Printf("Error setup prepare %s", err)
		return 2
	}

	_, err = stmt.Exec()
	if err != nil {
		fmt.Printf("Error setup create table: %s", err)
		return 2
	}

	lkey = LKey{Lat: 51.5034070, Long: -0.1275920, Id: "prime", Url: "http://gnslngr.us"}

	return 0
}

func teardown() {
	// need full path to db file for unlinking
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error td can't get working dir: %s", err)
		os.Exit(2)
	}

	// remove the test db file
	err = os.Remove(fmt.Sprintf("%s/%s", pwd, test_database))
	if err != nil {
		fmt.Printf("Error remove db: %s", err)
		os.Exit(2)
	}
}

func TestMain(m *testing.M) {
	retCode := setup()
	if retCode == 0 {
		retCode = m.Run()
	}
	teardown()
	os.Exit(retCode)
}
