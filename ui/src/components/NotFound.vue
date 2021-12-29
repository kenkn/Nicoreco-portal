<template>
  <div class="container">
    <h1 class="p-b3 display-5">404 Not Found</h1>
  </div>
</template>

<script>
import { watch } from 'vue'
import { useRoute } from 'vue-router'
import { useStore } from 'vuex'
export default {
  name: "NotFound",
  setup() {
    const store = useStore()
    const route = useRoute()
    watch(route, () =>
      store.dispatch("setIsNotFound", false)
    )
  }
}
</script>
