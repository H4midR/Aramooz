<template>
    <div  >
        <v-list-tile v-for="choice in choicesArray" :key="'ch_'+choice.num" class="add-question-choices">
              <v-list-tile-avatar>

                <v-btn small depressed :color="(choice.value)? 'success':'info'" fab @click="selectCheckbox(choice)">
                  {{choice.num}}
                </v-btn>
              </v-list-tile-avatar>
              <v-list-tile-content>
                <!--<v-list-tile-title :class=" (choicesArray.dir=='rtl') ? 'text-xs-right': 'text-xs-left' >-->
                <v-list-tile-title class="add-question-option" >
                  <v-text-field :counter="100" :label="'گزینه '+(choice.num)" required></v-text-field>
                </v-list-tile-title>
              </v-list-tile-content>
        </v-list-tile>
        <v-btn @click="addingOption()" :class="(choicesCount>3)? '':'success'" :disabled="(choicesCount>3)? true:false">
          <v-icon class="margin-left-10">add_circle</v-icon>
          افزودن گزینه
        </v-btn>
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
    choicesCount:Number
},
computed:{
  
},
methods:{
  addingOption(){
    this.$emit('addOption');
  }
  ,
  selectCheckbox(item){
    this.$emit('checkboxClicked',item);
  }
}
}
</script>

<style>

</style>
