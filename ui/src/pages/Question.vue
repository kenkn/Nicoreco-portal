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
        <button @click="updateQuestionLgtm(false)" id="lgtm" class="btn btn-outline-primary lgtm">
          <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-hand-thumbs-up" viewBox="0 0 16 16">
            <path d="M8.864.046C7.908-.193 7.02.53 6.956 1.466c-.072 1.051-.23 2.016-.428 2.59-.125.36-.479 1.013-1.04 1.639-.557.623-1.282 1.178-2.131 1.41C2.685 7.288 2 7.87 2 8.72v4.001c0 .845.682 1.464 1.448 1.545 1.07.114 1.564.415 2.068.723l.048.03c.272.165.578.348.97.484.397.136.861.217 1.466.217h3.5c.937 0 1.599-.477 1.934-1.064a1.86 1.86 0 0 0 .254-.912c0-.152-.023-.312-.077-.464.201-.263.38-.578.488-.901.11-.33.172-.762.004-1.149.069-.13.12-.269.159-.403.077-.27.113-.568.113-.857 0-.288-.036-.585-.113-.856a2.144 2.144 0 0 0-.138-.362 1.9 1.9 0 0 0 .234-1.734c-.206-.592-.682-1.1-1.2-1.272-.847-.282-1.803-.276-2.516-.211a9.84 9.84 0 0 0-.443.05 9.365 9.365 0 0 0-.062-4.509A1.38 1.38 0 0 0 9.125.111L8.864.046zM11.5 14.721H8c-.51 0-.863-.069-1.14-.164-.281-.097-.506-.228-.776-.393l-.04-.024c-.555-.339-1.198-.731-2.49-.868-.333-.036-.554-.29-.554-.55V8.72c0-.254.226-.543.62-.65 1.095-.3 1.977-.996 2.614-1.708.635-.71 1.064-1.475 1.238-1.978.243-.7.407-1.768.482-2.85.025-.362.36-.594.667-.518l.262.066c.16.04.258.143.288.255a8.34 8.34 0 0 1-.145 4.725.5.5 0 0 0 .595.644l.003-.001.014-.003.058-.014a8.908 8.908 0 0 1 1.036-.157c.663-.06 1.457-.054 2.11.164.175.058.45.3.57.65.107.308.087.67-.266 1.022l-.353.353.353.354c.043.043.105.141.154.315.048.167.075.37.075.581 0 .212-.027.414-.075.582-.05.174-.111.272-.154.315l-.353.353.353.354c.047.047.109.177.005.488a2.224 2.224 0 0 1-.505.805l-.353.353.353.354c.006.005.041.05.041.17a.866.866 0 0 1-.121.416c-.165.288-.503.56-1.066.56z"/>
          </svg>
        </button>
        <span class="pl-2">{{ question.lgtm }}件</span>
      </div>
      <div v-else class="mt-1">
        <button @click="updateQuestionLgtm(true)" id="lgtm" class="btn btn-primary lgtm">
          <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-hand-thumbs-up" viewBox="0 0 16 16">
            <path d="M8.864.046C7.908-.193 7.02.53 6.956 1.466c-.072 1.051-.23 2.016-.428 2.59-.125.36-.479 1.013-1.04 1.639-.557.623-1.282 1.178-2.131 1.41C2.685 7.288 2 7.87 2 8.72v4.001c0 .845.682 1.464 1.448 1.545 1.07.114 1.564.415 2.068.723l.048.03c.272.165.578.348.97.484.397.136.861.217 1.466.217h3.5c.937 0 1.599-.477 1.934-1.064a1.86 1.86 0 0 0 .254-.912c0-.152-.023-.312-.077-.464.201-.263.38-.578.488-.901.11-.33.172-.762.004-1.149.069-.13.12-.269.159-.403.077-.27.113-.568.113-.857 0-.288-.036-.585-.113-.856a2.144 2.144 0 0 0-.138-.362 1.9 1.9 0 0 0 .234-1.734c-.206-.592-.682-1.1-1.2-1.272-.847-.282-1.803-.276-2.516-.211a9.84 9.84 0 0 0-.443.05 9.365 9.365 0 0 0-.062-4.509A1.38 1.38 0 0 0 9.125.111L8.864.046zM11.5 14.721H8c-.51 0-.863-.069-1.14-.164-.281-.097-.506-.228-.776-.393l-.04-.024c-.555-.339-1.198-.731-2.49-.868-.333-.036-.554-.29-.554-.55V8.72c0-.254.226-.543.62-.65 1.095-.3 1.977-.996 2.614-1.708.635-.71 1.064-1.475 1.238-1.978.243-.7.407-1.768.482-2.85.025-.362.36-.594.667-.518l.262.066c.16.04.258.143.288.255a8.34 8.34 0 0 1-.145 4.725.5.5 0 0 0 .595.644l.003-.001.014-.003.058-.014a8.908 8.908 0 0 1 1.036-.157c.663-.06 1.457-.054 2.11.164.175.058.45.3.57.65.107.308.087.67-.266 1.022l-.353.353.353.354c.043.043.105.141.154.315.048.167.075.37.075.581 0 .212-.027.414-.075.582-.05.174-.111.272-.154.315l-.353.353.353.354c.047.047.109.177.005.488a2.224 2.224 0 0 1-.505.805l-.353.353.353.354c.006.005.041.05.041.17a.866.866 0 0 1-.121.416c-.165.288-.503.56-1.066.56z"/>
          </svg>
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
        <button @click="updateAnswerLgtm(false, answer.ID)" id="lgtm" class="btn btn-outline-primary lgtm">
          <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-hand-thumbs-up" viewBox="0 0 16 16">
            <path d="M8.864.046C7.908-.193 7.02.53 6.956 1.466c-.072 1.051-.23 2.016-.428 2.59-.125.36-.479 1.013-1.04 1.639-.557.623-1.282 1.178-2.131 1.41C2.685 7.288 2 7.87 2 8.72v4.001c0 .845.682 1.464 1.448 1.545 1.07.114 1.564.415 2.068.723l.048.03c.272.165.578.348.97.484.397.136.861.217 1.466.217h3.5c.937 0 1.599-.477 1.934-1.064a1.86 1.86 0 0 0 .254-.912c0-.152-.023-.312-.077-.464.201-.263.38-.578.488-.901.11-.33.172-.762.004-1.149.069-.13.12-.269.159-.403.077-.27.113-.568.113-.857 0-.288-.036-.585-.113-.856a2.144 2.144 0 0 0-.138-.362 1.9 1.9 0 0 0 .234-1.734c-.206-.592-.682-1.1-1.2-1.272-.847-.282-1.803-.276-2.516-.211a9.84 9.84 0 0 0-.443.05 9.365 9.365 0 0 0-.062-4.509A1.38 1.38 0 0 0 9.125.111L8.864.046zM11.5 14.721H8c-.51 0-.863-.069-1.14-.164-.281-.097-.506-.228-.776-.393l-.04-.024c-.555-.339-1.198-.731-2.49-.868-.333-.036-.554-.29-.554-.55V8.72c0-.254.226-.543.62-.65 1.095-.3 1.977-.996 2.614-1.708.635-.71 1.064-1.475 1.238-1.978.243-.7.407-1.768.482-2.85.025-.362.36-.594.667-.518l.262.066c.16.04.258.143.288.255a8.34 8.34 0 0 1-.145 4.725.5.5 0 0 0 .595.644l.003-.001.014-.003.058-.014a8.908 8.908 0 0 1 1.036-.157c.663-.06 1.457-.054 2.11.164.175.058.45.3.57.65.107.308.087.67-.266 1.022l-.353.353.353.354c.043.043.105.141.154.315.048.167.075.37.075.581 0 .212-.027.414-.075.582-.05.174-.111.272-.154.315l-.353.353.353.354c.047.047.109.177.005.488a2.224 2.224 0 0 1-.505.805l-.353.353.353.354c.006.005.041.05.041.17a.866.866 0 0 1-.121.416c-.165.288-.503.56-1.066.56z"/>
          </svg>
        </button>
        <span class="pl-2">{{ answerLgtmCount[answer.ID] }}件</span>
      </div>
      <div v-else class="mt-1">
        <button @click="updateAnswerLgtm(true, answer.ID)" id="lgtm" class="btn btn-primary lgtm">
          <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-hand-thumbs-up" viewBox="0 0 16 16">
            <path d="M8.864.046C7.908-.193 7.02.53 6.956 1.466c-.072 1.051-.23 2.016-.428 2.59-.125.36-.479 1.013-1.04 1.639-.557.623-1.282 1.178-2.131 1.41C2.685 7.288 2 7.87 2 8.72v4.001c0 .845.682 1.464 1.448 1.545 1.07.114 1.564.415 2.068.723l.048.03c.272.165.578.348.97.484.397.136.861.217 1.466.217h3.5c.937 0 1.599-.477 1.934-1.064a1.86 1.86 0 0 0 .254-.912c0-.152-.023-.312-.077-.464.201-.263.38-.578.488-.901.11-.33.172-.762.004-1.149.069-.13.12-.269.159-.403.077-.27.113-.568.113-.857 0-.288-.036-.585-.113-.856a2.144 2.144 0 0 0-.138-.362 1.9 1.9 0 0 0 .234-1.734c-.206-.592-.682-1.1-1.2-1.272-.847-.282-1.803-.276-2.516-.211a9.84 9.84 0 0 0-.443.05 9.365 9.365 0 0 0-.062-4.509A1.38 1.38 0 0 0 9.125.111L8.864.046zM11.5 14.721H8c-.51 0-.863-.069-1.14-.164-.281-.097-.506-.228-.776-.393l-.04-.024c-.555-.339-1.198-.731-2.49-.868-.333-.036-.554-.29-.554-.55V8.72c0-.254.226-.543.62-.65 1.095-.3 1.977-.996 2.614-1.708.635-.71 1.064-1.475 1.238-1.978.243-.7.407-1.768.482-2.85.025-.362.36-.594.667-.518l.262.066c.16.04.258.143.288.255a8.34 8.34 0 0 1-.145 4.725.5.5 0 0 0 .595.644l.003-.001.014-.003.058-.014a8.908 8.908 0 0 1 1.036-.157c.663-.06 1.457-.054 2.11.164.175.058.45.3.57.65.107.308.087.67-.266 1.022l-.353.353.353.354c.043.043.105.141.154.315.048.167.075.37.075.581 0 .212-.027.414-.075.582-.05.174-.111.272-.154.315l-.353.353.353.354c.047.047.109.177.005.488a2.224 2.224 0 0 1-.505.805l-.353.353.353.354c.006.005.041.05.041.17a.866.866 0 0 1-.121.416c-.165.288-.503.56-1.066.56z"/>
          </svg>
        </button>
        <span class="pl-2">{{ answerLgtmCount[answer.ID] }}件</span>
      </div>
      <hr>
      <template v-for="reply in replys" :key="reply.ID">
        <div v-if="reply.parent_id==answer.ID" class="border p-2 ml-2 mb-3">
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

        // ログインしていない場合LGTMボタンをdisabledにする
        console.log(store.state.auth)
        if(!store.state.auth){
          const lgtmButtons = document.getElementsByClassName('lgtm')
          for(let i=0; i < lgtmButtons.length; i++){
            console.log(lgtmButtons[i])
            lgtmButtons[i].classList.add("disabled")
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