<template>
  <div>
    <!-- PC用のナビ -->
    <div id="for-pc">
      <nav class="navbar navbar-expand-md bg-white fixed-top py-0">
        <ul class="navbar-nav mr-auto">
          <li class="nav-item">
            <router-link to="/" class="nav-link p-0">
              <img :src="logo" width="140" height="60" alt="Top">
            </router-link>
          </li>
        </ul>
        <ul class="navbar-nav d-none d-md-inline mr5">
          <SearchForm />
        </ul>
        <ul class="navbar-nav my-sm-0">
          <!-- ログイン済なら表示 -->
          <template v-if="auth">
            <!-- TODO 通知機能 -->
            <li class="nav-item p-2">
              <button class="btn p-1">
                <img :src="bell" alt="bell" width="30" height="30">
              </button>
            </li>
            <!-- メニュー -->
            <li class="nav-item p-2">
              <UserIconMenu @sendLogout="logout"/>
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
            <li class="nav-item">
              <router-link to="/links" class="nav-link" @click="closeNav">リンク集</router-link>
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
import { computed } from 'vue'
import { useStore } from 'vuex'
import { useRouter, useRoute } from 'vue-router'
import UserIconMenu from "@/components/UserIconMenu"
import SearchForm from '@/components/SearchForm'
import logo from '@/assets/img/logo.png'
import bell from '@/assets/img/bell.png'
import axios from 'axios'

export default {
  name: "Nav",
  components: {
    UserIconMenu,
    SearchForm
  },
  setup() {
    const store = useStore()
    const route = useRoute()
    const router = useRouter()
    const auth = computed(() => store.state.auth)
    const displayName = computed(() => store.state.displayName)
    store.dispatch("setAuth", localStorage.isLogin)
    store.dispatch('setDisplayName', localStorage.displayName)
    store.dispatch("setGrade", localStorage.grade)
    const logout = async () => {
      localStorage.isLogin     = false
      localStorage.displayName = null
      localStorage.userID      = null
      localStorage.grade       = null     
      localStorage.email       = null
      store.dispatch("setAuth", false)
      store.dispatch("setDisplayName", '')
      store.dispatch("setGrade", '')
      await axios.get("/logout")
      router.push("/login")
    }
    // ナビが開いている場合は閉じる
    const closeNav = () => {
      const navBtn = document.getElementById('nav-btn')
      if (navBtn.getAttribute('aria-expanded') === 'true') {
        navBtn.click()
      }
    }
    return {
      logo,
      bell,
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
    /*画面サイズが768px(BootstrapのMedium)からはここを読み込む*/
    #for-smart {
      display: none;
    }
    #for-pc {
      display: block;
    }
  }
</style>