<template>
  <div class="h-100">
    <Nav />
    <main class="container-fluid h-auto row p-0 pt-5 m-0">
      <div class="col-12 col-md-2">
        <Sidebar />
      </div>
      <div class="container col-12 col-md-10 h-auto border-left overflow-auto">
        <router-view />
      </div>
    </main>
  </div>
</template>
 
<script>
import Nav from "@/components/Nav";
import Sidebar from "@/components/Sidebar";
export default {
  components: {
    Nav,
    Sidebar
  },
};
</script>

<style>
  html {
    height: 100%;
  }
  body{
    height: 100%;
  }
  #app {
    height: 100%;
  }
</style>
