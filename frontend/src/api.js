
import * as config from './config.js'

//const AUTH_TOKEN = 'X-Auth-Token';

export default class BackendAPI {
    static handleErrorResponse(response) {
        if (response.status === 403) {
            window.location.reload();
        } else {
            console.log(response);
            throw new Error(response.status + " (" + response.statusText + ")");
        }
      }
  
    static callApiWithGET(endpoint) {
        //let token = BackendAPI.getAuthToken();

        console.log("BackendAPI::callApiWithGET"); //, token: ", token);

        return fetch(endpoint, {
            method: 'GET',
            cache: 'no-cache',
            headers: {
                'Content-Type': 'application/json; charset=utf-8',      
                'Accept': 'application/json',
            }
        }).then(response => {
            if (response.ok) {
                return response.json();
            } else {
                BackendAPI.handleErrorResponse(response);
            }
        })
    }
    
    static getGroup() {
        console.log("BackendAPI::getGroup")
        return BackendAPI.callApiWithGET(config.SERVER_API_ROOT + "/pouet/group/26")
    }
    static getProduct(prod) {
        console.log("BackendAPI::getProd")
        return BackendAPI.callApiWithGET(config.SERVER_API_ROOT + "/pouet/product/"+prod)
    }

}