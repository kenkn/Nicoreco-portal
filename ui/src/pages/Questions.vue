<template>
  <div class="container">
    <h1 class="p-3">{{ subject }}の質問一覧</h1>
    <div v-if="questions">
      <ul class="list-group">
        <li v-for="question in questions" :key="question.pk" class="list-group-item">
          <div class="row">
            <div class="col-6 col-md-8 col-lg-10 btn">
              <router-link tag="li" class="pageLink" v-bind:to="'/question/' + subject  + '/' + question.pk">
                <p>{{ question.title }}</p>
              </router-link>
              <p class="text-right text-secondary m-0">更新日時:{{ question.updated_at }}</p>
            </div>
            <div class="col-3 col-md-2 col-lg-1">
              <div>
                <p class="text-center text-secondary m-0 p-2">{{ question.good }}</p>
              </div>
              <div>
                <p class="text-center text-secondary m-0 p-2">Good</p>
              </div>
            </div>
            <div class="col-3 col-md-2 col-lg-1">
              <div>
                <p class="text-center text-secondary m-0 p-2">{{ question.answers }}</p>
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

export default {
  name: "Questions",
  setup() {
    const subject = "高度ものづくり創成演習Ⅰ"
    const questions = []
    const question = {
      pk: 1,
      title: "～について",
      good: 3,
      answers: 5,
      created_at: Date.now(),
      updated_at: Date(1999, 11, 31, 23, 59, 59),
    }
    questions.push(question) 
    questions.push({
      pk: 2,
      title: "～~~~~~~~~~~~~~~について",
      good: 3,
      answers: 5,
      created_at: Date.now(),
      updated_at: Date.now(),
    })

    return {
      subject,
      questions
    }
  }
};
</script>