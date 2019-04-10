<template>
  <v-container fluid fill-height>
     <v-layout fluid  justify-space-around align-center>
      <v-flex xs12 md5 >
        <!--<v-layout fluid wrap> -->
          <v-card xs12 md5>

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
            <v-text-field v-model="phone" label="موبایل" ></v-text-field>
          </v-flex>
          <v-flex xs11>
            <v-text-field v-model="password" type="password" label="کلمه عبور" ></v-text-field>
          </v-flex>
            </v-card-text>
            <v-divider></v-divider>
            <v-card-actions>
              <v-tooltip top>
                <v-btn @click="login" color="success" slot="activator" flat icon >
                  <v-icon>mdi-login</v-icon>
                </v-btn>
                <span>
                  تایید
                </span>
              </v-tooltip>
            </v-card-actions>
          </v-card>



        <!-- </v-layout>-->
      </v-flex>
    </v-layout>
  </v-container>
</template>
<script>

export default {
  data(){return{
    axios:require('axios'),
    phone:null,
    password:null,

  }},
  methods: {
    login(){
      this.axios.post("http://localhost:9090/user/login",JSON.stringify({
        mobile:this.phone,
        password:this.password
      }))
    }
  },
}
</script>
