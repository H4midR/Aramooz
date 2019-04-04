<template>
  <v-app>
    <v-content fill-height fluid>

      <router-view :User="User" />

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

    },
    created: function () {

    },

    watch: {

    }


  }

</script>
<style scoped>
</style>
