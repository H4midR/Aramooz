<template>
    <v-container @keyup.enter="signup" fluid full-height>
        <v-layout fluid  justify-space-around align-center>
            <v-flex xs12 md5>
                <v-card xs12 md5 v-if="User != null && resMessage == null">
                    <v-alert value="true" type="success" outline>
                    {{ User.name }} عزیز:
                    قبلا وارد سایت شده اید
                    <v-progress-linear indeterminate color="success" class="mb-0" />
                    </v-alert>
                </v-card>
                <v-card xs12 md5 v-if="User == null">
                    <v-toolbar color="pink darken-3" dark>
                        فرم عضویت
                        <v-spacer></v-spacer>
                        <v-tooltip top>
                        <v-btn color="secondary darken-3" slot="activator" flat icon  @click.native="$router.push('/login')">
                            <v-icon>mdi-login</v-icon>
                        </v-btn>
                        <span>
                            ورود
                        </span>
                        </v-tooltip>
                    </v-toolbar>


          <v-form ref="ProfileForm" >
                    <v-card-text>
                        <v-flex xs11>
                            <v-text-field :rules="[rules.required]" hint="نام خود را با حروف فارسی وارد نمایید" v-model="name" type="text" label="نام و نام خانوادگی"></v-text-field>
                        </v-flex>
                        <v-flex xs11>
                            <v-text-field :rules="[rules.required, rules.minMobile]" hint="شماره موبایل را با اعداد انگلیسی وارد نمایید" v-model="mobile" type="text" mask="###########" label="موبایل"></v-text-field>
                        </v-flex>
                        <v-flex xs11>
                            <v-text-field :append-icon="showPass ? 'visibility' : 'visibility_off'" :rules="[rules.required, rules.minpass]" :type="showPass ? 'text' : 'password'"  @click:append="showPass = !showPass" v-model="password" label="کلمه عبور"></v-text-field>
                        </v-flex>
                    </v-card-text>
                    <v-divider></v-divider>
                    <v-card-actions>
                        <v-tooltip top>
                            <v-btn block :loading="btnLoading" :disabled="btnLoading" color="success" @click="signup" slot="activator">
                                ثبت نام
                                <v-icon left>mdi-account-plus</v-icon>
                            </v-btn>
                            <span>ثبت نام در سایت</span>
                        </v-tooltip>
                    </v-card-actions>
                </v-form>
                    <v-snackbar v-model="snackbar" :color="snackbarColor" top="top" right="right" timeout="5000">
                    {{ response }}
                    <v-btn right flat fab @click="snackbar = false" :loading="redirectBtnLoging">
                        <v-icon>close</v-icon>
                    </v-btn>
                    </v-snackbar>
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
        name:null,
        response:null,
        btnLoading: false,
        redirectBtnLoging:false,
        snackbar: false,
        snackbarColor: null,
        showPass: false,
        rules: {
          required: value => !!value || 'اجباری!!',
          minpass: v => v.length >= 8 || 'حداقل 8 عدد یا حرف',
          minMobile: v => v.length >= 11 || 'شماره موبایل باید 11 عدد باشد!',
        }
    }},
    methods:{
        signup(){
          if(!this.$refs.ProfileForm.validate())
          return



            this.btnLoading = true
            /*
            required
            name
            mobile
            password
            */

            this.axios.post("http://localhost:9090/user",JSON.stringify({
                mobile:this.mobile,
                password:this.password,
                name:this.name,
            })).then(res=>{
              this.response=res.data.Message
              if(res.data.Code > 0){
                  this.snackbarColor = "success";
                  this.redirectBtnLoging = true;
                  setTimeout(() => this.$router.replace("/login") , 3000);
              }else{
                  this.snackbarColor = "warning";
              }
              this.snackbar = true
              this.btnLoading = false
            })
        }
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
    } //mounted
}
</script>
