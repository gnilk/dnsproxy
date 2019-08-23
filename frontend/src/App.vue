<template>
    <div class="root_container">
        <div class="header">
            <div class="toplogo">
                <img src="toplogo.gif" />
            </div>
            <div class="tabs">
                <span :class="[ subpage === 'news' ? 'tab is-active' : 'tab']">
                    <a @click="subpage='news'">News</a>
                </span>
                <span :class="[ subpage === 'releases' ? 'tab is-active' : 'tab']">
                    <a @click="subpage='releases'">Releases</a>
                </span>                
                <span :class="[ subpage === 'history' ? 'tab is-active' : 'tab']">
                    <a @click="subpage='history'">History</a>
                </span>                
                <span :class="[ subpage === 'bla' ? 'tab is-active' : 'tab']">
                    <a @click="onIntraClick()">Intranet</a>
                </span>                
            </div>
        </div>
        <div class="main_container">
            <div v-if="subpage === 'news'">
                THIS IS A NEWS PAGE
            </div>
            <div v-if="subpage === 'releases'">
                <ReleaseList />
            </div>
            <div v-if="subpage === 'history'">
                History
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
    //
    // NOTE: Tried this with Dynamic Components but never got them to work properly
    //       Was unable to get events to propagate properly up to the parent
    //
    // import GalleryListing from './Components/GalleryListing.vue'
    // import GalleryView from './Components/GalleryView.vue'
    // import ImageViewer from './Components/ImageViewer.vue'

    import ReleaseList from './Components/ReleaseList.vue'

    export default {
        components: {
            // GalleryListing,
            // GalleryView,
            // ImageViewer,
            ReleaseList
        },
   
        data: function() {
            return {
                subpage: 'news',
                greeting: "Hello"
            }
        },
        methods: {
            onIntraClick: function() {
                console.log("App::onIntraClick")
                window.location.replace("/intra_index.html");
            },
        }
    }   
</script>

