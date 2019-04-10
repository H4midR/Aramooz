<template>
  <v-app>

    <v-content fill-height fluid>

      <router-view @Login="userloggedin" :User="User" :BaseURL="baseUrl" :ACL="ACL" />

    </v-content>





  </v-app>
</template>

<script>
  export default {
    data() {
      return {
        User:null,
        baseUrl:"http://localhost:9090",
        ACL:{},

      }
    },

    name: 'App',
    computed: {

    },
    methods: {
      userloggedin(e){
        this.User=e;
        // TODO : get ACL list from server , depends on user.acl
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
