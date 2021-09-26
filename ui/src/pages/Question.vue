<template>
  <div>
    <div id="question" class="p-3 border">
      <h2>{{ question.title }}</h2>
            <p>
        {{ question.body }}
      </p>
      <span class="text-secondary m-0">質問者: {{ question.questionerID }} </span>
      <span class="text-secondary m-0 pl-3">質問日時: {{ question.CreatedAt }} </span><br>
      <div class="mt-1">
        <button id="lgtm" class="btn btn-outline-primary">
          私も知りたい
        </button>
        <span class="pl-2">{{ question.lgtm }}件</span>
      </div>
    </div>
    <h1 class="p-3">{{ answers.length }}件の回答</h1>
  
    <div v-for="answer in answers" :key="answer.ID" class="border p-3 mt-4">
      <p>{{ answer.body }}</p>
      <span class="text-secondary m-0">回答者: {{ answer.user_id }} </span>
      <span class="text-secondary m-0 pl-3">回答日時: {{ answer.CreatedAt }} </span><br>
      <div class="mt-1">
        <button id="lgtm" class="btn btn-outline-primary">
          参考になった
        </button>
        <span class="pl-2">{{ question.lgtm }}件</span>
      </div>
      <hr>
      <template v-for="reply in replys" :key="reply.ID">
        <div v-if="reply.parent_id==answer.ID" class="border p-2 ml-2">
          <p>{{ reply.body }}</p>
          <span class="text-secondary m-0">返信者: {{ reply.user_id }} </span>
          <span class="text-secondary m-0 pl-3">返信日時: {{ reply.CreatedAt }} </span><br>
          <div class="mt-1">
            <button id="lgtm" class="btn btn-outline-primary">
            参考になった
            </button>
            <span class="pl-2">{{ reply.lgtm }}件</span>
          </div>
        </div>
      </template>
      <button class="btn btn-outline-primary form-control w-100" @click="displayReplyForm(answer.ID)" v-bind:id='"reply-button-" + answer.ID'>返信を追加</button>
        <form action="" @submit.prevent="submitReply(answer.ID)" v-bind:id='"reply-form-" + answer.ID' style="display: none;">
          <div class="form-group ml-2">
            <textarea v-model="replyBody[answer.ID]" class="form-control p-1 my-4" placeholder='回答に返信' required />
            <button class="btn btn-outline-primary form-control w-100" type="submit">返信</button>
          </div>
        </form>
    </div>
    <div>
      <h2 class="p-3">回答する</h2>
      <form action="" @submit.prevent="submitAnswer">
        <div class="form-group">
        <textarea v-model="answerBody" class="form-control p-1 my-2 mb-4" rows="4" placeholder='回答を追加' required />
        <button class="btn btn-outline-primary w-100" type="submit">回答を送信</button>
        </div>
      </form>
    </div>

  </div>

</template>
 
<script>
import { onMounted, ref } from 'vue'
import axios from 'axios';
export default {
  name: "Question",
  data() {
    const answerBody = ref("")
    const question = ref({})
    const answers = ref({})
    const replys = ref({})
    const replyBody = ref([])

    onMounted(async () => {
      try {
        // 質問情報の取得
        const questionData = await axios.get(
          "/question/detail/" + this.$route.params.question_id
        )
        question.value = questionData.data

        // 回答情報の取得
        const answerData = await axios.get(
          "/answer/" + questionData.data.ID
        )
        answers.value = answerData.data

        // リプライ情報の取得
        const replyData = await axios.get(
          "/reply/" + questionData.data.ID
        )
        replys.value = replyData.data
      } catch (e) {
        console.log(e)
      }
    })

    const submitAnswer = async () => {
      try {
        const userData = await axios.get("user")
        await axios.post("answer/post", {
          parent_id : this.$route.params.question_id,
          user_id : userData.data.user_id,
          body : answerBody.value
        })
        // リロード
        this.$router.go({path: this.$router.currentRoute.path, force: true})
      } catch (e) {
        console.log(e)
      }
    }

    const submitReply = async (id) => {
      try {
        const userData = await axios.get("user")
        await axios.post("reply/post", {
          question_id : this.$route.params.question_id,
          parent_id : String(id),
          user_id : userData.data.user_id,
          body : replyBody.value[id]
        })
        // リロード
        this.$router.go({path: this.$router.currentRoute.path, force: true})
      } catch (e) {
        console.log(e)
      }
    }

    const displayReplyForm = (answerid) => {
      document.getElementById("reply-button-" + answerid).style.display = "none";
      document.getElementById("reply-form-" + answerid).style.display = "block";
    }

    return {
      question,
      answers,
      replys,
      answerBody,
      replyBody,
      submitAnswer,
      submitReply,
      displayReplyForm
    }
  }
};
</script>