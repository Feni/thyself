// Generated by CoffeeScript 1.4.0

/*
Copyright Feni Varughese
Thyself.io Version 0.06.00
Last Updated: August 3, 2013 - 2:35
*/


(function() {
  var DetailEditView, DetailSummaryView, DetailsListEditView, EntrySummaryView, ThyselfRouter,
    __hasProp = {}.hasOwnProperty,
    __extends = function(child, parent) { for (var key in parent) { if (__hasProp.call(parent, key)) child[key] = parent[key]; } function ctor() { this.constructor = child; } ctor.prototype = parent.prototype; child.prototype = new ctor(); child.__super__ = parent.prototype; return child; },
    __bind = function(fn, me){ return function(){ return fn.apply(me, arguments); }; };

  Thyself.Models.Detail = (function(_super) {

    __extends(Detail, _super);

    function Detail() {
      return Detail.__super__.constructor.apply(this, arguments);
    }

    Detail.prototype.defaults = {
      amount: "",
      type: ""
    };

    Detail.prototype.validate = function(attrs, options) {
      if (attrs.type === "") {
        return "Detail's type field cannot be empty";
      }
    };

    return Detail;

  })(Backbone.Model);

  Thyself.Models.Details = (function(_super) {

    __extends(Details, _super);

    function Details() {
      return Details.__super__.constructor.apply(this, arguments);
    }

    Details.prototype.model = Thyself.Models.Detail;

    return Details;

  })(Backbone.Collection);

  Thyself.Models.Entry = (function(_super) {

    __extends(Entry, _super);

    function Entry() {
      return Entry.__super__.constructor.apply(this, arguments);
    }

    Entry.prototype.defaults = {
      id: "",
      user_id: "",
      description: "",
      time: 0,
      metric: "",
      details: new Thyself.Models.Details()
    };

    Entry.prototype.validate = function(attrs, options) {
      if (attrs.id !== "" && attrs.metric === "") {
        return "Metric can't be null";
      }
    };

    Entry.prototype.timeObj = function() {
      if (this.get('time') === 0) {
        return new Date();
      } else {
        return new Date(this.get('time') * 1000);
      }
    };

    Entry.prototype.urlRoot = '/api/v0/entries';

    Entry.prototype.url = function() {
      if (this.get("user_id") === "demo") {
        return '/i/demo/m';
      } else {
        return this.urlRoot + "/" + this.get("id");
      }
    };

    Entry.prototype.pageUrl = function() {
      var cleanDesc, modelUrl, tempTimeObj;
      cleanDesc = this.get("description").replace(new RegExp("\\s", "g"), "-").replace(new RegExp("[^A-Za-z0-9_-]", "g"), "-").slice(0, 80);
      tempTimeObj = this.timeObj();
      modelUrl = ("/u/" + (this.get('user_id'))) + ("/" + (tempTimeObj.getFullYear())) + ("/" + (tempTimeObj.getMonth() + 1)) + ("/" + (tempTimeObj.getDate())) + ("/m/" + ($.trim(this.get('metric')))) + ("/e/" + (this.get('id'))) + ("/" + ($.trim(cleanDesc)));
      return modelUrl;
    };

    Entry.prototype.dateUrl = function() {
      var timeObj, urlDate;
      timeObj = this.timeObj();
      urlDate = ("/u/" + (this.get('user_id'))) + ("/" + (timeObj.getFullYear())) + ("/" + (timeObj.getMonth() + 1)) + ("/" + (timeObj.getDate()));
      if (this.get('user_id') === "demo") {
        urlDate = "/i/demo";
      }
      return urlDate;
    };

    return Entry;

  })(Backbone.Model);

  Thyself.Models.Entries = (function(_super) {

    __extends(Entries, _super);

    function Entries() {
      return Entries.__super__.constructor.apply(this, arguments);
    }

    Entries.prototype.model = Thyself.Models.Entry;

    Entries.prototype.url = "/api/v0/entries";

    Entries.prototype.comparator = function(e) {
      return -1 * e.get("time");
    };

    return Entries;

  })(Backbone.Collection);

  Thyself.Models.JournalEntry = (function(_super) {

    __extends(JournalEntry, _super);

    function JournalEntry() {
      this.id = __bind(this.id, this);
      return JournalEntry.__super__.constructor.apply(this, arguments);
    }

    JournalEntry.prototype.defaults = {
      user_id: "",
      text: "",
      time: 0
    };

    JournalEntry.prototype.id = function() {
      var timeObj;
      timeObj = new Date(time * 1000);
      return timeObj.getYear() + '/' + timeObj.getMonth() + "/" + timeObj.getDay();
    };

    JournalEntry.prototype.urlRoot = '/api/v0/journal';

    JournalEntry.prototype.timeObj = function() {
      if (this.get('time') === 0) {
        return new Date();
      } else {
        return new Date(this.get('time') * 1000);
      }
    };

    return JournalEntry;

  })(Backbone.Model);

  Thyself.Collections.JournalEntries = (function(_super) {

    __extends(JournalEntries, _super);

    function JournalEntries() {
      return JournalEntries.__super__.constructor.apply(this, arguments);
    }

    JournalEntries.prototype.model = Thyself.Models.JournalEntry;

    JournalEntries.prototype.url = "/api/v0/journal";

    return JournalEntries;

  })(Backbone.Collection);

  DetailSummaryView = (function(_super) {

    __extends(DetailSummaryView, _super);

    function DetailSummaryView() {
      this.initialize = __bind(this.initialize, this);
      return DetailSummaryView.__super__.constructor.apply(this, arguments);
    }

    DetailSummaryView.prototype.tagName = "li";

    DetailSummaryView.prototype.initialize = function(args) {
      return _.bindAll(this);
    };

    DetailSummaryView.prototype.render = function() {
      if (this.model.get("amount")) {
        $(this.el).html("<h5 class=\"val\">" + (this.model.get("amount")) + "</h5>\n<h5 class=\"key\">" + (this.model.get("type")) + "</h5>");
      } else {
        $(this.el).addClass("loneval");
        $(this.el).html("<h5 class=\"val\">" + (this.model.get("type")) + "</h5>\n<h5 class=\"key\">&nbsp</h5>");
      }
      return this;
    };

    return DetailSummaryView;

  })(Backbone.View);

  EntrySummaryView = (function(_super) {

    __extends(EntrySummaryView, _super);

    function EntrySummaryView() {
      this.initialize = __bind(this.initialize, this);
      return EntrySummaryView.__super__.constructor.apply(this, arguments);
    }

    EntrySummaryView.prototype.tagName = "li";

    EntrySummaryView.prototype.initialize = function(args) {
      return _.bindAll(this);
    };

    EntrySummaryView.prototype.render = function() {
      var detailListElem, editEntryLink;
      editEntryLink = $("<a href=\"" + (this.model.pageUrl()) + "\"></a>");
      $(editEntryLink).append("<h2 class='fixed-width-3 column actionHeading'>" + (this.model.get("metric")) + "</h2>");
      detailListElem = $("<ul class='fixed-width-8 column actionDetails'></ul>");
      _(this.model.get("details").models).each(function(detail) {
        var detailView;
        detailView = new DetailSummaryView({
          model: detail
        });
        return detailListElem.append(detailView.render().el);
      }, this);
      $(editEntryLink).append(detailListElem);
      $(this.el).html(editEntryLink);
      $(this.el).addClass("actionInstance");
      return this;
    };

    return EntrySummaryView;

  })(Backbone.View);

  Thyself.Views.EntrySummaryListView = (function(_super) {

    __extends(EntrySummaryListView, _super);

    function EntrySummaryListView() {
      this.initialize = __bind(this.initialize, this);
      return EntrySummaryListView.__super__.constructor.apply(this, arguments);
    }

    EntrySummaryListView.prototype.initialize = function(args) {
      _.bindAll(this);
      return this.render();
    };

    EntrySummaryListView.prototype.render = function() {
      var item, _i, _len, _ref;
      $(this.el).html("");
      _ref = this.collection.models;
      for (_i = 0, _len = _ref.length; _i < _len; _i++) {
        item = _ref[_i];
        $(this.el).append(new EntrySummaryView({
          model: item
        }).render().el);
      }
      return this;
    };

    return EntrySummaryListView;

  })(Backbone.View);

  DetailEditView = (function(_super) {

    __extends(DetailEditView, _super);

    function DetailEditView() {
      this.render = __bind(this.render, this);

      this.typeEdit = __bind(this.typeEdit, this);

      this.deleteDetail = __bind(this.deleteDetail, this);

      this.saveDetail = __bind(this.saveDetail, this);

      this.initialize = __bind(this.initialize, this);
      return DetailEditView.__super__.constructor.apply(this, arguments);
    }

    DetailEditView.prototype.tagName = "tr";

    DetailEditView.prototype.events = {
      "click .deleteDetailBtn": "deleteDetail",
      "change .detailType": "typeEdit"
    };

    DetailEditView.prototype.initialize = function() {
      $(this.el).unbind();
      return $(this.el).bind('save', this.saveDetail);
    };

    DetailEditView.prototype.saveDetail = function() {
      var newAmount, newType;
      newAmount = $.trim($(this.el).find(".detailAmount").val());
      newType = $.trim($(this.el).find(".detailType").val());
      if (newAmount !== this.model.get("amount")) {
        this.model.set("amount", newAmount);
      }
      if (newType !== this.model.get("type")) {
        this.model.set("type", newType);
        if (newType === "") {
          return deleteDetail();
        }
      }
    };

    DetailEditView.prototype.deleteDetail = function() {
      this.model.destroy();
      return this.remove();
    };

    DetailEditView.prototype.typeEdit = function() {
      var newType;
      newType = $.trim($(this.el).find(".detailType").val());
      if (newType === "") {
        return this.deleteDetail();
      }
    };

    DetailEditView.prototype.render = function() {
      $(this.el).addClass("detailRow");
      $(this.el).html("<td class=\"fixed-width-4 column\"><input type=\"number\" class=\"detailAmount fullInput\" maxlength=\"32\" value='" + (this.model.get("amount")) + "'/></td>\n<td class=\"fixed-width-4 column\"><input type=\"text\" class=\"detailType fullInput\" maxlength=\"120\" value='" + (this.model.get("type")) + "'/></td>\n<td class='fixed-width-2 tblBtnCol column deleteDetailBtn'><button>Delete</button></td>          ");
      return this;
    };

    return DetailEditView;

  })(Backbone.View);

  DetailsListEditView = (function(_super) {

    __extends(DetailsListEditView, _super);

    function DetailsListEditView() {
      this.render = __bind(this.render, this);

      this.tempDetails = __bind(this.tempDetails, this);

      this.addDetailsTypeChanged = __bind(this.addDetailsTypeChanged, this);

      this.initialize = __bind(this.initialize, this);
      return DetailsListEditView.__super__.constructor.apply(this, arguments);
    }

    DetailsListEditView.prototype.tagName = "table";

    DetailsListEditView.prototype.initialize = function() {
      return $(this.el).unbind();
    };

    DetailsListEditView.prototype.addDetailsTypeChanged = function() {
      var tempRow;
      tempRow = $(this.el).find("#tempRow");
      this.collection.add(new Thyself.Models.Detail({
        amount: "" + tempRow.find(".detailAmount").val(),
        type: tempRow.find(".detailType").val()
      }));
      return this.render();
    };

    DetailsListEditView.prototype.tempDetails = function() {
      var tempRow, tempTypeField;
      tempRow = $("<tr id='tempRow'>");
      tempRow.append("<td class=\"fixed-width-4 pad-1 column\"><input type=\"number\" placeholder=\"Quantity\" class=\"detailAmount fullInput\" maxlength=\"32\" value='" + "'/></td>");
      tempTypeField = $("<td class=\"fixed-width-4 pad-1 column\"><input type=\"text\" placeholder=\"Units/Type\" class=\"detailType fullInput\" maxlength=\"120\" value='" + "'/></td>");
      tempTypeField.bind('change', this.addDetailsTypeChanged);
      return tempRow.append(tempTypeField);
    };

    DetailsListEditView.prototype.render = function() {
      $(this.el).html("");
      $(this.el).addClass("width-full");
      $(this.el).addClass("dataSummaryTable");
      $(this.el).append("<thead>\n  <tr>\n    <th class=\"fixed-width-4 column\">Amount</th>\n    <th class=\"fixed-width-4 column\">Type</th>\n  </tr>\n</thead>");
      _(this.collection.models).each(function(detail) {
        var detailView;
        detailView = new DetailEditView({
          model: detail,
          collection: this.collection
        });
        return $(this.el).append(detailView.render().el);
      }, this);
      $(this.el).append(this.tempDetails());
      return this.el;
    };

    return DetailsListEditView;

  })(Backbone.View);

  Thyself.Views.EntryEditView = (function(_super) {

    __extends(EntryEditView, _super);

    function EntryEditView() {
      this.unrender = __bind(this.unrender, this);

      this.render = __bind(this.render, this);

      this.deleteEntry = __bind(this.deleteEntry, this);

      this.saveEntry = __bind(this.saveEntry, this);

      this.initialize = __bind(this.initialize, this);
      return EntryEditView.__super__.constructor.apply(this, arguments);
    }

    EntryEditView.prototype.el = $("#journal_entry");

    EntryEditView.prototype.events = {
      "click .entrySaveBtn": "saveEntry",
      "click .entryDelteBtn": "deleteEntry"
    };

    EntryEditView.prototype.initialize = function() {
      return $(this.el).unbind();
    };

    EntryEditView.prototype.saveEntry = function() {
      var newAction, newDescription,
        _this = this;
      newAction = $.trim($(this.el).find(".editAction").val());
      newDescription = $.trim($(this.el).find(".editDescription").val());
      $(this.el).find(".detailRow").trigger("save");
      this.model.save({
        id: this.model.get('id')
      }, {
        success: function(response) {
          var detailsCollection, newMessage;
          detailsCollection = new Thyself.Models.Details(_this.model.get("details"));
          _this.model.set("details", detailsCollection);
          newMessage = $("<li class='alert-box alert'>Entry saved successfully</li>");
          $(".message_flashes").append(newMessage);
          return newMessage.delay(3500).fadeOut(1200);
        },
        error: function(entry, response) {
          var newMessage;
          newMessage = $("<li class='alert-box alert'>Error saving entry: " + response + "</li>");
          $(".message_flashes").append(newMessage);
          return newMessage.delay(3500).fadeOut(1200);
        }
      });
      $(this.el).find(".editAction").val(this.model.get("metric"));
      return Thyself.Page.sidebarView.render();
    };

    EntryEditView.prototype.deleteEntry = function() {
      Thyself.router.navigate(this.model.dateUrl(), {
        trigger: true
      });
      this.model.destroy();
      return Thyself.Page.sidebarView.render();
    };

    EntryEditView.prototype.render = function() {
      var deleteButton, entryControlsDiv, saveButton, timeObj,
        _this = this;
      timeObj = this.model.timeObj();
      $(this.el).html("<a href=\"" + (this.model.dateUrl()) + "\"> <h4 class=\"date\">" + (timeObj.toDateString()) + "</h4></a>\n  <input type=\"text\" class=\"editAction\" placeholder=\"Action\" maxlength=\"32\" value='" + (this.model.get("metric")) + "'/>\n  <input type=\"text\" class=\"fullInput editDescription\" placeholder=\"Description\" maxlength=\"160\" value='" + (this.model.get("description")) + "'/>\n<p class=\"time\">" + (timeObj.toTimeString()) + "</p>\n</hr>");
      $(this.el).append(new DetailsListEditView({
        collection: this.model.get("details")
      }).render());
      entryControlsDiv = $("<div class='entryControls'>");
      deleteButton = $("<button class='flatButton pad-1 entryDelteBtn'>Delete</button>");
      saveButton = $("<button class='flatButton pad-1 entrySaveBtn'>Save</button>");
      $(deleteButton).bind('click', function() {});
      $(entryControlsDiv).append(deleteButton);
      $(entryControlsDiv).append(saveButton);
      $(this.el).append(entryControlsDiv);
      return this;
    };

    EntryEditView.prototype.unrender = function() {
      return $(this.el).remove();
    };

    return EntryEditView;

  })(Backbone.View);

  Thyself.Views.IndexView = (function(_super) {

    __extends(IndexView, _super);

    function IndexView() {
      return IndexView.__super__.constructor.apply(this, arguments);
    }

    IndexView.prototype.el = $("#journal_entry");

    IndexView.prototype.initialize = function() {
      return this.render();
    };

    IndexView.prototype.render = function() {
      return this;
    };

    return IndexView;

  })(Backbone.View);

  Thyself.Views.SettingsView = (function(_super) {

    __extends(SettingsView, _super);

    function SettingsView() {
      return SettingsView.__super__.constructor.apply(this, arguments);
    }

    SettingsView.prototype.el = $("#journal_entry");

    SettingsView.prototype.render = function() {
      $(this.el).html("Upgrade to Premium");
      return this;
    };

    return SettingsView;

  })(Backbone.View);

  Thyself.Views.JournalView = (function(_super) {

    __extends(JournalView, _super);

    function JournalView() {
      return JournalView.__super__.constructor.apply(this, arguments);
    }

    JournalView.prototype.el = $("#journal_entry");

    JournalView.prototype.render = function() {
      var timeObj, urlDate;
      if (this.model) {
        timeObj = this.model.timeObj();
        urlDate = ("/u/" + (this.model.get('user_id'))) + ("/" + (timeObj.getFullYear())) + ("/" + (timeObj.getMonth() + 1)) + ("/" + (timeObj.getDate()));
        if (this.model.get('user_id') === "demo") {
          urlDate = "/i/demo";
        }
        return $(this.el).html("<a href=\"" + urlDate + "\"> <h4 class=\"date\">" + (timeObj.toDateString()) + "</h4></a>\n<form id=\"journalEntryForm\" action=\"" + urlDate + "\" method=\"POST\">\n   <textarea id=\"journalEntryText\" placeholder=\"How was your day?\" name=\"text\" maxlength=\"4000\">" + (this.model.get("text")) + "</textarea>\n  <input type=\"submit\" class=\"flatButton\" id=\"loginButton\" value=\"Save\" onClick=\"$('#journalEntryForm').submit(); return false;\"/>\n</form>");
      } else {
        return $(this.el).html("");
      }
    };

    return JournalView;

  })(Backbone.View);

  ThyselfRouter = (function(_super) {

    __extends(ThyselfRouter, _super);

    function ThyselfRouter() {
      this.demoMain = __bind(this.demoMain, this);

      this.journal = __bind(this.journal, this);

      this.settings = __bind(this.settings, this);

      this.index = __bind(this.index, this);
      return ThyselfRouter.__super__.constructor.apply(this, arguments);
    }

    ThyselfRouter.prototype.routes = {
      "": "index",
      "u": "settings",
      "i/demo": "demoMain",
      "u/:user/:year/:month/:day": "journal",
      "u/:user/:year/:month/:day/m/:metric_name/e/:entry_id/:entry_desc": "entrySummary"
    };

    ThyselfRouter.prototype.index = function() {
      var indexView;
      return indexView = new Thyself.Views.IndexView();
    };

    ThyselfRouter.prototype.settings = function() {
      var settingsView;
      settingsView = new Thyself.Views.SettingsView;
      return settingsView.render();
    };

    ThyselfRouter.prototype.journal = function(user, year, month, day) {
      var journalView;
      journalView = new Thyself.Views.JournalView();
      return journalView.render();
    };

    ThyselfRouter.prototype.entrySummary = function(user, year, month, day, metric_name, entry_id, entry_desc) {
      var entry, entryView;
      entry = Thyself.Data.Entries.get(entry_id);
      entryView = new Thyself.Views.EntryEditView({
        model: entry
      });
      return entryView.render();
    };

    ThyselfRouter.prototype.demoMain = function() {
      var demoEntry, journalView;
      demoEntry = new Thyself.Models.JournalEntry({
        user_id: "demo",
        text: ""
      });
      journalView = new Thyself.Views.JournalView({
        model: demoEntry
      });
      return journalView.render();
    };

    return ThyselfRouter;

  })(Backbone.Router);

  $(document).delegate("a", "click", function(event) {
    var href;
    href = $(this).attr("href");
    if (!event.altKey && !event.ctrlKey && !event.metaKey && !event.shiftKey) {
      if (href.substring(0, 1) === '/' && href.substring(0, 3) !== "/a/" && href !== "/") {
        Thyself.router.navigate(href, {
          trigger: true
        });
        if ($("#journal_entry").html() === "") {
          return true;
        } else {
          event.preventDefault();
          return false;
        }
      }
    }
  });

  $("#mEntryForm").submit(function() {
    var actionUrl, descriptionField, entryFields, mergedTime, newEntry, timeNow,
      _this = this;
    actionUrl = $(this).attr('action');
    newEntry = new Thyself.Models.Entry();
    if (actionUrl === '/i/demo/m') {
      newEntry.url = '/i/demo/m';
    }
    descriptionField = $(this).find("#description");
    timeNow = new Date();
    mergedTime = new Date(Thyself.Data.ContextDate.getFullYear(), Thyself.Data.ContextDate.getMonth(), Thyself.Data.ContextDate.getDate(), timeNow.getHours(), timeNow.getMinutes(), timeNow.getSeconds(), timeNow.getMilliseconds());
    entryFields = {
      description: descriptionField.val(),
      time: Math.round(mergedTime.getTime() / 1000)
    };
    newEntry.save(entryFields, {
      success: function(entry) {
        var detailsCollection;
        console.log(entry.toJSON());
        detailsCollection = new Thyself.Models.Details(entry.get("details"));
        entry.set("details", detailsCollection);
        Thyself.Page.sidebarView.collection.add(newEntry);
        Thyself.Page.sidebarView.collection.sort();
        return Thyself.Page.sidebarView.render();
      },
      error: function(model, response) {
        console.log(model);
        return console.log(response);
      }
    });
    descriptionField.val("");
    return false;
  });

  $(".alert-box").delay(5500).fadeOut(1200);

  Backbone.history.start({
    pushState: true
  });

  Thyself.router = new ThyselfRouter();

}).call(this);
