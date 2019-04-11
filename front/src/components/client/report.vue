<template>
    <v-container grid-list-md text-xs-center >
        
        <h2 style="color:red;">نتیجه آزمون</h2><br>
        <v-layout fluid align-center justify-space-around row wrap>
            <v-flex xs12>
                <v-card color="white">
                    <v-card-text class="px-0">آزمون ادبیات فارسی</v-card-text>
                </v-card>
            </v-flex>
           
            <v-flex xs1>
                <v-card color="white">
                        <v-card-text class="px-0 myFont">صحیح <v-icon>check</v-icon></v-card-text>
                </v-card>
                <v-card>
                     <v-card-text class="px-0 myFont">11</v-card-text>
                 </v-card>
            </v-flex>
             <v-flex xs1>
                <v-card color="white">
                        <v-card-text class="px-0  myFont">غلط <v-icon>close</v-icon></v-card-text>
                </v-card>
                <v-card>
                     <v-card-text class="px-0 myFont">31</v-card-text>
                </v-card>
            </v-flex>
            <v-flex xs2>
                <v-card color="white">
                         <v-card-text class="px-0 myFont">بدون پاسخ <v-icon>error_outline</v-icon></v-card-text>
                </v-card>
                <v-card>
                     <v-card-text class="px-0 myFont">8</v-card-text>
                 </v-card>
            </v-flex>
            <v-flex xs2>
                <v-card color="white">
                        <v-card-text class="px-0 myFont">تعداد سوالات <v-icon>help_outline</v-icon></v-card-text>
                 </v-card>
                 <v-card>
                     <v-card-text class="px-0 myFont">50 </v-card-text>
                 </v-card>
            </v-flex>
            <v-flex xs2>
                <v-card color="white">
                        <v-card-text class="px-0">نتیجه <v-icon>book</v-icon></v-card-text>
                </v-card>
                <v-card>
                     <v-card-text class="px-0 myFont">30</v-card-text>
                </v-card>
            </v-flex>
            <v-flex xs2>
              <v-card color="white">
                 <v-card-text class="px-0 myFont">گزارش آزمون <v-icon>assessment</v-icon></v-card-text>
            </v-card>
              <v-card>
                <v-btn class="px-0 myFont" color="success">دانلود</v-btn>   
            </v-card>
            </v-flex>
              <v-flex xs2>
                <v-card color="white">
                  <v-card-text class="px-0 myFont">نمایش پاسخنامه <v-icon>assignment</v-icon></v-card-text>
                </v-card>
                <v-card color="white" style="color:white;">
                  <v-btn @click="show" class="px-0 myFont" color="primary"> رویت </v-btn>
              </v-card>
            </v-flex>
        </v-layout>

        <br>
        
        <!-- <v-layout align-center justify-space-between row fill-height>
                <v-flex xs12>
                    <v-card>
                        <v-card-text><h3>.برای نمایش تحلیل عملکردتان در هر سوال می توانید روی شماره سوال کلیک کنید</h3><br></v-card-text>
                        <v-card color="white">
                        <v-card-text class="px-0 myFont">
                            <p style="direction:rtl;margin-right:10px;padding:10px;text-align:center;">
                        			پاسخ صحیح داوطلب  <span class="greendark border lboption" style="color: green !important;"> 1 </span>
                        			&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
                        			پاسخ نادرست <span class="red border lboption" style="color: #F22613 !important;"> 1 </span>
                        			&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
                        			پاسخ صحیح  سوال <span class="green border lboption" style="color: #8BC34A !important;"> 1 </span>
                        			&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
                        		</p>

                        </v-card-text>
                    </v-card>
                    </v-card>
                    
                </v-flex>
        </v-layout> -->
    </v-container>
</template>

<script>
    export default {
        data() {return{
            axios: require('axios'),
            parameter: null,

        }},

        methods:{

        }
    }
</script>

<style scoped>
    .myFont {
        font-size: 12px;
    }
    .myFont2 {
        font-size: 11px;
    }
</style>

