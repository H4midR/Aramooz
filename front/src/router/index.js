import Vue from 'vue'
import Router from 'vue-router'

// add your components here
import HelloWorld from '@/components/HelloWorld'
import Login from '@/components/login'
import signup from '@/components/signup'
import addExam from '@/components/addExam'
import addQuestion from '@/components/addQuestion'

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
    },
    {
      path : '/signup',
      component: signup
    },
    {
      path : '/addExam',
      component: addExam
    }
    ,
    {
      path : '/addQuestion',
      component: addQuestion
    }
  ]
})
