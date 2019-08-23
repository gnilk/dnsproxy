<template>
    <div class="list-container">
        <div class="list-header">
            <h1>Releases</h1>
        </div>
        <div class="tabs">
                <span :class="[ filter === 'pc' ? 'tab is-active' : 'tab']">
                    <a @click="setFilter('pc')">PC</a>
                </span>
                <span :class="[ filter === 'amiga' ? 'tab is-active' : 'tab']">
                    <a @click="setFilter('amiga')">Amiga</a>
                </span>                
                <span :class="[ filter === 'c64' ? 'tab is-active' : 'tab']">
                    <a @click="setFilter('c64')">C64</a>
                </span>                
                <span :class="[ filter === 'other' ? 'tab is-active' : 'tab']">
                    <a @click="setFilter('other')">Other</a>
                </span>                
                <span :class="[ filter === '' ? 'tab is-active' : 'tab']">
                    <a @click="setFilter('')">All</a>
                </span>                
            
        </div>
        <div class="list-body" v-if="releases.length">
            <ul v-for="n in years()">
                <h2>{{ n }} </h2>
                <ul v-for="r in relyear(n)">
                    <ReleaseListItem
                        :key="r.id"
                        :release=r
                        @OnViewRelease="OnViewRelease"
                    />                   
                </ul>
            </ul>            
            <ul v-if="unknown.size">
                <h2>Unknown</h2>
                <ul v-for="r in unknowns()">
                    <ReleaseListItem
                        :key="r.id"
                        :release=r
                        @OnViewRelease="OnViewRelease"
                    />                   
                </ul>
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
    import ReleaseListItem from './ReleaseListItem.vue'
    import moment from 'moment'
    export default {
        components: {
            ReleaseListItem,
        },
        data:  function() {
            return {
                releases: [],
                peryear: new Map(),
                unknown: new Set(),
                filter: "",
            }
        },
        created: function() {
            console.log("releases listing created")
            BackendAPI.getGroup().then((resp) => {
                this.releases = resp.group.prods;
                console.log("Releases downloaded")
                //console.log(resp)
                this.SortReleases();
                return resp
            }).catch(reason => {
                console.log("ERROR: " + reason);
            })
        },
        methods: {
            years: function() {
                // need special handling of 'unknown' release dates
                let data = Array.from(this.peryear.keys()).sort();
                data.reverse();
                return data;
                
            },
            unknowns: function() {
                return Array.from(this.unknown);
            },
            relyear: function(year) {
                let data = Array.from(this.peryear.get(year));
                return data;
            },
            setFilter: function(filter) {
                this.filter = filter;
                console.log("new filter: "+filter);
                this.SortReleases();
            },
            isPc(slug) {
                if (slug === "windows") {
                    return true;
                }
                if (slug === "msdos") {
                    return true;
                }
                if (slug === "msdosgus") {
                    return true;
                }
                return false;            
            },
            isAmiga(slug) {
                if (slug === "amigaocsecs") {
                    return true;
                }
                if (slug === "amigaaga") {
                    return true;
                }
                return false;
            },
            isC64(slug) {
                if (slug === "c64dtv") {
                    return true;
                }
                if (slug === "commodore64") {
                    return true;
                }
                return false;
            },

            isValidForFilter(r) {
                if (this.filter === "") {
                    return true;
                }
                var isok = false;
                Object.entries(r.platforms).forEach(([key,value]) => {
                    if (this.filter === "pc") {
                        isok = this.isPc(value.slug);
                    }
                    if (this.filter === "amiga") {
                        isok = this.isAmiga(value.slug);
                    }
                    if (this.filter === "c64") {
                        isok = this.isC64(value.slug);
                    }
                    if (this.filter === "other") {
                        if (!this.isC64(value.slug) && !this.isAmiga(value.slug) && !this.isPc(value.slug)) {
                            isok = true;
                        }
                    }

                })
                return isok;
            },
            SortReleases: function() {    
                this.peryear = new Map();
                this.unknown = new Set();
                this.releases.forEach((r) => {
                    if (this.isValidForFilter(r)) {
                        let year = "unknown";
                        if (r.releaseDate != null) {                          
                            let mo = moment(r.releaseDate);
                            // Is date valid????
                            if (mo.isValid()) {
                                year = mo.year();
                            } else {                            
                                // In the Pouet data the release date month is zero...
                                let fakeDate = r.releaseDate.substring(0,4) + "-01-01";
                                mo = moment(fakeDate);
                                if (mo.isValid()) {
                                    year = mo.year();
                                }
                            }
                        }
                        if (year !== "unknown") {
                            if (this.peryear.has(year)) {
                                let list = this.peryear.get(year);
                                list.add(r);
                            } else {
                                let list = new Set();
                                list.add(r);
                                this.peryear.set(year,list);
                            }
                        } else {
                            this.unknown.add(r);
                        }
                    }
                })
                console.log(this.peryear)
            },
            OnViewRelease: function(release) {
                console.log("ReleaseListing::OnViewRelease: ", release.name)
            }
        }
    }
</script>