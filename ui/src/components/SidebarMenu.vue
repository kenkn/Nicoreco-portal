<template>
  <ul class="sidebar-menu">
    <li class="header">カテゴリ</li>
    <router-link tag="li" class="pageLink" to="/question">
      <a>
        <i class="fa fa-clipboard"></i>
        <span class="page">科目の質問</span>
        <small class="label pull-right bg-green">New</small>
      </a>
    </router-link>
    <router-link tag="li" class="pageLink" to="/laboreview">
      <a>
        <i class="fa fa-flask"></i>
        <span class="page">研究室レビュー</span>
        <small class="label pull-right bg-green">New</small>
      </a>
    </router-link>

    <router-link tag="li" class="pageLink" to="/chat">
      <a>
        <i class="fa fa-comments"></i>
        <span class="page">雑談</span>
        <small class="label pull-right bg-green">New</small>
      </a>
    </router-link>

    <li class="header">プライベート</li>
    <!-- TODO 動的スレッド表示にする(もしくはなくす？) -->
    <li class="treeview">
      <a href="#">
        <i class="fa fa-star"></i>
        <span class="treeview-title">お気に入りスレッド</span>
        <span class="pull-right-container pull-right">
          <i class="fa fa-angle-left fa-fw"></i>
        </span>
      </a>
      <ul class="treeview-menu">
        <li>
          <a href="#"> <i class="fa fa-file-word-o"></i> スレ1 </a>
        </li>
        <li>
          <a href="#"> <i class="fa fa-file-picture-o"></i> スレ2 </a>
        </li>
        <li>
          <a href="#"> <i class="fa fa-file-pdf-o"></i> スレ3 </a>
        </li>
      </ul>
    </li>

    <li class="header">その他リンク</li>

    <router-link tag="li" class="pageLink" to="/">
      <a>
        <i class="fa fa-book"></i>
        <span class="page">AtCoder</span>
      </a>
    </router-link>

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
