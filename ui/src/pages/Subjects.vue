<template>
  <div class="container">
    <h1 class="pb-3 display-5">講義一覧</h1>
    <div class="row">
      <SubjectButton 
        v-for="subject in subjects" 
        :key="subject.name"
        :subject-name="subject.name" 
        :subject-code="subject.code"
      />
    </div>
  </div>
</template>
 
<script>
import { onBeforeUnmount } from 'vue'
import { useStore } from 'vuex'
import data from '../data/subject-data.json'
import SubjectButton from '@/components/SubjectButton'

export default {
  name: "Subjects",
  components: {
    SubjectButton
  },
  setup() {
    const store = useStore()
    // 見出しの処理
    store.dispatch("setJumbotron", "講義一覧")
    onBeforeUnmount(() =>
      store.dispatch("setJumbotron", "")
    )
    return {
      subjects: data
    }
  },
};
</script>
