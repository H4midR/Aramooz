<template>
    <v-app light>       
        <drawerBlock :drawer="drawer" :clipped="clipped"/>
        
        <toolbarBlock :drawer="drawer" :clipped="clipped" @drawerChenge="drawerChenge"/>

        <v-content>
            <v-container fluid>

            </v-container>
        </v-content>
  </v-app>
</template>

<script>
import drawerBlock from '@/components/manager/dashboard/drawer_block'
import toolbarBlock from '@/components/manager/dashboard/toolbar_block'
export default {
    data () {
        return {
            title: "پنل مدیریت آزمون یار", //Page or Site Title
            drawer: true,
            clipped: false,
        }
    },
    components:{
        drawerBlock,toolbarBlock
    },
    methods:{
        drawerChenge(){
            this.drawer = !this.drawer
            this.clipped = !this.clipped
        }
    },
    prop:{
        User:Object
    },
    mounted(){
        if (this.User == null){
            this.$router.replace("/");
        }
    }

}
</script>