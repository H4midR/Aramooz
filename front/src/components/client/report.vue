<template>
    <v-container fluid align-center justify-space-around row wrap>
       <v-layout>
           <h2>نتیجه آزمون</h2>
       </v-layout>
       
       <v-layout row wrap>
            <v-flex xs12 sm12 lg12>
                <v-card> 
                    <v-card-title primary-title> 
                        <span>آزمون ادبیات و زبان فارسی</span>
                    </v-card-title>
                    <v-card-text >
                            <v-layout row wrap >
                                <v-flex lg-1 border><v-layout border><v-icon>done</v-icon>صحیح</v-layout><v-layout >10</v-layout></v-flex>
                                <v-flex lg-1 border><v-layout border><v-icon>clear</v-icon>غلط</v-layout><v-layout >15</v-layout></v-flex>
                                <v-flex lg-1 border><v-layout border><v-icon>trip_origin</v-icon>بدون پاسخ</v-layout><v-layout >0</v-layout></v-flex>
                                <v-flex lg-1 border><v-layout border><v-icon>help</v-icon>سوال های آزمون</v-layout><v-layout >25</v-layout></v-flex>
                                <v-flex lg-2 border><v-layout border><v-icon>remove_circle</v-icon>نتیجه با نمره منفی</v-layout><v-layout >5%</v-layout></v-flex>
                                <v-flex lg-2 border><v-layout border><v-icon>add_circle</v-icon>نتیجه بدون نمره منفی</v-layout><v-layout >1%</v-layout></v-flex>
                                <v-flex lg-2 border><v-layout border><v-icon>flip_to_front</v-icon>گزارش آزمون</v-layout><v-layout ><v-btn class="btn primary">دانلود کارنامه</v-btn></v-layout></v-flex>
                                <v-flex lg-2 border><v-layout border><v-icon>description</v-icon>نمایش پاسخنامه</v-layout><v-layout ><v-btn class="btn success">نمایش پاسخنامه</v-btn></v-layout></v-flex>
                            </v-layout>
                           
                    </v-card-text>
                   
                </v-card>
            </v-flex>
        </v-layout> 
        
        <v-layout row wrap>
            <v-flex lg12>
                <v-card>
                    <v-card-title primary>
                        <span>برای نمایش تحلیل عملکردتان در هر سوال می توانید روی شماره سوال کلیک کنید.</span>
                    </v-card-title>
                    <v-card-title>    
                        <v-layout offset-lg2>
                            <v-flex lg3>پاسخ صحیح داوطلب<span class="border lboption greendark">1</span></v-flex>
                            <v-flex lg2>پاسخ نادرست     <span class="border lboption red"      >1</span></v-flex>
                            <v-flex lg2>پاسخ صحیح سوال  <span class="border lboption green"    >1</span></v-flex>
                        </v-layout>
                    </v-card-title>
                    <v-card-text>
                        <v-data-table :headers="headers" :items="desserts" class="elevation-1" >
                            <template slot="items" slot-scope="props">
                                <td>{{ props.item.name }}</td>
                                <td class="text-xs-right">{{ props.item.calories }}</td>
                                <td class="text-xs-right">{{ props.item.fat }}</td>
                                <td class="text-xs-right">{{ props.item.carbs }}</td>
                                <td class="text-xs-right">{{ props.item.protein }}</td>
                                <td class="text-xs-right">{{ props.item.iron }}</td>
                            </template>
                        </v-data-table>
                    </v-card-text>
                </v-card>
            </v-flex>
        </v-layout>
        <v-layout>
            <v-flex xs12 sm12 lg12>
            <v-card>
                <v-card-title>
                    <v-flex lg6>
                        <v-layout>
                            <v-flex text-center>زمان پاسخ دهی</v-flex>
                        </v-layout>
                        <v-layout>
                            <v-flex lg6>
                                <v-progress-circular :rotate="-90" :size="100" :width="15" :value="80" color="primary" >80</v-progress-circular>
                            </v-flex>
                            <v-flex lg6>
                                <v-progress-circular :rotate="-90" :size="100" :width="15" :value="24" color="deep-orange lighten-2" >24</v-progress-circular>
                            </v-flex>
                        </v-layout>
                    </v-flex>
                    <v-flex lg6>
                        <v-layout >
                            <v-flex text-center>نتیجه آزمون</v-flex>
                        </v-layout>
                        <v-layout>
                            <v-flex lg6>
                                <v-progress-circular :rotate="-90" :size="100" :width="15" :value="10" color="lime" >10</v-progress-circular>
                            </v-flex>
                            <v-flex lg6>
                                <v-progress-circular :rotate="-90" :size="100" :width="15" :value="100" color="indigo darken-2" >100</v-progress-circular>
                            </v-flex>
                        </v-layout>
                    </v-flex>
            </v-card-title>
            </v-card>
            </v-flex>
        </v-layout>
        <v-layout>
            <v-flex>
                <v-card>
                    <v-layout>
                        <v-flex><v-text>نمودار تجزیه و تحلیل آزمون </v-text></v-flex>
                    </v-layout>
                    <v-layout row wrap >
                        <v-flex lg4>
                            <v-progress-circular :rotate="-90" :size="100" :width="40" :value="99" color="lime lighten-1">99</v-progress-circular>
                        </v-flex>
                        <v-flex lg8>
                            <v-sparkline
                                :value="value"
                                :gradient="gradient"
                                :smooth="radius || false"
                                :padding="padding"
                                :line-width="width"
                                :stroke-linecap="lineCap"
                                :gradient-direction="gradientDirection"
                                auto-draw
                                ></v-sparkline>

                        </v-flex>
                    </v-layout>
                </v-card>
            </v-flex>
        </v-layout>
        
        <v-layout>
            <v-flex>
                <v-card>
                    <v-card-title>
                        <ul >
                            <li><v-icon>info</v-icon>شما در موضوع زبان فارسی نیاز به تمرین بیشتری دارید، آزمون های موضوعی زبان فارسی را به شما توصیه می کنیم.</li>
                            <li><v-icon>info</v-icon>شما در موضوع زبان فارسی نیاز به تمرین بیشتری دارید، آزمون های موضوعی زبان فارسی را به شما توصیه می کنیم.</li>
                            <li><v-icon>info</v-icon>شما در موضوع زبان فارسی نیاز به تمرین بیشتری دارید، آزمون های موضوعی زبان فارسی را به شما توصیه می کنیم.</li>
                            <li><v-icon>info</v-icon>شما در موضوع زبان فارسی نیاز به تمرین بیشتری دارید، آزمون های موضوعی زبان فارسی را به شما توصیه می کنیم.</li>
                        </ul>
                    </v-card-title>
                     <v-card-actions>
                        <v-btn color="pink darken-1">بازگشت</v-btn>
                        
                    </v-card-actions>
                </v-card>
            </v-flex>
        </v-layout>
        
    </v-container> 
   
</template>
<script>
const gradients = [
    ['#222'],
    ['#42b3f4'],
    ['red', 'orange', 'yellow'],
    ['purple', 'violet'],
    ['#00c6ff', '#F0F', '#FF0'],
    ['#f72047', '#ffd200', '#1feaea']
  ]
    export default {
        data() {return{
            axios: require('axios'),
            width: 2,
            radius: 10,
            padding: 8,
            lineCap: 'round',
            gradient: gradients[5],
            value: [0, 2, 5, 9, 5, 10, 3, 5, 0, 0, 1, 8, 2, 9, 0],
            gradientDirection: 'top',
            gradients
        }},
        methods:{
        }
    }
</script>



<style>
.green {
    background: #8BC34A;
    color: #8BC34A !important;
}
.red {
    background: #F22613;
    color: #F22613 !important;
}
.greendark {
   background: #1b5e20;
   color: #1b5e20 !important;
}
.lboption {
    width: 29px !important;
    display: inline;
    padding-right: 27px;
}
.v-card{
    max-width: 1000px;
}
li {list-style-type: none;}
.border{
    border: 1px solid grey;
}
</style>




