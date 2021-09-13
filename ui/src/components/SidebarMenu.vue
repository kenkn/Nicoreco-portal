<template>
  <ul class="list-group pt-2">
    <router-link tag="li" class="pageLink list-group-item" to="/question">講義質問</router-link>
    <router-link tag="li" class="pageLink list-group-item" to="/">研究室</router-link>
    <router-link tag="li" class="pageLink list-group-item" to="/">雑談</router-link>
  </ul>
</template>
<script>
  export default {
    name: 'SidebarMenu'
  }
</script>
<style>
  /* override default */
  .sidebar-menu>li>a {
    padding: 12px 15px 12px 15px;
  }

  .sidebar-menu li.active>a>.fa-angle-left,
  .sidebar-menu li.active>a>.pull-right-container>.fa-angle-left {
    animation-name: rotate;
    animation-duration: 0.2s;
    animation-fill-mode: forwards;
  }

  .treeview-title {
    z-index: 1;
  }

  @keyframes rotate {
    0% {
      transform: rotate(0deg);
    }

    100% {
      transform: rotate(-90deg);
    }
  }
</style>
