<template>
    <div v-if="editingProp">
Mahdi
    </div>
    <div v-else >
      <v-btn @click="addingOption()" flat class="success">
        <v-icon class="margin-left-10">add_circle</v-icon>
        افزودن گزینه</v-btn>
        <v-list-tile v-for="(choice,n) in choicesArray" :key="'ch_'+n" class="add-question-choices">
              <v-list-tile-avatar>

                <v-btn small depressed color="info" fab @click="selectCheckbox(choice)">
                  {{n+1}}
                </v-btn>
              </v-list-tile-avatar>
              <v-list-tile-content>
                <!--<v-list-tile-title :class=" (choicesArray.dir=='rtl') ? 'text-xs-right': 'text-xs-left' >-->
                <v-list-tile-title class="add-question-option" >
                  <!--<v-text-field :counter="100" :label="'گزینه '+(n+1)" required></v-text-field>-->
                  <v-checkbox v-modal="choice.value" hide-details class="shrink mr-2"></v-checkbox>
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
    editingProp:Boolean,
},
updated(){
}
,
methods:{
  addingOption(){
    this.$emit('addOption')
  }
  ,
  selectCheckbox(val){
    this.$emit('checkboxClicked',val);
  }
}
}
</script>

<style>

</style>
