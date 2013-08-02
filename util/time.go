package util

import (
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

// given some time within a day
// give the start and end of that day
func GetDayEndpoints(t time.Time) (int64, int64) {
	startTime := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	endTime := time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 999999999, t.Location())
	return startTime.Unix(), endTime.Unix()
}

// Get a time from a request in whatever way possible.
// If a form parameter for time is given, use that.
// Else, get the time from the url using whatever chunks are given
// Else, get the time from the user's request
// Else use the server time
func GetTime(r *http.Request) int64 {
	vars := mux.Vars(r)
	var tNow = time.Now()
	// Parse the time from the request header
	hTimeRaw := r.Header.Get("Date")
	if hTimeRaw != "" {
		hTime, err := http.ParseTime(hTimeRaw)
		if err == nil {
			tNow = hTime
		}
	}
	// Then use the url date modules to construct a date
	resolvedYr := ConvertToInt(vars["year"], int64(tNow.Year()))
	resolvedMth := ConvertToInt(vars["month"], int64(tNow.Month()))
	resolvedDay := ConvertToInt(vars["day"], int64(tNow.Day()))

	resolvedTime := time.Date(int(resolvedYr), time.Month(resolvedMth), int(resolvedDay), tNow.Hour(), tNow.Minute(), tNow.Second(), tNow.Nanosecond(), tNow.Location())

	// Use form time if it's specified. else what we just computed
	// I know.. I know... Doing this here after all this work is wasted effort
	// (could've checked in beginning and returned it)
	// but the code is sooo much cleaner this way
	return ConvertToInt(r.FormValue("time"), resolvedTime.Unix())
}
