<template>
  <div class='card m-2' style="width: 12rem;">
    <img :src="subjectLogo" alt="subject-logo" class="bd-placeholder-img card-img-top border-bottom" width="100%" style="height: 11rem;">
    <div class="card-body">
      <h5 class="card-title text-center">{{ subjectName }}</h5>
      <router-link v-bind:to="'/question/' + subjectCode" class="stretched-link"></router-link>
    </div>
  </div>
</template>

<script>
import { ref, onMounted } from 'vue'

export default {
  name: "SubjectButton",
  props: {
    subjectName: {
      type: String,
      require: true,
    },
    subjectCode: {
      type: String,
      require: true,
    },
  },
  setup(props) {
    const subjectLogo = ref()
    subjectLogo.value = require('@/assets/img/subject_icons/' + props.subjectCode + '.png')
    return {
      subjectLogo,
    }

  }
}
</script>

<style scoped>
  /* 共通 */
  .stretched-link:hover::after {
    background: black;
    opacity: 0.1;
  }
</style>
