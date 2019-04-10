<template>
    <v-container fluid>
        <v-card class="exam-card">
            <examBasicInfo :examInfoProp="examInfo"></examBasicInfo>
            <questionsNumber :questionsNumberProp="examContent" @question_selected="showQuestion"></questionsNumber>
            <v-divider></v-divider>
            <questionContent :questionContentProp="currentQuestion" ></questionContent>
        </v-card>
    </v-container>
</template>

<script>
import examBasicInfo from '@/components/client/exam/examBasicInfo'
import questionContent from '@/components/client/exam/questionContent'
import questionsNumber from '@/components/client/exam/questionsNumber'
export default {
data(){return{
    currentQuestion:Object,
    examInfo:{
        title:'آزمون مقدماتی ریاضی',
        time:"20",
    },
    examContent:[
        // fetch from DB
            {
                dir:'rtl',
                num:'1',
                title: "<p>کدام گزینه یک عدد زوج است؟</p>",
                choices:[
                    {title: "<p>asasdasdasd</p>",value:1},
                    {title: "3",value:0},
                    {title: "4",value:1},
                    {title: "5",value:0}
                ]
            },
            {
                dir:'rtl',
                num:'2',
                title: "کدام یک عدد است؟",
                choices:[
                    {title: "2",value:1},
                    {title: "3",value:0},
                    {title: "4",value:1},
                    {title: "5",value:0}
                ]
            },
            {
                dir:'rtl',
                num:'3',
                title: "کدام یک حرف الفبا است؟",
                choices:[
                    {title: "2",value:1},
                    {title: "3",value:0},
                    {title: "4",value:1},
                    {title: "5",value:0}
                ]
            },
            {
                dir:'ltr',
                num:'4',
                title: "کدام گزینه یک عدد فرد است؟",
                choices:[
                    {title: "2",value:1},
                    {title: "3",value:0},
                    {title: "4",value:1},
                    {title: "5",value:0}
                ]
            }
        ],
        
}},
components:{
    questionContent,
    questionsNumber,
    examBasicInfo
},
created(){
    this.currentQuestion=this.examContent[0];
},
mounted(){
},
methods:{
    showQuestion(passedItem){
        this.currentQuestion=passedItem;
    }
}
}
</script>

<style>
#timer{
    font-family:yekan;
    font-size:28px;
    padding-left:20px;
}
.v-btn{
    font-family: iransans;
    font-size:12px;
}
#question_selector .v-btn{
    font-family: iransans;
    font-weight:bold
}
.question-bar{
    background-color:#ffffff;
    border-bottom:2px solid #dadada;
}
#question_selector{
    padding:3px 0;
}
#question_selector .v-btn{
padding:0 5px !important;
min-width:50px;
}
.mini-bottom-margin{
margin-bottom: 2px;
padding-bottom: 0;
}
.zero-top-margin{
    margin-top: 0;
    padding-top:0
}
.inline{
    display:inline;
}
.container.fluid{
    width:90%;
}
.exam-card{
    background-image: url(//gitiazmoon.com/public/images/bg/bgtest/bgexam8.png);
    background-position: bottom left;
    background-size: 70% 60%;
    background-repeat: no-repeat;
}
.exam-card hr.v-divider{
    margin-top:0;
    margin-bottom:0;
}
.exam-card .v-list{
    background:none;
}
</style>
