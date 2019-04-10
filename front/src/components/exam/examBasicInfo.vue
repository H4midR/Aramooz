<template>
<v-layout class="black" >
    <v-layout>
        <v-btn class="gray darken-1">
        {{examInfoProp.title}}
        </v-btn>
        <v-btn class="teal lighten-1 white--text">
        <v-icon>create</v-icon>
        گزارش اشکال سوال
        </v-btn>
        <v-btn class="orange white--text" >
            <v-icon>create</v-icon>
        مرور سوال ها
        </v-btn>
        <v-btn class="pink lighten-1" dark @click.stop="dialog = true">
            <v-icon>help</v-icon>
        راهنما
        </v-btn>
        <v-btn class="blue darken-2 white--text">
            <v-icon>access_time</v-icon>
        زمان
        </v-btn>
    </v-layout>
    <v-spacer></v-spacer>
    <v-layout class="white--text" justify-end>
        <timer></timer>
    </v-layout>
</v-layout>
</template>
<script>
import timer from '@/components/exam/timeCountdown'
export default {data(){return{}},
    props:{
        examInfoProp:Object
    },
    components:{
        timer
    }

}
</script>

<style>

</style>
