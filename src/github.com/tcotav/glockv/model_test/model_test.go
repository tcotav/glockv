package model_test

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/tcotav/glockv/model"
	"os"
	"testing"
)

var db *sql.DB
var lkey model.LKey

func TestCreateKV(t *testing.T) {
	_, err := model.CreateKV(lkey.Lat, lkey.Long, lkey.Url, lkey.Id)
	if err != nil {
		t.Error(err)
	}
}
func TestGetKV(t *testing.T) {
	keyList, err := model.GetKV(lkey.Lat, lkey.Long)
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
		fmt.Printf("Error prepare %s", err)
		return 2
	}

	_, err = stmt.Exec()
	if err != nil {
		fmt.Printf("Error create table: %s", err)
		return 2
	}

	lkey = model.LKey{Lat: 51.5034070, Long: -0.1275920, Id: "prime", Url: "http://gnslngr.us"}

	return 0
}

func teardown() {
	// close the connection
	//db.Close()

	// remove the test file
	err := os.Remove(test_database)
	if err != nil {
		fmt.Printf("Error create table: %s", err)
		os.Exit(2)
	}
}

func TestMain(m *testing.M) {
	// your func
	retCode := setup()

	if retCode == 0 {
		retCode = m.Run()
	}
	// your func
	teardown()
	// call with result of m.Run()
	os.Exit(retCode)
}
