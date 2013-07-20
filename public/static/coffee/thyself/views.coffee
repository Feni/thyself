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
    $(@el).bind('change', @save);    
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
        <td class="fixed-width-2 column"><button>Delete</button></td>
      """);
    return @

class Thyself.Views.EntryEditView extends Backbone.View
  el: $("#journal_entry")
  initialize: () =>
    $(@el).unbind(); # Remove attachments to previous renderings
    $(@el).bind('change', @save);
  save: () =>
    newAction = $.trim($(@el).find(".editAction").text())
    newDescription = $.trim($(@el).find(".editDescription").text())
    if newAction != ""
      @model.set("metric", newAction)
    else
      $(@el).find(".editAction").text(@model.get("metric"))
      alert("Action cannot be empty")
    if newDescription != @model.get("description")
      @model.set("description", newDescription)
    Thyself.Page.sidebarView.render()
  render: () =>
    urlDate = "/u/#{@model.get('user_id')}" +   # define url from base. else it will append on exiting page url
      "/#{@model.get('time').getFullYear()}"+
      "/#{@model.get('time').getMonth() + 1}" +
      "/#{@model.get('time').getDate()}"

    $(@el).html("""
      <a href="#{urlDate}"> <h4 class="date">#{@model.get('time').toDateString()}</h4></a>
        <input type="text" class="editAction" placeholder="Action" maxlength="32" value='#{@model.get("metric")}'/>
        <input type="text" class="editDescription" placeholder="Description" maxlength="160" value='#{@model.get("description")}'/>
      <p class="time">#{@model.get("time").toTimeString()}</p>
      </hr>
    """);
    detailListElem = $("""<table class='width-full'>
        <thead>
        <tr>
        <th class="fixed-width-3 column">Amount</th>
        <th class="fixed-width-3 column">Type</th>
        <th class="fixed-width-3 column">Group</th>
          </tr>
          </thead>
      </table>""")
    _(@model.get("details").models).each((detail) ->
      detailView = new DetailEditView({ model: detail });
      detailListElem.append(detailView.render().el);
    , @);
    $(@el).append(detailListElem)
    $(@el).append("<button>Add Details</button>")
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
  initialize: (user, year, month, day) ->
    @user = user
    @year = year
    @month = month
    @day = day
    @render()
  render: () ->
    $(@el).html("Journal Entry for ")


