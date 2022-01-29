<template>
  <div class="container">
    <h1 class="pb-3 display-5 d-md-none">リンク集</h1>
    <ul>
      <li v-for="link in links" :key="link.ID" class="m-2"> 
        <a class="m-2" :href="link.url" target="_blank">
            {{link.name}}
            <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-arrow-up-right-square" viewBox="0 0 16 16">
              <path fill-rule="evenodd" d="M15 2a1 1 0 0 0-1-1H2a1 1 0 0 0-1 1v12a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1V2zM0 2a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2v12a2 2 0 0 1-2 2H2a2 2 0 0 1-2-2V2zm5.854 8.803a.5.5 0 1 1-.708-.707L9.243 6H6.475a.5.5 0 1 1 0-1h3.975a.5.5 0 0 1 .5.5v3.975a.5.5 0 1 1-1 0V6.707l-4.096 4.096z"/>
            </svg>
        </a>
      </li>
    </ul>
  </div>
</template>
 
<script>
import { onBeforeUnmount } from 'vue'
import { useStore } from 'vuex'
import links from '../data/links.json'

export default {
  name: "Links",
  setup() {
    const store = useStore()

    // 見出しの処理
    store.dispatch("setJumbotron", "リンク集")
    onBeforeUnmount(() =>
      store.dispatch("setJumbotron", "")
    )
    return {
      links
    }
  },
};
</script>
