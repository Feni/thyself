# Hook to all link-clicks on the page so we can route it using push-state
$(document).delegate "a", "click", (event) -> 
  # Get the anchor href and protcol
  href = $(this).attr("href");
  #protocol = this.protocol + "//";  # http://
  # Ensure the protocol is not part of URL, meaning its relative.
  if !event.altKey and !event.ctrlKey and !event.metaKey and !event.shiftKey
    if href.substring(0, 1) == '/'  # only catch urls starting with /. abs urls are treated normally
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
 # alert("Sure, the form was submitted. but I'm not going to do anything. Go to : "+actionUrl);
  newEntry = new Thyself.Models.Entry(); 
  if actionUrl == '/i/demo/m' 
    newEntry.url = '/i/demo/m'

  descriptionField = $(this).find("#description")
  
  entryFields = { description: descriptionField.val(), time: Math.round(new Date().getTime() / 1000) }; 
  newEntry.save(entryFields, { success: (entry) => 
    console.log(entry.toJSON()); 
    #timeObj = new Date(entry.get("time") * 1000)
    #entry.set("time", timeObj )
    detailsCollection = new Thyself.Models.Details(entry.get("details"))
    entry.set("details", detailsCollection)
    Thyself.Page.sidebarView.collection.add(newEntry)
    Thyself.Page.sidebarView.render()
  })
  descriptionField.val("")
  return false;
);

# TODO: journal entry form



Backbone.history.start({ pushState: true })
Thyself.router = new ThyselfRouter();
