import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router'
import Home from '../pages/Home.vue'
import Login from '../pages/Login.vue'
import Register from '../pages/Register.vue'
import Forgot from '../pages/Forgot.vue'
import Reset from '../pages/Reset.vue'
import Question from '../pages/Question.vue'

const routes: Array<RouteRecordRaw> = [
  { path: '/', component: Home },
  { path: '/login', component: Login },
  { path: '/register', component: Register },
  { path: '/forgot', component: Forgot },
  { path: '/reset/:token', component: Reset },
  { path: '/question', component: Question },
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
