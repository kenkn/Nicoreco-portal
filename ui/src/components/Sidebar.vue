<template>
  <aside class="main-sidebar">
    <!-- サイドバー -->
    <section class="sidebar">
      <!-- ユーザ情報 -->
      <div class="user-panel">
        <!-- <div class="pull-left image"> -->
          <!-- <img :src="user.avatar" /> -->
        <!-- </div> -->
        <div class="pull-left info">
          <!-- TODO ユーザのdisplayNameに置き換える -->
          <div>
            <p class="white">陰キャワイ</p>
          </div>
          <a href="javascript:;">
            <i class="fa fa-circle text-success"></i> ONLINE
          </a>
        </div>
      </div>

      <!-- 検索ウインドウ -->
      <form v-on:submit.prevent class="sidebar-form" id="searchForm">
        <div class="input-group" id="searchContainer">
          <span class="input-group-btn">
            <input type="text" name="search" id="search" class="search form-control" data-toggle="hideseek"
              placeholder="検索" data-list=".sidebar-menu" />
            <button type="submit" name="search" id="search-btn" class="btn btn-flat">
              <i class="fa fa-search"></i>
            </button>
          </span>
        </div>
      </form>

      <!-- Sidebar Menu -->
      <sidebar-menu />
      <!-- /.sidebar-menu -->
    </section>
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
