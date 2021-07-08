import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import axios from 'axios'

// /api/registerにリクエスト飛ばすためのやつ
axios.defaults.baseURL = 'http://localhost:80/api/'
axios.defaults.withCredentials = true

createApp(App).use(store).use(router).mount('#app')
