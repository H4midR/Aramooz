import Vue from 'vue'
import Router from 'vue-router'

// add your components here
import Main from '@/components/public/Main'
import Login from '@/components/public/login'
import signup from '@/components/public/signup'
import addExam from '@/components/manager/addExam'
import addQuestion from '@/components/manager/addQuestion'
import startExam from '@/components/client/startExam'
import examsList from '@/components/manager/examsList'
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
    },
    {
      path : '/examsList',
      component: examsList
    }
  ]
})
