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
import data from '../data/subject-data.json'
import SubjectButton from '@/components/SubjectButton'

export default {
  name: "Subjects",
  components: {
    SubjectButton
  },
  setup() {
    return {
      subjects: data
    }
  },
};
</script>
