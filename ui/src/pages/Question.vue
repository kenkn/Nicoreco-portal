<template>
  <div>
    <div id="question" class="p-3 mt-2 border">
      <h2>{{ question.title }}</h2>
            <p>
        {{ question.body }}
      </p>
      <!-- デバッグ用 TODO 消す -->
      <span class="text-secondary m-0">ID: {{ question.ID }} </span>
      <span class="text-secondary m-0">質問者: {{ question.questionerID }} </span>
      <span class="text-secondary m-0 pl-3">質問日時: {{ question.CreatedAt }} </span><br>
      <div v-if='!questionLgtm' class="mt-1">
        <button @click="updateQuestionLgtm(false)" id="lgtm" class="btn btn-outline-primary">
          私も知りたい
        </button>
        <span class="pl-2">{{ question.lgtm }}件</span>
      </div>
      <div v-else class="mt-1">
        <button @click="updateQuestionLgtm(true)" id="lgtm" class="btn btn-primary">
          私も知りたい
        </button>
        <span class="pl-2">{{ question.lgtm }}件</span>
      </div>
    </div>
    <h1 class="p-3">{{ answers.length }}件の回答</h1>
  
    <div v-for="answer in answers" :key="answer.ID" class="border p-3 mt-4">
      <p>{{ answer.body }}</p>
      <!-- デバッグ用 TODO 消す -->
      <span class="text-secondary m-0">ID: {{ answer.ID }} </span>
      <span class="text-secondary m-0">回答者: {{ answer.user_id }} </span>
      <span class="text-secondary m-0 pl-3">回答日時: {{ answer.CreatedAt }} </span><br>
      <div v-if="!answerLgtm[answer.ID]" class="mt-1">
        <button @click="updateAnswerLgtm(false, answer.ID)" id="lgtm" class="btn btn-outline-primary">
          参考になった
        </button>
        <span class="pl-2">{{ answerLgtmCount[answer.ID] }}件</span>
      </div>
      <div v-else class="mt-1">
        <button @click="updateAnswerLgtm(true, answer.ID)" id="lgtm" class="btn btn-primary">
          参考になった
        </button>
        <span class="pl-2">{{ answerLgtmCount[answer.ID] }}件</span>
      </div>
      <hr>
      <template v-for="reply in replys" :key="reply.ID">
        <div v-if="reply.parent_id==answer.ID" class="border p-2 ml-2">
          <p>{{ reply.body }}</p>
          <span class="text-secondary m-0">返信者: {{ reply.user_id }} </span>
          <span class="text-secondary m-0 pl-3">返信日時: {{ reply.CreatedAt }} </span><br>
        </div>
      </template>
      <form action="" @submit.prevent="submitReply(answer.ID)">
        <router-link v-if="!auth" class="pageLink d-inline form-control btn btn-outline-primary" to="/login">
          ログインして返信
        </router-link>
        <div v-else class="form-group ml-2">
          <button class="btn btn-outline-primary form-control w-100" @click="displayReplyForm(answer.ID)" v-bind:id='"reply-disply-button-" + answer.ID'>返信を追加</button>
          <div v-bind:id='"reply-form-" + answer.ID' style="display: none;">
            <textarea v-model="replyBody[answer.ID]" class="form-control p-1 my-4" placeholder='回答に返信' required />
            <button class="btn btn-outline-primary form-control w-100" type="submit">返信</button>
          </div>
        </div>
      </form>
    </div>
    <div>
      <h2 class="p-3">回答する</h2>
      <form action="" @submit.prevent="submitAnswer">
        <router-link v-if="!auth" class="pageLink d-inline btn btn-outline-primary" to="/login">
          ログインして回答を送信
        </router-link>
        <div v-else class="form-group">
          <textarea v-model="answerBody" class="form-control p-1 my-2 mb-4" rows="4" placeholder='回答を追加' required />
          <button class="btn btn-outline-primary w-100" type="submit">回答を送信</button>
        </div>
      </form>
    </div>
  </div>
</template>
 
<script>
import { computed, onMounted, ref } from 'vue'
import axios from 'axios';
import { useStore } from 'vuex';
export default {
  name: "Question",
  data() {
    const store           = useStore()
    const auth            = computed(() => store.state.auth)
    const question        = ref({}) // questionの内容
    const answers         = ref({}) // 投稿されているanswerの集合
    const answerBody      = ref("") // 投稿時のanswerの文章
    const answerLgtmCount = ref([]) // answerのLGTM数
    const replys          = ref({}) // 投稿されているreplyの集合
    const replyBody       = ref([]) // 投稿時のreplyの文章
    const questionLgtm    = ref()   // ユーザがquestionをLGTMしているかどうか
    const answerLgtm      = ref([]) // ユーザがquestionをLGTMしているかどうか

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
        for (const i in answerData.data) {
          answerLgtmCount.value[answerData.data[i].ID] = answerData.data[i].lgtm
        }

        // リプライ情報の取得
        const replyData = await axios.get(
          "/reply/" + questionData.data.ID
        )
        replys.value = replyData.data

        // LGTM情報の取得
        if (localStorage.isLogin) {
          const questionLgtmData = await axios.get(
            "/lgtm/question/" + questionData.data.ID + "/" + localStorage.userID
          )
          // ユーザはLGTMしているか?
          if (questionLgtmData.data.user_id == "") {
            questionLgtm.value = false
          } else {
            questionLgtm.value = true
          }

          // answerに対してユーザがLGTMしているか?
          for (const d in answerData.data) {
            const id = answerData.data[d].ID
            const answerLgtmData = await axios.get(
              "/lgtm/answer/" + id + "/" + localStorage.userID
            )
            if (answerLgtmData.data.user_id == "") {
              answerLgtm.value[id] = false
            } else {
              answerLgtm.value[id] = true
            }
          }
        }
      } catch (e) {
        console.log(e)
      }
    })

    const submitAnswer = async () => {
      try {
        await axios.post("answer/post", {
          parent_id : this.$route.params.question_id,
          user_id : localStorage.userID,
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
        await axios.post("reply/post", {
          question_id : this.$route.params.question_id,
          parent_id : String(id),
          user_id : localStorage.userID,
          body : replyBody.value[id]
        })
        // リロード
        this.$router.go({path: this.$router.currentRoute.path, force: true})
      } catch (e) {
        console.log(e)
      }
    }

    // questionのLGTMボタンが押された時
    // TODO destroy時までLGTM情報を持っておいて，一度に送信させる
    const updateQuestionLgtm = async (lgtmed) => {
      questionLgtm.value = !questionLgtm.value
      if (!lgtmed) {
        question.value.lgtm++
      } else {
        question.value.lgtm--
      }
      try {
        await axios.post("/lgtm/question", {
          question_id : this.$route.params.question_id,
          user_id : localStorage.userID
        })
      } catch (e) {
        console.log(e)
      }
    }

    // answerのLGTMボタンが押された時
    const updateAnswerLgtm = async (lgtmed, id) => {
      answerLgtm.value[id] = !answerLgtm.value[id]
      if (!lgtmed) {
        answerLgtmCount.value[id]++
      } else {
        answerLgtmCount.value[id]--
      }
      try {
        await axios.post("/lgtm/answer", {
          answer_id : String(id),
          user_id : localStorage.userID
        })
      } catch (e) {
        console.log(e)
      }
    }

    // 「返信を追加」ボタンを非表示にし, リプライフォームを表示する
    const displayReplyForm = (answerid) => {
      document.getElementById("reply-disply-button-" + answerid).style.display = "none";
      document.getElementById("reply-form-" + answerid).style.display = "block";
    }

    return {
      auth,
      question,
      answers,
      answerLgtmCount,
      replys,
      answerBody,
      replyBody,
      questionLgtm,
      answerLgtm,
      submitAnswer,
      submitReply,
      updateQuestionLgtm,
      updateAnswerLgtm,
      displayReplyForm
    }
  }
};
</script>