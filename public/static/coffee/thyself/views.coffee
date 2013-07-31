class DetailSummaryView extends Backbone.View
  tagName: "li"
  initialize: (args) =>
    _.bindAll(@)
  render: () ->
    if @model.get("amount")
      $(@el).html """
        <h5 class="val">#{@model.get("amount")}</h5>
        <h5 class="key">#{@model.get("type")}</h5>
      """
    else
      $(@el).addClass("loneval")
      $(@el).html """
        <h5 class="val">#{@model.get("type")}</h5>
        <h5 class="key">&nbsp</h5>
      """
    return @

class EntrySummaryView extends Backbone.View
  tagName: "li"
  initialize: (args) =>
    _.bindAll(@)
  render: () ->
    editEntryLink = $("""<a href="#{@model.pageUrl()}"></a>""")
    $(editEntryLink).append("<h2 class='fixed-width-3 column actionHeading'>#{@model.get("metric")}</h2>")    
    detailListElem = $("<ul class='fixed-width-8 column actionDetails'></ul>")
    _(@model.get("details").models).each((detail) ->
      detailView = new DetailSummaryView({ model: detail });
      detailListElem.append(detailView.render().el);
    , @);
    $(editEntryLink).append(detailListElem)
    $(@el).html(editEntryLink)
    $(@el).addClass("actionInstance")
    return @

# for the sidebar
class Thyself.Views.EntrySummaryListView extends Backbone.View
  initialize: (args) =>
    _.bindAll(@)
    @render()
  render: ->
    $(@el).html("")
    $(@el).append(new EntrySummaryView({ model: item }).render().el) for item in @collection.models;
    return @


# Edit views

class DetailEditView extends Backbone.View
  tagName: "tr",
  events: {
    "click .deleteDetailBtn" : "deleteDetail",
    "change .detailType": "typeEdit"
  }
  initialize: () =>
    $(@el).unbind(); # Remove attachments to previous renderings
    $(@el).bind('save', @saveDetail);
  saveDetail: () =>
    newAmount = $.trim($(@el).find(".detailAmount").val())
    newType = $.trim($(@el).find(".detailType").val())
    if newAmount != @model.get("amount")
      @model.set("amount", newAmount)
    if newType != @model.get("type")
      @model.set("type", newType)
      if newType == ""
        deleteDetail()
  deleteDetail: () =>
      @model.destroy() # Delete this element from all collections. 
      @remove() # Remove view from dom
  typeEdit: () =>
    newType = $.trim($(@el).find(".detailType").val())
    if newType == ""
      @deleteDetail()

  render: () =>
    $(@el).addClass("detailRow")
    $(@el).html("""
        <td class="fixed-width-4 column"><input type="number" class="detailAmount fullInput" maxlength="32" value='#{@model.get("amount")}'/></td>
        <td class="fixed-width-4 column"><input type="text" class="detailType fullInput" maxlength="120" value='#{@model.get("type")}'/></td>
        <td class='fixed-width-2 tblBtnCol column deleteDetailBtn'><button>Delete</button></td>          
      """);
    return @

class DetailsListEditView extends Backbone.View
  tagName: "table"
  initialize: () =>
    $(@el).unbind();
  addDetailsTypeChanged: () =>
    tempRow = $(@el).find("#tempRow")
    @collection.add(new Thyself.Models.Detail({
      amount: "" + tempRow.find(".detailAmount").val()
      type: tempRow.find(".detailType").val()
      }))
    @render()
  tempDetails: () =>
    tempRow = $("<tr id='tempRow'>")
    tempRow.append("""<td class="fixed-width-4 pad-1 column"><input type="number" placeholder="Quantity" class="detailAmount fullInput" maxlength="32" value='#{}'/></td>""")
    tempTypeField = $("""<td class="fixed-width-4 pad-1 column"><input type="text" placeholder="Units/Type" class="detailType fullInput" maxlength="120" value='#{}'/></td>""")
    tempTypeField.bind('change', @addDetailsTypeChanged)
    tempRow.append(tempTypeField)
  render: () =>
    $(@el).html("")
    $(@el).addClass("width-full")
    $(@el).addClass("dataSummaryTable")
    $(@el).append("""<thead>
        <tr>
          <th class="fixed-width-3 column">Amount</th>
          <th class="fixed-width-3 column">Type</th>
        </tr>
      </thead>
    """)
    _(@collection.models).each((detail) ->
      detailView = new DetailEditView({ model: detail, collection: @collection });
      $(@el).append(detailView.render().el);
    , @);
    $(@el).append(@tempDetails())
    return @el

#    $(@el).append(@tempDetails())



class Thyself.Views.EntryEditView extends Backbone.View
  el: $("#journal_entry")
  events: {
    "click .entrySaveBtn": "saveEntry",
    "click .entryDelteBtn": "deleteEntry"
  }
  initialize: () =>
    $(@el).unbind();

  saveEntry: () =>
    newAction = $.trim($(@el).find(".editAction").val())
    newDescription = $.trim($(@el).find(".editDescription").val())
    $(@el).find(".detailRow").trigger("save")
    @model.save({
        id: @model.get('id')
      }, 
      success: (response) =>
        # needa change it from an obj to a proper type. else it bugs out when sidebar re-renders
        detailsCollection = new Thyself.Models.Details(@model.get("details"))
        @model.set("details", detailsCollection)
        newMessage = $("<li class='alert-box alert'>Entry saved successfully</li>")
        $(".message_flashes").append(newMessage)
        newMessage.delay(3500).fadeOut(1200);
      error: (entry, response) =>
        newMessage = $("<li class='alert-box alert'>Error saving entry: "+response+"</li>")
        $(".message_flashes").append(newMessage)
        newMessage.delay(3500).fadeOut(1200);
    )
    $(@el).find(".editAction").val(@model.get("metric"))
    Thyself.Page.sidebarView.render()
  deleteEntry: () =>
    Thyself.router.navigate(@model.dateUrl(), { trigger: true })  # todo: fix this so new page actually renders something
    @model.destroy()
    Thyself.Page.sidebarView.render()

  render: () =>
    timeObj = @model.timeObj()
    $(@el).html("""
      <a href="#{@model.dateUrl()}"> <h4 class="date">#{timeObj.toDateString()}</h4></a>
        <input type="text" class="editAction" placeholder="Action" maxlength="32" value='#{@model.get("metric")}'/>
        <input type="text" class="fullInput editDescription" placeholder="Description" maxlength="160" value='#{@model.get("description")}'/>
      <p class="time">#{timeObj.toTimeString()}</p>
      </hr>
    """);
    # Details table
    $(@el).append(new DetailsListEditView({collection: @model.get("details")}).render())
    #$(@el).append("<hr>")
    entryControlsDiv = $("<div class='entryControls'>")
    deleteButton = $("<button class='flatButton pad-1 entryDelteBtn'>Delete</button>")
    saveButton = $("<button class='flatButton pad-1 entrySaveBtn'>Save</button>")
    $(deleteButton).bind('click', () => 
      
      #@el.html("")
      # TODO ; sidebar re-render
      # reroute journal entry to something else
    );
    
    $(entryControlsDiv).append(deleteButton)
    $(entryControlsDiv).append(saveButton)

    


    $(@el).append(entryControlsDiv)


    return @
  unrender: () =>
    $(@el).remove();


# Top level views

class Thyself.Views.IndexView extends Backbone.View
  el: $("#journal_entry")
  initialize: () ->
    @render()
  render: () ->
    return @

class Thyself.Views.SettingsView extends Backbone.View
  el: $("#journal_entry")
  render: () ->
    $(@el).html("Upgrade to Premium")
    return @

class Thyself.Views.JournalView extends Backbone.View
  el: $("#journal_entry")
#  initialize: (user, year, month, day) ->
#    @user = user
#    @year = year
#    @month = month
#    @day = day
#    @render()
  render: () ->
    if @model
      timeObj = @model.timeObj()
      urlDate = "/u/#{@model.get('user_id')}" +   # define url from base. else it will append on exiting page url
        "/#{timeObj.getFullYear()}"+
        "/#{timeObj.getMonth() + 1}" +
        "/#{timeObj.getDate()}"
      if @model.get('user_id') == "demo"
        urlDate = "/i/demo"
      $(@el).html("""<a href="#{urlDate}"> <h4 class="date">#{timeObj.toDateString()}</h4></a>
    <form id="journalEntryForm" action="#{urlDate}" method="POST">
       <textarea id="journalEntryText" placeholder="How was your day?" name="text" maxlength="4000">#{@model.get("text")}</textarea>
      <input type="submit" class="flatButton" id="loginButton" value="Save" onClick="$('#journalEntryForm').submit(); return false;"/>
    </form>
    """)
    else
      $(@el).html("")

