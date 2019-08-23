import * as config from '../config.js'


export default class AuthAPI {

  static callApiWithPOST(endpoint, data) {
    return fetch(endpoint, {
      method: 'POST',
      cache: 'no-cache',
      headers: {
        'Content-Type': 'application/json; charset=utf-8',      
        'Accept': 'application/json',
      },
      body:JSON.stringify(data)
    }).then(response => response.json())
  }

  static callApiWithGET(endpoint) {
    return fetch(endpoint, {
      method: 'GET',
      cache: 'no-cache',
      headers: {
        'Content-Type': 'application/json; charset=utf-8',      
        'Accept': 'application/json',
      }
    }).then(response => response.json())
  }

  static postAuthenticate(user, password) {
    return AuthAPI.callApiWithPOST(config.SERVER_AUTH_ROOT+"/authenticate",{UserName:user, Password: password})
  }

  static authenticate(user, password) {
    return AuthAPI.callApiWithPOST(config.SERVER_AUTH_ROOT+"/authenticate",{UserName:user, Password: password})
  }


  static async isUserAuthorized() {
      var token = AuthAPI.getToken();
      if (token != null) {
        //console.log("AuthAPI::isUserAuthorized, found token: ", token);
        let isVerified = await AuthAPI.verifyToken(token);
        if (isVerified === false) {
          AuthAPI.deleteToken();  // Delete this token - it's not valid!
        }
        return isVerified
      }
      return false
  }

  static async verifyToken(token) {
    let response = await AuthAPI.callApiWithGET(config.SERVER_AUTH_ROOT+"/verifytoken/"+token);
    if (response === true) {
      console.log("AuthAPI::verifyToken, token is valid: ", response);
    } else {
      console.log("AuthAPI::verifyToken, invalid access token: ", response); 
    }
    return response;
  }

  static getToken() {
      return window.localStorage.getItem(config.AUTH_TOKEN);
  }

  static setToken(token) {
      if (token) {
          window.localStorage.setItem(config.AUTH_TOKEN, token);
      }
  }

  static deleteToken() {
      window.localStorage.removeItem(config.AUTH_TOKEN);
  }
}
