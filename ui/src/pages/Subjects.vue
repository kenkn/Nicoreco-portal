<template>
  <div class="container">
    <h1 class="p-3">講義一覧</h1>
  <div class="row">
    <div v-for="subject in subjects" :key="subject.name" class="col col-12 col-md-3 col-lg-2 m-2 border text-center">
      <router-link tag="li" class="pageLink btn btn-link" v-bind:to="'/question/' + subject.code">
        <p>{{ subject.name }}</p>
      </router-link>
    </div>
  </div>
  </div>
</template>
 
<script>
import data from '../data/subject-data.json'
export default {
  data() {
    return {
      subjects: data
    }
  },
};
</script>