
import * as config from './config.js'
import Auth from './Auth/AuthApi'

//const AUTH_TOKEN = 'X-Auth-Token';

export default class IntraAPI {
    static handleErrorResponse(response) {
        console.log(response);
        throw new Error(response.status + " (" + response.statusText + ")");
      }
  
    static callApiWithGET(endpoint) {
        let token = Auth.getToken();


        console.log("IntraAPI::callApiWithGET"); //, token: ", token);

        return fetch(endpoint, {
            method: 'GET',
            cache: 'no-cache',
            headers: {
                'Content-Type': 'application/json; charset=utf-8',      
                'Accept': 'application/json',
                'X-Auth-Token': token,
            }
        }).then(response => {
            if (response.ok) {
                return response.json();
            } else {
                IntraAPI.handleErrorResponse(response);
            }
        })
    }
    
    static getCurrentUser() {
        console.log("IntraAPI::getCurrentUser")
        return IntraAPI.callApiWithGET(config.INTRA_API_ROOT+"/sessionuser")
    }
}