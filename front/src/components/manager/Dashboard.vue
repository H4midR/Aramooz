<template>
    <v-app light>       
        <v-navigation-drawer :clipped="clipped" v-model="drawer" enable-resize-watcher app >
            <v-list>
            </v-list>
        </v-navigation-drawer>      
        <v-toolbar fixed app :clipped-left="clipped">
            <v-toolbar-side-icon @click.stop="drawer = !drawer" />
            <v-toolbar-title v-text="title"/>
            <v-spacer/>
            <v-btn icon>
                <v-icon>more_vert</v-icon>
            </v-btn>
        </v-toolbar>      
        <v-content>
            <v-container fluid>
            </v-container>
        </v-content>
  </v-app>
</template>

<script>
export default {
    data () {
        return {
            title: "پنل مدیریت آزمون یار", //Page or Site Title
            drawer: true,
            clipped: false,
            items: [
                {
                action: 'local_activity',
                title: 'Attractions',
                path: '/',
                items: [],
                },
                {
                action: 'restaurant',
                title: 'Breakfast',
                path: '/breakfast',
                items: []
                },       
            ] //items
        }

}

}
</script>