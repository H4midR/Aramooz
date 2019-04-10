<template>
    <v-app light>       
        <v-navigation-drawer :clipped="clipped" v-model="drawer" enable-resize-watcher app right>

            <v-toolbar flat class="transparent">
                <v-list class="pa-0">
                <v-list-tile avatar>
                    <v-list-tile-avatar>
                    <img src="https://avatars2.githubusercontent.com/u/47215007?s=40&v=4">
                    </v-list-tile-avatar>

                    <v-list-tile-content>
                    <v-list-tile-title>کاربر</v-list-tile-title>
                    </v-list-tile-content>
                </v-list-tile>
                </v-list>
            </v-toolbar>

            <v-divider></v-divider>

            <v-list dense>
                <v-list-tile v-for="drawerItem in drawerList" :key="drawerItem.title" router :to="drawerItem.url">
                    <v-list-tile-action>
                    <v-icon>{{ drawerItem.icon }}</v-icon>
                    </v-list-tile-action>

                    <v-list-tile-content>
                    <v-list-tile-title>{{ drawerItem.title }}</v-list-tile-title>
                    </v-list-tile-content>
                </v-list-tile>
            </v-list>
        </v-navigation-drawer>      
  </v-app>
</template>

<script>
export default {
    data () {
        return {
            
            drawerList: [
                {
                title: 'داشبرد',
                icon:"dashboard",
                url: 'Dashboard',
                drawerList: [],
                },
                {
                title: 'آزمون ها',
                icon:"event_note",
                url: 'examsList',
                drawerList: [],
                },
                {
                title: 'کاربران',
                icon:"people",
                url: 'users',
                drawerList: [],
                },
            ] //drawerList
        }

    },
    props:{
        drawer: Boolean,
        clipped: Boolean,
    }

}
</script>