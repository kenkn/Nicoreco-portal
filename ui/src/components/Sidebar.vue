<template>
  <aside class="main-sidebar p-2">
    <!-- sidebar: style can be found in sidebar.less -->
    <nav class="sidebar ">

      <!-- search form (Optional) -->
      <form v-on:submit.prevent class="sidebar-form p-2 pb-2 border-bottom" id="searchForm">
        <div class="input-group" id="searchContainer">
          <input type="text" name="search" id="search" class="search form-control" data-toggle="hideseek" 
          placeholder="Search Menus" data-list=".sidebar-menu" />
          <!-- <button type="submit" name="search" id="search-btn" class="btn bbtn-primary">Search</button> -->
        </div>
      </form>
      <!-- /.search form -->

      <!-- Sidebar Menu -->
      <sidebar-menu />
      <!-- /.sidebar-menu -->
    </nav>
    <!-- /.sidebar -->
  </aside>
</template>
<script>
  import SidebarMenu from './SidebarMenu'

  export default {
    name: 'Sidebar',
    props: ['user'],
    components: { SidebarMenu },
    mounted: function () {
      window.jQuery('[data-toggle="hideseek"]').off().hideseek()
    }
  }
</script>
<style scope="local">
  .user-panel .image img {
    border-radius: 50%;
  }

  #searchForm {
    padding-left: 0em;
    padding-right: 0em;
  }

  #searchContainer {
    height: 100%;
    padding-bottom: 0em;
  }

  #search {
    width: 80%;
    float: right;
  }

  #search-btn {
    width: 20%;
  }
</style>
