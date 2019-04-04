<template>
    <v-container fluid full-height>
        <v-layout  fluid justify-space-around align-center>
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
                            <v-text-field v-model="fullname" type="text" label="نام و نام خانوادگی"></v-text-field>
                        </v-flex>
                        <v-flex xs11>
                            <v-text-field v-model="username" type="text" label="نام کاربری"></v-text-field>
                        </v-flex>
                        <v-flex xs11>
                            <v-text-field v-model="password" type="password" label="کلمه عبور"></v-text-field>
                        </v-flex>
                        <v-flex xs11>
                            <v-text-field v-model="email" type="email" label="ایمیل"></v-text-field>
                        </v-flex>
                        <v-flex xs11>
                            <v-text-field v-model="phone" type="phone" label="تلفن"></v-text-field>
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
        username:null,
        password:null,
        fullname:null,
        email:null,
        phone:null
    }},
    methods:{
        signup(){
            this.axios.post("http://localhost:8585/user/signup",JSON.stringify({
                username:this.username,
                password:this.password,
                fullname:this.fullname,
                email:this.email,
                phone:this.phone
            }))

        }
    }
}
</script>
 <style  scoped>

</style>


