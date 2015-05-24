ajax = function(uri, method, data) {

    var request = {
        url: uri,
        type: method,
        contentType: "application/json",
        accepts: "application/json",
        cache: false,
        dataType: 'json',
        data: JSON.stringify(data),
            // beforeSend: function (xhr) {
            //     xhr.setRequestHeader("Authorization", 
            //         "Basic " + btoa(self.username + ":" + self.password));
            // },
        error: function(jqXHR) {
                console.log("ajax error " + jqXHR.status);
        }
    };
    return $.ajax(request);
}