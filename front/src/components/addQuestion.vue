<template>
    <v-container>
        <v-layout class="add-question-form">
            <v-flex xs12 md2 lg3></v-flex>
            <v-flex xs12 md8 lg6>
                <v-card>
                    <v-toolbar class="light-blue lighten-1" >
                        <v-icon color="white">help</v-icon>
                        &nbsp;<span class="white--text " >افزودن سوال</span>
                    </v-toolbar>
                    <v-card-title primary-title>
                        <v-layout>
                            <v-flex xs12 md2>
                                <v-text-field style="margin-left:10px;" label="شماره سوال" required ></v-text-field>
                            </v-flex>
                            <v-flex xs12 md10>
<v-text-field :counter="100" label="عنوان سوال" required></v-text-field>
                            </v-flex>
                        </v-layout>
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
    edit:null,
    emptyChoices:{
            title:'',
            value:0
        },
    emptyQuestion:{
        title:'',
        rank:0,
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
</style>
