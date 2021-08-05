import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router'
import AppLayout from '@/layout/app-layout.vue'

const routes: Array<RouteRecordRaw> = [
  // { path: '/', component: Home },
  // { path: '/login', component: Login },
  // { path: '/register', component: Register },
  // { path: '/forgot', component: Forgot },
  // { path: '/reset/:token', component: Reset },
  // { path: '/question', component: Question },
  {
    name: 'login',
    path: '/login',
    component: () => import('@/pages/Login.vue'),
  },
  {
    name: 'reset',
    path: '/reset/:token',
    component: () => import('@/pages/Reset.vue'),
  },
  {
    name: 'register',
    path: '/register',
    component: () => import('@/pages/Register.vue'),
  },
  {
    path: '/',
    component: AppLayout,
    children: [
      {
        name: 'home',
        path: 'home',
        component: () => import('@/pages/Home.vue'),
      },
      {
        name: 'question',
        path: 'question',
        component: () => import('@/pages/Question.vue'),
      },
      {
        name: 'laboreview',
        path: 'laboreview',
        component: () => import('@/pages/LaboReview.vue'),
      },
      {
        name: 'chat',
        path: 'chat',
        component: () => import('@/pages/ChatPage.vue'),
      },
    ]
  }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
