package web

import (
	"fmt"
	"net/http"
	"time"
)

func UrlDay(user_id string, t time.Time) string {
	//t := time.Now() 		// TODO: fix time-zone issues. Use may not be in same time zone as us.
	return fmt.Sprintf("/u/%s/%d/%d/%d", user_id, t.Year(), t.Month(), t.Day())
}

func HomepageHandler(w http.ResponseWriter, r *http.Request) {
	if isAuth(r) {
		http.Redirect(w, r, UrlDay(GetLoggedInUser(r), time.Now()), 302) // TODO: change this url to today's date-time
	} else {
		fmt.Fprintln(w, AnonHomePage)
	}
}
