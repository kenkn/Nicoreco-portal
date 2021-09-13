<template>
  <div class="h-100 overflow-hidden">
    <Nav />
    <div class="container-fluid h-100 row p-0">
      <div class="col-12 col-md-2">
        <Sidebar />
      </div>
      <div class="col-12 col-md-10 h-100 border-left overflow-auto">
        <router-view />
      </div>
    </div>
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
