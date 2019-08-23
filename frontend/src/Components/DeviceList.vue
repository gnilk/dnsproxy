<template>
    <div class="list-container">
        <div class="list-body" v-if="devices.length">
            <ul v-for="d in devices">
                <DeviceListItem 
                    :key="d.MAC" 
                    :device=d />
            </ul>
        </div>
    </div>
</template>
<style scoped lang="scss">
    @import '../commonstyles.scss';

.list-container {
        display: flex;
        flex-direction: column;
    }
    .list-actions {
        display: flex;
        flex-direction: row;
        justify-content: center;
    }
    .list-header {
        display: flex;
        flex-direction: row;
        justify-content: center;
    }
    .list-body {
        display: flex;
        flex-direction: column;
    }


</style>
<script>

            // <button v-on:click="setFilter('pc')">pc</button>
            // <button v-on:click="setFilter('amiga')">amiga</button>
            // <button v-on:click="setFilter('c64')">c64</button>
            // <button v-on:click="setFilter('other')">other</button>
            // <button v-on:click="setFilter('')">all</button>


    import BackendAPI from '../api.js'
    import DeviceListItem from './DeviceListItem.vue'
    import moment from 'moment'
    export default {
        components: {
            DeviceListItem,
        },
        data:  function() {
            return {
                devices: [],
            }
        },
        created: function() {
            console.log("device listing created")
            BackendAPI.getDevices().then((resp) => {
                this.devices = resp;
                console.log("Devices downloaded")
                //console.log(resp)
                return resp
            }).catch(reason => {
                console.log("ERROR: " + reason);
            })
        },
        methods: {
            OnViewRelease: function(release) {
                console.log("ReleaseListing::OnViewRelease: ", release.name)
            }
        }
    }
</script>