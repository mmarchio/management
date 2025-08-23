async function apiGet(uri) {
    return await fetch(uri)
        .then(function(response) {
            if (!response.ok) {
                throw new Error(`HTTP error status: ${response.status}`);
            }
            return response.json();
        })
        .then(function(data) {
            console.log(data);
            return data;
        })
        .catch(error => {
            console.error(`error: ${error}`)
        });
}

function xapiGet(method, uri, body) {
    var xReq = new XMLHttpRequest()
    xReq.onreadystatechange = function(){
        if (this.readyState == 4 && this.status >= 200) {
            let resp = this.responseText
            let json = this.JSON.parse(resp)
            console.log(json)
            return json
        } else {
            console.log(this.responseType)
        }
    }
    xReq.open(method, uri, true)
    xReq.send()
}
