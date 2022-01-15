<template>
  <div class="container">
    <template v-if="auth">
      <div class="mb-5 position-relative">
        <UserIcon :width="60" :height="60" class="d-inline float-left mr-3" />
        <h1 class="display-5 d-inline">
          ようこそ! {{displayName}}さん
        </h1>
      </div>
      <div class="d-flex justify-content-around text-center">
        <div v-for="menu in menus" v-bind:key="menu.id" class="btn position-relative" @click="current=menu.id">
          {{menu.name}}
          <div v-if="menu.id==current" class="position-absolute bottom-0 start-0 bg-danger w-100" style="height: 2px;"/>
        </div>
      </div>
      <hr>
      <MyPosts :target="current" :userId="userId" />
    </template>

    <!-- ログインしていない場合 -->
    <template v-else>
      <div>
        <h1 class="display-5 mb-3">
          ログインしていません
        </h1>
        <div class="mb-2 ml-2">
          <router-link to="/login">
            ログインはこちら
            <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-arrow-up-right-square" viewBox="0 0 16 16">
              <path fill-rule="evenodd" d="M15 2a1 1 0 0 0-1-1H2a1 1 0 0 0-1 1v12a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1V2zM0 2a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2v12a2 2 0 0 1-2 2H2a2 2 0 0 1-2-2V2zm5.854 8.803a.5.5 0 1 1-.708-.707L9.243 6H6.475a.5.5 0 1 1 0-1h3.975a.5.5 0 0 1 .5.5v3.975a.5.5 0 1 1-1 0V6.707l-4.096 4.096z"/>
            </svg>
          </router-link>
        </div>
        <div class="mb-2 ml-2">
          <router-link to="/register">
            新規登録はこちら
            <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-arrow-up-right-square" viewBox="0 0 16 16">
              <path fill-rule="evenodd" d="M15 2a1 1 0 0 0-1-1H2a1 1 0 0 0-1 1v12a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1V2zM0 2a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2v12a2 2 0 0 1-2 2H2a2 2 0 0 1-2-2V2zm5.854 8.803a.5.5 0 1 1-.708-.707L9.243 6H6.475a.5.5 0 1 1 0-1h3.975a.5.5 0 0 1 .5.5v3.975a.5.5 0 1 1-1 0V6.707l-4.096 4.096z"/>
            </svg>
          </router-link>
        </div>
      </div>
    </template>
  </div>
</template>
 
<script>
import { ref, computed } from 'vue'
import axios from 'axios'
import { useStore } from 'vuex'
import UserIcon from "@/components/UserIcon"
import MyPosts from "@/components/MyPosts"

export default {
  name: "Home",
  components: {
    UserIcon,
    MyPosts
  },
  setup() {
    const store       = useStore()
    const auth        = computed(() => store.state.auth)
    const userId      = localStorage.userID
    const displayName = computed(() => store.state.displayName)
    const menus       = [
      { "id": "home", "name": "ホーム" },
      { "id": "question", "name": "質問" },
      { "id": "answer", "name": "回答" },
      { "id": "review", "name": "レビュー" },
      { "id": "good", "name": "Good" },
    ]
    const current = ref("home")
    return {
      auth,
      userId,
      displayName,
      menus,
      current
    }
  }
}
</script>
