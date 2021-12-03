package utils

import (
	"strconv"
	"time"
)

func TimeStringToSec(text string)(step int){
	ts, _ := strconv.Atoi(text[:len(text)-1])
	unit := text[len(text)-1]
	switch string(unit){
	case "y":
		step = ts*365*24*60*60
	case "w":
		step = ts*7*24*60*60
	case "d":
		step =ts*24*60*60
	case "h":
		step = ts*60*60
	case "m":
		step = ts*60
	case "s":
		step = ts
	}
	return
}


func UnixTsFormat(ts int64) string {
	return time.Unix(ts, 0).Format("2006-01-02 15:04:05")
}
