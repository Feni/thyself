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
      $(@el).html """
        <h5 class="val">#{@model.get("type")}</h5>
        <h5 class="key">#{@model.get("group")}</h5>
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
  tagName: "tr"
  initialize: () =>
    $(@el).unbind(); # Remove attachments to previous renderings
    #$(@el).bind('change', @save);    
  save: () =>
    newAmount = $.trim($(@el).find(".detailAmount").val())
    newType = $.trim($(@el).find(".detailType").val())
    newGroup = $.trim($(@el).find(".detailGroup").val())
    if newAmount != @model.get("amount")
      @model.set("amount", newAmount)
    if newType != @model.get("type")
      @model.set("type", newType)
    if newGroup != @model.get("group")
      @model.set("group", newGroup)

  render: () =>
    $(@el).html("""
        <td class="fixed-width-3 column"><input type="text" class="detailAmount fullInput" maxlength="32" value='#{@model.get("amount")}'/></td>
        <td class="fixed-width-3 column"><input type="text" class="detailType fullInput" maxlength="120" value='#{@model.get("type")}'/></td>
        <td class="fixed-width-3 column"><input type="text" class="detailGroup fullInput" maxlength="32" value='#{@model.get("group")}'/></td>
        <td class="fixed-width-2 tblBtnCol column"><button class="">Delete</button></td>
      """);
    return @

class DetailsListEditView extends Backbone.View
  tagName: "table"
  addDetailsTypeChanged: () =>
    #alert("Type field changed. Adding extra detail")
    tempRow = $(@el).find("#tempRow")
    @collection.add(new Thyself.Models.Detail({
      amount: "" + tempRow.find(".detailAmount").val()
      type: tempRow.find(".detailType").val()
      group: tempRow.find(".detailGroup").val()
      }))
    @render()
    #$(@el).append(@tempDetails())
  tempDetails: () =>
    tempRow = $("<tr id='tempRow'>")
    tempRow.append("""<td class="fixed-width-3 column"><input type="number" placeholder="Quantity" class="detailAmount fullInput" maxlength="32" value='#{}'/></td>""")
    tempTypeField = $("""<td class="fixed-width-3 column"><input type="text" placeholder="Units/Type" class="detailType fullInput" maxlength="120" value='#{}'/></td>""")
    tempTypeField.bind('change', @addDetailsTypeChanged)
    tempRow.append(tempTypeField)
    tempRow.append("""<td class="fixed-width-3 column"><input type="text" placeholder="Type Category" class="detailGroup fullInput" maxlength="32" value='#{}'/></td>""")  
  render: () =>
    $(@el).html("")
    $(@el).addClass("width-full")
    $(@el).append("""<thead>
        <tr>
          <th class="fixed-width-3 column">Amount</th>
          <th class="fixed-width-3 column">Type</th>
          <th class="fixed-width-3 column">Group</th>
        </tr>
      </thead>
    """)
    _(@collection.models).each((detail) ->
      detailView = new DetailEditView({ model: detail });
      $(@el).append(detailView.render().el);
    , @);
    $(@el).append(@tempDetails())
    return @el

#    $(@el).append(@tempDetails())



class Thyself.Views.EntryEditView extends Backbone.View
  el: $("#journal_entry")
  initialize: () =>
    $(@el).unbind(); # Remove attachments to previous renderings
    #$(@el).bind('change', @save);
  save: () =>
    newAction = $.trim($(@el).find(".editAction").text())
    newDescription = $.trim($(@el).find(".editDescription").text())
    if newAction != ""
      @model.set("metric", newAction)
    else
      $(@el).find(".editAction").text(@model.get("metric"))
    if newDescription != @model.get("description")
      @model.set("description", newDescription)
    @model.save()
    Thyself.Page.sidebarView.render()

  render: () =>
    timeObj = @model.timeObj()
    urlDate = "/u/#{@model.get('user_id')}" +   # define url from base. else it will append on exiting page url
      "/#{timeObj.getFullYear()}"+
      "/#{timeObj.getMonth() + 1}" +
      "/#{timeObj.getDate()}"
    if @model.get('user_id') == "demo"
      urlDate = "/i/demo"

    $(@el).html("""
      <a href="#{urlDate}"> <h4 class="date">#{timeObj.toDateString()}</h4></a>
        <input type="text" class="editAction" placeholder="Action" maxlength="32" value='#{@model.get("metric")}'/>
        <!--<button class="flatButton">Delete</button>  
        <button class="flatButton">Save</button>-->
        <input type="text" class="fullInput editDescription" placeholder="Description" maxlength="160" value='#{@model.get("description")}'/>
      <p class="time">#{timeObj.toTimeString()}</p>
      </hr>
    """);
    # Details table
    $(@el).append(new DetailsListEditView({collection: @model.get("details")}).render())
    #$(@el).append("<hr>")
    entryControlsDiv = $("<div class='entryControls'>")
    deleteButton = $("<button class='flatButton pad-1'>Delete</button>")
    saveButton = $("<button class='flatButton pad-1'>Save</button>")

    $(saveButton).bind('click', @save);
    
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

