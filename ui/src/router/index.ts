import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router'
import Home from '../pages/Home.vue'
import Login from '../pages/Login.vue'
import Register from '../pages/Register.vue'
import Forgot from '../pages/Forgot.vue'
import Reset from '../pages/Reset.vue'
import Subjects from '../pages/Subjects.vue'
import Questions from '../pages/Questions.vue'
import Question from '../pages/Question.vue'
import CreateQuestion from '../pages/CreateQuestion.vue'
import Labs from '../pages/Labs.vue'
import LabReviews from '../pages/LabReviews.vue'
import LabReview from '../pages/LabReview.vue'
import CreateLabReview from '../pages/CreateLabReview.vue'
import Profile from '../pages/Profile.vue'
import Links from '../pages/Links.vue'
import Search from '../pages/Search.vue'
import NotFound from '../components/NotFound.vue'

const routes: Array<RouteRecordRaw> = [
  { path: '/', component: Home },
  { path: '/login', component: Login },
  { path: '/register', component: Register },
  { path: '/forgot', component: Forgot },
  { path: '/reset/:token', component: Reset },
  // 講義一覧ページ
  { path: '/question', component: Subjects },
  // 講義の質問一覧ページ
  { path: '/question/:subject', component: Questions },
  // 新規質問の作成ページ
  { path: '/question/:subject/create', component: CreateQuestion },
  // 質問の詳細ページ
  { path: '/question/:subject/:question_id', component: Question },
  // ラボの一覧ページ
  { path: '/lab', component: Labs },
  // ラボレビューの一覧ページ
  { path: '/lab/:professor', component: LabReviews },
  // 新規レビューの作成ページ
  { path: '/lab/:professor/create', component: CreateLabReview },
  // ラボのレビュ詳細ーページ
  { path: '/lab/:professor/:labReviewID', component: LabReview },
  // プロフィールページ
  { path: '/profile', component: Profile, 
    beforeEnter: (to, from, next) => {
      if(localStorage.isLogin=='true') next()
      else router.push("/login")
    } 
  },
  // リンクページ
  { path: '/links', component: Links }, 
  // 検索結果表示ページ
  { path: '/search', component: Search} ,
  // 404ページ
  { path: '/:catchAll', component: NotFound, name: 'NotFound'}
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
