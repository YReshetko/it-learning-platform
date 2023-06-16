import * as cookies from './cookies.js';

function parseQuery() {
    let queryString = window.location.href;
    let query = {};
    let pairs = (queryString[0] === '?' ? queryString.substr(1) : queryString).split('&');
    for (let i = 0; i < pairs.length; i++) {
        let pair = pairs[i].split('=');
        query[decodeURIComponent(pair[0])] = decodeURIComponent(pair[1] || '');
    }
    return query;
}

export function verifyToken() {
    let query = parseQuery();
    let accessToken = query["access_token"]
    if (accessToken) {
        let expires = query["expires_in"]
        let tokenType = query["token_type"]
        cookies.setCookie("access_token", accessToken, {'max-age': expires})
        cookies.setCookie("token_type", tokenType, {'max-age': expires})

        /* document.cookie = "access_token=" + accessToken + "; max-age=" + expires;
         document.cookie = "token_type=" + tokenType + "; max-age=" + expires;*/
        window.location.href = "/";
    }
}