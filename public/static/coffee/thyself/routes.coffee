# /u/{user}/2013/07/27/m/{metric_id}/metric-name/here/e/{entry_id}/entry-description-here-which-may-be-long
class ThyselfRouter extends Backbone.Router
  routes:
    "":          "index"
    "u":         "settings" 
    "i/demo":    "demoMain" 
    "u/:user/:year/:month/:day" : "journal"
    "u/:user/:year/:month/:day/m/:metric_name/e/:entry_id/:entry_desc": "entrySummary"

  index: =>
    indexView = new Thyself.Views.IndexView() # Just bind to registration clicks, etc.
  settings: =>
    settingsView = new Thyself.Views.SettingsView
    settingsView.render()
  journal: (user, year, month, day) =>
    journalView = new Thyself.Views.JournalView()
    journalView.render()
  entrySummary: (user, year, month, day, metric_name, entry_id, entry_desc) ->
    entry = Thyself.Data.Entries.get(entry_id)
    entryView = new Thyself.Views.EntryEditView({model: entry})
    entryView.render()
  demoMain: () =>
    demoEntry = new Thyself.Models.JournalEntry({
      user_id: "demo",
      text: ""
    })
    journalView = new Thyself.Views.JournalView({model: demoEntry})
    journalView.render()
