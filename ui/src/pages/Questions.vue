<template>
  <div class="container">
    <h1 class="p-3">{{ subject }}の質問一覧</h1>
    <div v-if="questions">
      <ul class="list-group">
        <li v-for="question in questions" :key="question.id" class="list-group-item">
          <div class="row">
            <div class="col-6 col-md-8 col-lg-10 btn">
              <router-link tag="li" class="pageLink" v-bind:to="'/question/' + subject  + '/' + question.id">
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
                <p class="text-center text-secondary m-0 p-2">{{ answer }}</p>
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
      <h3 class="m-5">
        質問はまだありません
      </h3>
    </div>

    
    <h2>

    </h2>
  </div>
</template>
 
<script>
import { ref, onMounted } from 'vue'
import axios from 'axios'
import subjectData from '../data/subject-data.json'

export default {
  name: "Questions",
  data() {
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
        const url = "question/" + code
        const { data } = await axios.get(url)
        questions.value = data
      } catch (e) {
        console.log(e)
      }
    })

    // TODO 回答数をモデルに追加した後実装する
    answer.value = 77777

    return {
      subject,
      questions,
      answer
    }
  }
};
</script>