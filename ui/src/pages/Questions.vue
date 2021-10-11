<template>
  <div class="container">
    <div class="">
      <h1 class="pb-3 d-inline-block display-5">{{ subject }}の質問一覧</h1>
      <router-link v-if="!auth" class="pageLink d-inline p-3" to="/login">
        <button type="button" class="btn btn-primary">ログインして質問する</button>
      </router-link>
      <router-link v-else class="pageLink d-inline p-3" :to="'/question/' + code + '/create'">
        <button type="button" class="btn btn-primary">質問する</button>
      </router-link>
    </div>
    
    <div v-if="questions!=null">
      <router-link v-for="question in questions" :key="question.ID" v-bind:to="$route.path  + '/' + question.ID" class="list-group-item link-box p-0 text-dark">
        <div class="row m-0">
          <div class="col-6 col-md-8 col-xl-10 btn">
            <p class="fw-bolder">{{ question.title }}</p>
            <p class="text-right text-secondary m-0">更新日時:{{ question.created_at }}</p>
          </div>
          <div class="col-3 col-md-2 col-xl-1">
            <div>
              <p class="text-center text-secondary m-0 p-2">{{ question.lgtm }}</p>
            </div>
            <div>
              <p class="text-center text-secondary m-0 p-2">Good</p>
            </div>
          </div>
          <div class="col-3 col-md-2 col-xl-1 p-0">
            <div>
              <p class="text-center text-secondary m-0 p-2">{{ question.answer_count }}</p>
            </div>
            <div>
              <p class="text-center text-secondary m-0 p-2">回答数</p>
            </div>
          </div>
        </div>
      </router-link>
    </div>
    <div v-else>
      <h3 class="p-3">
        質問はまだありません
      </h3>
    </div>
  </div>
</template>
 
<script>
import { computed, ref, onMounted } from 'vue'
import { useStore } from 'vuex'
import axios from 'axios'
import subjectData from '../data/subject-data.json'

export default {
  name: "Questions",
  data() {
    const store = useStore()
    const auth = computed(() => store.state.auth)
    let code = ''
    const subject = ref("")
    const questions = ref({})
    const answer = ref()
    for (const d of subjectData) {
      if (d.code === this.$route.params.subject) {
        subject.value = d.name
        code = d.code
      }
    }
    onMounted(async () => {
      try {
        const url = "questions/" + code
        const { data } = await axios.get(url)
        questions.value = data
      } catch (e) {
        console.log(e)
      }
    })
    
    return {
      auth,
      subject,
      code,
      questions,
      answer
    }
  }
};
</script>

<style scoped>
  .link-box {
    border-width: 1px;
    border-color: black;
    background-color: white;
  }
  .link-box :hover{
    background: rgb(230, 229, 227);
  }
</style>