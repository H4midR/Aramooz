<template>
  <v-container>
  <v-text-field label="نام" v-model="mylocalUser.firstName"></v-text-field>

  <v-text-field label="username" v-model="mylocalUser.username"></v-text-field>
  <v-text-field label="password" v-model="mylocalUser.password"></v-text-field>
  <v-btn @click="saveName" >ذخیره</v-btn>

  </v-container>
</template>
<script>
export default {
  data(){return{
    mylocalUser:null,
  }},
  methods:{
    saveName(){
      console.log("log from children",this.mylocalUser)
      this.$emit('mysave',this.mylocalUser)
    }
  },
  props:{
    question:Object,// سوال ورودی
    myuser:Object,

  },
  mounted(){
    this.mylocalUser=Object.assign({},this.myuser)
  }


}
</script>

