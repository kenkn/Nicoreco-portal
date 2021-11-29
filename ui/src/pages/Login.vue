<template>
  <div class="form-login">
    <form @submit.prevent="login">
      <h1 class="h3 mb-3 fw-normal">ログイン</h1>
      <input
        v-model="email"
        name="email"
        class="form-control email d-inline"
        placeholder="Email"
        required
      />
      <span class="d-inline lead borde">
        @yamaguchi-u.ac.jp
      </span>

      <input
        v-model="password"
        type="password"
        class="form-control"
        placeholder="パスワード"
        required
      />
      <p>{{ errMessage }}</p>
      <div class="mb-2">
        <router-link to="/forgot">パスワードを忘れましたか？</router-link>
      </div>

      <button v-if="!loading" class="w-100 btn-lg btn-primary " type="submit">ログイン</button>
      <button v-else class="w-100 btn-lg btn-primary disabled" disabled>
        <div class="spinner-border" role="status">
          <span class="sr-only">Loading...</span>
        </div>
      </button>
    </form>
  </div>
</template>
 
<script>
import { ref } from 'vue'
import axios from 'axios'
import { useRouter } from 'vue-router'
import { useStore } from 'vuex'

export default {
  name: "Login",
  setup() {
    const email      = ref("")
    const password   = ref("")
    const router     = useRouter()
    const store      = useStore()
    const errMessage = ref("")
    const loading    = ref(false)
    const login = async () => {
      const emailYamaguchi = email.value + "@yamaguchi-u.ac.jp"
      loading.value = true
      try {
        const jwtToken = await axios.post("login", {
          email: emailYamaguchi,
          password: password.value
        })
        localStorage.authToken = jwtToken.data.jwt
      } catch (e) {
        errMessage.value = "メールアドレスかパスワードが間違っています．"
        loading.value = false
        console.log(e)
      }

      try {
        const userData = await axios.post("user", {
          jwt: localStorage.authToken
        })
        // localStorageの情報を更新
        localStorage.displayName = userData.data.display_name
        localStorage.userID = userData.data.user_id
        localStorage.grade = userData.data.Grade
        localStorage.email = userData.data.email
        localStorage.isLogin = true
        store.dispatch("setAuth", true)
        store.dispatch('setDisplayName', localStorage.displayName)
      } catch (e) {
        console.log(e)
      }
      
      // 前のページに遷移する
      if (localStorage.prevRoute == '/register') {
        router.push('/')
      } else {
        router.push(localStorage.prevRoute)
      }
    }

    return {
      prevRoute : null,
      errMessage,
      email,
      password,
      loading,
      login
    }
  },
  beforeRouteEnter(to, from, next) {
    localStorage.prevRoute = from.path
    next()
  }
}
</script>
 
<style>
.form-login {
  width: 100%;
  max-width: 330px;
  padding: 15px;
  margin: auto;
}

.form-login .form-control {
  position: relative;
  box-sizing: border-box;
  height: auto;
  padding: 10px;
  font-size: 16px;
}
.form-login .form-control:focus {
  z-index: 2;
}
.form-login input[name="email"] {
  width: 40%;
  margin-bottom: -1px;
  border-bottom-right-radius: 0;
  border-bottom-left-radius: 0;
}
.form-login input[type="password"] {
  margin-bottom: 10px;
  border-top-left-radius: 0;
  border-top-right-radius: 0;
}
</style>