<template>
  <v-container fill-height fluid grid-list-xl>
    <v-layout justify-center wrap>
      <v-flex xs12 md4>
        <v-card class="v-card-profile" flat>
          <v-avatar class="mx-auto d-block" size="130">
            <img src="https://cdn.vuetifyjs.com/images/cards/girl.jpg">
             
          </v-avatar>
          <v-card-text class="text-xs-center">
            <h6 class="category text-gray font-weight-thin mb-3">برنامه نویس</h6>
            <h4 class="card-title font-weight-light">مریم پهلون</h4>
            <p class="card-description font-weight-light">...</p>
            
          </v-card-text>
        </v-card>
      </v-flex>
      <v-flex xs12 md8>
        <v-card color="" title="Edit Profile" text="Complete your profile">
          <v-form>
            <v-container py-0>
              <v-layout wrap>
                <v-flex xs12 md4 >
                  <v-text-field label="نام و نام خانوادگی" disabled/>
                </v-flex>
                <v-flex xs12 md4 >
                  <v-text-field class="purple-input" label="نام کاربری" v-model="profileData.userName"/>
                </v-flex>
                <v-flex xs12 md4 >
                  <v-text-field label="ایمیل" class="purple-input" v-model="profileData.email"/>
                </v-flex>
                <v-flex xs12 md12>
                  <v-text-field label="آدرس" class="purple-input" v-model="profileData.address"/>
                </v-flex>
                <v-flex xs12 md4>
                  <v-text-field label="شهر" class="purple-input" v-model="profileData.city"/>
                </v-flex>
                <v-flex xs12 md4>
                  <v-text-field label="کشور" class="purple-input" v-model="profileData.country"/>
                </v-flex>
                <v-flex xs12 md4>
                  <v-text-field class="purple-input" label="کد پستی" type="number" v-model="profileData.postalCode"/>
                </v-flex>
                <v-flex xs12>
                  <v-textarea class="purple-input" label="درباره من" value="" v-model="profileData.aboutMe"/>
                </v-flex>
                <v-flex xs12 text-xs-right >
                  <v-btn class="mx-0 font-weight-light" color="success">
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
  data:() => ({
      axio: require('axios'),
      profileData:{
          name: null,
          userName: null,
          email: null,
          address: null,
          city: null,
          country: null,
          postalCode: null,
          aboutMe: null
      },
      response: null
  }),
  methods:{
      update(){
          this.axios("http://localhost:9090/user",JSON.stringify(this.profileData)).then(res=>{
              this.response=res.data.Message
              this.snackbar = true
              this.btnLoading = false
            })

      }
  }
}
</script>
<style>
.v-card-profile{
    background:none !important;
    border:none !important;
    
}
</style>