<template>
    <div class="list-item">
        <div class="header">
            {{ release.name }}
        </div>
        <div class="infopane">
            <div class="screenshot_container">
                <img class="screenshot" :src="release.screenshot"  />
            </div>
            <div class="details">
                <div class="detail">
                    <span class="detail_c1">type:</span>
                    <span class="detail_c2">{{ release.type }}</span>
                </div>
                <div class="detail">
                    <span class="detail_c1">date:</span>
                    <span class="detail_c2">{{ release.releaseDate }}</span>
                </div>
                <div class="detail" v-if="release.party">
                    <span class="detail_c1">party:</span>
                    <span class="detail_c2">{{ release.party.name }}</span>
                </div>
                <div class="detail" v-if="release.party">
                    <span class="detail_c1">rank:</span>
                    <span class="detail_c2">{{ release.party_place }}</span>
                </div>
                <div class="detail" v-if="release.party">
                    <span class="detail_c1">compo:</span>
                    <span class="detail_c2">{{ release.party_compo_name }}</span>
                </div>                
            </div>
        </div>
    </div>
</template>
<style scoped lang="scss">
    @import '../commonstyles.scss';

    .list-item {
        @extend .noice_release_container;
        display: flex;
        flex-direction: column;
        justify-content: flex-start;
    }
    .header {
        display: flex;
        background-color: $noice_bg_color_dark;
        flex-direction: row;
        justify-content: flex-start;
        padding-left: 2px;
    }
    .infopane {
        @extend .noice_default_container;
        display: flex;
        flex-direction: row;
        justify-content: flex-start;
//        padding-left: 4px;
    }
    .screenshot_container {
        display: flex;
        flex-direction: column;
        justify-content: flex-start;
    }
    .screenshot {
        width: 256px;
        height: auto;

    }
    .details {
        @extend .noice_default_container;
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
        release: Object
    },
    // data: function() {
    //     return {
    //         release: {}
    //     }
    // },
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
