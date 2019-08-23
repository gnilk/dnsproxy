<template>
    <div class="list-item">
        <div class="header">
            Block
        </div>
        <div class="header">
            Unblock
        </div>
        <div class="header">
            {{ device.Host.Name }}
        </div>
        <div class="header">
            {{ device.Host.DefaultRule.Type }}
        </div>
    </div>
</template>
<style scoped lang="scss">
    @import '../commonstyles.scss';

    .list-item {
        @extend .device_container;
        display: flex;
        flex-direction: row;
        justify-content: flex-start;
    }
    .header {
        display: flex;
        background-color: $noice_bg_color_dark;
        border-right: 1px solid black;
        flex-direction: row;
        justify-content: flex-start;
        padding: 8px;
    }
    .details {
        @extend .device_container;
        display: flex;
        flex-direction: column;
        justify-content: flex-start;
        
    }
    .detail {
        display: flex;
        flex-direction: row;
        justify-content: flex-start;
        width: 100%;
        padding-left: 8px;
        padding-right: 8px;
    }
    .detail_c1 {
        display: flex;
        flex-direction: row;
        flex-grow: 2;
        justify-content: flex-start;
        margin-right: 8px;
    }
    .detail_c2 {
        display: flex;
        justify-content: flex-end;
        flex-direction: row;
        flex-grow: 1;
    }



</style>
<script>
export default {
    props: {
        device: Object
    },
    created: function() {
        // get product details - youtube stuff is not in the main overview
    },
    ready: function () {
        //console.log("Ready: "+this.props.release);
    },
    methods: {
        getYouTubeURL() {            
            console.log("youtube for: ", this.release)
            //const youtube = this.release.downloadLinks.filter((link) => { link.type === "youtube" });
            //return youtube;
            return "wedfwef";
        }
    }
}
</script>
