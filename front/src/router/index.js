import Vue from 'vue'
import Router from 'vue-router'

// add your components here
import Main from '@/components/Main'
import Login from '@/components/login'
import signup from '@/components/signup'
import addExam from '@/components/addExam'
import addQuestion from '@/components/addQuestion'
import startExam from '@/components/startExam'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'root',
      component: Main
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
    },
    {
      path : '/addquestion',
      component: addQuestion
    },
    {
      path : '/startExam',
      component: startExam
    }
  ]
})
