package model

import (
	"database/sql"
	"fmt"
	"github.com/TomiHiltunen/geohash-golang"
	_ "github.com/mattn/go-sqlite3"
	"github.com/tcotav/gobase/logr"
)

const ltagsrc = "glmod"

//  dbUrl := fmt.Sprintf("%s:%s@%s", uuser, passwd, dbname)
var dbUrl = ""

type LKey struct {
	Lat    float64
	Long   float64
	Id     string
	Url    string
	Encloc string
}

const createKVQuery = "INSERT INTO LocValMap (loc, lat, lng, extrakey, url) VALUES (?, ?, ?, ?, ?)"

func CreateKV(lat float64, long float64, url string, extraId string) (string, error) {
	db, err := sql.Open("sqlite3", "./geoloc.db")
	if err != nil {
		logr.LogLine(logr.Lerror, ltagsrc, fmt.Sprintf("Error getting db connection %s", err))
		return "", err
	}
	defer db.Close()

	encLoc := geohash.Encode(lat, long)

	stmt, err := db.Prepare(createKVQuery)
	if err != nil {
		logr.LogLine(logr.Lerror, ltagsrc, err.Error())
		return "", err
	}

	res, err := stmt.Exec(encLoc, lat, long, extraId, url)
	if err != nil {
		logr.LogLine(logr.Lerror, ltagsrc, err.Error())
		return "", err
	}
	sourceHostId, err := res.LastInsertId()
	if err != nil {
		logr.LogLine(logr.Lerror, ltagsrc, err.Error())
		return "", err
	}

	return fmt.Sprintf("%s", sourceHostId), nil
}

/*

https://en.wikipedia.org/wiki/Geohash

The table below shows the dimensions of geohash cells at the worst-case scenario at the equator.

geohash length	width	height
1	5,009.4km	4,992.6km
2	1,252.3km	624.1km
3	156.5km	156km
4	39.1km	19.5km
5	4.9km	4.9km
6	1.2km	609.4m
7	152.9m	152.4m
8	38.2m	19m
9	4.8m	4.8m
10	1.2m	59.5cm
11	14.9cm	14.9cm
12	3.7cm	1.9cm


https://github.com/davetroy/geohash-js

However, 'dqcjqcp84c6e' is not centered inside 'dqcjqc', and searching within 'dqcjqc' may miss some desired targets.

So instead, we can use the mathematical properties of the Geohash to quickly calculate the neighbors of 'dqcjqc';
we find that they are: 'dqcjqf','dqcjqb','dqcjr1','dqcjq9','dqcjqd','dqcjr4','dqcjr0','dqcjq8'

This gives us a bounding box around 'dqcjqcp84c6e' roughly 2km x 1.5km and allows for a database search on just 9 keys:

SELECT *
FROM table
WHERE LEFT(geohash,6)
	IN ('dqcjqc', 'dqcjqf','dqcjqb','dqcjr1','dqcjq9','dqcjqd','dqcjr4','dqcjr0','dqcjq8');

SELECT id, loc, lat, lng, url, extrakey
FROM LocValMap
WHERE LEFT(loc, ?)
	-- bounding
	IN (?,?,?,?,?,?,?,?,?)
*/

func scaleAdjacents(encLoc string, matchSize int) []string {
	locAdjacents := geohash.CalculateAllAdjacent(encLoc)
	retAdjacents := make([]string, len(locAdjacents))
	// now scale them down to the match size that we're going to use
	for i := range locAdjacents {
		retAdjacents[i] = locAdjacents[i][:matchSize]
	}
	// now dedupe?
	return retAdjacents
}

const selectKVQuery = "SELECT id, loc, lat, lng, url, extrakey FROM LocValMap WHERE loc = ?"
const selectKVRangeQuery = "SELECT id, loc, lat, lng, url, extrakey FROM LocValMap WHERE substr(loc, 0, ?) IN (?,?,?,?,?,?,?,?,?)"

func GetKV(lat float64, long float64) ([]LKey, error) {
	return GetKVWithScale(lat, long, 6)
}

func GetKVWithScale(lat float64, long float64, matchSize int) ([]LKey, error) {
	db, err := sql.Open("sqlite3", "./geoloc.db")
	if err != nil {
		logr.LogLine(logr.Lerror, ltagsrc, fmt.Sprintf("Error getting db connection %s", err))
		return nil, err
	}
	defer db.Close()

	// encode lat,long
	encLoc := geohash.Encode(lat, long)
	// get adjacents
	locAdjacents := scaleAdjacents(encLoc, matchSize)
	//fmt.Println(locAdjacents)
	/*
		x-x-x
		x-o-x
		x-x-x

		search point 'o' as well all adjacent 'x' boxes.  total of 9 values to search for.
	*/
	rows, err := db.Query(selectKVRangeQuery, matchSize+1, encLoc, locAdjacents[0], locAdjacents[1], locAdjacents[2], locAdjacents[3], locAdjacents[4], locAdjacents[5], locAdjacents[6], locAdjacents[7])
	//rows, err := db.Query(selectKVQuery, encLoc)
	if err != nil {
		logr.LogLine(logr.Lerror, ltagsrc, err.Error())
		return nil, err
	}
	defer rows.Close()

	var retList []LKey
	for rows.Next() {
		var url string
		var id string
		var lat float64
		var long float64
		var encLoc string
		var extraId string
		if err := rows.Scan(&id, &encLoc, &lat, &long, &url, &extraId); err != nil {
			logr.LogLine(logr.Lerror, ltagsrc, err.Error())
			return nil, err
		}
		retList = append(retList, LKey{Id: id, Url: url, Lat: lat, Long: long, Encloc: encLoc})
		//centerLatLong := geohash.Decode(encLoc).Center()
		//fmt.Println(centerLatLong.Lat(), centerLatLong.Lng())
	}
	return retList, nil
}
