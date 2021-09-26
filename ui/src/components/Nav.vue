<template>
  <nav class="navbar navbar-expand-md navbar-dark bg-dark fixed-top">
    <ul class="navbar-nav mr-auto">
      <li class="nav-item">
        <router-link to="/" class="nav-link">Nicoreco</router-link>
      </li>
    </ul>

    <ul class="navbar-nav my-sm-0">
      <!-- ログイン済なら表示 -->
      <template v-if="auth">
        <li class="nav-item">
          <router-link to="/login" class="nav-link" @click="logout">ログアウト</router-link>
        </li>
      </template>
      <!-- 未ログインなら表示 -->
      <template v-if="!auth">
        <li class="nav-item">
          <router-link to="/login" class="nav-link">ログイン</router-link>
        </li>
        <li class="nav-item">
          <router-link to="/register" class="nav-link">新規登録</router-link>
        </li>
      </template>
    </ul>
  </nav>
</template>

<script>
import { computed } from 'vue'
import { useStore } from 'vuex'
import { useRouter } from 'vue-router'
import axios from 'axios'

export default {
  name: "Nav",
  setup() {
    const store = useStore()
    const router = useRouter()
    const auth = computed(() => store.state.auth)
    store.dispatch("setAuth", localStorage.isLogin)
    const logout = async () => {
      await axios.get("logout", {})
      localStorage.isLogin = false
      store.dispatch("setAuth", false)
      await router.push("/login")
    }
    return {
      auth,
      logout
    }
  }
}
</script>
