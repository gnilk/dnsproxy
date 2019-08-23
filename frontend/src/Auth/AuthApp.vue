<template>
    <div class="root_container">
        <div class="auth_container">
            <h1><center>Authenticate</center></h1>
            <div class="auth_form">
                <span v-if="error !== ''" class="auth_error"><p>{{ error }}</p></span>
                <input v-model="username" v-on:input="onChangeUsername" placeholder="username" />  
                <input v-model="password" v-on:input="onChangePassword" type="password" placeholder="" />  
                <button @click="onAuthClick()">Login</button>
            </div>
        </div>
    </div>
</template>
<style scoped lang="scss">
    @import '../commonstyles.scss';
    .auth_container {
        display: flex;
        flex-direction: column;
        margin-top: 8px;
    }
    .auth_error {
        color: $noice_fg_color_error;
    }
    .auth_form {
        display: flex;
        flex-direction: column;
        align-items: center;        
        margin-top: 8px;
    }
</style>
<script>
import AuthApi from '../Auth/AuthApi'
export default {
    components: {

    },
    data: function() {
        return {
            username: "",
            password: "",
            error: ""
        }
    },
    methods: {
        onAuthClick: function () {
            AuthApi.authenticate(this.username, this.password).then((resp) => {
                console.log("auth response: ", resp);
                AuthApi.setToken(resp);
                window.location.replace("/intra_index.html");
            }).catch(reason => {
                console.log("AuthApp::onAuthClick, user rejected");
                this.error = "bad login attempt";
            });
        },
        onChangeUsername: function() {
            this.error = "";
        },
        onChangePassword: function() {
            console.log("OnChange")
            this.error = "";
        }

    }
}
</script>
