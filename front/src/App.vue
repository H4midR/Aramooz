<template>
  <v-app>

    <v-content fill-height fluid>

      <router-view @Login="userloggedin" :User="User" />

    </v-content>





  </v-app>
</template>

<script>
  export default {
    data() {
      return {
        User:null,
      }
    },

    name: 'App',
    computed: {

    },
    methods: {
      userloggedin(e){
        this.User=e
      }

    },
    created: function () {

    },

    watch: {

    }


  }

</script>
<style scoped>
</style>
