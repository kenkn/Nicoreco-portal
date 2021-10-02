<template>
  <main class="form-login">
    <form @submit.prevent="login">
      <h1 class="h3 mb-3 fw-normal">ログイン</h1>
      <input
        v-model="email"
        type="email"
        class="form-control"
        placeholder="Email"
        required
      />

      <input
        v-model="password"
        type="password"
        class="form-control"
        placeholder="Password"
        required
      />

      <div class="mb-2">
        <router-link to="/forgot">パスワードを忘れましたか？</router-link>
      </div>

      <button class="w-100 btn-lg btn-primary" type="submit">ログイン</button>
    </form>
  </main>
</template>
 
<script>
import { ref } from 'vue'
import axios from 'axios'
import { useRouter } from 'vue-router'
import { useStore } from 'vuex'

// TODO ログイン失敗時の挙動を設定する
export default {
  name: "Login",
  data() {
    const email = ref("")
    const password = ref("")
    const router = useRouter()
    const store = useStore()

    const login = async () => {
      await axios.post("login", {
        email: email.value,
        password: password.value
      })
      const userData = await axios.get("user")
      console.log(userData.data)
      localStorage.displayName = await userData.data.display_name
      localStorage.userID = await userData.data.user_id
      localStorage.grade = await userData.data.Grade
      localStorage.email = await userData.data.email
      localStorage.isLogin = true
      store.dispatch("setAuth", true)
      // 前のページに遷移する
      await router.push(this.prevRoute.path)
    }

    return {
      prevRoute : null,
      email,
      password,
      login
    }
  },
  beforeRouteEnter(to, from, next) {
    next(vm => {
      vm.prevRoute = from
    })
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
.form-login input[type="email"] {
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