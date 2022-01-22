<template>
  <div class='lab-btn position-relative bg-Secondary m-2 p-0'>
    <router-link v-bind:to="'/lab/' + labCode" class="stretched-link">
      <img :src="labLogo" alt="lab-logo" class="lab-img position-absolute top-50 start-50 translate-middle">
      <p class="position-absolute top-50 start-50 translate-middle text-light text-nowrap fw-bold fs-5">{{ labName }}</p>
    </router-link>
  </div>
</template>

<script>
import { ref } from 'vue'

export default {
  name: "SubjectButton",
  props: {
    labName: {
      type: String,
      require: true,
    },
    labCode: {
      type: String,
      require: true,
    },
    labCourse: {
      type: String,
      require: true,
    },
  },
  setup(props) {
    const labLogo = ref()
    labLogo.value = require('@/assets/img/lab_icons/' + props.labCourse + '.png')
    return {
      labLogo,
    }

  }
}
</script>

<style scoped>
  .stretched-link:hover::after {
    background: black;
    opacity: 0.1;
  }
  .lab-btn {
    width: 12rem;
    height: 12rem;
    background-color: gray;
  }
  .lab-img {
    width: 12rem;
    height: 12rem;
    /* ぼかし具合 */
    filter: blur(4px);
  }
</style>
