<template>
  <div class="container">
    <h1 class="p-b3 display-5">404 Not Found</h1>
  </div>
</template>

<script>
import { useStore } from 'vuex'
export default {
  name: "NotFound",
  data() {
    const store = useStore()
    return{
      store
    }
  },
  watch: {
    $route(){
      this.store.dispatch("setIsNotFound", false)
    }
  }
}
</script>
