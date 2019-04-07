<template>
  <v-container>
    مجموع نمرات : {{total}}
    <question v-for="q in exam" :key="q.title" :qestion="q" @mysignal="myselected"></question>
  </v-container>
</template>
<script>
import question from '@/components/question'
export default {
  data(){return{
    total:0,
    exam:[
      {
        title:"کدام یک عدد زوج است؟",
        choise:[
          {
            title:"2",
            value:1,
          },
           {
            title:"3",
            value:0,
          },
           {
            title:"5",
            value:0,
          },
        ]
      },
      {
        title:"کدام یک عدد فرد است؟",
        choise:[
          {
            title:"2",
            value:0,
          },
           {
            title:"3",
            value:1,
          },
           {
            title:"4",
            value:0,
          },
        ]
      },
    ]
  }},
  methods:{
    myselected(e){
      this.total += e
    }
  },
  components:{
    question
  }
}
</script>
