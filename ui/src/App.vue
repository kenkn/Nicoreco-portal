<template>
  <div>
    <Nav id="nav"/>
    <img :src="codeImg" id="code" alt="code">
    <div id="sidebar">
      <Sidebar />
    </div>
    <main class="p-4">
      <NotFound v-if="isNotFound" />
      <router-view v-else />
    </main>
  </div>
</template>
 
<script>
import { computed } from 'vue'
import { useStore } from 'vuex'
import Nav from "@/components/Nav";
import Sidebar from "@/components/Sidebar";
import NotFound from "@/components/NotFound";
import codeImg from "@/assets/img/code.png";
export default {
  components: {
    Nav,
    Sidebar,
    NotFound
  },
  setup(){
    const store = useStore()
    // 404エラーの状態管理
    const isNotFound = computed(() => store.state.isNotFound)
    return {
      codeImg,
      isNotFound
    }
  }
};
</script>

<style>
  /* 共通 */
  body {
    word-wrap: break-word;
  }
  #nav {
    position: relative;
    /* code部分より前に出すため */
    z-index: 1031
  }
  /* スマホ版 */
  #code {
    display: none;
  }
  #sidebar {
    display: none;
  }
  main {
    position: relative;
    top: 60px;
  }
  /* PC版 */
  /*画面サイズが768px(BootstrapのMedium)からはここを読み込む*/
  @media screen and (min-width:768px) {
    #code {
      display: block;
      object-fit: cover;
      position: fixed;
      top: 60px;
      width: 100vw;
      height: 160px;
      z-index: 1030;
    }
    #sidebar {
      display: block;
      position: fixed;
      top: 220px;
      width: 250px;
      height: 100vh;
      background-color: rgba(227, 225, 230, 0.555);
    }
    main{
      position: relative;
      top: 220px;
      margin-left: 250px;
    }
  }
</style>
