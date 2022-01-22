<template>
  <img :src="userIcon" alt="icon" :width="width" :height="height">
</template>

<script>
import { ref, computed } from 'vue'
import { useStore } from 'vuex'
import b1Icon from "@/assets/img/user_icons/B1.png";
import b2Icon from "@/assets/img/user_icons/B2.png";
import b3Icon from "@/assets/img/user_icons/B3.png";
import b4Icon from "@/assets/img/user_icons/B4.png";
import m1Icon from "@/assets/img/user_icons/M1.png";
import m2Icon from "@/assets/img/user_icons/M2.png";

export default {
  name: "UserIcon",
  props: {
    width: {
      type: Number,
      require: true,
    },
    height: {
      type: Number,
      require: true,
    },
  },
  setup(props, context) {
    const store  = useStore()
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

    return {
      userIcon,
    }

  }
};
</script>
