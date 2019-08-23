<template>
    <div class="root_container">
        <div class="header">
            <DeviceList />
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

    import DeviceList from './Components/DeviceList.vue'

    export default {
        components: {
            DeviceList
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

