package web

import (
	"fmt"
	"net/http"
	"time"
	//"thyself/util"
)

func UrlDay(user_id string, t time.Time) string {
	return fmt.Sprintf("/u/%s/%d/%d/%d", user_id, t.Year(), t.Month(), t.Day())
}

// If our server time doesn't match with their local time, redirect to proper page
var TimeRedirectScript = `
<script type="text/javascript">
	var timeNow = new Date();
	if(Thyself.Data.ContextDate == undefined || Thyself.Data.ContextDate.getDate() != timeNow.getDate()){
		window.location = "/u/" + Thyself.Data.ContextUser + "/" + timeNow.getFullYear() +"/" + (timeNow.getMonth() + 1)  + "/" + timeNow.getDate();
	}
</script>
`

func HomepageHandler(w http.ResponseWriter, r *http.Request) {
	LoadTemplates(true)	//	 todo : remove for prod
	if user_id := GetLoggedInUser(r); user_id != "" { 
		//http.Redirect(w, r, UrlDay(GetLoggedInUser(r), time.Unix(util.GetTime(r), 0)), 302) // TODO: change this url to today's date-time
		JournalHelper(w, r, user_id, TimeRedirectScript)
	} else {
		fmt.Fprintln(w, AnonHomePage)
	}
}
