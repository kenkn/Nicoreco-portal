<template>
  <div class="h-100">
    <Nav />
    <main class="container-fluid h-100 m-0">
      <div class="row h-100">
        <div id="sidebar" class="col-md-2 h-auto">
          <Sidebar />
        </div>
        <div id="contents" class="col-md-10 h-100 border-left border-dark overflow-auto px-5 pt-5">
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
  body {
    height: 100%;
  }
  #app {
    height: 100%;
  }
  main > div {
    padding-top: 3.5em;
  }
  #sidebar {
    background-color: rgba(227, 225, 230, 0.555);
  }
  #contents {
    background-color: rgba(213, 236, 236, 0.726);
  }
</style>
