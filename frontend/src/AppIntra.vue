<template>
    <div class="root_container">
        <div class="header">
            <div class="toplogo">                
                <img src="intralogo.gif" />
            </div>
            <div class="tabs">
                <span :class="[ subpage === 'news' ? 'tab is-active' : 'tab']">
                    <a @click="subpage='news'">News</a>
                </span>
                <span :class="[ subpage === 'releases' ? 'tab is-active' : 'tab']">
                    <a @click="subpage='releases'">Releases</a>
                </span>               
                <!-- don't show user admin unless you are admin - additional checks are done on the backend --> 
                <span v-if="user.Role==='UserRoleAdmin'" :class="[ subpage === 'users' ? 'tab is-active' : 'tab']">
                    <a @click="subpage='users'">Users</a>
                </span>                
                <span :class="[ subpage === 'profile' ? 'tab is-active' : 'tab']">
                    <a @click="subpage='profile'">Profile</a>
                </span>                
                <span :class="[ subpage === 'public' ? 'tab is-active' : 'tab']">
                    <a @click="onPublicClick()">PublicSite</a>
                </span>                
            </div>
        </div>
        <div class="main_container">
            <div v-if="subpage === 'news'">
                News admin should go here
            </div>
            <div v-if="subpage === 'releases'">
                Release admin should go here
            </div>
            <div v-if="subpage === 'users'">
                User admin should go here
            </div>
            <div v-if="subpage === 'profile'">
                Your profile goes here
            </div>
        </div>
    </div>
</template>

<style scoped lang="scss">
    @import "./commonstyles.scss";
    .main_container {
        display: flex;
        flex-direction: column;
        height: 100%;
        margin-top: 8px;
        border-top: 1px solid black;
    }
    .header {
        display: flex;
        flex-direction: column;
        justify-content: center;
        margin-top: 16px;
    }
    .toplogo {
        display: flex;
        flex-direction: row;
        justify-content: center;
    }
</style>

<script>
    import IntraAPI from './api_intra';

    export default {
        components: {
        },
   
        data: function() {
            return {
                subpage: 'news',
                greeting: "Hello",
                user: Object,
            }
        },
        created: async function() {
            console.log("AppIntra, created");
            this.user = await IntraAPI.getCurrentUser()
            console.log("Current user: ", this.user)
        },
        methods: {
            onPublicClick: function() {            
                console.log("AppIntra::onPublicClick")
                window.location.replace("/");
            },
        }
    }   
</script>

