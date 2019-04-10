<template>
  <div :dir="questionContentProp.dir">
        <v-list tow-line>
            <v-subheader>
              <h4 class="blue--text text--darken-4" v-html="questionContentProp.title" ></h4>
            </v-subheader>
            <v-list-tile v-for="(choice,n) in (questionContentProp.choices)" :key="'q_'+n+'_ch'">
              <v-list-tile-avatar>
                <v-btn small depressed color="success" fab>
                  {{n+1}}
                </v-btn>
              </v-list-tile-avatar>
              <v-list-tile-content>
                <v-list-tile-title :class=" (questionContentProp.dir=='rtl') ? 'text-xs-right': 'text-xs-left'" v-html="choice.title"></v-list-tile-title>
              </v-list-tile-content>
            </v-list-tile>
        </v-list>
        <v-spacer></v-spacer>
        <v-layout justify-center>
          <v-btn color="grey" class="lighten-2" >
            سوال بعد <v-icon >navigate_before</v-icon>
          </v-btn>
          <v-btn color="info" class=" " >
            پایان آزمون
          </v-btn>
        </v-layout>
      </div>
</template>
<script>
export default {
  data(){
    return{
      a:[],
    }
  },
  methods:{
    //a.indexOf()
  },
    props:{
        questionContentProp:Object,
    }
}
</script>
<style >
</style>


