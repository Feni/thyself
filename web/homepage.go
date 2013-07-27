package web

import (
	"fmt"
	"net/http"
	"time"
	"thyself/util"
)

func UrlDay(user_id string, t time.Time) string {
	return fmt.Sprintf("/u/%s/%d/%d/%d", user_id, t.Year(), t.Month(), t.Day())
}

func HomepageHandler(w http.ResponseWriter, r *http.Request) {
	//LoadTemplates()	//	 todo : remove for prod
	if isAuth(r) {
		http.Redirect(w, r, UrlDay(GetLoggedInUser(r), time.Unix(util.GetTime(r), 0)), 302) // TODO: change this url to today's date-time
	} else {
		fmt.Fprintln(w, AnonHomePage)
	}
}
