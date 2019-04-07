<template>
    <v-container @keyup.enter="signup" fluid full-height>
        <v-layout fluid  justify-space-around align-center>
            <v-flex xs12 md5>
                <v-card xs12 md5>
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

                    <v-snackbar v-model="snackbar" color="success" :bottom="y === 'bottom'" :left="x === 'left'" :multi-line="mode === 'multi-line'" :right="x === 'right'" :timeout="timeout" :top="y === 'top'" :vertical="mode === 'vertical'" >
                    {{ response }}
                    <v-btn right flat fab @click="snackbar = false">
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
      /*
      axios:require('axios').create({
              baseURL: 'http://localhost:9090/',
              timeout: 1000,
              withCredentials: false,
              headers:{
                'Authorization-Token':this.myuser.token,
                'X-USER':this.myuser.uid
              }
            }),
          */
        axios:require('axios'),
        mobile:null,
        password:null,
        name:null,
        response:null,

        btnLoading: false,

        snackbar: false,
        y: 'top',
        x: 'right',
        mode: '',
        timeout: 4000,

        showPass: false,
        rules: {
          required: value => !!value || 'اجباری!!',
          minpass: v => v.length >= 8 || 'حداقل 8 عدد یا حرف',
          minMobile: v => v.length >= 11 || 'شماره موبایل باید 11 عدد باشد!',
        }
    }},
    methods:{
        signup(){
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
              this.snackbar = true
              this.btnLoading = false
            })

        }
    }
}
</script>


