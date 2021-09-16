<template>
  <div class="h-100">
    <Nav />
    <main class="container-fluid h-100 pt-5 m-0">
      <div class="row h-100">
        <div class="col-md-2 h-auto">
          <Sidebar />
        </div>
        <div class="col-md-10 h-100 border-left">
          <router-view />
        </div>
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
