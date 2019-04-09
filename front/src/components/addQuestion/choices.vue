<template>
    <div  >
        <v-list-tile v-for="choice in choicesArray" :key="'ch_'+choice.num" class="add-question-choices">
              <v-list-tile-avatar>

                <v-btn small depressed :class="(choice.value)? 'success':'grey lighten-2'" fab @click="selectCheckbox(choice)">
                  {{choice.num}}
                </v-btn>
              </v-list-tile-avatar>
              <v-list-tile-content>
                <!--<v-list-tile-title :class=" (choicesArray.dir=='rtl') ? 'text-xs-right': 'text-xs-left' >-->
                <v-list-tile-title class="add-question-option" >
                  <v-layout>
                    <v-flex xs11>
<v-text-field :counter="100" :label="'گزینه '+(choice.num)" required></v-text-field>
                    </v-flex>
                    <v-flex xs1 justify-center>
                      <v-btn flat icon class="grey--text text--darken-1" style="margin-top:27px">
                        <v-icon >delete</v-icon>
                      </v-btn>
                    </v-flex>
                  </v-layout>
                </v-list-tile-title>
              </v-list-tile-content>
        </v-list-tile>
        <v-btn depressed large icon @click="addingOption()" :class="(choicesCount>3)? '':'light-blue white--text'" :disabled="(choicesCount>3)? true:false">
          <v-icon large >add</v-icon>
        </v-btn>
    </div>
</template>

<script>
export default {
  data(){
    return{
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
