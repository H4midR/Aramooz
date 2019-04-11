<template>
    <v-container>
        <v-layout class="add-question-form" justify-center>
            <v-flex xs12 sm11 md10 lg8>
                <v-card>
                    <v-form>
                        <v-toolbar class="pink" >
                            <v-icon color="white">help</v-icon>
                            &nbsp;<span class="white--text" >افزودن سوال</span>
                            <v-spacer>
                            </v-spacer>
                            <v-flex shrink class="dir-switch">
                                <v-switch label="چپ چین" color="yellow" v-model="editingQuestion.ltr"></v-switch>
                            </v-flex>
                        </v-toolbar>
                        <div :class="(editingQuestion.ltr) ? 'ltr text-xs-left' : 'rtl text-xs-right'">
                        <v-card-title primary-title >
                            <v-layout>
                                <v-flex xs12 sm2>
                                    <v-text-field v-model="editingQuestion.qnum" style="margin:4px 10px;" label="شماره سوال" required >{{editingQuestion.qnum}}</v-text-field>
                                </v-flex>
                                <v-flex xs12 sm10>
    <v-text-field :counter="200" label="عنوان سوال" required v-model="editingQuestion.title"></v-text-field>
                                </v-flex>
                            </v-layout>
                        </v-card-title>
                        <v-card-text >
    <choices :choicesCount="numberOfChoices" :choicesArray="editingQuestion.choices" @request2deleteItem="itemDeleted" @request2addOption="optionAdded" @request2selectItem="itemSelected"></choices>
                        </v-card-text>
                        </div>
                        <v-divider></v-divider>
                        <v-btn class="success" @click="submitQuestion">
                            <v-icon >check</v-icon>&nbsp;ثبت
                        </v-btn>
                    </v-form>
                </v-card>
            </v-flex>
        </v-layout>
    </v-container>
</template>
<script>
import choices from '@/components/manager/addQuestion/choices'
export default {
data(){
return{
    axios:require('axios'),
    editing:false,
    emptyChoice:{
            num:0,
            title:'',
            value:false
        },
    emptyQuestion:{
        title:null,
        qnum:null,
        ltr:false,
        choices:[{
            num:1,
            title:'',
            value:false
        }]
    },
    editingQuestion:{},
    numberOfChoices:0,
}
},
/*
editing: true | false ==> editing a question or add new one?
item2edit => Which item is being edited?

props:{
editing:Boolean,
item2edit:Object,
},*/
mounted(){
    if(this.editing){
        this.editingQuestion= Object.assign({},item2edit)
    }else
        this.editingQuestion= Object.assign({},this.emptyQuestion)
    this.numberOfChoices=(this.editingQuestion.choices.length);
},
methods:{
    optionAdded(){
        var oneChoice=Object.assign({},this.emptyChoice);
        var lengthOfChoices=this.editingQuestion.choices.length;
        oneChoice.num=(this.editingQuestion.choices[lengthOfChoices-1].num)+1;
        this.editingQuestion.choices.push(oneChoice);
        this.numberOfChoices=(this.editingQuestion.choices.length);
    },
    itemDeleted(item){
var item_index=this.editingQuestion.choices.indexOf(item);
confirm('این عمل قابل برگشت نمی باشد، آیا مطمئن هستید؟') && this.editingQuestion.choices.splice(item_index, 1)
var lengthOfChoices=(this.editingQuestion.choices.length);
this.numberOfChoices=lengthOfChoices;
    },
    itemSelected(item){
    var item_index=this.editingQuestion.choices.indexOf(item);
    if((this.editingQuestion.choices[item_index].value)==false){
        this.editingQuestion.choices[item_index].value=true;
    }else{
        this.editingQuestion.choices[item_index].value=false;
    }
    },
    submitQuestion(){
        console.log(this.editingQuestion);
        this.axios.post("http://localhost:9090/addquestion",JSON.stringify({
            question:this.editingQuestion
        })).then(res=>{
            console.log(res);
        });
    }
},
components:{
    choices
},
computed:{
    
}

}
</script>

<style>
.add-question-form .dir-switch{
    margin-top:17px;
}
.add-question-form .dir-switch *{
    color:#fff;
}
.add-question-form .ltr *{
    direction:ltr;
    text-align:left;
}
.add-question-form .rtl *{
    direction:rtl;
    text-align:right;
}
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
