import Vue from 'vue'
import AppIntra from './AppIntra'
import Auth from './Auth/AuthApi'
import AuthApp from './Auth/AuthApp'
import IntraAPI from './api_intra';


Vue.config.productionTip = false

  
async function startApp() {

    let auth = await Auth.isUserAuthorized();
    if (auth) {
        new Vue({
            el: '#app',
            template: '<AppIntra/>',
            components: { AppIntra }
        })
    } else {
        // Start auth app here
        new Vue({
            el: '#app',
            template: '<AuthApp/>',
            components: { AuthApp }
        })
    }

}

startApp();