<template>
    <div v-if="editingProp">
Mahdi
    </div>
    <div v-else >
      <v-btn @click="addingOption()" flat class="success">
        <v-icon>add_circle</v-icon>
        افزودن گزینه</v-btn>
        <v-list-tile v-for="(choice,n) in choicesArray" :key="'ch_'+n" class="add-question-choices">
              <v-list-tile-avatar>

                <v-btn small depressed color="info" fab>
                  {{n+1}}
                </v-btn>
              </v-list-tile-avatar>
              <v-list-tile-content>
                <!--<v-list-tile-title :class=" (choicesArray.dir=='rtl') ? 'text-xs-right': 'text-xs-left' >-->
                <v-list-tile-title class="add-question-option" >
                  <v-text-field :counter="100" label="پاسخ" required></v-text-field>
                </v-list-tile-title>
              </v-list-tile-content>
        </v-list-tile>
    </div>
</template>

<script>
export default {
  data(){
    return{
      //choicesArray:[],
      //editingProp:false,
    }
  },
props:{
    choicesArray:Array,
    editingProp:Boolean
},
methods:{
  addingOption(){
    this.$emit('addOption')
  }
}
}
</script>

<style>

</style>
