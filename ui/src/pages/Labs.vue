<template>
  <div class="container">
    <h1 class="p-3">研究室一覧</h1>
  <div class="row">
    <div v-for="lab in labs" :key="lab.code" class="col col-12 col-md-3 col-lg-2 m-2 text-center border border-3">
      <router-link tag="li" class="pageLink btn btn-link" v-bind:to="'/lab/' + lab.code">
        <p>{{ lab.name }}</p>
      </router-link>
    </div>
  </div>
  </div>
</template>
 
<script>
import data from '../data/lab-data.json'
export default {
  name: "Labs",
  data() {
    return {
      labs: data
    }
  },
};
</script>