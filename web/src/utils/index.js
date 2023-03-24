export function getCookie(cname) {
    let name = cname + "=";
    let decodedCookie = decodeURIComponent(document.cookie);
    let ca = decodedCookie.split(';');
    for (let i = 0; i < ca.length; i++) {
        let c = ca[i];
        while (c.charAt(0) == ' ') {
            c = c.substring(1);
        }
        if (c.indexOf(name) == 0) {
            return c.substring(name.length, c.length);
        }
    }
    return "";
}

export function setCookie(cname, cvalue, days) {
    let date = new Date();
    date.setDate(date.getDate() + days);
    let value = cvalue + ((days == null) ? "" : "; expires=" + date.toUTCString());
    document.cookie = cname + "=" + value;
}

export function delUser() {
    setCookie('loginUser', '', -1)
}

export function getUser() {
    let obj = getCookie('loginUser');
    if (obj) {
        return JSON.parse(obj);
    }
}
