<template>
  <div>
    <!-- PC用のナビ -->
    <div id="for-pc">
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
              <router-link to="/profile" class="nav-link">{{ displayName }}</router-link>
            </li>
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
    </div>

    <!-- スマホ用のナビ -->
    <div id="for-smart"> 
      <nav class="navbar navbar-expand-md navbar-dark bg-dark fixed-top">
        <ul class="navbar-nav mr-auto">
          <li class="nav-item">
            <router-link to="/" class="nav-link" @click="closeNav">Nicoreco</router-link>
          </li>
        </ul>
        <button 
          id="nav-btn" 
          class="navbar-toggler" 
          type="button" 
          data-bs-toggle="collapse" 
          data-bs-target="#navbarToggler" 
          aria-controls="navbarToggler" 
          aria-expanded="flase" 
          aria-label="Toggle navigation"
        >
          <span class="navbar-toggler-icon"></span>
        </button>
        <div class="collapse navbar-collapse" id="navbarToggler">
          <ul class="navbar-nav me-auto mb-2 mb-lg-0">
            <li class="nav-item">
              <router-link to="/question" class="nav-link" @click="closeNav">講義質問</router-link>
            </li>
            <li class="nav-item">
              <router-link to="/lab" class="nav-link" @click="closeNav">研究室</router-link>
            </li>
            <li class="nav-item">
              <router-link to="/" class="nav-link" @click="closeNav">雑談</router-link>
            </li>
            <!-- ログイン済なら表示 -->
            <template v-if="auth">
              <li class="nav-item">
                <router-link to="/profile" class="nav-link" @click="closeNav">マイページ</router-link>
              </li>
              <li class="nav-item">
                <router-link to="/login" class="nav-link" @click="logout(), closeNav()">ログアウト</router-link>
              </li>
            </template>
            <!-- 未ログインなら表示 -->
            <template v-if="!auth">
              <li class="nav-item">
                <router-link to="/login" class="nav-link" @click="closeNav">ログイン</router-link>
              </li>
              <li class="nav-item">
                <router-link to="/register" class="nav-link" @click="closeNav">新規登録</router-link>
              </li>
            </template>
          </ul>
        </div>
        </nav>
    </div>
  </div>
</template>

<script>
import { computed, watch } from 'vue'
import { useStore } from 'vuex'
import { useRouter, useRoute } from 'vue-router'

export default {
  name: "Nav",
  setup() {
    const store = useStore()
    const route = useRoute()
    const router = useRouter()
    const auth = computed(() => store.state.auth)
    const displayName = computed(() => store.state.displayName)
    store.dispatch("setAuth", localStorage.isLogin)
    store.dispatch('setDisplayName', localStorage.displayName)
    const logout = async () => {
      localStorage.isLogin     = false
      localStorage.displayName = null
      localStorage.userID      = null
      localStorage.grade       = null     
      localStorage.email       = null
      localStorage.authToken   = null
      store.dispatch("setAuth", false)
      store.dispatch("setAuth", '')
      router.push("/login")
    }
    // ナビが開いている場合は閉じる
    const closeNav = () => {
      const navBtn = document.getElementById('nav-btn')
      if(navBtn.getAttribute('aria-expanded') === 'true'){
        navBtn.click()
      }
    }
    return {
      auth,
      displayName,
      logout,
      closeNav,
    }
  }
}
</script>

<style>
  #for-smart {
    display: block;
  }
  #for-pc {
    display: none;
  }
  @media screen and (min-width:768px) {
    /*画面サイズが768px(BootstraoのMedium)からはここを読み込む*/
    #for-smart {
      display: none;
    }
    #for-pc {
      display: block;
    }
  }
</style>