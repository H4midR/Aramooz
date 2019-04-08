<template>
    <v-container>
        <v-layout class="add_question">
            <v-flex xs12 md2 lg3></v-flex>
            <v-flex xs12 md8 lg6>
                <v-card>
                    <v-toolbar class="light-blue lighten-1" >
                        <v-icon color="white">help</v-icon>
                        &nbsp;<span class="white--text " >افزودن سوال</span>                       
                    </v-toolbar>
                    <v-card-title primary-title>
                        <v-text-field :counter="100" label="عنوان سوال" required></v-text-field>
                    </v-card-title>
                    <v-card-text>
                        <choices :choicesList="choicesArray"></choices>
                    </v-card-text>
                </v-card>
            </v-flex>
            <v-flex xs12 md2 lg3></v-flex>
        </v-layout>
    </v-container>
</template>

<script>
import choices from '@/components/addQuestion/choices'
export default {
data(){
return{
    emptyQuestion:{
        title:'',
        rank:0,
        choices:[
            {
                title:'',
                value:0
            }
        ]
    }    
}
},
components:{
    choices
}
}
</script>

<style>
.add_question .v-list{
    background:none;
}
</style>
