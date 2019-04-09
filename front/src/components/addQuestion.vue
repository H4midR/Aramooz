<template>
    <v-container>
        {{editingQuestion}}
        <v-layout class="add-question-form" justify-center>
            <v-flex xs12 sm11 md10 lg8>
                <v-card>
                    <v-toolbar class="light-blue lighten-1" >
                        <v-icon color="white">help</v-icon>
                        &nbsp;<span class="white--text " >افزودن سوال</span>
                    </v-toolbar>
                    <v-card-title primary-title>
                        <v-layout>
                            <v-flex xs12 sm2 md2 lg2>
                                <v-text-field style="margin-left:10px;" label="شماره سوال" required ></v-text-field>
                            </v-flex>
                            <v-flex xs12 sm10 md10 lg10>
<v-text-field :counter="200" label="عنوان سوال" required></v-text-field>
                            </v-flex>
                        </v-layout>
                    </v-card-title>
                    <v-card-text>
<choices :choicesArray="editingQuestion.choices" :editingProp="editing" @addOption="addOption2question" @checkboxClicked="checkboxChenged"></choices>
                    </v-card-text>
                </v-card>
            </v-flex>
        </v-layout>
    </v-container>
</template>

<script>
import choices from '@/components/addQuestion/choices'
export default {
data(){
return{
    editing:false,
    emptyChoice:{
            num:0,
            title:'',
            value:false
        },
    emptyQuestion:{
        title:'',
        rank:0,
        choices:[{
            num:1,
            title:'1',
            value:false
        }]
    },
    editingQuestion:{}
}
},
mounted(){
    if(this.editing){

    }else
        this.editingQuestion=this.emptyQuestion
}
,

methods:{
    addOption2question(){
        var oneChoice=this.emptyChoice;
        var lengthOfChoices=(this.editingQuestion.choices.length)+1;
        oneChoice.title=lengthOfChoices.toString();
        oneChoice.num=lengthOfChoices;
        console.log(this.editingQuestion.choices.push(oneChoice));
        //this.emptyChoice.title='';
    },
    checkboxChenged(item){
    //console.log(item);
    var item_index=this.editingQuestion.choices.indexOf(item);
    //console.log(this.item_index);
    if((this.editingQuestion.choices[item_index].value)==false){
        this.editingQuestion.choices[item_index].value=true;
    }else{
        this.editingQuestion.choices[item_index].value=false;
    }
    }
},
components:{
    choices
}
}
</script>

<style>
.add-question-form .v-list{
    background:none;
}
.add-question-form .v-label{
    font-size:12px;
}
.add-question-form .add-question-choices>*{
    height:70px;
}
.add-question-form .add-question-choices{ 
    margin-bottom:5px;
}
.add-question-form .add-question-option{
    height:inherit;
}
.add-question-form .add-question-option>*{
    height:60px;
    margin-top:5px;
}
.margin-left-10{
    margin-left:10px;
}
</style>
