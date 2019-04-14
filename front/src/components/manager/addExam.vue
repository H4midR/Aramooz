<!-- Add Exam -->
<template>
<v-container justify-center>
    <v-layout class="addExamForm" row wrap >
        <v-flex xs12 sm11 md10 lg8>
            <v-card>
                <v-toolbar class="primary"> 
                    <v-icon color=""></v-icon>
                    <span style="color:#FFF">افزودن آزمون</span>
                    <v-spacer></v-spacer>
                </v-toolbar>
                <v-card-title>
                    <v-layout>
                        <p>آزمون / افزودن آزمون</p>
                    </v-layout>
                </v-card-title>
                <v-card-text>
                    <v-layout>
                        <v-flex xs12 sm8 md8 lg8>
                            <v-layout>
                                <v-flex xs12 md4>
                                    <v-text-field label="عنوان آزمون" v-model="ExamData.title" style="padding:10px"></v-text-field>
                                </v-flex>
                                <v-flex xs12 md4>
                                    <v-text-field label="هزینه شرکت در آزمون" v-model="ExamData.cost" style="padding:10px"></v-text-field>
                                </v-flex>
                                <v-flex xs12 md4>
                                    
                                </v-flex>
                            </v-layout>

                            <v-layout row wrap>
                                
                                <v-flex xs12 md4>
                                    <v-text-field label="انتخاب درس" v-model="ExamData.selectCourse" style="padding:10px"></v-text-field>
                                </v-flex>
                                <v-flex xs12 md4>
                                    <v-text-field label="تعداد سوالات نهایی" v-model="ExamData.numberofFinalQuestion" style="padding:10px"></v-text-field>
                                </v-flex>
                                <v-flex xs12 md4>
                                    <v-text-field  type="number" mask="##" label="مدت آزمون" v-model="ExamData.duration" style="padding:10px"></v-text-field>
                                </v-flex>
                            </v-layout>
                            <v-layout row wrap>
                                <v-flex xs12 sm6 md3>
                                    <v-text-field label="تاریخ شروع آزمون" v-model="ExamData.startDate" style="padding:10px"></v-text-field>
                                    <!--<date-picker v-model="date"></date-picker>-->
                                </v-flex>
                                <v-flex xs12 md3>
                                    <v-text-field label="تاریخ پایان آزمون" v-model="ExamData.endDate" style="padding:10px"></v-text-field>
                                </v-flex>
                                <v-flex xs12 md3>
                                    <v-text-field label="ساعت شروع آزمون" v-model="ExamData.startTime" style="padding:10px"></v-text-field>
                                </v-flex>
                                <v-flex xs12 md3>
                                    <v-text-field label="ساعت پایان آزمون" v-model="ExamData.endTime" style="padding:10px"></v-text-field>
                                </v-flex>
                            </v-layout>
                            <v-layout row wrap>
                                <v-flex xs12 md6 style="padding:15px">
                                    <v-textarea label="پیغام قبل از آزمون" v-model="ExamData.beforeMessage"></v-textarea>
                                </v-flex>
                                
                                <v-flex xs12 md6 style="padding:15px">
                                    <v-textarea label="متن پایانی زیر سوالات" v-model="ExamData.afterMessage"></v-textarea>
                                </v-flex>
                                
                            </v-layout>
                        </v-flex>
                        <v-flex xs12 sm4 md4 lg4 >
                            <v-layout>
                                <v-flex shrink class="dir-switch" style="direction:rtl;">
                                    <v-switch color="green" v-model="ExamData.examStatus" :label="(ExamData.examStatus)? 'وضعیت آزمون: فعال':'وضعیت آزمون: غیر فعال'" ></v-switch>
                                </v-flex>
                            </v-layout>
                            <v-layout>
                                <v-flex shrink class="dir-switch" style="direction:rtl;">
                                    <v-switch color="green" v-model="ExamData.multiCourse" :label="(ExamData.multiCourse)? 'آزمون چند درسه است؟ بلی':'آزمون چند درسه است؟ خیر'"></v-switch>
                                </v-flex>
                            </v-layout>
                            <v-layout>
                                <v-flex shrink class="dir-switch" style="direction:rtl;">
                                    <v-switch color="green" v-model="ExamData.negativeScore" :label="(ExamData.negativeScore)? 'آزمون نمره منفی دارد؟ بلی':'آزمون نمره منفی دارد؟ خیر'"></v-switch>
                                </v-flex>
                            </v-layout>
                            <v-layout>
                                <v-flex shrink class="dir-switch" style="direction:rtl;">
                                    <v-switch color="green" v-model="ExamData.randomShow" :label="(ExamData.randomShow)? 'آیا سوالات به صورت رندم نمایش داده شود؟ بلی':'آیا سوالات به صورت رندم نمایش داده شود؟ خیر'"></v-switch>
                                </v-flex>
                            </v-layout>
                            <v-layout>
                                <v-flex shrink class="dir-switch" style="direction:rtl;">
                                    <v-switch color="green" v-model="ExamData.specificTime" :label="(ExamData.specificTime)? 'آیا آزمون در مقطع خاصی فعال شود؟ بلی':'آیا آزمون در مقطع خاصی فعال شود؟ خیر'"></v-switch>
                                </v-flex>
                            </v-layout>
                        </v-flex>
                    </v-layout>
                    <v-divider></v-divider>
                   <v-layout row wrap>
                        <v-flex md2 lg2 xs2 sm2>
                            <v-btn @click="addExam" color="success">ذخیره</v-btn>
                        </v-flex>
                        <v-flex md2 lg2 xs2 sm2>
                            <v-btn >بازگشت</v-btn>
                        </v-flex>
                        <v-flex md8 lg8 xs8 sm8>
                        </v-flex>
                    </v-layout>
                </v-card-text>
            </v-card>
        </v-flex>
    </v-layout>
    
</v-container>
</template>
<script>
export default {
    data: () => ({
        axios:require("axios"),
        ExamData:{
            title:'',
            cost:'',
            examStatus:'',
            startDate:'',
            endDate:'',
            startTime:'',
            endTime:'',
            specificTime:'',
            randomShow:'',
            negativeScore:'',
            beforeMessage:'',
            afterMessage:'',
            duration:'',
            multiCourse:'',
            selectCourse:'',
            numberofFinalQuestion:'',
            
        },
        response:null,
        date: '1397/02/02',
        rules :{
            
            //maxDuration: v >= v.length <= 2 || 'طول بیشتر از 2 مجار نمیباشد'
        }  
    }),
    methods:{
        addExam(){
            this.axios.post("http://localhost:9090/exam",JSON.stringify(this.ExamData)).then(res=>{this.response = res.data.messages})
           
        }
    },
    mounted(){

    },
    computed:{

    },
    components:{
        //DatePicker: VuePersianDatetimePicker
    }
}
</script>

<style>
.dir-switch{
    margin-top: 15px;
    color:#FFF;
}
</style>
