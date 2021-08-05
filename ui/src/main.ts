import { VuesticPlugin } from 'vuestic-ui'
import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import axios from 'axios'
import vuesticGlobalConfig from './vuestic/global-config'

// /api/registerにリクエスト飛ばすためのやつ
axios.defaults.baseURL = 'http://localhost:80/api/'
axios.defaults.withCredentials = true

const app = createApp(App)
app.use(store)
app.use(router)
// app.use(VuesticPlugin, vuesticGlobalConfig)
app.mount('#app')

// createApp(App).use(store).use(router).mount('#app')
