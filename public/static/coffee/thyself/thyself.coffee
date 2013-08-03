
# Hook to all link-clicks on the page so we can route it using push-state
$(document).delegate "a", "click", (event) -> 
  # Get the anchor href and protcol
  href = $(this).attr("href");
  #protocol = this.protocol + "//";  # http://
  # Ensure the protocol is not part of URL, meaning its relative.
  if !event.altKey and !event.ctrlKey and !event.metaKey and !event.shiftKey
    if href.substring(0, 1) == '/' && href.substring(0,3) != "/a/" && href != "/"  # only catch urls starting with /. abs urls are treated normally
      # Try to render the page
      Thyself.router.navigate href, { trigger: true }
      # If content could not be loaded by JS, force a server side reload
      if $("#journal_entry").html() == ""
        return true # Enable default behavior
      else
        event.preventDefault();
        return false  # Disable default behavior


$("#mEntryForm").submit( () -> 
  actionUrl = $(this).attr('action')
  newEntry = new Thyself.Models.Entry(); 
  if actionUrl == '/i/demo/m'
    newEntry.url = '/i/demo/m'
  descriptionField = $(this).find("#description")
  timeNow = new Date()
  mergedTime = new Date(Thyself.Data.ContextDate.getFullYear(), Thyself.Data.ContextDate.getMonth(), Thyself.Data.ContextDate.getDate(), 
    timeNow.getHours(), timeNow.getMinutes(), timeNow.getSeconds(), timeNow.getMilliseconds())

  entryFields = { 
    description: descriptionField.val(),
    time: Math.round(mergedTime.getTime() / 1000)
  };
  newEntry.save(entryFields, { 
  success: (entry) => 
    console.log(entry.toJSON()); 
    detailsCollection = new Thyself.Models.Details(entry.get("details"))
    entry.set("details", detailsCollection)
    Thyself.Page.sidebarView.collection.add(newEntry)
    Thyself.Page.sidebarView.collection.sort()
    Thyself.Page.sidebarView.render()
  error: (model, response) =>
    console.log(model)
    console.log(response)
  }
  )
  descriptionField.val("")
  return false;
);

$(".alert-box").delay(5500).fadeOut(1200);
# TODO: journal entry form


Backbone.history.start({ pushState: true })
Thyself.router = new ThyselfRouter();

