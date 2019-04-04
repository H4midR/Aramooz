import Vue from 'vue'
import Router from 'vue-router'

// add your components here
import HelloWorld from '@/components/HelloWorld'
import Login from '@/components/login'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'root',
      component: HelloWorld
    },
    {
      path : '/login',
      component: Login
    }
  ]
})
