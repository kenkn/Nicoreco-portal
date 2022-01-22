<template>
  <div class="container">
    <h1 class="pb-3 display-5">研究室一覧</h1>
    <div class="row">
      <LabButton
        v-for="lab in labs" 
        :key="lab.name"
        :lab-name="lab.name" 
        :lab-code="lab.code"
        :lab-course="lab.course"
      />
    </div>
  </div>
</template>
 
<script>
import { onBeforeUnmount } from 'vue'
import { useStore } from 'vuex'
import LabButton from '@/components/LabButton'
import data from '../data/lab-data.json'

export default {
  name: "Labs",
  components: {
    LabButton
  },
  setup() {
    const store = useStore()
    store.dispatch("setJumbotron", "研究室一覧")
    onBeforeUnmount(() =>
      store.dispatch("setJumbotron", "")
    )
    return {
      labs: data
    }
  },
};
</script>
