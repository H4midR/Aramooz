<template>
    <div  >
        <v-list-tile v-for="(choice,n) in choicesArray" :key="'ch_'+choice.num" class="add-question-choices">
              <v-list-tile-avatar>
                <v-tooltip bottom>
                <v-btn slot="activator" small depressed :class="(choice.value)? 'success':'grey lighten-2'" fab @click="selectingOption(choice)">
                  {{n+1}}
                </v-btn>
                <span v-if="choice.value">حذف پاسخ</span>
                <span v-else>انتخاب به عنوان پاسخ</span>
                </v-tooltip>
              </v-list-tile-avatar>
              <v-list-tile-content>
                <!--<v-list-tile-title :class=" (choicesArray.dir=='rtl') ? 'text-xs-right': 'text-xs-left' >-->
                <v-list-tile-title class="add-question-option" >
                  <v-layout>
                    <v-flex xs11>
<v-text-field :counter="100" :label="'گزینه '+(n+1)" required v-model="choice.title"></v-text-field>
                    </v-flex>
                    <v-flex xs1 justify-center>
                    <v-tooltip bottom v-if="(choicesCount)>1">
                      <v-btn flat icon slot="activator" class="grey--text text--darken-1" style="margin-top:27px" @click="deletingOption(choice)">
                        <v-icon >delete</v-icon>
                      </v-btn>
                      <span>حذف گزینه</span>
                    </v-tooltip>
                      
                    </v-flex>
                  </v-layout>
                </v-list-tile-title>
              </v-list-tile-content>
        </v-list-tile>
        <v-tooltip bottom>
        <v-btn slot="activator" depressed large icon @click="addingOption()" :class="(choicesCount>3)? '':'light-blue white--text'" :disabled="(choicesCount>3)? true:false">
          <v-icon large >add</v-icon>
        </v-btn>
        <span v-if="(choicesCount)>3">حداکثر گزینه ها ایجاد شده است</span>
        <span v-else>افزودن گزینه</span>
        </v-tooltip>
        
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
    choicesCount:Number
},
computed:{
  
},
methods:{
  addingOption(){
    this.$emit('request2addOption');
  }
  ,
  selectingOption(item){
    this.$emit('request2selectItem',item);
  },
  deletingOption(item){
this.$emit('request2deleteItem',item);
  }
}
}
</script>

<style>

</style>
