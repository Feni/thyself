class Thyself.Models.Detail extends Backbone.Model
  defaults:
    amount: "",
    type: "",
    group: ""

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

class Thyself.Models.Entries extends Backbone.Collection
  model: Thyself.Models.Entry
  url: "/api/v0/entries"
  # Create a tree grouping toghether thigns in similar
  # categories > then types > then date
  groupData: () =>
    groups = {}
    for line in @models
      myGroup = undefined
      if line.hasOwnProperty("group")
        myGroup = line.group
      if !(group.hasOwnProperty(myGroup))
        group[myGroup] = {}
      myType = undefined
      if line.hasOwnProperty("type")
        myType = line.type
      if !(group[myGroup].hasOwnProperty(myType))
        group[myGroup][myType] = []
      group[myGroup][myType].push(line)
    return group


