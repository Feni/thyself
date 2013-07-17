package web

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"thyself/data"
	"thyself/util"
	"time"
)

func JournalHandler(w http.ResponseWriter, r *http.Request) {
	if user_id := GetLoggedInUser(r); user_id != "" {
		jeText := ""
		rawtime := time.Unix(util.GetTime(r), 0)
		if r.Method == "POST" {
			if jeText = r.FormValue("text"); jeText != "" {
				jeText = util.Slice(jeText, 4000) // limit to 4000 chars
				data.UpsertJournalEntry(user_id, rawtime, jeText)
			}
		} else {
			je, _ := data.GetJournalEntry(user_id, rawtime)
			jeText = je.Je_Text
		}

		vars := mux.Vars(r)
		startTime, endTime := util.GetDayEndpoints(rawtime)
		urlToday := UrlDay(user_id, rawtime)
		// "magic" date based on golang example. don't change it
		// Format is what's produced by javascript toDateString()
		dateStr := rawtime.Format("Mon Jan 2 2006")

		metrics_list := data.GetMetricsByDate(user_id, int(startTime), int(endTime))
		preData := CreatePrefetch(metrics_list, vars["entry_id"])

		rendered := RenderJournal(jeText, preData, urlToday, dateStr, BuildMessages(w, r), "")
		fmt.Fprintln(w, rendered)
	} else {
		http.Redirect(w, r, "/a/login", 302) // TODO: change this url to today's date-time
	}
}

func DemoHandler(w http.ResponseWriter, r *http.Request) {
	rendered := RenderJournal("", PrefetchExample, "/i/demo", time.Unix(util.GetTime(r), 0).Format("Mon Jan 2 2006"), BuildMessages(w, r), PartialRegisterForm)
	fmt.Fprintln(w, rendered)
}

func DemoParseHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" || r.Method == "PUT" {
		structuredRep := CreateEntry(r)
		structuredRep.User_ID = "demo"
		RespondEntry(w, structuredRep)
	} else {
		DemoHandler(w, r)
	}
}

func RenderJournal(journalText, preData, urlDate, dateStr, errors, sidebarExtra string) string {
	actionRendered := string(TemplateActionEntry.Render(map[string]string{
		"urlDate": urlDate}))
	textEmpty := "true"
	if journalText != "" && len(journalText) > 0 {
		textEmpty = "false"
	}
	journalRendered := string(TemplateJE.Render(map[string]string{
		"date":        dateStr,
		"journalText": journalText,
		"empty":       textEmpty,
		"urlDate":     urlDate}))
	pageRendered := string(TemplateMain.Render(map[string]string{
		"sidebar":        errors + actionRendered + sidebarExtra,
		"journalEntry":   journalRendered,
		"prefetchedData": preData,
		"urlDate":        urlDate}))
	return pageRendered
}

func CreatePrefetch(entries_list []data.MetricEntry, entry_id string) string {
	preData := `<script type="text/javascript">`

	renderedEntries := "var defaultEntries = new Thyself.Models.Entries(["
	for i, entry := range entries_list {
		if i != 0 {
			renderedEntries += ","
		}
		renderedEntries += entry.RenderJS()
	}
	renderedEntries += "]);"

	preData += renderedEntries

	preData += `  Thyself.Data.prefetch = defaultEntries;
  	Thyself.Page.sidebarView = new Thyself.Views.EntrySummaryListView({
      collection: defaultEntries,
      el: $('#sidebarActionList')
    }); `

	if entry_id != "" {
		preData += `
			var entryView = new Thyself.Views.EntryEditView({model: Thyself.Data.prefetch.get("` + entry_id + `")});
			entryView.render();
		`
	}
	preData += `</script>`
	return preData
}
