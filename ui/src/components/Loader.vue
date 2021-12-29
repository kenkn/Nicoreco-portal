<template>
  <div>
    <div id="loader-bg">
    </div>
    <div id="loader">
      <div class="spinner-border" role="status">
        <span class="visually-hidden">Loading...</span>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "Loader",
}
</script>
<style scoped>
#loader-bg {
  position: fixed;
  background-color: grey;
  opacity: 0.7;
  top: 0;
  left: 0;
  height: 100vh;
  width: 100vw;
}
#loader {
  position: fixed;
  top: 50vh;
  left: 50vw;
}

</style>
