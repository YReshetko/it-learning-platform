import * as auth from "./auth";

export function get(url, query) {
    let request = prepareRequest("GET");
    // build full url with query
    return doRequest(url, request)
}

export function post(url, body) {
    let request = prepareRequest("POST", body);
    return doRequest(url, request)
}

function prepareRequest(method, body = null) {
    let request = {
        method: method,
        redirect: 'follow',
        headers: {
            'Accept': 'application/json'
        }
    }
    if (body) {
        request.body = JSON.stringify(body);
        request.headers['Content-Type'] = 'application/json';
    }
    let accessToken = window.localStorage.getItem("access_token")
    if (accessToken) {
        let expiryTime = window.localStorage.getItem("expiry_time")
        if (expiryTime) {
            const now = new Date()
            if (now.getTime() < expiryTime) {
                request.headers['Authorization'] = 'Bearer ' + accessToken;
            } else {
                auth.clean();
            }
        } else {
            request.headers['Authorization'] = 'Bearer ' + accessToken;
        }

    }
    return request;
}

function doRequest(url, request) {
    return fetch(url, request).then(res => {
        if (res.redirected) {
            window.location.replace(res.url);
        }
        // TODO Verify response statuses

        return res.json();
    })
}