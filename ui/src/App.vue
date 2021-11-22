<template>
  <div class="h-100">
    <Nav />
    <div class="h-100 m-0 wrapper">
      <div id="sidebar">
        <Sidebar />
      </div>
      <main id="contents" class="border-left border-dark overflow-auto">
        <NotFound v-if="isNotFound" />
        <router-view v-else />
      </main>
    </div>
  </div>
</template>
 
<script>
import { computed } from 'vue'
import { useStore } from 'vuex'
import Nav from "@/components/Nav";
import Sidebar from "@/components/Sidebar";
import NotFound from "@/components/NotFound";
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
      isNotFound
    }
  }
};
</script>

<style>
  html {
    height: 100%;
  }
  body {
    height: 100%;
    word-wrap: break-word;
  }
  #app {
    height: 100%;
  }
  .wrapper {
    display: flex;
  }
  #sidebar {
    background-color: rgba(227, 225, 230, 0.555);
    display: none;
  }
  #contents {
    padding-top: 5rem;
    padding-inline: 2rem;
    background-color: rgba(255, 255, 255, 0.966);
    flex: 1;
  }
  @media screen and (min-width:768px) {
    /*画面サイズが768px(BootstraoのMedium)からはここを読み込む*/
    #menu {
      display: none;
    }
    #sidebar {
      display: block;
      width: 250px;
      background-color: rgba(227, 225, 230, 0.555);
    }
    #contents {
      padding-top: 5rem;
      padding-inline: 5rem;
      height: 100%;
    }
  }
</style>
