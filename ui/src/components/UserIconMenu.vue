<template>
  <div @click="menu=!menu" ref="menuEl" class="position-relative p-1" style="cursor: pointer;">
    <UserIcon :width="35" :height="35" />
    <div v-if="menu" id="dropdown-menu">
      <ul class="list-group">
        <li class="list-group-item list-group-item-action p-0">
          <router-link to="/profile" class="nav-link">プロフィール編集</router-link>
        </li>
        <li class="list-group-item list-group-item-action p-0">
          <router-link to="/login" class="nav-link" @click="sendLogout">ログアウト</router-link>
        </li>
      </ul>
    </div>
  </div>
</template>

<script>
import { computed, onMounted, onUnmounted, ref } from 'vue'
import { useStore } from 'vuex'
import UserIcon from "@/components/UserIcon"

export default {
  name: "UserIconMenu",
  components: {
    UserIcon
  },
  setup(props, context) {
    const store  = useStore()
    const menu   = ref(false) //ドロップダウンメニューのon/off
    const menuEl = ref(null)
    const closeMenu = (event) => {
      // @clickより後に実行される
      // メニューが開いていて、メニュー要素以外の部部分をクリックしたときにメニューを閉じる
      if(menu.value && !menuEl.value.contains(event.target)) {
        menu.value = false
      }
    }
    onMounted(() => {
      window.addEventListener('click', closeMenu);
    })
    onUnmounted(() => {
      window.removeEventListener('click', closeMenu);
    })

    const sendLogout = () => {
      context.emit('sendLogout')
    }

    return {
      menu,
      menuEl,
      sendLogout,
    }

  }
};
</script>
<style scoped>
#dropdown-menu {
  position: absolute;
  left: -5rem;
  top: 50px;
}
</style>
