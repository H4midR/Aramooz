<template>
    <v-app light>       
        <v-toolbar fixed app :clipped-left="clipped">
            <v-toolbar-side-icon @click="drawerChenge" />
            <v-toolbar-title v-text="title" />
            <v-spacer/>

            <v-menu>
                <v-icon slot="activator">more_vert</v-icon>
                <v-card>
                <v-list>
                <v-list-tile avatar>
                    <v-list-tile-avatar>
                        <img src="https://avatars2.githubusercontent.com/u/47215007?s=40&v=4" alt="John">
                    </v-list-tile-avatar>

                    <v-list-tile-content>
                        <v-list-tile-title>کاربر</v-list-tile-title>
                        <v-list-tile-sub-title>09156535299</v-list-tile-sub-title>
                    </v-list-tile-content>
                </v-list-tile>
                </v-list>

                <v-divider></v-divider>

                <v-list>
                    <v-list-tile router to="profile">
                        <v-list-tile-content>
                            <v-list-tile-title>پروفایل</v-list-tile-title>
                        </v-list-tile-content>

                        <v-list-tile-action>
                            <v-icon>person</v-icon>
                        </v-list-tile-action>
                    </v-list-tile>

                    <v-list-tile router to="login">
                        <v-list-tile-content>
                            <v-list-tile-title>خروج</v-list-tile-title>
                        </v-list-tile-content>

                        <v-list-tile-action>
                            <v-icon>mdi-login</v-icon>
                        </v-list-tile-action>
                    </v-list-tile>
                </v-list>
            </v-card>
            </v-menu>
        </v-toolbar>      
  </v-app>
</template>

<script>
export default {
    data () {
        return {
            title: "پنل مدیریت آزمون یار", //Page or Site Title
            drawer: true,
            clipped: false,
        }
    },
    props:{
        drawer: Boolean,
        clipped: Boolean,
    },
    methods:{
        drawerChenge(){
            this.$emit('drawerChenge');
        }
    }

}
</script>