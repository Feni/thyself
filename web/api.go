package web

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
	"thyself/data"
	"thyself/log"
	"thyself/nlp"
	"thyself/util"
)

func EntryListHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := cookieStore.Get(r, defaultSessionName)
	user_id := fmt.Sprintf("%s", session.Values["user_id"])

	switch r.Method {
	case "GET":
		{
			if user_id == "" {
				//401 Unauthorized
				// TODO: return public items
				http.Error(w, "401 : Unauthorized!! Login before you send a request", 400)
			} else {
				//				fmt.Println("Getting stuff. has user id. ", r.FormValue("until"), r.FormValue("since"))
				var metrics_list []data.MetricEntry
				if r.FormValue("since") != "" && r.FormValue("until") != "" {
					start_date, sd_err := strconv.Atoi(r.FormValue("since"))
					end_date, ed_err := strconv.Atoi(r.FormValue("until"))
					if sd_err != nil || ed_err != nil {
						http.Error(w, "400 : Baaaddd Request!! 'since' & 'until' should be valid unix-time integers", 400)
						return
					} else {
						metrics_list = data.GetMetricsByDate(user_id, start_date, end_date)
					}
				} else {
					// get ALL Metrics.
				}

				jsonRep, err := json.MarshalIndent(metrics_list, "", "\t") // 4 spaces works too, but tab is more byte-efficient.

				if err == nil {
					fmt.Fprintln(w, string(jsonRep))
				} else {
					http.Error(w, fmt.Sprintf("500: Ooops. Json Encoding Error %s", err), 500) // internal server error
				}

				//  week_day := r.FormValue("week_day") // 0 to 6. 0 = sunday
				//  week := r.FormValue("week")

			}
		}
	case "POST", "PUT": { // Create
		structuredRep := CreateEntry(r)
		RespondEntry(w, structuredRep)
	} 
	default:
		http.Error(w, "400 : Baaaddd Request!! String parameter 'description' is required. Try /m?description=I+fail+at+http", 400) // Client error
	}
}

func EntryItemHandler(w http.ResponseWriter, r *http.Request) {
//	session, _ := cookieStore.Get(r, defaultSessionName)
//	user_id := fmt.Sprintf("%s", session.Values["user_id"])
	vars := mux.Vars(r)
	entry_id := vars["entry_id"]

	if entry_id != "" {
		switch r.Method {
		case "GET": {
			// TODO : retrieve a specific metric
			log.Info("Is get method")
		}
		case "PUT": { // Update 
			// Check if the user is the owner of this entry

			// Update that specific metric
			log.Info("Got API request for entry id ", entry_id)
			var metric data.MetricEntry
			body, err := ioutil.ReadAll(r.Body)
			log.Debug(err, "EntryListHandler - entry_id - Error reading create entry content body")
			log.Info("Raw content body ", string(body))
			if len(body) > 0 {
				err := json.Unmarshal(body, &metric)
				log.Debug(err, "ERROR; body Unmarhsall error ")
			}
			log.Info("Unmarshalled api request body ", metric)

			if GetLoggedInUser(r) == metric.User_ID {
				valid, errMsg := metric.Validate()
				if valid{
					log.Info("Auth matches and data is valid. Updating ", metric)
					data.UpdateMetric(&metric)
					/* TODO: 
					maybe return string(body) Make sure this isn't a security risk 
					I get a bad feeling returning user defined stuff...
					*/
					fmt.Fprintln(w, "{}" )
				}else {
					log.Info("Update failed. Error is " + errMsg )
					http.Error(w, "400 : Baaaddd Request!! " + errMsg, 400) // Client error
				}
			}else {
				http.Error(w, "401 : Unauthorized!! Login before you send a request", 400)
			}
		}
		case "DELETE":{
			data.DeleteMetric(GetLoggedInUser(r), entry_id)
		}
		}
	}
}

func CreateEntry(r *http.Request) *data.MetricEntry {
	var metric data.MetricEntry;
	body, err := ioutil.ReadAll(r.Body)
	log.Debug(err, "Error reading create entry content body")
	if len(body) > 0 {
		err := json.Unmarshal(body, &metric)
		log.Debug(err, "ERROR; body Unmarhsall error " + string(body))
	}
	// Overwrite with form fields if not specified in content body
	if metric.Description == "" {
		metric.Description = r.FormValue("description")
		log.Info("Form value description is ",  metric.Description)
	}
	if metric.UnixTime == 0 {
		metric.UnixTime = util.GetTime(r)
	}
	if metric.Privacy == ""{
		metric.Privacy = "PUBLIC"
	}
	if metric.Description != "" {
		structuredRep := nlp.Parse(metric.Description)
		if structuredRep != nil {
			structuredRep.UnixTime = metric.UnixTime
			structuredRep.ID = util.GenID(8)
			user_id := GetLoggedInUser(r)
			// id, unixTime and desc are already specified
			if user_id != "" && structuredRep.Metric != "" {
				structuredRep.User_ID = user_id
				data.AddMetric(user_id, structuredRep) // TODO ; Uncomment
			}
			return structuredRep
		}
	}
	// TODO: return metric?
	return nil
}

func RespondEntry(w http.ResponseWriter, structuredRep *data.MetricEntry) {
	if structuredRep != nil {
		jsonRep, err := json.MarshalIndent(structuredRep, "", "\t") // 4 spaces works too, but tab is more byte-efficient.
		if err == nil {
			fmt.Fprintln(w, string(jsonRep))
			log.Info("Json rep is " + string(jsonRep))
		} else {
			http.Error(w, fmt.Sprintf("500: Ooops. Json Encoding Error %s", err), 500) // internal server error
		}
	} else {
		// description was null
		http.Error(w, "400 : Baaaddd Request!! String parameter 'description' is required. Try /m?description=I+fail+at+http", 400) // Client error
	}
}
