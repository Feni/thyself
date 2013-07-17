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
    time: new Date(), # May cause bugs
    metric: "",
    details: new Thyself.Models.Details()
  urlRoot: '/api/v0/entries'
  pageUrl: () ->
    # remove special chars and create a url our of description
    # g for global match. Else stop after first find
    cleanDesc = @get("description").replace(
     new RegExp("\\s","g"),"-").replace( new RegExp("[^A-Za-z0-9_-]", "g"), "-").slice(0, 80)    
    modelUrl = "/u/#{@get('user_id')}" +   # define url from base. else it will append on exiting page url
      "/#{@get('time').getFullYear()}"+
      "/#{@get('time').getMonth() + 1}" +
      "/#{@get('time').getDate()}" +
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


