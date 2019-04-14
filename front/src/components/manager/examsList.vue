<template>
<v-container>
{{post}}
</v-container>
<!--
  <v-data-table :headers="headers" :items="post" class="elevation-1">
    <template slot="items" slot-scope="props">
      <td>{{ props.item.name }}</td>
      <td class="text-xs-right">{{ props.item.examTitle }}</td>
      <td class="text-xs-right">{{ props.item.examCost }}</td>
      <td class="text-xs-right">{{ props.item.selectCourse }}</td>
      <td class="text-xs-right">{{ props.item.numberofFinalQuestion }}</td>
      <td class="text-xs-right">{{ props.item.examDuration }}</td>
    </template>
  </v-data-table>
  -->
</template>

<script>
  export default {
    data() {
      return {
        loading: false,
        post:{},
        error: null,
        axios:require('axios')
      }
    }, 
    created () {
    // fetch the data when the view is created and the data is
    // already being observed
    this.fetchData()
    },
     watch: {
    // call again the method if the route changes
    '$route': 'fetchData'
    },
    methods:{
      fetchData () {
      this.axios.get('http://localhost:9090/exam/list').then(response => (this.post = response))
      }

    }
  }
</script>