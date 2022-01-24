<template>
  <div>
    <Nav id="nav"/>
    <div id="sidebar">
      <Sidebar />
    </div>
    <main class="">
      <NotFound v-if="isNotFound" />
      <div v-else>
        <CodeBar id="code-bar"></CodeBar>
        <router-view class="p-4" />
      </div>
    </main>
  </div>
</template>
 
<script>
import { computed } from 'vue'
import { useStore } from 'vuex'
import Nav from "@/components/Nav";
import Sidebar from "@/components/Sidebar";
import NotFound from "@/components/NotFound";
import CodeBar from "@/components/CodeBar";
export default {
  components: {
    Nav,
    Sidebar,
    CodeBar,
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
  /* 共通 */
  body {
    word-wrap: break-word;
  }
  /* スマホ版 */
  #code-bar {
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
    #code-bar {
      display: block;
    }
    #sidebar {
      display: block;
      position: fixed;
      top: 60px;
      width: 250px;
      height: 100vh;
      background-color: rgba(227, 225, 230, 0.555);
    }
    main{
      margin-left: 250px;
    }
  }
</style>
