<template>
  <div class="container">
    <div class="">
      <h1 class="p-3 d-inline-block">{{ subject }}の質問一覧</h1>
      <router-link v-if="!auth" class="pageLink d-inline p-3" to="/login">
        <button type="button" class="btn btn-primary">ログインして質問する</button>
      </router-link>
      <router-link v-else class="pageLink d-inline p-3" :to="'/question/' + code + '/create'">
        <button type="button" class="btn btn-primary">質問する</button>
      </router-link>
    </div>
    
    <div v-if="questions!=null">
      <ul class="list-group">
        <li v-for="question in questions" :key="question.ID" class="list-group-item">
          <div class="row">
            <div class="col-6 col-md-8 col-lg-10 btn">
              <router-link tag="li" class="pageLink" v-bind:to="$route.path  + '/' + question.ID">
                <p>{{ question.title }}</p>
              </router-link>
              <p class="text-right text-secondary m-0">更新日時:{{ question.created_at }}</p>
            </div>
            <div class="col-3 col-md-2 col-lg-1">
              <div>
                <p class="text-center text-secondary m-0 p-2">{{ question.lgtm }}</p>
              </div>
              <div>
                <p class="text-center text-secondary m-0 p-2">Good</p>
              </div>
            </div>
            <div class="col-3 col-md-2 col-lg-1">
              <div>
                <p class="text-center text-secondary m-0 p-2">{{ question.answer_count }}</p>
              </div>
              <div>
                <p class="text-center text-secondary m-0 p-2">回答数</p>
              </div>
            </div>
          </div>

        </li>
      </ul>
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