function parseQuery() {
    let queryString = window.location.href;
    let query = {};
    let pairs = (queryString[0] === '?' ? queryString.substr(1) : queryString).split('&');
    for (let i = 0; i < pairs.length; i++) {
        let pair = pairs[i].split('=');
        query[decodeURIComponent(pair[0])] = decodeURIComponent(pair[1] || '');
    }
    if (query["access_token"]) {
        return {
            access_token : query["access_token"],
            expires_in: query["expires_in"],
            token_type: query["token_type"]
        }
    } else {
        return null;
    }
}

export function verifyToken() {
    let access = parseQuery();
    if (access) {

        const currentTime = new Date()
        let expiry_time = currentTime.getTime() + access.expires_in;

        window.localStorage.setItem("access_token", access.access_token);
        window.localStorage.setItem("expires_in", access.expires_in);
        window.localStorage.setItem("expiry_time", expiry_time);
        window.localStorage.setItem("token_type", access.token_type);
        window.location.replace("/");
    }
}

export function clean(){
    window.localStorage.removeItem('access_token')
    window.localStorage.removeItem('expires_in')
    window.localStorage.removeItem('expiry_time')
    window.localStorage.removeItem('token_type')
}