<template>

<v-container>
  {{post[0]['_predicate_']}}

  <v-data-table :headers="post[0]['_predicate_']" :items="post" class="elevation-1">
    <template slot="items" slot-scope="props">
      <td>{{ props.item.name }}</td>
      <td class="text-xs-right">{{ props.item.Title }}</td>
      <td class="text-xs-right">{{ props.item.Cost }}</td>
      <td class="text-xs-right">{{ props.item.selectCourse }}</td>
      <td class="text-xs-right">{{ props.item.numberofFinalQuestion }}</td>
      <td class="text-xs-right">{{ props.item.duration }}</td>
    </template>
  </v-data-table>
</v-container>

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
      this.axios.get('http://localhost:9090/exam/list').then(response => (this.post = response.data.Data))
      }

    }
  }
</script>