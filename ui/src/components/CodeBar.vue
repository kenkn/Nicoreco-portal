<template>
  <div class=position-relative>
    <img :src="codeImg" height="200" alt="code-bar" class="w-100 p-1">
    <div class="position-absolute top-0 start-0 text-light fw-bold p-3"> 
        {{ message }}
    </div>
    <div class="position-absolute top-0 end-0 text-light text-center p-5">
        <h3><em>Nicoreco Portal</em></h3>
        <p class="m-0">知情生のための</p>
        <p >情報共有ポータルサイト</p>
    </div>
  </div>
</template>

<script>
import { computed } from 'vue'
import { useStore } from 'vuex'
import UserIconMenu from "@/components/UserIconMenu"
import codeImg from '@/assets/img/code.png'

export default {
  name: "CodeBar",
  setup() {
    const store = useStore()
    const auth = computed(() => store.state.auth)
    const message = "講義一覧"
    return {
      message,
      codeImg,
    }
  }
}
</script>

<style>
</style>