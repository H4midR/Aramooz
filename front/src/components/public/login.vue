<template>
  <v-container @keyup.enter="login" fluid fill-height>
     <v-layout fluid  justify-space-around align-center>
      <v-flex xs12 md5 >
          <v-card xs12 md5 v-if="User != null && resMessage == null">
            <v-alert value="true" type="success" outline>
              {{ User.name }} عزیز: 
               قبلا وارد سایت شده اید
              <v-progress-linear indeterminate color="success" class="mb-0" />
            </v-alert>
          </v-card>
          <v-alert v-model="alertStatus" :type="alertColor" outline dismissible>
            {{resMessage}}
            <v-progress-linear v-if="alertColor == 'success' " indeterminate color="success" class="mb-0" />
          </v-alert>
          <v-card xs12 md5 v-if="User == null">

              <v-toolbar color="accent" dark>
                فرم ورود
                <v-spacer></v-spacer>
                <v-tooltip top>
                <v-btn color="primary darken-3" slot="activator" flat icon router to="signup" >
                  <v-icon>mdi-account-plus</v-icon>
                </v-btn>
                <span>
                  عضویت
                </span>
              </v-tooltip>
              </v-toolbar>

            <v-card-text>
               <v-flex xs11>
            <v-text-field v-model="mobile" label="موبایل" ></v-text-field>
          </v-flex>
          <v-flex xs11>
            <v-text-field v-model="password" type="password" label="کلمه عبور" ></v-text-field>
          </v-flex>
            </v-card-text>
            <v-divider></v-divider>
            <v-card-actions>
              <v-tooltip top>
                <v-btn @click="login" color="success" slot="activator" flat >
                  <v-icon>mdi-login</v-icon> ورود
                </v-btn>
                <span>
                  ورود
                </span>
              </v-tooltip>
            </v-card-actions>
          </v-card>
      </v-flex>
    </v-layout>
  </v-container>
</template>
<script>

export default {
  data(){return{
    axios:require('axios'),
    mobile:null,
    password:null,

    alertStatus:false,
    alertColor:"success",
    resMessage:null,

  }},//data
  methods: {
    login(){
      this.axios.post( this.BaseURL +"/user/login",JSON.stringify({
        mobile:this.mobile,
        password:this.password
      })).then(res=>{
        this.resMessage = res.data.Message;
        this.alertStatus = true;
        if(res.data.Code > 0 ){
          this.alertColor = "success";
          this.$emit("Login",res.data.Data);
          setTimeout(() => this.$router.replace("/") , 1500);
        }else{
          this.alertColor = "warning";
        }
      })
    }, //login
  }, // methods
  props:{
    User:Object,
    BaseURL:String,
    ACL:Object,
  }, //props
  mounted(){
    if (this.User != null ){
      setTimeout(() => this.$router.replace("/") , 1500);
    }
  }//mounted
}
</script>
