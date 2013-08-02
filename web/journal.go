package web

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"thyself/data"
	"thyself/util"
	"time"
)

func GetJournalText(r *http.Request, user_id string, rawtime time.Time) string{
	jeText := ""
	if r.Method == "POST" {
		if jeText = r.FormValue("text"); jeText != "" {
			jeText = util.Slice(jeText, 4000) // limit to 4000 chars
			data.UpsertJournalEntry(user_id, rawtime, jeText)
		}
	} else {
		je, _ := data.GetJournalEntry(user_id, rawtime)
		jeText = je.Je_Text
	}
	return jeText
}

func JournalHandler(w http.ResponseWriter, r *http.Request) {
	if user_id := GetLoggedInUser(r); user_id != "" {
		JournalHelper(w, r, user_id, "")
	} else {
		http.Redirect(w, r, "/a/login", 302) 
	}
}

func JournalHelper(w http.ResponseWriter, r *http.Request, user_id, extraScript string){
	rawtime := time.Unix(util.GetTime(r), 0)
	jeText := GetJournalText(r, user_id, rawtime)

	vars := mux.Vars(r)
	preData := CreatePrefetch(rawtime, jeText, user_id, vars["entry_id"])
	
	urlToday := UrlDay(user_id, rawtime)
	// "magic" date based on golang example. don't change it
	// Format is what's produced by javascript toDateString()
	dateStr := rawtime.Format("Mon Jan 2 2006")

	rendered := RenderJournal(jeText, preData + extraScript, urlToday, dateStr, BuildMessages(w, r), "")
	fmt.Fprintln(w, rendered)
}

var DemoInputRotator = `
	<script type="text/javascript">

	document.exampleIndex = 0;
	document.exampleInputs=["I slept for 7.5 hours", "I played basketball for 2 hours and scored 10 points", 
	"I watered the cactus",	"I took my vitamins", 
	"I ate pasta for lunch with Paul at the mall","I bought new glasses", 
	"Drank 2 cups of water", "I studied for my calculus exam for 3 hours at the library"];
 	setInterval(function () {
 		document.exampleIndex = (document.exampleIndex + 1) % (document.exampleInputs.length);
        $("#mEntryForm #description").attr("placeholder", document.exampleInputs[document.exampleIndex]);
    },3500);

		
	</script>
`

func DemoHandler(w http.ResponseWriter, r *http.Request) {
	rendered := RenderJournal("", PrefetchExample + DemoInputRotator, "/i/demo", time.Unix(util.GetTime(r), 0).Format("Mon Jan 2 2006"), BuildMessages(w, r), PartialRegisterForm)
	fmt.Fprintln(w, rendered)
}

func DemoParseHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" || r.Method == "PUT" {
		structuredRep := CreateEntry(r)
		if structuredRep != nil{
			structuredRep.User_ID = "demo"
		}
		RespondEntry(w, structuredRep)
	} else {
		DemoHandler(w, r)
	}
}

func RenderJournal(journalText, preData, urlDate, dateStr, flashes, sidebarExtra string) string {
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
		"sidebar":        actionRendered + sidebarExtra,
		"journalEntry":   journalRendered,
		"prefetchedData": preData,
		"urlDate":        urlDate, 
		"flashes" : flashes}))
	return pageRendered
}

// Returns rendered prefetch and raw journal entry text
// TODO : escape quotes properly
func CreatePrefetch(rawtime time.Time, jeText, user_id, entry_id string) string {
	startTime, endTime := util.GetDayEndpoints(rawtime)
	entries_list := data.GetMetricsByDate(user_id, int(startTime), int(endTime))

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

	//journalEntry := "Thyself.Data.Journal = new Thyself.Models.JournalEntry();"
	// Because this is after the thyself.js, the main application doesn't know who the user or
	// time is until after this module is loaded
	preData += `Thyself.Data.ContextDate = new Date(` + fmt.Sprintf("%d",rawtime.Unix()) + `*1000);`
	preData += `Thyself.Data.ContextUser = "` + user_id + `";`

	preData += `  Thyself.Data.Entries = defaultEntries;
  	Thyself.Page.sidebarView = new Thyself.Views.EntrySummaryListView({
      collection: defaultEntries,
      el: $('#sidebarActionList')
    }); `

	if entry_id != "" {
		preData += `
			var entryView = new Thyself.Views.EntryEditView({model: Thyself.Data.Entries.get("` + entry_id + `")});
			entryView.render();
		`
	}
	preData += `</script>`
	return preData
}
