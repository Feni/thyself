class Thyself.Models.Detail extends Backbone.Model
  defaults:
    amount: "",
    type: ""
  validate: (attrs, options) ->
    if attrs.type == ""
      return "Detail's type field cannot be empty"

class Thyself.Models.Details extends Backbone.Collection
  model: Thyself.Models.Detail

# TODO: implement the built-in 'validate' function
class Thyself.Models.Entry extends Backbone.Model
  defaults:
    id: "",
    user_id: "",
    description: "",
    time: 0, # May cause bugs
    metric: "",
    details: new Thyself.Models.Details()
  validate: (attrs, options) -> 
    # metric can only be null if we haven't parsed it yet
    if attrs.id != "" && attrs.metric == ""
      return "Metric can't be null"
    #if attrs.time  is not a number
  timeObj: () ->  # Calling this many times may be inefficient but whatever
    if @get('time') == 0 
      return new Date()
    else  
      return new Date(@get('time') * 1000)
  urlRoot: '/api/v0/entries'
  pageUrl: () ->
    # remove special chars and create a url our of description
    # g for global match. Else stop after first find
    cleanDesc = @get("description").replace(
     new RegExp("\\s","g"),"-").replace( new RegExp("[^A-Za-z0-9_-]", "g"), "-").slice(0, 80)    
    tempTimeObj = @timeObj()
    modelUrl = "/u/#{@get('user_id')}" +   # define url from base. else it will append on exiting page url
      "/#{tempTimeObj.getFullYear()}"+
      "/#{tempTimeObj.getMonth() + 1}" +
      "/#{tempTimeObj.getDate()}" +
      "/m/#{$.trim(@get('metric'))}" +
      "/e/#{@get('id')}"+
      "/#{$.trim(cleanDesc)}"
    return modelUrl
  dateUrl: () ->
    timeObj = @timeObj()
    urlDate = "/u/#{@get('user_id')}" +   # define url from base. else it will append on exiting page url
      "/#{timeObj.getFullYear()}"+
      "/#{timeObj.getMonth() + 1}" +
      "/#{timeObj.getDate()}"
    if @get('user_id') == "demo"
      urlDate = "/i/demo"
    return urlDate

class Thyself.Models.Entries extends Backbone.Collection
  model: Thyself.Models.Entry
  url: "/api/v0/entries"
  # Create a tree grouping toghether thigns in similar
  # categories > then types > then date
#  groupData: () =>
#    groups = {}
#    for line in @models
#      myGroup = undefined
#      if line.hasOwnProperty("group")
#        myGroup = line.group
#      if !(group.hasOwnProperty(myGroup))
#        group[myGroup] = {}
#      myType = undefined
#      if line.hasOwnProperty("type")
#        myType = line.type
#      if !(group[myGroup].hasOwnProperty(myType))
#        group[myGroup][myType] = []
#      group[myGroup][myType].push(line)
#    return group

class Thyself.Models.JournalEntry extends Backbone.Model
  defaults:
    user_id: "",
    text: "",
    time: 0 # May cause bugs
  id: () =>
    timeObj = new Date(time * 1000)
    return timeObj.getYear() + '/' + timeObj.getMonth() + "/" + timeObj.getDay()
    
  urlRoot: '/api/v0/journal'

  timeObj: () ->  # Calling this many times may be inefficient but whatever
    if @get('time') == 0 
      return new Date()
    else  
      return new Date(@get('time') * 1000)
  

class Thyself.Collections.JournalEntries extends Backbone.Collection
  model: Thyself.Models.JournalEntry
  url: "/api/v0/journal"