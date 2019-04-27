<template>
    <v-container  >
       <v-layout >
           <h2>نتیجه آزمون</h2>
       </v-layout>
       <v-layout>
           <table border justify_align >
               <thead><tr><span>آزمون ادبیات و زبان فارسی</span></tr></thead>
               <thead>
                   <tr>
                       <td> <v-icon>done</v-icon>صحیح</td>
                       <td> <v-icon>clear</v-icon>غلط</td>
                       <td> <v-icon>trip_origin</v-icon>بدون پاسخ</td>
                       <td> <v-icon>help</v-icon>سوال های آزمون</td>
                       <td> <v-icon>remove_circle</v-icon>نتیجه با نمره منفی</td>
                       <td> <v-icon>add_circle</v-icon>نتیجه بدون نمره منفی</td>
                       <td> <v-icon>flip_to_front</v-icon>گزارش آزمون</td>
                       <td> <v-icon>description</v-icon>نمایش پاسخنامه</td>
                   </tr>
               </thead>
               <tbody>
                   <tr>
                       <td>10</td>
                       <td>15</td>
                       <td>0</td>
                       <td>25</td>
                       <td>0%</td>
                       <td>1%</td>
                       <td><v-btn class="btn primary">دانلود کارنامه</v-btn></td>
                       <td><v-btn class="btn success">نمایش پاسخنامه</v-btn></td>

                   </tr>
               </tbody>
          
           </table>
        </v-layout> 
        <v-layout >
            <table border mt-30>
                <thead>
                    <tr><span>برای نمایش تحلیل عملکردتان در هر سوال می توانید روی شماره سوال کلیک کنید.</span></tr>
                </thead>
                <thead>
                    <tr justify-content-center>
                        <td>پاسخ صحیح داوطلب<span class="greendark border lboption" style="color: #1b5e20 !important;">1</span></td>
                        <td>پاسخ نادرست     <span class="red border lboption"       style="color: #F22613 !important;"> 1</span></td>
                        <td>پاسخ صحیح سوال  <span class="green border lboption"     style="color: #8BC34A !important;"> 1</span></td>
                    </tr>
                </thead>
                <tbody>
                    <tr>
                        
                    </tr>
                </tbody>
            </table>
        </v-layout>
        <v-layout>
            <v-flex lg3>
                <v-progress-circular :rotate="-90" :size="100" :width="15" :value="80" color="primary" >
                    80
                </v-progress-circular>
            </v-flex>
            <v-flex lg3>
                <v-progress-circular :rotate="-90" :size="100" :width="15" :value="24" color="deep-orange lighten-2" >
                    24
                </v-progress-circular>
            </v-flex>
            <v-flex lg3>
                <v-progress-circular :rotate="-90" :size="100" :width="15" :value="10" color="lime" >
                    10
                </v-progress-circular>
            </v-flex>
            <v-flex lg3>
                <v-progress-circular :rotate="-90" :size="100" :width="15" :value="100" color="indigo darken-2" >
                    100
                </v-progress-circular>
            </v-flex>
        </v-layout>
    </v-container>
    
</template>

<script>
    export default {
        data() {return{
            

    }
    }
    }
</script>
<style>
.green {
    background: #8BC34A;
}
.red {
    background: #F22613;
}
.greendark {
   background: #1b5e20;
}
.lboption {
    width: 29px !important;
    display: inline;
    padding-right: 27px;
}

</style>



