<template>
  <div class="container">
    <h1 class="p-3">講義一覧</h1>
  <div class="row">
    <div v-for="subject in subjects" :key="subject.name" class="col col-12 col-md-3 col-lg-2 m-2 border text-center">
      <router-link tag="li" class="pageLink btn btn-link" v-bind:to="'/question/' + subject.name">
        <p>{{ subject.name }}</p>
      </router-link>
    </div>
  </div>
  </div>
</template>
 
<script>

export default {
  name: "Subjects",
  data() {
    return {
      subjects: [
        {name: "高度ものづくり創成演習Ⅰ"},
        {name: "高度ものづくり創成演習Ⅱ"},
        {name: "高度ものづくり創成演習Ⅲ"},
        {name: "高度ものづくり創成演習Ⅳ"},
        {name: "高度ものづくり創成演習Ⅴ"},
        {name: "高度ものづくり創成演習Ⅵ"},
        ]
    }
  }
};
</script>