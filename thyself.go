package main

// mux, session, schema, context
import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"thyself/data"
	"thyself/log"
	"thyself/web"
	"flag"
)

var isDev bool;

func initServer() {
	web.LoadTemplates(isDev) // Loads the template for all of the pages into memory
	globalRouter := mux.NewRouter()

	// Pages
	globalRouter.HandleFunc("/", web.HomepageHandler)
	//globalRouter.HandleFunc("/journal", web.JournalHandler)

	// View some detail about a user.
	globalRouter.HandleFunc("/u/{user_id}/{year}/{month}/{day}", web.JournalHandler)                                           // POST for updating journal
	globalRouter.HandleFunc("/u/{user_id}/{year}/{month}/{day}/m", web.EntriesHandler)                                         // POST for adding/parsing entries. Will redirecto to newly created entry page
	globalRouter.HandleFunc("/u/{user_id}/{year}/{month}/{day}/m/{metric_name}/e/{entry_id}/{entry_desc}", web.JournalHandler) // TODO - Entry summary

	// Account Management - authorized for self only
	globalRouter.HandleFunc("/a", web.HomepageHandler)
	globalRouter.HandleFunc("/a/login", web.LoginHandler)
	globalRouter.HandleFunc("/a/logout", web.LogoutHandler)
	globalRouter.HandleFunc("/a/register", web.RegisterHandler)

	// i = informational. For site-support pages like
	// demo, api docs, blog, etc.
	globalRouter.HandleFunc("/i/demo", web.DemoHandler)
	globalRouter.HandleFunc("/i/demo/m", web.DemoParseHandler)
	globalRouter.HandleFunc("/i/terms", web.TermsHandler)

	// Parse will be handled here too
	apiRouter := globalRouter.PathPrefix("/api/v0/").Subrouter()
	apiRouter.HandleFunc("/entries", web.EntryListHandler)
	apiRouter.HandleFunc("/entries/", web.EntryListHandler)
	apiRouter.HandleFunc("/entries/{entry_id}", web.EntryItemHandler)

	//	Queries("key", "value")
	http.Handle("/", globalRouter)
	http.ListenAndServe(":8080", nil)
}

func main() {
	fmt.Println("Starting Thyself.io Server")
	//startTime := time.Now()
	log.InitLog()

	devPtr := flag.Bool("dev",false,"Toggles dev mode (use local scripts, headers and enable debug code)")
	flag.Parse()
	isDev = *devPtr

	if isDev{
		fmt.Println("RUNNING IN DEVELOPMENT MODE")
		log.Info("WARNING: Running in Development Mode")		
	}else  {
		fmt.Println("Running in production mode")
		log.Info("Running in production mode")		
	}

	data.RedisInit()
	data.SqlInit()
	log.Info("Server started")
	initServer()
	//shared.Log.Println("\n Server ran for %s \n ", time.Since(startTime))
}
