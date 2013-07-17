#Note that the @ makes it global



@thyModule = angular.module('thyModule', ["restangular"]);
# Jinja and angular template tags conflict. So set angualr to use [[]]
@thyModule.config( ($interpolateProvider) ->
  $interpolateProvider.startSymbol('[[');
  $interpolateProvider.endSymbol(']]');
);

@thyModule.config( (RestangularProvider) -> 
    RestangularProvider.setBaseUrl('/api/v0')
  )


@thyModule.controller('ThyselfCtrl', ThyselfCtrl = ($scope, $http, Restangular) -> 
  baseMetrics = Restangular.all('metrics');
  $scope.metrics = baseMetrics.getList({since: 1371194419, until: 1471498419})
  $scope._ = _
  $scope.journal_page = "JOURNAL"
  $scope.selected_metric = {}
  $("#mEntryForm").submit( () -> 
    baseMetrics.post("", {description: $("#metricEntry").val(), time: Math.round(new Date().getTime() / 1000)}).then((response) -> 
      $scope.metrics.push(response)
      #$scope.$apply();     # Force a redraw
    , (err) -> 
      alert("There was an error saving " + err);
    );
    $("#metricEntry").val("")
    return false;
  );
  $scope.getMetricByID = (metricID) ->
    _.find($scope.metrics, (metricInstance) -> metric.ID == metricID );
  $scope.parseDate = (dateText) ->
    newDate = Date.parse(dateText)
    if $scope.selected_metric
      $scope.selected_metric.UnixTime = newDate.getTime()
    return newDate

  $scope.encodeAsURL = (actionText) ->
    # IF sluggified length > 80, shrink it to the first 80 chars
    return actionText.slugify()


  # TODO: keep the original till we get a response
  # TODO: Also make other versions where you can update a specific KEY or field
  $scope.postMetricUpdates = (mIndex) ->
    $http({
      "url": "api/0/metrics/" + mIndex.id, 
      "method": "POST",
      "data": mIndex,
      'Content-Type' : 'application/x-www-form-urlencoded; charset=UTF-8'
     }).success((data, status, headers, config) -> 
      #$scope.data = data;     # It should echo the data back...
      $scope.metrics[mIndex.id] = data
      # TODO: This may have a bug of if mIndex is inaccessible from this scope
    ).error((data, status, headers, config) -> 
      $scope.status = status;
      alert("Error sending " + status + " : " + data);
    );
    return mIndex


  $scope.deleteMetric = (mIndex) ->
    delProtocol = (mIndex) ->
      foundIndex = $scope.metrics.indexOf(mIndex)
      if foundIndex != -1  # Splice will properly remove the element. delete will set it to null
        $scope.metrics.splice(foundIndex, 1);
      return mIndex
    $http({
      "url": "/api/0/metrics/" + mIndex.id, 
      "method": "DELETE",
      "data": mIndex,
      'Content-Type' : 'application/x-www-form-urlencoded; charset=UTF-8'
     }).success((data, status, headers, config) -> 
      delProtocol(mIndex)
    ).error((data, status, headers, config) -> 
      $scope.status = status;
      alert("Error sending " + status + " : " + data);
    );


  $scope.addDetails = (mIndex) ->
    mIndex.details["category or units"] = "item/qty"
    $scope.postMetricUpdates(mIndex)

  $scope.replaceDetail = (mIndex, oldDetail, newDetail) ->
    if oldDetail != newDetail
      if $.trim(newDetail)                 # Delete the metric if it's empty
        mIndex.details[newDetail] = mIndex.details[oldDetail]
      delete mIndex.details[oldDetail]
      $scope.postMetricUpdates(mIndex)
    return mIndex.details

);  # End metric list ctrl definition

