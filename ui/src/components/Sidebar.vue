<template>
  <nav class="sidebar position-sticky">
    <ul class="list-group pt-2">
      <router-link
        tag="li"
        class="link-box list-group-item text-dark p-0 position-relative"
        to="/"
      >
        <div v-if="appName=='home'" class="position-absolute top-0 end-0 selected"></div>
        <div class="w-100 h-100 px-3 py-2">
          <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-house" viewBox="0 0 16 16">
            <path fill-rule="evenodd" d="M2 13.5V7h1v6.5a.5.5 0 0 0 .5.5h9a.5.5 0 0 0 .5-.5V7h1v6.5a1.5 1.5 0 0 1-1.5 1.5h-9A1.5 1.5 0 0 1 2 13.5zm11-11V6l-2-2V2.5a.5.5 0 0 1 .5-.5h1a.5.5 0 0 1 .5.5z"/>
            <path fill-rule="evenodd" d="M7.293 1.5a1 1 0 0 1 1.414 0l6.647 6.646a.5.5 0 0 1-.708.708L8 2.207 1.354 8.854a.5.5 0 1 1-.708-.708L7.293 1.5z"/>
          </svg>
          ホーム
        </div>
      </router-link>
      <router-link
        tag="li"
        class="link-box list-group-item text-dark p-0 position-relative"
        to="/question"
      >
        <div v-if="appName=='question'" class="position-absolute top-0 end-0 selected"></div>
        <div class="w-100 h-100 px-3 py-2">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            width="16"
            height="16"
            fill="currentColor"
            class="bi bi-book"
            viewBox="0 0 16 16"
          >
            <path
              d="M1 2.828c.885-.37 2.154-.769 3.388-.893 1.33-.134 2.458.063 3.112.752v9.746c-.935-.53-2.12-.603-3.213-.493-1.18.12-2.37.461-3.287.811V2.828zm7.5-.141c.654-.689 1.782-.886 3.112-.752 1.234.124 2.503.523 3.388.893v9.923c-.918-.35-2.107-.692-3.287-.81-1.094-.111-2.278-.039-3.213.492V2.687zM8 1.783C7.015.936 5.587.81 4.287.94c-1.514.153-3.042.672-3.994 1.105A.5.5 0 0 0 0 2.5v11a.5.5 0 0 0 .707.455c.882-.4 2.303-.881 3.68-1.02 1.409-.142 2.59.087 3.223.877a.5.5 0 0 0 .78 0c.633-.79 1.814-1.019 3.222-.877 1.378.139 2.8.62 3.681 1.02A.5.5 0 0 0 16 13.5v-11a.5.5 0 0 0-.293-.455c-.952-.433-2.48-.952-3.994-1.105C10.413.809 8.985.936 8 1.783z"
            />
          </svg>
          講義質問
        </div>
      </router-link>
      <router-link
        tag="li"
        class="link-box list-group-item text-dark p-0 position-relative"
        to="/lab"
      >
        <div v-if="appName=='lab'" class="position-absolute top-0 end-0 selected"></div>
        <div class="w-100 h-100 px-3 py-2">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            width="16"
            height="16"
            fill="currentColor"
            class="bi bi-bar-chart"
            viewBox="0 0 16 16"
          >
            <path
              d="M4 11H2v3h2v-3zm5-4H7v7h2V7zm5-5v12h-2V2h2zm-2-1a1 1 0 0 0-1 1v12a1 1 0 0 0 1 1h2a1 1 0 0 0 1-1V2a1 1 0 0 0-1-1h-2zM6 7a1 1 0 0 1 1-1h2a1 1 0 0 1 1 1v7a1 1 0 0 1-1 1H7a1 1 0 0 1-1-1V7zm-5 4a1 1 0 0 1 1-1h2a1 1 0 0 1 1 1v3a1 1 0 0 1-1 1H2a1 1 0 0 1-1-1v-3z"
            />
          </svg>
          研究室
        </div>
      </router-link>
      <router-link
        tag="li"
        class="link-box list-group-item text-dark p-0 position-relative"
        to="/"
      >
        <div v-if="appName=='chat'" class="position-absolute top-0 end-0 selected"></div>
        <div class="w-100 h-100 px-3 py-2">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            width="16"
            height="16"
            fill="currentColor"
            class="bi bi-chat-dots"
            viewBox="0 0 16 16"
          >
            <path
              d="M5 8a1 1 0 1 1-2 0 1 1 0 0 1 2 0zm4 0a1 1 0 1 1-2 0 1 1 0 0 1 2 0zm3 1a1 1 0 1 0 0-2 1 1 0 0 0 0 2z"
            />
            <path
              d="m2.165 15.803.02-.004c1.83-.363 2.948-.842 3.468-1.105A9.06 9.06 0 0 0 8 15c4.418 0 8-3.134 8-7s-3.582-7-8-7-8 3.134-8 7c0 1.76.743 3.37 1.97 4.6a10.437 10.437 0 0 1-.524 2.318l-.003.011a10.722 10.722 0 0 1-.244.637c-.079.186.074.394.273.362a21.673 21.673 0 0 0 .693-.125zm.8-3.108a1 1 0 0 0-.287-.801C1.618 10.83 1 9.468 1 8c0-3.192 3.004-6 7-6s7 2.808 7 6c0 3.193-3.004 6-7 6a8.06 8.06 0 0 1-2.088-.272 1 1 0 0 0-.711.074c-.387.196-1.24.57-2.634.893a10.97 10.97 0 0 0 .398-2z"
            />
          </svg>
          雑談
        </div>
      </router-link>
      <router-link
        tag="li"
        class="link-box list-group-item text-dark p-0 position-relative"
        to="/links"
      >
        <div v-if="appName=='links'" class="position-absolute top-0 end-0 selected"></div>
        <div class="w-100 h-100 px-3 py-2">
          <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-link-45deg" viewBox="0 0 16 16">
            <path d="M4.715 6.542 3.343 7.914a3 3 0 1 0 4.243 4.243l1.828-1.829A3 3 0 0 0 8.586 5.5L8 6.086a1.002 1.002 0 0 0-.154.199 2 2 0 0 1 .861 3.337L6.88 11.45a2 2 0 1 1-2.83-2.83l.793-.792a4.018 4.018 0 0 1-.128-1.287z"/>
            <path d="M6.586 4.672A3 3 0 0 0 7.414 9.5l.775-.776a2 2 0 0 1-.896-3.346L9.12 3.55a2 2 0 1 1 2.83 2.83l-.793.792c.112.42.155.855.128 1.287l1.372-1.372a3 3 0 1 0-4.243-4.243L6.586 4.672z"/>
          </svg>
          リンク集
        </div>
      </router-link>
    </ul>
  </nav>
</template>

<script>
import { ref, watch } from 'vue'
import { useRoute } from 'vue-router'

export default {
  name: "Sidebar",
  setup() {
    const route = useRoute()
    const appName = ref(null)
    // 赤いバーを表示するために、urlからappNameを判定
    watch(route, () => {
      if(route.path==="/") {
        appName.value = "home"
      } else if(route.path.indexOf("/question") != -1) {
        appName.value = "question"
      } else if(route.path.indexOf("/lab") != -1) {
        appName.value = "lab"
      } else if(route.path.indexOf("/chat") != -1) {
        appName.value = "chat"
      } else if(route.path.indexOf("/links") != -1) {
        appName.value = "links"
      } else {
        appName.value = null
      }
    })
    return {
      appName
    }
  }
};
</script>

<style scoped>
  .selected {
    background-color: red;
    height: 100%;
    width: 2px;
  }
  .link-box {
    background-color: white;
  }
  .link-box :hover{
    background-color: rgb(230, 229, 227);
  }
</style>
