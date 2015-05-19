function BeaconsViewModel() 
{
    var self = this;
    self.usersURI = 'getUsers';
    self.beaconsURI = 'getBeacons';
    self.users = ko.observableArray();

    self.beginAdd = function() {
        alert("Add");
    }
    self.beginEdit = function(beacon) {
        alert("Edit: " + beacon.id());
    }
    self.remove = function(beacon) {
        alert("Remove: " + beacon.id());
    }

    ajax(self.beaconsURI, 'GET').done(function(data) {
        for (var i = 0; i < data.length; i++) {
            self.users.push({
                    // uri: ko.observable(data.users[i].uri),
                    latitude: ko.observable(data[i].latitude),
                    userId: ko.observable(data[i].userId),
                    id: ko.observable(data[i].id),
                    longitude: ko.observable(data[i].longitude),
                    time: ko.observable(data[i].time)
                });
        }
    });
}
   

