<template>
    <v-container fluid full-height>
        <v-layout  fluid justify-space-around align-center>
          <v-flex xs12>
            {{response}}
          </v-flex>
            <v-flex xs12 md5>
                <v-card xs12 md5>
                    <v-toolbar color="pink darken-3" dark>
                        فرم عضویت
                        <v-spacer></v-spacer>
                        <v-tooltip top>
                        <v-btn color="secondary darken-3" slot="activator" flat icon >
                        <v-icon>mdi-login</v-icon>
                        </v-btn>
                        <span>
                        ورود
                        </span>
                        </v-tooltip>
                    </v-toolbar>

                    <v-card-text>
                        <v-flex xs11>
                            <v-text-field v-model="name" type="text" label="نام و نام خانوادگی"></v-text-field>
                        </v-flex>
                        <v-flex xs11>
                            <v-text-field v-model="mobile" type="text" label="نام کاربری"></v-text-field>
                        </v-flex>
                        <v-flex xs11>
                            <v-text-field v-model="password" type="password" label="کلمه عبور"></v-text-field>
                        </v-flex>
                    </v-card-text>
                    <v-card-actions>
                        <v-tooltip top>
                            <v-btn color="success" @click="signup" slot="activator">ثبت نام</v-btn>
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
        name:null,
        response:null,
    }},
    methods:{
        signup(){
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
            })

        }
    }
}
</script>
 <style  scoped>

</style>


