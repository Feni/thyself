package web

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func EntriesHandler(w http.ResponseWriter, r *http.Request) {
	auth := isAuth(r)
	if auth {
		vars := mux.Vars(r)
		if url_user_id := vars["user_id"]; url_user_id == GetLoggedInUser(r) { // User is at her own url
			switch r.Method {
			case "GET":
				{
					fmt.Println("entry.go Is a get method")
				}
			case "POST":
				{
					// {user_id}/{year}/{month}/{day}/m
					createdEntry := CreateEntry(r) // data.MetricEntry
					if createdEntry != nil {
						http.Redirect(w, r, createdEntry.GetEntryURL(), 302)
					} else {
						// header.get Referer
						http.Redirect(w, r, "/error", 302)
					}
				}
			}
		} else {
			fmt.Fprintln(w, "User auth fail")
		}

	} else { // not auth
		http.Redirect(w, r, "/a/login", 302)
	}

}
