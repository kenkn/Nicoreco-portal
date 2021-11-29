<template>
  <div class="container">
    <h1 class="pb-3 display-5">研究室一覧</h1>
  <div class="row">
    <router-link v-for="lab in labs" :key="lab.code"  v-bind:to="'/lab/' + lab.code" tag="li" class="btn btn-link col col-12 col-md-3 col-lg-2 m-3 p-0 text-center text-dark rounded shadow-sm link-box">
      <p class="text-center fw-bolder p-3 mt-0 my-2 w-100 h-100">{{ lab.name }}</p>
    </router-link>
  </div>
  </div>
</template>
 
<script>
import data from '../data/lab-data.json'
export default {
  name: "Labs",
  setup() {
    return {
      labs: data
    }
  },
};
</script>

<style scoped>
  .link-box {
    border-width: 1px;
    border-color: rgb(167, 167, 167);
    background-color: white;
  }
  .link-box :hover{
    background: rgb(230, 229, 227);
  }
</style>