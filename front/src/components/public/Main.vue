<template>
    <v-app>
        <v-toolbar>
            <v-toolbar-title
            class="tertiary--text font-weight-light" v-text="title">
            </v-toolbar-title>

            <v-spacer />

            <v-toolbar-items>
                <v-menu offset-y content-class="dropdown-menu" transition="slide-y-transition" v-for="menu in menus" :key="menu.title"> 
                    <v-btn flat slot="activator" >
                        <v-icon right v-text="menu.icon"></v-icon>
                        {{menu.title}}
                    </v-btn>
                    <v-card v-if="menu.child">
                        <v-list dense>
                        <v-list-tile v-for="child in menu.child" :key="child.title" router :to="child.url" >
                            <v-icon right v-text="child.icon"></v-icon>
                            <v-list-tile-title v-text="child.title"/>
                        </v-list-tile>
                        </v-list>
                    </v-card>
                </v-menu>
            </v-toolbar-items>
        </v-toolbar>


        <v-card>
            <v-card-title>
                <h2>لیست آزمون ها</h2>
            </v-card-title>
            <v-container fluid grid-list-md>
            <v-layout row wrap>
                <v-flex v-for="exam in Exams" :key="exam.title" v-bind="{ [`md${exam.flex}`]: true }">
                <v-card>
                    <v-img :src="exam.src" height="250px" />

                    <v-card-actions>
                        <span class=" black--text" v-text="exam.title"></span>
                        <v-spacer></v-spacer>

                        <v-tooltip right>
                            <v-btn icon @click="exam.show = !exam.show" slot="activator">
                                <v-icon>{{ exam.show ? 'keyboard_arrow_up' : 'keyboard_arrow_down' }}</v-icon>
                            </v-btn>
                            <span>جزئیات آزمون</span>
                        </v-tooltip>
                        
                        <v-tooltip top>
                            <v-btn icon router to="startExam" slot="activator" color="pink white--text">
                                <v-icon>link</v-icon>
                            </v-btn>
                            <span>شروع آزمون</span>
                        </v-tooltip>
                        
                    </v-card-actions>

                    <v-slide-x-transition>
                        <v-card-text v-show="exam.show" v-text="exam.desc" />
                    </v-slide-x-transition>
                </v-card>
                </v-flex>
            </v-layout>
            </v-container>
        </v-card>


    </v-app>
</template>

<script>
export default {
    data(){
        return{
            title: "آزمون یار", //Page or Site Title
            menus:[
                {
                    title:"آشنایی با ما",
                    icon:"account_balance",
                    url:"",
                    child:[
                        {
                            title:"تماس با ما",
                            icon:"call",
                            url:"",
                        },
                        {
                            title:"درباره ما",
                            icon:"business",
                            url:"",
                        },
                    ],
                }, //Menu Item
                {
                    title:"ورود و عضویت",
                    icon:"person",
                    url:"",
                    child:[
                        {
                            title:"ورود به سایت",
                            icon:"mdi-login",
                            url:"login",
                        },
                        {
                            title:"عضویت در سایت",
                            icon:"mdi-account-plus",
                            url:"signup",
                        },
                    ],
                }//Menu Item
            ], //Menus
            Exams: [
                { title: 'آزمون یک',    desc:"توضیحات آزمون به شرح ذیل می باشد. شما می توانید در این آزمون شرکت کنید، به شرط اینکه بچه خوبی باشید.",    src: require("@/assets/images/test1.jpg"), flex: 3,uid: "0x1",  show:false, },
                { title: 'آزمون دو',    desc:"توضیحات آزمون به شرح ذیل می باشد. شما می توانید در این آزمون شرکت کنید، به شرط اینکه بچه خوبی باشید.",    src: require("@/assets/images/test2.jpg"), flex: 3,uid: "0x2",  show:false, },
                { title: 'آزمون سه',    desc:"توضیحات آزمون به شرح ذیل می باشد. شما می توانید در این آزمون شرکت کنید، به شرط اینکه بچه خوبی باشید.",    src: require("@/assets/images/test3.jpg"), flex: 3,uid: "0x3",  show:false, },
                { title: 'آزمون چهار',  desc:"توضیحات آزمون به شرح ذیل می باشد. شما می توانید در این آزمون شرکت کنید، به شرط اینکه بچه خوبی باشید.",    src: require("@/assets/images/test4.jpg"), flex: 3,uid: "0x4",  show:false, },
            ], //Exams

        }
    }
}
</script>

<style>

</style>
