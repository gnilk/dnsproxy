<template>
    <div class="list-item">
        <div class="header">
            <div :class="[ blockState() === 'Blocked'? 'red':'green' ]">
                {{ blockState() }}
            </div>
            <SwitchButton color="#21F521" v-model="state" v-on:toggle="toggleState()"> {{ toggleActionFromState() }} </SwitchButton>
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
        align-items: center;
    }
    .header {
        display: flex;
        background-color: $noice_bg_color_dark;
        border-right: 1px solid black;
        flex-direction: row;
        justify-content: flex-start;
        padding: 8px;
    }
    .red {
        background-color: red;
        margin-right: 8px;
        width: 80px;
    }
    .green {
        background-color: green;
        margin-right: 8px;
        width: 80px;
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
    import SwitchButton from './SwitchButton.vue'
    import BackendAPI from '../api.js'

    export default {
        components: {
            SwitchButton,
        },
        data: function() {
            return {
                state: this.device.Host.DefaultRule.Type==="ActionTypeBlockedDevice"?false:true
            }
        },
        props: {
            device: Object
        },
        created: function() {
            // get product details - youtube stuff is not in the main overview
            
        },
        ready: function () {
            console.log("Ready: "+this.props.device);
        },
        methods: {
            toggleState() {
                console.log("toggleState, state is: ", this.state)
                let device = this.device;
                console.log("Device: ", device.Host.Name)
                if (this.state) {
                    BackendAPI.unblockDevice(device).then((resp) => {
                        console.log("Device unblocked")
                        this.device.Host.DefaultRule.Type = "ActionTypePass";
                    }).catch(reason => {
                        console.log("ERROR: " + reason);
                    })


                } else {
                    BackendAPI.blockDevice(device).then((resp) => {
                        console.log("Device blocked")
                        this.device.Host.DefaultRule.Type = "ActionTypeBlockedDevice";
                    }).catch(reason => {
                        console.log("ERROR: " + reason);
                    })
                }
            },
            toggleActionFromState() {
                return (this.state?"Block":"Unblock");
            },
            blockState() {
                console.log("BlockState for: ", this.device);
                console.log("BlockState is: ", this.device.Host.DefaultRule.Type);
                if (this.device.Host.DefaultRule.Type == "ActionTypeBlockedDevice") {
                    return "Blocked";
                }
                return "Free";
            }
        }
    }
</script>
