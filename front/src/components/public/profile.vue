<template>
  <v-container fill-height fluid grid-list-xl>
    <v-layout justify-center wrap>

        <v-snackbar v-model="snackbar" :color="snackbarColor" top="top" right="right" timeout="5000">
        {{ response }}
        <v-btn right flat fab @click="snackbar = false" :loading="redirectBtnLoging">
            <v-icon>close</v-icon>
        </v-btn>
        </v-snackbar>


      <v-flex xs12 md5  v-if="User == null">
        <v-card xs12 md5>
            <v-alert value="true" type="red" outline>
            دسترسی شما به این صفحه مجاز نمی باشد!!
            <v-progress-linear indeterminate color="red" class="mb-0" />
            </v-alert>
        </v-card>
      </v-flex>

      <v-flex xs12 md4 v-if="User != null">
        <v-card class="v-card-profile" flat>
          <v-avatar class="mx-auto d-block" size="130">
            <img src="https://cdn.vuetifyjs.com/images/cards/girl.jpg">

          </v-avatar>
          <v-card-text class="text-xs-center">
            <h6 class="category text-gray font-weight-thin mb-3">برنامه نویس</h6>
            <h4 class="card-title font-weight-light">{{User.name}}</h4>
            <p class="card-description font-weight-light" v-html="User.bio"></p>

          </v-card-text>
        </v-card>
      </v-flex>
      <v-flex xs12 md8 v-if="User != null">
        <v-card color="" title="Edit Profile" text="Complete your profile">
          <v-form >
            <v-container py-0>
              <v-layout wrap>
                <v-flex xs12 md4 >
                  <v-text-field  label="نام و نام خانوادگی" v-model="name" @keyup.enter="update"/>
                </v-flex>
                <v-flex xs12 md4 >
                  <v-text-field label="شماره موبایل" v-model="mobile" @keyup.enter="update"/>
                </v-flex>
                <v-flex xs12 md4 >
                  <v-text-field label="ایمیل" v-model="email" @keyup.enter="update"/>
                </v-flex>
                <v-flex xs12 md12>
                  <v-text-field label="آدرس" v-model="address" @keyup.enter="update"/>
                </v-flex>
                <v-flex xs12>
                  <v-textarea label="درباره من" value="" v-model="bio" />
                </v-flex>
                <v-flex xs12 text-xs-right >
                  <v-btn class="mx-0 font-weight-light" color="success" @click="update">
                    به روز رسانی
                  </v-btn>
                </v-flex>
              </v-layout>
            </v-container>
          </v-form>
        </v-card>
      </v-flex>

    </v-layout>
  </v-container>
</template>

<script>
export default {
  data(){return{
      axios:require('axios').create({
              baseURL: this.BaseURL+'/',
              timeout: 1000,
              withCredentials: false,
              headers:{
                'Authorization-Token':this.User.token,
                'X-USER':this.User.uid
              }
            }),
      name: this.User.name,
      mobile: this.User.mobile,
      email: this.User.email,
      address: this.User.address,
      bio: this.User.bio,

      snackbar:false,
      snackbarColor:"success",
      response:null,
  }},
  methods:{
    noActons(){

    },
      update(){
          this.axios.put("user",JSON.stringify({
              name: this.name,
              mobile: this.mobile,
              email: this.email,
              address: this.address,
              bio: this.bio,
              token: this.User.token,
            })).then(res=>{
              this.response = res.data.Message;
              this.snackbar = true;
              if(res.data.Code > 0 ){
                this.$emit("Login",res.data.Data);
                this.snackbarColor = "success";
              }else{
                this.snackbarColor = "warning";
              }
            })

      },//update
  }, // methods
  props:{
    User:Object,
    BaseURL:String,
    ACL:Object,
  }, //props
  mounted(){
    if (this.User == null ){
      setTimeout(() => this.$router.replace("/login") , 1500);
    }
  }//mounted
}
</script>
<style>
.v-card-profile{
    background:none !important;
    border:none !important;

}
</style>
