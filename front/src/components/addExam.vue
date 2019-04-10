<template>
<v-container>

    <v-layout row wrap>
        <v-flex xs12 md4>
            <v-text-field label="عنوان آزمون" v-model="title"></v-text-field>
        </v-flex>
        <v-flex xs12 md4>
            <v-text-field label="هزینه شرکت در آزمون" v-model="cost"></v-text-field>
        </v-flex>
        <v-flex xs12 md4>
            <v-radio-group v-model="status" row>
                <span>وضعیت آزمون</span>
                <v-radio label="فعال" value="active"></v-radio>
                <v-radio label="غیرفعال" value="deactive"></v-radio>
            </v-radio-group>
        </v-flex>
    </v-layout>
    <v-layout row wrap>
        <v-flex xs12 md3>
            <v-radio-group v-model="multiCourse" row>
                <span>آزمون چند درسه است؟</span>
                <v-radio label="خیر" value="yes"></v-radio>
                <v-radio label="بلی" value="no"></v-radio>
            </v-radio-group>
        </v-flex>
        <v-flex xs12 md3>
            <v-text-field label="انتخاب درس" v-model="selectCourse"></v-text-field>
        </v-flex>
        <v-flex xs12 md3>
            <v-text-field label="تعداد سوالات نهایی" v-model="numberofFinalQuestion"></v-text-field>
        </v-flex>
        <v-flex xs12 md3>
            <v-text-field label="مدت آزمون" v-model="duration"></v-text-field>
        </v-flex>
    </v-layout>
    <v-layout row wrap>
        <v-flex xs12 md4>
            <v-radio-group v-model="negative" row>
                <span>آزمون نمره منفی دارد؟</span>
                <v-radio label="خیر" value="yes"></v-radio>
                <v-radio label="بلی" value="no"></v-radio>
            </v-radio-group>
        </v-flex>
        <v-flex xs12 md4>
            <v-radio-group v-model="randomShow" row>
                <span>آیا سوالات به صورت رندم (تصادفی( نمایش داده شود؟</span>
                <v-radio label="خیر" value="yes"></v-radio>
                <v-radio label="بلی" value="no"></v-radio>
            </v-radio-group>
        </v-flex>
        <v-flex xs12 md4>
            <v-radio-group v-model="specTime" row>
                <span>آیا آزمون در مقطع خاصی فعال شود؟</span>
                <v-radio label="خیر" value="yes"></v-radio>
                <v-radio label="بلی" value="no"></v-radio>
            </v-radio-group>
        </v-flex>
    </v-layout>
    <v-layout row wrap>
        <v-flex xs12 md3>
            <v-text-field label="تاریخ شروع آزمون" v-model="startDate" box></v-text-field>
        </v-flex>
        <v-flex xs12 md3>
            <v-text-field label="تاریخ پایان آزمون" v-model="endDate"></v-text-field>
        </v-flex>
        <v-flex xs12 md3>
            <v-text-field label="ساعت شروع آزمون" v-model="startTime"></v-text-field>
        </v-flex>
        <v-flex xs12 md3>
            <v-text-field label="ساعت پایان آزمون" v-model="endTime"></v-text-field>
        </v-flex>
    </v-layout>
    <v-layout row wrap>
        <v-flex xs12 md5>
            <v-textarea label="پیغام قبل از آزمون" v-model="beforeMessage"></v-textarea>
        </v-flex>
        <v-flex md1></v-flex>
        <v-flex xs12 md5>
            <v-textarea label="متن پایانی زیر سوالات" v-model="afterMessage"></v-textarea>
        </v-flex>
        <v-flex md1></v-flex>
    </v-layout>
    <v-layout row wrap>
    </v-layout>
    <v-layout row wrap>
        <v-flex md1>
            <v-btn @click="addExam" color="success">ذخیره</v-btn>
        </v-flex>
        <v-flex md1>
            <v-btn @click="ReturnE">بازگشت</v-btn>
        </v-flex>
        <v-flex md10>
        </v-flex>
    </v-layout>

</v-container>
</template>
<script>
export default {
    data(){return{
        axios:require("axios"),
        title:null,
        cost:null,
        status:null,
        response:null,
    }},
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

