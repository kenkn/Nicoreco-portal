import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router'
import HomeView from '../pages/Home.vue'
import LoginView from '../pages/Login.vue'
import RegisterView from '../pages/Register.vue'
import ForgotView from '../pages/Forgot.vue'
import ResetView from '../pages/Reset.vue'
import QuestionView from '../pages/Question.vue'
import DashView from '../components/Dash.vue'
import LaboReviewView from '../pages/LaboReview.vue'
import Chat from '../pages/Chat.vue'

// const router = new VueRouter({
//   routes: [
//       {
//     path: '/login',
//     component: LoginView
//   },
//   {
//     path: '/register',
//     component: RegisterView
//   },
//   {
//     path: '/forgot',
//     component: ForgotView
//   },
//   {
//     path: '/reset/:token',
//     component: ResetView
//   },
//   {
//     path: '/',
//     components: DashView,
//     children: [
//       {
//         path: 'home',
//         alias: '',
//         component: HomeView,
//         meta: {description: 'Home'}
//       },
//       {
//         path: 'question',
//         component: QuestionView,
//         meta: {description: 'Question'}
//       },
//       {
//         path: 'laboreview',
//         component: LaboReviewView,
//         meta: {descriprion: 'labreview'}
//       },
//       {
//         path: 'chat',
//         component: Chat,
//         meta: {description: 'Chat'}
//       }
//     ]
//   }
//   ]
// })

const routes: Array<RouteRecordRaw> = [
  // { path: '/', component: Home },
  // { path: '/login', component: Login },
  // { path: '/register', component: Register },
  // { path: '/forgot', component: Forgot },
  // { path: '/reset/:token', component: Reset },
  // { path: '/question', component: Question },
  {
    path: '/login',
    component: LoginView
  },
  {
    path: '/register',
    component: RegisterView
  },
  {
    path: '/forgot',
    component: ForgotView
  },
  {
    path: '/reset/:token',
    component: ResetView
  },
  {
    path: '/',
    component: DashView,
    children: [
      {
        path: 'home',
        alias: '',
        component: HomeView,
        meta: {description: 'Home'}
      },
      {
        path: 'question',
        component: QuestionView,
        meta: {description: 'Question'}
      },
      {
        path: 'laboreview',
        component: LaboReviewView,
        meta: {descriprion: 'labreview'}
      },
      {
        path: 'chat',
        component: Chat,
        meta: {description: 'Chat'}
      }
    ]
  }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
