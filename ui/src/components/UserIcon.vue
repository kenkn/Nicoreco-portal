<template>
  <div v-if="userIcon" @click="menu=!menu" ref="menuEl" class="position-relative p-1">
    <img :src="userIcon" alt="icon" id="user-icon">
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
import b1Icon from "@/assets/img/user_icons/B1.png";
import b2Icon from "@/assets/img/user_icons/B2.png";
import b3Icon from "@/assets/img/user_icons/B3.png";
import b4Icon from "@/assets/img/user_icons/B4.png";
import m1Icon from "@/assets/img/user_icons/M1.png";
import m2Icon from "@/assets/img/user_icons/M2.png";

export default {
  name: "UserIcon",
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

    const grade = computed(() => store.state.grade)
    const userIcon  = ref(null)
    switch (grade.value) {
      case '学部1年':
        userIcon.value = b1Icon
        break
      case '学部2年':
        userIcon.value = b2Icon
        break
      case '学部3年':
        userIcon.value = b3Icon
        break
      case '学部4年':
        userIcon.value = b4Icon
        break
      case '修士1年':
        userIcon.value = m1Icon
        break
      case '修士2年':
        userIcon.value = m2Icon
        break
    }

    const sendLogout = () => {
      context.emit('sendLogout')
    }

    return {
      menu,
      menuEl,
      userIcon,
      sendLogout,
    }

  }
};
</script>
<style scoped>
#user-icon {
  width: 35px;
  height: 35px;
  cursor: pointer;
}
#dropdown-menu {
  position: absolute;
  left: -5rem;
  top: 50px;
}
</style>
