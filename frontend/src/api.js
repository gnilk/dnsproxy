
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

        console.log("BackendAPI::callApiWithGET, endpoint: " + endpoint); //, token: ", token);

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
    
    static getDevices() {
        console.log("BackendAPI::getDevices")
        return BackendAPI.callApiWithGET(config.SERVER_API_ROOT + "/devices")
    }
    static getDeviceRules(device) {
        console.log("BackendAPI::getDeviceRules")
        return BackendAPI.callApiWithGET(config.SERVER_API_ROOT + "/pouet/product/"+prod)
    }
    static blockDevice(device) {
        console.log("BackendAPI::blockDevice")
        return BackendAPI.callApiWithGET(config.SERVER_API_ROOT + "/device/"+device.Name+"/block")
    }
    static releaseDevice(device) {
        console.log("BackendAPI::unblockDevice")
        return BackendAPI.callApiWithGET(config.SERVER_API_ROOT + "/device/"+device.Name+"/release")
    }

}