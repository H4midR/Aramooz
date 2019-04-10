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
                    <v-flex xs12 md4>
                        <v-text-field label="عنوان آزمون" v-model="title" style="padding:10px"></v-text-field>
                    </v-flex>
                    <v-flex xs12 md4>
                        <v-text-field label="هزینه شرکت در آزمون" v-model="cost" style="padding:10px"></v-text-field>
                    </v-flex>
                    <v-flex xs12 md4>
                        <div>وضعیت آزمون</div>
                        <v-radio-group v-model="status" row>
                            
                            <v-radio label="فعال" value="active"></v-radio>
                            <v-radio label="غیرفعال" value="deactive"></v-radio>
                        </v-radio-group>
                    </v-flex>
                </v-layout>
                <v-layout row wrap>
                    <v-flex xs12 md3>
                        <div>آزمون چند درسه است؟</div>
                        <v-radio-group v-model="multiCourse" row>
                            
                            <v-radio label="خیر" value="yes"></v-radio>
                            <v-radio label="بلی" value="no"></v-radio>
                        </v-radio-group>
                    </v-flex>
                    <v-flex xs12 md3>
                        <v-text-field label="انتخاب درس" v-model="selectCourse" style="padding:10px"></v-text-field>
                    </v-flex>
                    <v-flex xs12 md3>
                        <v-text-field label="تعداد سوالات نهایی" v-model="numberofFinalQuestion" style="padding:10px"></v-text-field>
                    </v-flex>
                    <v-flex xs12 md3>
                        <v-text-field label="مدت آزمون" v-model="duration" style="padding:10px"></v-text-field>
                    </v-flex>
                </v-layout>
                
                <v-layout row wrap>
                    <v-flex xs12 md4>
                    <div>آزمون نمره منفی دارد؟</div>
                        <v-radio-group label="" v-model="negative" row>
                            <v-radio name="negative" label="خیر" :value="0"></v-radio>
                            <v-radio name="negative" label="بلی" :value="1"></v-radio>                
                        </v-radio-group>
                    </v-flex>

                    <v-flex xs12 md4>
                        <div>آیا سوالات به صورت رندم (تصادفی( نمایش داده شود؟</div>
                        <v-radio-group v-model="randomShow" row>
                            
                            <v-radio label="خیر" value="yes"></v-radio>
                            <v-radio label="بلی" value="no"></v-radio>
                        </v-radio-group>
                    </v-flex>
                    <v-flex xs12 md4>
                        <div>آیا آزمون در مقطع خاصی فعال شود؟</div>
                        <v-radio-group v-model="specTime" row>
                            
                            <v-radio label="خیر" value="yes"></v-radio>
                            <v-radio label="بلی" value="no"></v-radio>
                        </v-radio-group>
                    </v-flex>
                </v-layout>
                <v-layout row wrap>
                    <v-flex xs12 sm6 md3>
                    <v-menu ref="menu" v-model="menu" :close-on-content-click="false" :nudge-right="40" :return-value.sync="date" lazy transition="scale-transition" offset-y full-width min-width="290px">
                        <template slot:activator="{ on }">
                            <v-text-field v-model="date" label="Picker in menu" prepend-icon="event" readonly v-on="on"></v-text-field>
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
                <v-layout row wrap>
                    <v-flex md2 lg2 xs2 sm2>
                        <v-btn @click="addExam" color="success">ذخیره</v-btn>
                    </v-flex>
                    <v-flex md2 lg2 xs2 sm2>
                        <v-btn @click="ReturnE">بازگشت</v-btn>
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
        title:null,
        cost:null,
        status:null,
        startDate:null,
        endDate:null,
        startTime:null,
        endTime:null,
        specTime:null,
        randomShow:null,
        negative:null,
        beforeMessage:null,
        afterMessage:null,
        duration:null,
        multiCourse:null,
        selectCourse:null,
        numberofFinalQuestion:null,
        ReturnE:null,
        response:null,
        date: new Date().toISOString().substr(0, 10),
        menu: false,
        modal: false,
        menu2: false
      
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

