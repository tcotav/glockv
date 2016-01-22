package main

import (
	"encoding/json"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/tcotav/glockv/model"
	"github.com/tcotav/gobase/config"
	"github.com/tcotav/gobase/logr"
	"net/http"
	"os"
)

const ltagsrc = "glkvm"

//  dbUrl := fmt.Sprintf("%s:%s@%s", uuser, passwd, dbname)
var dbUrl = ""

func CreateKV(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var lkey model.LKey
	err := decoder.Decode(&lkey)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	retString, err := model.CreateKV(lkey.Lat, lkey.Long, lkey.Url, lkey.Id)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(retString))
}

func main() {
	config, err := config.ParseConfig("glockv.cfg")
	logr.LogLine(logr.Linfo, ltagsrc, fmt.Sprintf("%+v", config))
	if err != nil {
		fmt.Print(err)
		os.Exit(2)
	}

	dbUrl = config["db_url"]
	if dbUrl != "" {
		http.HandleFunc("/kv/create/", CreateKV)
		//serviceMap["/kv/create"] = "Create a key-value pair"
	}

	listenPort := config["web_listen_port"]
	if listenPort == "" {
		// use a default listen port
		listenPort = "8080"
	}

	testMode := config["test_mode"]
	if testMode == "1" {
		lkey := model.LKey{Lat: 51.5034070, Long: -0.1275920, Id: "prime", Url: "http://gnslngr.us"}

		/*
			_, err := model.CreateKV(lkey.Lat, lkey.Long, lkey.Url, lkey.Id)
			if err != nil {
				fmt.Print(err)
				os.Exit(2)
			}*/

		//fmt.Println(str)
		keyList, err := model.GetKV(lkey.Lat, lkey.Long)
		if err != nil {
			fmt.Print(err)
			os.Exit(2)
		}
		for i := 0; i < len(keyList); i++ {
			fmt.Println(keyList[i])
		}
		os.Exit(0) // normal exit
	}

	//http.HandleFunc("/", services)
	logr.LogLine(logr.Linfo, ltagsrc, fmt.Sprintf("webservice starting on port:%s", listenPort))
	http.ListenAndServe(fmt.Sprintf(":%s", listenPort), nil)

}
