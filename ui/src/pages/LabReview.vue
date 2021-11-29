<template>
  <div>
    <h1 class="display-5">{{ labName }}のレビュー</h1>
    <div class="mt-5 border border-dark bg-white rounded">
      <p class="p-4 display-6 border-bottom border-dark">{{ reviews.length }}件のレビュー</p>
      <div v-for="review in reviews" :key="review.ID" class="border-bottom border-dark p-4 mt-2">
        <div class="border p-3 mb-2 shadow-sm">
          <h4>{{ review.body }}</h4>
          <span class="text-secondary m-0">レビュー者: {{ review.lab_reviewer_id }} </span>
          <span class="text-secondary m-0 pl-3">レビュー日時: {{ review.CreatedAt }} </span>
          <div class="d-none lgtm-info">
            <!-- ここにlgtmボタンの分岐 -->
            <div v-if='!reviewLgtm[review.ID]' class="mt-1">
              <button @click="updateReviewLgtm(review.ID)" id="lgtm" class="btn btn-outline-primary lgtm">
                <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-hand-thumbs-up" viewBox="0 0 16 16">
                  <path d="M8.864.046C7.908-.193 7.02.53 6.956 1.466c-.072 1.051-.23 2.016-.428 2.59-.125.36-.479 1.013-1.04 1.639-.557.623-1.282 1.178-2.131 1.41C2.685 7.288 2 7.87 2 8.72v4.001c0 .845.682 1.464 1.448 1.545 1.07.114 1.564.415 2.068.723l.048.03c.272.165.578.348.97.484.397.136.861.217 1.466.217h3.5c.937 0 1.599-.477 1.934-1.064a1.86 1.86 0 0 0 .254-.912c0-.152-.023-.312-.077-.464.201-.263.38-.578.488-.901.11-.33.172-.762.004-1.149.069-.13.12-.269.159-.403.077-.27.113-.568.113-.857 0-.288-.036-.585-.113-.856a2.144 2.144 0 0 0-.138-.362 1.9 1.9 0 0 0 .234-1.734c-.206-.592-.682-1.1-1.2-1.272-.847-.282-1.803-.276-2.516-.211a9.84 9.84 0 0 0-.443.05 9.365 9.365 0 0 0-.062-4.509A1.38 1.38 0 0 0 9.125.111L8.864.046zM11.5 14.721H8c-.51 0-.863-.069-1.14-.164-.281-.097-.506-.228-.776-.393l-.04-.024c-.555-.339-1.198-.731-2.49-.868-.333-.036-.554-.29-.554-.55V8.72c0-.254.226-.543.62-.65 1.095-.3 1.977-.996 2.614-1.708.635-.71 1.064-1.475 1.238-1.978.243-.7.407-1.768.482-2.85.025-.362.36-.594.667-.518l.262.066c.16.04.258.143.288.255a8.34 8.34 0 0 1-.145 4.725.5.5 0 0 0 .595.644l.003-.001.014-.003.058-.014a8.908 8.908 0 0 1 1.036-.157c.663-.06 1.457-.054 2.11.164.175.058.45.3.57.65.107.308.087.67-.266 1.022l-.353.353.353.354c.043.043.105.141.154.315.048.167.075.37.075.581 0 .212-.027.414-.075.582-.05.174-.111.272-.154.315l-.353.353.353.354c.047.047.109.177.005.488a2.224 2.224 0 0 1-.505.805l-.353.353.353.354c.006.005.041.05.041.17a.866.866 0 0 1-.121.416c-.165.288-.503.56-1.066.56z"/>
                </svg>
              </button>
              <span class="pl-2">{{ reviewLgtmCount[review.ID] }}件</span>
            </div>
            <div @click="updateReviewLgtm(review.ID)" v-else class="mt-1">
              <button id="lgtm" class="btn btn-primary lgtm">
                <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-hand-thumbs-up" viewBox="0 0 16 16">
                  <path d="M8.864.046C7.908-.193 7.02.53 6.956 1.466c-.072 1.051-.23 2.016-.428 2.59-.125.36-.479 1.013-1.04 1.639-.557.623-1.282 1.178-2.131 1.41C2.685 7.288 2 7.87 2 8.72v4.001c0 .845.682 1.464 1.448 1.545 1.07.114 1.564.415 2.068.723l.048.03c.272.165.578.348.97.484.397.136.861.217 1.466.217h3.5c.937 0 1.599-.477 1.934-1.064a1.86 1.86 0 0 0 .254-.912c0-.152-.023-.312-.077-.464.201-.263.38-.578.488-.901.11-.33.172-.762.004-1.149.069-.13.12-.269.159-.403.077-.27.113-.568.113-.857 0-.288-.036-.585-.113-.856a2.144 2.144 0 0 0-.138-.362 1.9 1.9 0 0 0 .234-1.734c-.206-.592-.682-1.1-1.2-1.272-.847-.282-1.803-.276-2.516-.211a9.84 9.84 0 0 0-.443.05 9.365 9.365 0 0 0-.062-4.509A1.38 1.38 0 0 0 9.125.111L8.864.046zM11.5 14.721H8c-.51 0-.863-.069-1.14-.164-.281-.097-.506-.228-.776-.393l-.04-.024c-.555-.339-1.198-.731-2.49-.868-.333-.036-.554-.29-.554-.55V8.72c0-.254.226-.543.62-.65 1.095-.3 1.977-.996 2.614-1.708.635-.71 1.064-1.475 1.238-1.978.243-.7.407-1.768.482-2.85.025-.362.36-.594.667-.518l.262.066c.16.04.258.143.288.255a8.34 8.34 0 0 1-.145 4.725.5.5 0 0 0 .595.644l.003-.001.014-.003.058-.014a8.908 8.908 0 0 1 1.036-.157c.663-.06 1.457-.054 2.11.164.175.058.45.3.57.65.107.308.087.67-.266 1.022l-.353.353.353.354c.043.043.105.141.154.315.048.167.075.37.075.581 0 .212-.027.414-.075.582-.05.174-.111.272-.154.315l-.353.353.353.354c.047.047.109.177.005.488a2.224 2.224 0 0 1-.505.805l-.353.353.353.354c.006.005.041.05.041.17a.866.866 0 0 1-.121.416c-.165.288-.503.56-1.066.56z"/>
                </svg>
              </button>
              <span class="pl-2">{{ reviewLgtmCount[review.ID] }}件</span>
            </div>
            <!-- ここまで -->
          </div>
        </div>
        <template v-for="reply in replys[review.ID]" :key="reply.ID">
          <div v-if="reply.lab_review_id==review.ID">
            <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-three-dots-vertical  my-2 ml-5" viewBox="0 0 16 16">
              <path d="M5.921 11.9 1.353 8.62a.719.719 0 0 1 0-1.238L5.921 4.1A.716.716 0 0 1 7 4.719V6c1.5 0 6 0 7 8-2.5-4.5-7-4-7-4v1.281c0 .56-.606.898-1.079.62z"/>
            </svg>
            <div class="border p-2 ml-5 mb-2 shadow-sm">
              <p>{{ reply.body }}</p>
              <span class="text-secondary m-0">返信者: {{ reply.user_id }} </span>
              <span class="text-secondary m-0 pl-3">返信日時: {{ reply.CreatedAt }} </span>
            </div>
          </div>
        </template>
        <form action="" @submit.prevent="submitReply(review.ID)">
          <router-link v-if="!auth" class="pageLink d-inline form-control btn btn-outline-primary" to="/login">
            ログインして返信
          </router-link>
          <div v-else class="form-group ml-2">
            <button class="btn btn-outline-primary form-control w-100" @click="displayReplyForm(review.ID)" v-bind:id='"reply-disply-button-" + review.ID'>
              <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-plus-square-dotted" viewBox="0 0 16 16">
                <path d="M2.5 0c-.166 0-.33.016-.487.048l.194.98A1.51 1.51 0 0 1 2.5 1h.458V0H2.5zm2.292 0h-.917v1h.917V0zm1.833 0h-.917v1h.917V0zm1.833 0h-.916v1h.916V0zm1.834 0h-.917v1h.917V0zm1.833 0h-.917v1h.917V0zM13.5 0h-.458v1h.458c.1 0 .199.01.293.029l.194-.981A2.51 2.51 0 0 0 13.5 0zm2.079 1.11a2.511 2.511 0 0 0-.69-.689l-.556.831c.164.11.305.251.415.415l.83-.556zM1.11.421a2.511 2.511 0 0 0-.689.69l.831.556c.11-.164.251-.305.415-.415L1.11.422zM16 2.5c0-.166-.016-.33-.048-.487l-.98.194c.018.094.028.192.028.293v.458h1V2.5zM.048 2.013A2.51 2.51 0 0 0 0 2.5v.458h1V2.5c0-.1.01-.199.029-.293l-.981-.194zM0 3.875v.917h1v-.917H0zm16 .917v-.917h-1v.917h1zM0 5.708v.917h1v-.917H0zm16 .917v-.917h-1v.917h1zM0 7.542v.916h1v-.916H0zm15 .916h1v-.916h-1v.916zM0 9.375v.917h1v-.917H0zm16 .917v-.917h-1v.917h1zm-16 .916v.917h1v-.917H0zm16 .917v-.917h-1v.917h1zm-16 .917v.458c0 .166.016.33.048.487l.98-.194A1.51 1.51 0 0 1 1 13.5v-.458H0zm16 .458v-.458h-1v.458c0 .1-.01.199-.029.293l.981.194c.032-.158.048-.32.048-.487zM.421 14.89c.183.272.417.506.69.689l.556-.831a1.51 1.51 0 0 1-.415-.415l-.83.556zm14.469.689c.272-.183.506-.417.689-.69l-.831-.556c-.11.164-.251.305-.415.415l.556.83zm-12.877.373c.158.032.32.048.487.048h.458v-1H2.5c-.1 0-.199-.01-.293-.029l-.194.981zM13.5 16c.166 0 .33-.016.487-.048l-.194-.98A1.51 1.51 0 0 1 13.5 15h-.458v1h.458zm-9.625 0h.917v-1h-.917v1zm1.833 0h.917v-1h-.917v1zm1.834-1v1h.916v-1h-.916zm1.833 1h.917v-1h-.917v1zm1.833 0h.917v-1h-.917v1zM8.5 4.5a.5.5 0 0 0-1 0v3h-3a.5.5 0 0 0 0 1h3v3a.5.5 0 0 0 1 0v-3h3a.5.5 0 0 0 0-1h-3v-3z"/>
              </svg>
              返信を追加
            </button>
            <div v-bind:id='"reply-form-" + review.ID' style="display: none;">
              <textarea v-model="replyBody[review.ID]" class="form-control p-1 my-4" placeholder='回答に返信' required />
              <button class="btn btn-outline-primary form-control w-100" type="submit">返信</button>
            </div>
          </div>
        </form>
      </div>
      <div class="p-4">
        <h1 class="mt-3 display-6 d-inline">レビューする</h1>
        <form action="" @submit.prevent="submitReview" class="p-3">
          <router-link v-if="!auth" class="pageLink d-inline btn btn-outline-primary" to="/login">
            ログインしてレビューを送信
          </router-link>
          <div v-else class="form-group">
            <textarea v-model="reviewBody" class="form-control p-1 my-2 mb-4" rows="4" placeholder='レビューを追加' required />
            <button class="btn btn-outline-primary w-100" type="submit">レビューを送信</button>
          </div>
        </form>
      </div>
    </div>
  </div>

</template>
 
<script>
import { computed, onMounted, ref } from 'vue'
import axios from 'axios'
import { useRoute, useRouter } from 'vue-router'
import { useStore } from 'vuex'
import labData from '../data/lab-data.json'
export default {
  name: "LabReview",
  setup() {
    const store           = useStore()
    const route           = useRoute()
    const router          = useRouter()
    const auth            = computed(() => store.state.auth)
    const labCode         = route.params.professor // 研究室コード
    const labName         = ref({}) // 研究室名
    const reviews         = ref([]) // 投稿されているreviewの集合
    const reviewBody      = ref("") // reviewの文章
    const replys          = ref([]) // 投稿されているreplyの集合
    const replyBody       = ref([]) // replyの文章
    const reviewLgtm      = ref([]) // ユーザがreviewをLGTMしているかどうか
    const reviewLgtmCount = ref([]) // reviewのLGTM数

    onMounted(async () => {
      // URLから研究室を取得
      const lab = labData.find((lab) => lab.code == labCode)
      // labDataの中に一致するlabがない場合は404
      if(lab === undefined){
        store.dispatch("setIsNotFound", true)
      }
      else {
        labName.value = lab.name
        try {
          // 研究室レビューの取得
          const labReviewData = await axios.get(
            "/lab/reviews/" + labCode
          ).catch(error => {
            console.log(error)
          })
          reviews.value = labReviewData.data
          for (const i in labReviewData.data) {
            reviewLgtmCount.value[labReviewData.data[i].ID] = labReviewData.data[i].lgtm
          }
          // 研究室レビューに対するリプライの取得
          for (const i in labReviewData.data) {
            const labReplyData = await axios.get(
              "/lab/reply/" + labReviewData.data[i].ID
            )
            replys.value[labReviewData.data[i].ID] = labReplyData.data
          }
          // LGTM情報の取得
          if (localStorage.isLogin) {
            for (const i in labReviewData.data) {
              const id = labReviewData.data[i].ID
              const reviewLgtmData = await axios.get(
                "/lgtm/lab/" + id + "/" + localStorage.userID
              )
              // ユーザはLGTMしているか?
              if (reviewLgtmData.data.length == 0) {
                reviewLgtm.value[id] = false
              } else {
                reviewLgtm.value[id] = true
              }
            }
          }
          // lgtm情報の表示(lgtmボタン+件数)
          const lgtmInfos = document.getElementsByClassName('lgtm-info')
          for (let i = 0; i < lgtmInfos.length; i++) {
            lgtmInfos[i].classList.remove("d-none")
          }
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
      }
    })

    const submitReview = async () => {
      try {
        await axios.post("lab/review/post", {
          jwt             : localStorage.authToken,
          lab             : labCode,
          lab_reviewer_id : localStorage.userID,
          body            : reviewBody.value
        })
        // リロード
        router.go({path: router.currentRoute.path, force: true})
      } catch (e) {
        console.log(e)
      }
    }

    const submitReply = async (id) => {
      try {
        await axios.post("lab/reply/post", {
          jwt           : localStorage.authToken,
          lab_review_id : String(id),
          user_id       : localStorage.userID,
          body          : replyBody.value[id]
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

    const updateReviewLgtm = async (id) => {
      // UIのLGTM数の変更
      if (!reviewLgtm.value[id]) {
        reviewLgtmCount.value[id]++
      } else {
        reviewLgtmCount.value[id]--
      }
      reviewLgtm.value[id] = !reviewLgtm.value[id]

      try {
        await axios.post("/lgtm/lab", {
          lab_review_id : String(id),
          user_id : localStorage.userID
        })
      } catch (e) {
        console.log(e)
      }
    }

    return {
      auth,
      labName,
      reviews,
      replys,
      reviewBody,
      replyBody,
      reviewLgtm,
      reviewLgtmCount,
      submitReview,
      submitReply,
      displayReplyForm,
      updateReviewLgtm,
    }
  }
};
</script>