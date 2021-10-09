import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import axios from 'axios'
import { useStore } from 'vuex'
import "bootstrap/dist/css/bootstrap.min.css"
import "bootstrap"

// /api/registerにリクエスト飛ばすためのやつ
axios.defaults.baseURL = 'http://localhost:80/api/'
axios.defaults.withCredentials = true

createApp(App).use(store).use(router).mount('#app')

const loginCheck = async () => {
    try {
        const userData = await axios.get(axios.defaults.baseURL + "user")
        localStorage.displayName = await userData.data.display_name
        localStorage.userID      = await userData.data.user_id
        localStorage.grade       = await userData.data.Grade
        localStorage.email       = await userData.data.email
        localStorage.isLogin     = true
        store.dispatch("setAuth", true)
    } catch (e) {
        localStorage.displayName = null
        localStorage.userID      = null
        localStorage.grade       = null
        localStorage.email       = null
        localStorage.isLogin     = false
        store.dispatch("setAuth", false)
    }
}

loginCheck()