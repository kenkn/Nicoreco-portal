<template>
  <div class="container">
    <h1 class="p-b3 display-5">{{ message }}</h1>
  </div>
</template>
 
<script>
import { ref, onMounted } from 'vue'
import axios from 'axios'
import { useStore } from 'vuex'

export default {
  name: "Home",
  setup() {
    const message = ref("ログインしていません")
    const store   = useStore()

    onMounted(async () => {
      try {
        const { data } = await axios.post("user", {
          jwt: localStorage.authToken
        })
        message.value = `こんにちは．${data.display_name} さん`

        // index.tsのactionsに設定したパラメータ名を設定
        store.dispatch("setAuth", true)
      } catch (e) {
        console.log(e)
        store.dispatch("setAuth", false)
      }
    })

    return {
      message
    }
  }
}
</script>