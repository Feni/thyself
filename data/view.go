package data
import (
	"encoding/json"	
	"strconv"
	"fmt"
)



func (d *MetricDetail) RenderJS() string {
	rendered := "new Thyself.Models.Detail("
	jsonResp, err := json.MarshalIndent(d, "", "\t")
	if err == nil {
		rendered += string(jsonResp)
	}
	rendered += ")"
	return rendered
}


func (e *MetricEntry) RenderJS() string {
	detailsView := "new Thyself.Models.Details(["
	for di, detail := range(e.Details) {
		if di != 0 {
			detailsView += ", "
		}
		detailsView += detail.RenderJS()
	}
	detailsView += "])"
	rendered := fmt.Sprintf(`new Thyself.Models.Entry({
		"id": %s, 
		"user_id": %s,
		"description": %s,
		"time": %d,
		"metric": %s,
		"details": %s})`, strconv.Quote(e.ID), strconv.Quote(e.User_ID), strconv.Quote(e.Description), 
			e.UnixTime , strconv.Quote(e.Metric), detailsView)
	return rendered
}
