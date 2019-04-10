<template >
        <div class="mini-bottom-margin" id="question_selector">
        <v-layout align-center justify-center>
            <div class="inline" v-for="item in questionsNumberProp" :key="item.num" >
<v-btn depressed :id="'key_'+item.num" @click="select_question(item)" :class="(selected_item==item)? 'warning':'teal lighten-3'" >{{item.num}}</v-btn>
            </div>
            
        </v-layout>
        </div>
</template>

<script>
export default {
    data(){
        return{
            selected_item:false,
        }
    },
props:{
    questionsNumberProp:Array
},
methods:{
select_question(question){
    this.selected_item=question;
    this.$emit('question_selected',question );
}
},
}
</script>

<style>

</style>
