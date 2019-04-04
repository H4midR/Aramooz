<template>
 <v-container>
 سلام <h1 :class="color">  {{fullName}}</h1>
      <v-text-field v-model="color" hint="کلاس را مشخص کنید" label="کلاس"  ></v-text-field>
      <v-text-field v-model="user.firstName"></v-text-field>
      <input :value="color"  />

 <v-btn class="primary" @click="chname" >تغییر نام</v-btn>
 <hr>
 child


  <myuser :myuser="user" @mysave="logChildData" ></myuser>


 </v-container>
</template>
<script>
import myuser from '@/components/user'
export default {
  data(){return{
    firstName:"Hamid",
    lastName:"Reza",
    color:"",
    user:{
      firstName:"hamid",
      username:"hamid",
      password:"reza"
    }
  }},
  methods:{
    chname(){
      this.firstName = "Hello from Aramooz";
      this.color="blue";
    },
    logChildData(e){
      console.log("from parent",e)
      this.firstName=e.firstName
    }
  },
  computed: {
    fullName(){
      return this.firstName + this.lastName + this.firstName.length
    }
  },
  components:{
    myuser
  }

}
</script>
<style scoped>
.red{
  color: red;
}
.blue{
  color: blue;
}
</style>

