<template>
  <div class="container">
    <!-- ロード画面 -->
    <Loader v-if="loading"></Loader>
    <!-- コンテンツ -->
    <div v-else>
      <h1 class="pb-3 d-inline-block display-5 d-md-none">{{ labName }}のレビュー</h1>
      <!-- 一覧ページへのリンク -->
      <div class="mb-2 ml-2">
        <router-link :to="'/lab/' + labCode">
         <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-arrow-left" viewBox="0 0 16 16">
           <path fill-rule="evenodd" d="M15 8a.5.5 0 0 0-.5-.5H2.707l3.147-3.146a.5.5 0 1 0-.708-.708l-4 4a.5.5 0 0 0 0 .708l4 4a.5.5 0 0 0 .708-.708L2.707 8.5H14.5A.5.5 0 0 0 15 8z"/>
         </svg>
        {{ labName }}のレビュー一覧に戻る
        </router-link>
      </div>
      <div class="mt-5 border border-dark bg-white rounded">
        <div class="border-dark p-4 mt-2">
          <div class="border p-3 mb-2 shadow-sm">
            <h4>{{ review.labreview.body }}</h4>
            <span class="text-secondary m-0">レビュー者: {{ review.labreview.labreviewer_id }} </span>
            <span class="text-secondary m-0 pl-3">レビュー日時: {{ formatDate(review.labreview.CreatedAt) }} </span>
            <!-- ここにlgtmボタンの分岐 -->
            <div v-if='!review.islgtmed' class="mt-1">
              <button @click="updateReviewLgtm" id="lgtm" class="btn btn-outline-primary lgtm">
                <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-hand-thumbs-up" viewBox="0 0 16 16">
                  <path d="M8.864.046C7.908-.193 7.02.53 6.956 1.466c-.072 1.051-.23 2.016-.428 2.59-.125.36-.479 1.013-1.04 1.639-.557.623-1.282 1.178-2.131 1.41C2.685 7.288 2 7.87 2 8.72v4.001c0 .845.682 1.464 1.448 1.545 1.07.114 1.564.415 2.068.723l.048.03c.272.165.578.348.97.484.397.136.861.217 1.466.217h3.5c.937 0 1.599-.477 1.934-1.064a1.86 1.86 0 0 0 .254-.912c0-.152-.023-.312-.077-.464.201-.263.38-.578.488-.901.11-.33.172-.762.004-1.149.069-.13.12-.269.159-.403.077-.27.113-.568.113-.857 0-.288-.036-.585-.113-.856a2.144 2.144 0 0 0-.138-.362 1.9 1.9 0 0 0 .234-1.734c-.206-.592-.682-1.1-1.2-1.272-.847-.282-1.803-.276-2.516-.211a9.84 9.84 0 0 0-.443.05 9.365 9.365 0 0 0-.062-4.509A1.38 1.38 0 0 0 9.125.111L8.864.046zM11.5 14.721H8c-.51 0-.863-.069-1.14-.164-.281-.097-.506-.228-.776-.393l-.04-.024c-.555-.339-1.198-.731-2.49-.868-.333-.036-.554-.29-.554-.55V8.72c0-.254.226-.543.62-.65 1.095-.3 1.977-.996 2.614-1.708.635-.71 1.064-1.475 1.238-1.978.243-.7.407-1.768.482-2.85.025-.362.36-.594.667-.518l.262.066c.16.04.258.143.288.255a8.34 8.34 0 0 1-.145 4.725.5.5 0 0 0 .595.644l.003-.001.014-.003.058-.014a8.908 8.908 0 0 1 1.036-.157c.663-.06 1.457-.054 2.11.164.175.058.45.3.57.65.107.308.087.67-.266 1.022l-.353.353.353.354c.043.043.105.141.154.315.048.167.075.37.075.581 0 .212-.027.414-.075.582-.05.174-.111.272-.154.315l-.353.353.353.354c.047.047.109.177.005.488a2.224 2.224 0 0 1-.505.805l-.353.353.353.354c.006.005.041.05.041.17a.866.866 0 0 1-.121.416c-.165.288-.503.56-1.066.56z"/>
                </svg>
              </button>
              <span class="pl-2">{{ review.labreview.lgtm }}件</span>
            </div>
            <div @click="updateReviewLgtm" v-else class="mt-1">
              <button id="lgtm" class="btn btn-primary lgtm">
                <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-hand-thumbs-up" viewBox="0 0 16 16">
                  <path d="M8.864.046C7.908-.193 7.02.53 6.956 1.466c-.072 1.051-.23 2.016-.428 2.59-.125.36-.479 1.013-1.04 1.639-.557.623-1.282 1.178-2.131 1.41C2.685 7.288 2 7.87 2 8.72v4.001c0 .845.682 1.464 1.448 1.545 1.07.114 1.564.415 2.068.723l.048.03c.272.165.578.348.97.484.397.136.861.217 1.466.217h3.5c.937 0 1.599-.477 1.934-1.064a1.86 1.86 0 0 0 .254-.912c0-.152-.023-.312-.077-.464.201-.263.38-.578.488-.901.11-.33.172-.762.004-1.149.069-.13.12-.269.159-.403.077-.27.113-.568.113-.857 0-.288-.036-.585-.113-.856a2.144 2.144 0 0 0-.138-.362 1.9 1.9 0 0 0 .234-1.734c-.206-.592-.682-1.1-1.2-1.272-.847-.282-1.803-.276-2.516-.211a9.84 9.84 0 0 0-.443.05 9.365 9.365 0 0 0-.062-4.509A1.38 1.38 0 0 0 9.125.111L8.864.046zM11.5 14.721H8c-.51 0-.863-.069-1.14-.164-.281-.097-.506-.228-.776-.393l-.04-.024c-.555-.339-1.198-.731-2.49-.868-.333-.036-.554-.29-.554-.55V8.72c0-.254.226-.543.62-.65 1.095-.3 1.977-.996 2.614-1.708.635-.71 1.064-1.475 1.238-1.978.243-.7.407-1.768.482-2.85.025-.362.36-.594.667-.518l.262.066c.16.04.258.143.288.255a8.34 8.34 0 0 1-.145 4.725.5.5 0 0 0 .595.644l.003-.001.014-.003.058-.014a8.908 8.908 0 0 1 1.036-.157c.663-.06 1.457-.054 2.11.164.175.058.45.3.57.65.107.308.087.67-.266 1.022l-.353.353.353.354c.043.043.105.141.154.315.048.167.075.37.075.581 0 .212-.027.414-.075.582-.05.174-.111.272-.154.315l-.353.353.353.354c.047.047.109.177.005.488a2.224 2.224 0 0 1-.505.805l-.353.353.353.354c.006.005.041.05.041.17a.866.866 0 0 1-.121.416c-.165.288-.503.56-1.066.56z"/>
                </svg>
              </button>
              <span class="pl-2">{{ review.labreview.lgtm }}件</span>
            </div>
            <!-- ここまで -->
          </div>
          <template v-for="reply in review.replies" :key="reply.ID">
            <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-three-dots-vertical  my-2 ml-5" viewBox="0 0 16 16">
              <path d="M5.921 11.9 1.353 8.62a.719.719 0 0 1 0-1.238L5.921 4.1A.716.716 0 0 1 7 4.719V6c1.5 0 6 0 7 8-2.5-4.5-7-4-7-4v1.281c0 .56-.606.898-1.079.62z"/>
            </svg>
            <div class="border p-2 ml-5 mb-2 shadow-sm">
              <p>{{ reply.body }}</p>
              <span class="text-secondary m-0">返信者: {{ reply.replyer_id }} </span>
              <span class="text-secondary m-0 pl-3">返信日時: {{ formatDate(reply.CreatedAt) }} </span>
            </div>
          </template>
          <form action="" @submit.prevent="submitReply">
            <router-link v-if="!auth" class="pageLink d-inline form-control btn btn-outline-primary" to="/login">
              ログインして返信
            </router-link>
            <div v-else class="form-group ml-2">
              <button class="btn btn-outline-primary form-control w-100" @click="displayReplyForm(review.labreview.ID)" v-bind:id='"reply-disply-button-" + review.labreview.ID'>
                <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-plus-square-dotted" viewBox="0 0 16 16">
                  <path d="M2.5 0c-.166 0-.33.016-.487.048l.194.98A1.51 1.51 0 0 1 2.5 1h.458V0H2.5zm2.292 0h-.917v1h.917V0zm1.833 0h-.917v1h.917V0zm1.833 0h-.916v1h.916V0zm1.834 0h-.917v1h.917V0zm1.833 0h-.917v1h.917V0zM13.5 0h-.458v1h.458c.1 0 .199.01.293.029l.194-.981A2.51 2.51 0 0 0 13.5 0zm2.079 1.11a2.511 2.511 0 0 0-.69-.689l-.556.831c.164.11.305.251.415.415l.83-.556zM1.11.421a2.511 2.511 0 0 0-.689.69l.831.556c.11-.164.251-.305.415-.415L1.11.422zM16 2.5c0-.166-.016-.33-.048-.487l-.98.194c.018.094.028.192.028.293v.458h1V2.5zM.048 2.013A2.51 2.51 0 0 0 0 2.5v.458h1V2.5c0-.1.01-.199.029-.293l-.981-.194zM0 3.875v.917h1v-.917H0zm16 .917v-.917h-1v.917h1zM0 5.708v.917h1v-.917H0zm16 .917v-.917h-1v.917h1zM0 7.542v.916h1v-.916H0zm15 .916h1v-.916h-1v.916zM0 9.375v.917h1v-.917H0zm16 .917v-.917h-1v.917h1zm-16 .916v.917h1v-.917H0zm16 .917v-.917h-1v.917h1zm-16 .917v.458c0 .166.016.33.048.487l.98-.194A1.51 1.51 0 0 1 1 13.5v-.458H0zm16 .458v-.458h-1v.458c0 .1-.01.199-.029.293l.981.194c.032-.158.048-.32.048-.487zM.421 14.89c.183.272.417.506.69.689l.556-.831a1.51 1.51 0 0 1-.415-.415l-.83.556zm14.469.689c.272-.183.506-.417.689-.69l-.831-.556c-.11.164-.251.305-.415.415l.556.83zm-12.877.373c.158.032.32.048.487.048h.458v-1H2.5c-.1 0-.199-.01-.293-.029l-.194.981zM13.5 16c.166 0 .33-.016.487-.048l-.194-.98A1.51 1.51 0 0 1 13.5 15h-.458v1h.458zm-9.625 0h.917v-1h-.917v1zm1.833 0h.917v-1h-.917v1zm1.834-1v1h.916v-1h-.916zm1.833 1h.917v-1h-.917v1zm1.833 0h.917v-1h-.917v1zM8.5 4.5a.5.5 0 0 0-1 0v3h-3a.5.5 0 0 0 0 1h3v3a.5.5 0 0 0 1 0v-3h3a.5.5 0 0 0 0-1h-3v-3z"/>
                </svg>
                返信を追加
              </button>
              <div v-bind:id='"reply-form-" + review.labreview.ID' style="display: none;">
                <textarea v-model="replyBody" class="form-control p-1 my-4" placeholder='回答に返信' required />
                <button class="btn btn-outline-primary form-control w-100" type="submit">返信</button>
              </div>
            </div>
          </form>
        </div>
      </div>
    </div>
  </div>

</template>
 
<script>
import { computed, onMounted, onBeforeUnmount, ref } from 'vue'
import axios from 'axios'
import { useRoute, useRouter } from 'vue-router'
import { useStore } from 'vuex'
import labData from '../data/lab-data.json'
import Loader from "@/components/Loader"
import formatDate from "../functions/formatDate.js"

export default {
  name: "LabReview",
  components: {
    Loader
  },
  setup() {
    const store           = useStore()
    const route           = useRoute()
    const router          = useRouter()
    const auth            = computed(() => store.state.auth)
    const labCode         = route.params.professor // 研究室コード
    const labName         = ref({}) // 研究室名
    const review          = ref({}) // 投稿されているreview
    const replyBody       = ref("") // replyの文章
    const loading         = ref(true) // ロード中であるか(mountedの最後にロード画面を解除)

    // URLから科目を取得
    const lab = labData.find((lab) => lab.code == labCode)
    // labDataの中に一致するlabがない場合は404
    if (lab === undefined) {
      store.dispatch("setIsNotFound", true)
    }
    else {
      labName.value = lab.name
    }

    onMounted(async () => {
      try {
        const labReviewData = await axios.get(
          "/lab/" + route.params.labReviewID
        ).catch(error => {
          console.log(error)
        })
        review.value = labReviewData.data
      
        // ログインしていない場合LGTMボタンをdisabledにする
        if (!store.state.auth) {
          const lgtmButtons = document.getElementsByClassName('lgtm')
          for (let i = 0; i < lgtmButtons.length; i++) {
            lgtmButtons[i].classList.add("disabled")
          }
        }
      } catch (e) {
        console.log(e)
      }
      // ロード画面を解除
      loading.value = false
    })

    const submitReply = async () => {
      try {
        await axios.post("lab/reply/post", {
          lab_review_id : route.params.labReviewID,
          user_id       : localStorage.userID,
          body          : replyBody.value
        })
        // リロード
        router.go({path: router.currentRoute.path, force: true})
      } catch (e) {
        console.log(e)
      }
    }

    // 「返信を追加」ボタンを非表示にし, リプライフォームを表示する
    const displayReplyForm = (answerid) => {
      document.getElementById("reply-disply-button-" + answerid).style.display = "none";
      document.getElementById("reply-form-" + answerid).style.display = "block";
    }

    const updateReviewLgtm = async () => {
      // UIのLGTM数の変更
      if (!review.value.islgtmed) {
        review.value.labreview.lgtm++
      } else {
        review.value.labreview.lgtm--
      }
      review.value.islgtmed = !review.value.islgtmed

      try {
        await axios.put("/lgtm/lab/" + route.params.labReviewID)
      } catch (e) {
        console.log(e)
      }
    }

    // 見出しの処理
    store.dispatch("setJumbotron", labName.value + "のレビュー")
    onBeforeUnmount(() =>
      store.dispatch("setJumbotron", "")
    )

    return {
      auth,
      labCode,
      labName,
      review,
      replyBody,
      loading,
      submitReply,
      displayReplyForm,
      updateReviewLgtm,
      formatDate,
    }
  }
}
</script>