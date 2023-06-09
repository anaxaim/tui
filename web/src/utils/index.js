export function getCookie(cname) {
  const name = `${cname}=`;
  const decodedCookie = decodeURIComponent(document.cookie);
  const ca = decodedCookie.split(';');
  for (let i = 0; i < ca.length; i += 1) {
    let c = ca[i];
    while (c.charAt(0) === ' ') {
      c = c.substring(1);
    }
    if (c.indexOf(name) === 0) {
      return c.substring(name.length, c.length);
    }
  }
  return '';
}

export function setCookie(cname, cvalue, days) {
  const date = new Date();
  date.setDate(date.getDate() + days);
  const value = cvalue + ((days === null) ? '' : `; expires=${date.toUTCString()}`);
  document.cookie = `${cname}=${value}`;
}

export function delUser() {
  setCookie('loginUser', '', -1);
}

export function getUser() {
  const obj = getCookie('loginUser');
  if (obj) {
    return JSON.parse(obj);
  }
  return null;
}
