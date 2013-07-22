package data

import (
	"database/sql"
	"fmt"
	"regexp"
	"strings"
	"thyself/log"
	"time"
	"strconv"
)

type MetricDetail struct {
	Amount string `json:"amount,omitempty"` // Quantity. float64 on trasmission. But string here so we can omit empty without omitting zero
	Type   string `json:"type,omitempty"`   // Units. Length = 32 chars
	Group  string `json:"group,omitempty"`  // Property / attribute. Length = 32 chars
}

type MetricEntry struct {
	ID          string `json:"id"`
	User_ID     string `json:"user_id"`
	Description string `json:"description"`
	UnixTime    int64  `json:"time"` // TODO: rename in code
	//	LinkedData  string         `json:"linked_data,omitempty"`
	Metric  string         `json:"metric"`
	Details []MetricDetail `json:"details"`
}

func AddMetric(user_id string, metric *MetricEntry) {
	log.Info("data : Add Metric : ", metric, " : USER : ", user_id)
	// (user_id, me_time, me_id, action, description, details)
	_, err := SQL_ADD_METRIC.Exec(user_id, metric.ID, metric.UnixTime, metric.Metric, metric.Description)
	if err == nil {
		for _, detail := range metric.Details {
			if detail.Amount != "" {
				_, err := SQL_ADD_DETAIL.Exec(metric.ID, detail.Group, detail.Type, detail.Amount)
				log.Debug(err, "DATA: Detail : Insertion : FAILED : ")
			} else {
				_, err := SQL_ADD_DETAIL_NO_AMT.Exec(metric.ID, detail.Group, detail.Type)
				log.Debug(err, "DATA: Detail - no amt : Insertion : FAILED : ")
			}
		}
	} else {
		log.Debug(err, "DATA : Metric : Insertion : FAILED ")
	}
}

func DeleteMetric(user_id, metric_id string){
	_, err := SQL_DELETE_ME.Exec(user_id, metric_id)
	log.Debug(err, "DATA: Delete metric FAILED ")
	
}

// For now just delete the old metric and re-add it
func UpdateMetric(metric *MetricEntry) {
	if metric.ID != "" && metric.User_ID != "" {
		DeleteMetric(metric.User_ID, metric.ID)
		AddMetric(metric.User_ID, metric)
	}
}

func GetMetricsByDate(user_id string, start_date, end_date int) []MetricEntry {
	select_metrics := make([]MetricEntry, 0, 5)
	var rows, err = SQL_RETRIEVE_METRIC_DATE_RANGE.Query(user_id, start_date, end_date)
	if err == nil {
		for rows.Next() {
			var rowEntry = MetricEntry{User_ID: user_id}
			if err := rows.Scan(&(rowEntry.ID), &(rowEntry.UnixTime), &(rowEntry.Metric), &(rowEntry.Description)); err != nil {
				log.Debug(err, "DATA : Metric : Retrieval Scan : FAILED : ")
			} else {
				rowEntry.Details = make([]MetricDetail, 0, 5)
				details_rows, d_err := SQL_RETRIEVE_DETAILS.Query(rowEntry.ID)
				if d_err != nil {
					log.Debug(d_err, "DATA : All Detail : Retrieval Query : FAILED : ")
				} else {
					for details_rows.Next() {
						en_detail := MetricDetail{}
						var amt, grp sql.NullString
						err := details_rows.Scan(&grp, &(en_detail.Type), &amt)
						log.Debug(err, "DATA : Detail : Retrieval Scan : FAILED  ")
						en_detail.Amount = amt.String // amt.Valid
						en_detail.Group = grp.String
						rowEntry.Details = append(rowEntry.Details, en_detail)
					}
				}
			}
			select_metrics = append(select_metrics, rowEntry)
		}
	} else {
		log.Debug(err, "DATA : Metric : Retrieval Query : FAILED : ")
	}
	return select_metrics
}

func (e *MetricEntry) GetDateURL() string {
	createTime := time.Unix(e.UnixTime, 0)
	userAtTime := fmt.Sprintf("/u/%s/%d/%d/%d", e.User_ID, createTime.Year(), createTime.Month(), createTime.Day())
	return userAtTime
}

func (e *MetricEntry) GetMetricURL() string {
	metricUrl := e.GetDateURL() + "/m/" + strings.TrimSpace(e.Metric)
	return metricUrl
}

var r_non_alpha_num, _ = regexp.Compile("[^A-Za-z0-9_-]")
var r_space, _ = regexp.Compile("\\s")

func (e *MetricEntry) GetEntryURL() string {
	// replace all spaces with -
	spaceLess := r_space.ReplaceAllString(e.Description, "-")
	cleanDesc := r_non_alpha_num.ReplaceAllString(spaceLess, "-") // delete non-alpha num
	if len(cleanDesc) > 80 {
		cleanDesc = cleanDesc[0:80]
	}
	entryUrl := e.GetMetricURL() + "/e/" + e.ID + "/" + strings.TrimSpace(cleanDesc)
	return entryUrl
}


func (e *MetricEntry) Validate() (bool, string){
	if len(e.ID) != 8 {
		return false, "Invalid Entry ID. Must be a string of length 8"
	}
	if len(e.User_ID) != 5 {
		return false, "Invalid User ID. Must be a string of length 5"
	}
	if e.UnixTime == 0 {
		return false, "Invalid Time. Not defined."
	}
	if e.Metric == "" {
		return false, "Invalid Metric. Not defined."
	}
	if len(e.Description) > 160 {
		return false, "Invalid Description. Max length is 160 chars"	
	}
	for _, detail := range e.Details {	
		if detail.Amount != ""{ // If a string is defined, it must be a number
			_, err := strconv.ParseFloat(detail.Amount, 64)
			if err != nil{
				return false, "Amount must be a float"
			}
		}
		if len(detail.Group) > 32{
			return false, "Max group name length is 32"
		}
		if detail.Type == "" || len(detail.Type) > 160{
			return false, "Detail type must be defined and be < 160 chars"
		}
	}
	return true, ""
}
