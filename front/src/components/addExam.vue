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
        <v-flex >
            <v-btn @click="addExam"></v-btn>
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

