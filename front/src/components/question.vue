<template>
  <v-container>
    <ul>
      <li v-for="item in question.choise" :key="'key'+item.title" @click="liselected(item)">
        {{item.title}}
      </li>
    </ul>
  </v-container>
</template>
<script>
export default {
  data(){return{

  }},
  methods:{
    liselected(item){
      this.$emit("mysignal",item.value)

    }
  },
  props:{
    question:Object,
  }
}
</script>
