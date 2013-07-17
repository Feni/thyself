package web

import (
	"fmt"
	"github.com/gorilla/sessions"
	"github.com/hoisie/mustache"
	"io/ioutil"
	"net/http"
	"strings"
	"thyself/log"
)

//var TemplateBase *mustache.Template
var TemplateMessage *mustache.Template
var TemplateMain *mustache.Template

var TemplateJE *mustache.Template
var TemplateActionEntry *mustache.Template

var PartialLoginForm, PartialRegisterForm string
var PartialScripts string

var AnonHomePage string

const MSG_SUCCESS = "success"
const MSG_ERROR = "alert"
const MSG_NORMAL = ""

// TODO : Write these values down before deployment
var cookieStore = sessions.NewCookieStore(
	[]byte("94}2*=ND!PI:{ztNb3p~M:Bf_zAT&K.*1FenimblxAsdfft0FtkEtg2iNT4361mb"), // Cookie auth
	[]byte("JGVfc~wGiSnZTj[^@[]ITl[Ne)qp#Qkf"))                                 // Encryption

const defaultSessionName = "thyself_private"

func LoadTemplates() {
	templateLoc := "/var/www/thyself/src/thyself/public/"

	loginForm, err := ioutil.ReadFile(templateLoc + "partials/anon/loginForm.html")
	log.Debug(err, "Error loading login form")
	PartialLoginForm = string(loginForm)

	registerForm, err := ioutil.ReadFile(templateLoc + "partials/anon/registerForm.html")
	log.Debug(err, "Error loading registration form")
	PartialRegisterForm = string(registerForm)

	scripts, err := ioutil.ReadFile(templateLoc + "partials/scripts.html")
	log.Debug(err, "Error loading scripts")
	PartialScripts = string(scripts)

	actionEntry, err := ioutil.ReadFile(templateLoc + "partials/actionEntry.html")
	log.Debug(err, "Error loading action entry")
	TemplateActionEntry, _ = mustache.ParseString(string(actionEntry))

	jeTmpl, err := ioutil.ReadFile(templateLoc + "partials/journalEntry.html")
	log.Debug(err, "Error loading journal entry")
	TemplateJE, _ = mustache.ParseString(string(jeTmpl))

	base, err := ioutil.ReadFile(templateLoc + "templates/base.html")
	log.Debug(err, "Error loading base template")

	messageTmpl, err := ioutil.ReadFile(templateLoc + "templates/message.html")
	log.Debug(err, "Error loading message template")
	msgTemplTemp := string(mustache.Render(string(base),
		map[string]string{
			"content": string(messageTmpl),
			"scripts": ""}))

	TemplateMessage, err = mustache.ParseString(msgTemplTemp)
	log.Debug(err, "Error rendering message template")

	mainTmpl, err := ioutil.ReadFile(templateLoc + "templates/main.html")
	log.Debug(err, "Error loading main template")
	mainTmplTemp := string(
		mustache.Render(string(base),
			map[string]string{
				"content": string(mainTmpl),
				"scripts": string(PartialScripts) + "{{{prefetchedData}}}"}))
	TemplateMain, err = mustache.ParseString(mainTmplTemp)
	log.Debug(err, "Error rendering main template")

	homepageTmpl, err := ioutil.ReadFile(templateLoc + "templates/homepage.html")
	log.Debug(err, "Error loading homepage template")
	homepageTmplTemp := string(
		mustache.Render(string(base),
			map[string]string{
				"content": string(homepageTmpl),
				"scripts": ""}))

	// Most of the time, this is what we'll be serving up.
	// So just cache it and return it.
	AnonHomePage = string(mustache.Render(homepageTmplTemp,
		map[string]string{
			"register": PartialRegisterForm,
			"login":    PartialLoginForm}))

	// Note that escaping is space sensitive {{var}} != {{ var }}. Only the first one works
}

func RenderMessage(message, msgtype string) string {
	return `<li class="alert-box ` + msgtype + `" > ` +
		message + `
    </li>`
}

// Builds a list of messages from previously flashed messages
// This is probably inefficient. Doesn't matter for now, but later
// use a different structure for flashed messages
func BuildMessages(w http.ResponseWriter, r *http.Request) string {
	session, _ := cookieStore.Get(r, defaultSessionName)
	allMessages := ""

	flashes := session.Flashes()
	//	fmt.Printf("Flashes are %v", flashes)
	for _, message := range flashes {
		parts := strings.SplitN(fmt.Sprintf("%s", message), ":", 2)
		if len(parts) == 2 {
			allMessages += RenderMessage(parts[1], parts[0])
		}
	}
	session.Save(r, w) // Is this necessary to just read the flashes?

	//	fmt.Printf("All messages are : %s", allMessages)
	if allMessages != "" {
		return "<ul class='row message_flashes'>" + allMessages + "</ul>"
	} else {
		return ""
	}
}

const PrefetchExample = ` 
<script type="text/javascript">
 defaultEntries = new Thyself.Models.Entries([
    new Thyself.Models.Entry({
      "id": "XEScxaet",
      "user_id": "demo",
      "description": "I slept for 7.5 hours",
      "time": new Date(1372195416 * 1000),
      "metric": "sleep",
      "details": new Thyself.Models.Details([
        new Thyself.Models.Detail({
          "amount": "7.5",
          "type": "hours",
          "group": "time"
        })
      ])
    }), new Thyself.Models.Entry({
      "id": "oaOR5OlY",
      "user_id": "demo",
      "description": "I ran for 2.5 miles in 15 minutes",
      "time": new Date(1372495416 * 1000),
      "metric": "run",
      "details": new Thyself.Models.Details([
        new Thyself.Models.Detail({
          "amount": "15",
          "type": "minutes",
          "group": "time"
        }), new Thyself.Models.Detail({
          "amount": "2.5",
          "type": "miles",
          "group": "distance"
        })
      ])
    }), new Thyself.Models.Entry({
      "id": "oQdR5OlY",
      "user_id": "demo",
      "description": "Washed half load of blue jeans at medium temperature",
      "time": new Date(1372495418 * 1000),
      "metric": "Laundry",
      "details": new Thyself.Models.Details([
      	new Thyself.Models.Detail({
          "type": "blue",
          "group": "color"
        }),
      	new Thyself.Models.Detail({
          "type": "jeans",
          "group": "clothes"
        }),
        new Thyself.Models.Detail({
          "amount": "0.5",
          "type": "load"
        }), 
		new Thyself.Models.Detail({
          "type": "medium",
          "group": "temperature"
        })
      ])
    }), new Thyself.Models.Entry({
      "id": "7YRakSmr",
      "user_id": "demo",
      "description": "I ate 4 cookies",
      "time": new Date(1372199416 * 1000),
      "metric": "eat",
      "details": new Thyself.Models.Details([
        new Thyself.Models.Detail({
          "amount": "4",
          "type": "cookies",
          "group": "food"
        }), new Thyself.Models.Detail({
          "amount": "274",
          "type": "calories",
          "group": "nutrition"
        })
      ])
    })
  ]);

  Thyself.Data.prefetch = defaultEntries;
  Thyself.Page.sidebarView = new Thyself.Views.EntrySummaryListView({
      collection: defaultEntries,
      el: $('#sidebarActionList'),
    });
</script>`
