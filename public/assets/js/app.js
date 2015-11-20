var goShareApp = angular.module('goShareApp', ['ngRoute']);

goShareApp.config(function($routeProvider) {
    $routeProvider
        .when('/', {
            templateUrl: 'pages/form_upload.html',
            controller: 'uploadController'
        })
        .when('/download', {
            templateUrl: 'pages/form_download.html',
            controller: 'downloadController'
        })
        .when('/info', {
            templateUrl: 'pages/form_info.html',
            controller: 'infoController'
        })
        .when('/contacts', {
            templateUrl: 'pages/form_contacts.html',
            controller: 'contactsController'
        })
        .otherwise({
            redirectTo: '/'
        })
});

goShareApp.controller('uploadController', function($scope, $http) {
    $scope.uploadFile = function(files) {
        var formData = new FormData();
        formData.append("file", files[0]);
        $http.post("/api/uploadFile", formData, {
            withCredentials: true,
            headers: {'Content-Type': undefined },
            transformRequest: angular.identity
        })
        .success(function(response) {
            location.href = "#download?code=" + response
        })
        .error(function() {
            // TODO
        });
    };
});

goShareApp.controller('downloadController', function($scope) {
    $scope.message = 'Download';
});

goShareApp.controller('infoController', function($scope) {
    $scope.message = 'Info';
});

goShareApp.controller('contactsController', function($scope) {
    $scope.message = 'Contacts';
});
