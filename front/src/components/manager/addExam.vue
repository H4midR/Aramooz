<template>
<v-container justify-center>
    <v-layout row wrap >
        <v-flex xs12 sm11 md10 lg8>
            <v-card>
                <v-toolbar class="primary"> 
                    <v-icon color=""></v-icon>
                    <span style="color:#FFF">افزودن آزمون</span>
                    <v-spacer></v-spacer>
                </v-toolbar>
                <v-card-title>
                    <v-layput>
                        <v-text>آزمون / افزودن آزمون</v-text>
                    </v-layput>
                </v-card-title>
                <v-card-text>
                    <v-layout>
                        <v-flex xs12 sm8 md8 lg8>
                            <v-layout>
                                <v-flex xs12 md4>
                                    <v-text-field label="عنوان آزمون" v-model="title" style="padding:10px"></v-text-field>
                                </v-flex>
                                <v-flex xs12 md4>
                                    <v-text-field label="هزینه شرکت در آزمون" v-model="cost" style="padding:10px"></v-text-field>
                                </v-flex>
                                <v-flex xs12 md4>
                                    
                                </v-flex>
                            </v-layout>

                            <v-layout row wrap>
                                
                                <v-flex xs12 md4>
                                    <v-text-field label="انتخاب درس" v-model="selectCourse" style="padding:10px"></v-text-field>
                                </v-flex>
                                <v-flex xs12 md4>
                                    <v-text-field label="تعداد سوالات نهایی" v-model="numberofFinalQuestion" style="padding:10px"></v-text-field>
                                </v-flex>
                                <v-flex xs12 md4>
                                    <v-text-field label="مدت آزمون" v-model="duration" style="padding:10px"></v-text-field>
                                </v-flex>
                            </v-layout>
                            <v-layout row wrap>
                                <v-flex xs12 sm6 md3>
                                    <v-menu
                                        ref="menu"
                                        v-model="menu"
                                        :close-on-content-click="false"
                                        :nudge-right="40"
                                        :return-value.sync="date"
                                        lazy
                                        transition="scale-transition"
                                        offset-y
                                        full-width
                                        min-width="290px"
                                    >
                                        <template v-slot:activator="{ on }">
                                        <v-text-field
                                            v-model="date"
                                            label="Picker in menu"
                                            prepend-icon="event"
                                            readonly
                                            v-on="on"
                                        ></v-text-field>
                                        </template>
                                        <v-date-picker v-model="date" no-title scrollable>
                                        <v-spacer></v-spacer>
                                        <v-btn flat color="primary" @click="menu = false">Cancel</v-btn>
                                        <v-btn flat color="primary" @click="$refs.menu.save(date)">OK</v-btn>
                                        </v-date-picker>
                                    </v-menu>
                                </v-flex>
                                <v-flex xs12 md3>
                                    <v-text-field label="تاریخ پایان آزمون" v-model="endDate" style="padding:10px"></v-text-field>
                                </v-flex>
                                <v-flex xs12 md3>
                                    <v-text-field label="ساعت شروع آزمون" v-model="startTime" style="padding:10px"></v-text-field>
                                </v-flex>
                                <v-flex xs12 md3>
                                    <v-text-field label="ساعت پایان آزمون" v-model="endTime" style="padding:10px"></v-text-field>
                                </v-flex>
                            </v-layout>
                            <v-layout row wrap>
                                <v-flex xs12 md6 style="padding:15px">
                                    <v-textarea label="پیغام قبل از آزمون" v-model="beforeMessage"></v-textarea>
                                </v-flex>
                                
                                <v-flex xs12 md6 style="padding:15px">
                                    <v-textarea label="متن پایانی زیر سوالات" v-model="afterMessage"></v-textarea>
                                </v-flex>
                                
                            </v-layout>
                        </v-flex>
                        <v-flex xs12 sm4 md4 lg4 >
                            <v-layout>
                                <v-flex shrink class="dir-switch" style="direction:rtl;">
                                    <v-switch color="green" v-model="examStatus" :label="(examStatus)? 'وضعیت آزمون: فعال':'وضعیت آزمون: غیر فعال'" value="false"></v-switch>
                                </v-flex>
                            </v-layout>
                            <v-layout>
                                <v-flex shrink class="dir-switch" style="direction:rtl;">
                                    <v-switch color="green" v-model="multiCourse" :label="(multiCourse)? 'آزمون چند درسه است؟ بلی':'آزمون چند درسه است؟ خیر'" value="false"></v-switch>
                                </v-flex>
                            </v-layout>
                            <v-layout>
                                <v-flex shrink class="dir-switch" style="direction:rtl;">
                                    <v-switch color="green" v-model="negativeScore" :label="(negativeScore)? 'آزمون نمره منفی دارد؟ بلی':'آزمون نمره منفی دارد؟ خیر'" value="false"></v-switch>
                                </v-flex>
                            </v-layout>
                            <v-layout>
                                <v-flex shrink class="dir-switch" style="direction:rtl;">
                                    <v-switch color="green" v-model="randomShow" :label="(randomShow)? 'آیا سوالات به صورت رندم نمایش داده شود؟ بلی':'آیا سوالات به صورت رندم نمایش داده شود؟ خیر'" value="false"></v-switch>
                                </v-flex>
                            </v-layout>
                            <v-layout>
                                <v-flex shrink class="dir-switch" style="direction:rtl;">
                                    <v-switch color="green" v-model="specificTime" :label="(specificTime)? 'آیا آزمون در مقطع خاصی فعال شود؟ بلی':'آیا آزمون در مقطع خاصی فعال شود؟ خیر'" value="false"></v-switch>
                                </v-flex>
                            </v-layout>
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
        title:null,
        cost:null,
        examStatus:false,
        startDate:null,
        endDate:null,
        startTime:null,
        endTime:null,
        specificTime:false,
        randomShow:false,
        negativeScore:false,
        beforeMessage:null,
        afterMessage:null,
        duration:null,
        multiCourse:false,
        selectCourse:null,
        numberofFinalQuestion:null,
        ReturnE:null,
        response:null,
        date: new Date().toISOString().substr(0, 10),
        menu: false,
        modal: false,      
    }),
    methods:{
        addExam(){
            this.axios.post("",JSON.stringify({
                title:this.title,
                cost:this.cost,
                status:this.status
            })).then(res=>{this.response = res.data.messages})
           
        }
    }
}
</script>

<style>
.dir-switch{
    margin-top: 15px;
    color:#FFF;
}
</style>
