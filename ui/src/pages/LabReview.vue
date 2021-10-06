<template>
  <div>
    <h1 class="p-3">{{ lab.name }}のレビュー</h1>
    <h1 class="p-3">{{ reviews.length }}件のレビュー</h1>
  
    <div v-for="review in reviews" :key="review.ID" class="border p-3 mt-4">
      <p>{{ review.body }}</p>
      <span class="text-secondary m-0">レビュー者: {{ review.lab_reviewer_id }} </span>
      <span class="text-secondary m-0 pl-3">レビュー日時: {{ review.CreatedAt }} </span><br>
      <div class="mt-1">
        <button id="lgtm" class="btn btn-outline-primary lgtm">
          <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-hand-thumbs-up" viewBox="0 0 16 16">
            <path d="M8.864.046C7.908-.193 7.02.53 6.956 1.466c-.072 1.051-.23 2.016-.428 2.59-.125.36-.479 1.013-1.04 1.639-.557.623-1.282 1.178-2.131 1.41C2.685 7.288 2 7.87 2 8.72v4.001c0 .845.682 1.464 1.448 1.545 1.07.114 1.564.415 2.068.723l.048.03c.272.165.578.348.97.484.397.136.861.217 1.466.217h3.5c.937 0 1.599-.477 1.934-1.064a1.86 1.86 0 0 0 .254-.912c0-.152-.023-.312-.077-.464.201-.263.38-.578.488-.901.11-.33.172-.762.004-1.149.069-.13.12-.269.159-.403.077-.27.113-.568.113-.857 0-.288-.036-.585-.113-.856a2.144 2.144 0 0 0-.138-.362 1.9 1.9 0 0 0 .234-1.734c-.206-.592-.682-1.1-1.2-1.272-.847-.282-1.803-.276-2.516-.211a9.84 9.84 0 0 0-.443.05 9.365 9.365 0 0 0-.062-4.509A1.38 1.38 0 0 0 9.125.111L8.864.046zM11.5 14.721H8c-.51 0-.863-.069-1.14-.164-.281-.097-.506-.228-.776-.393l-.04-.024c-.555-.339-1.198-.731-2.49-.868-.333-.036-.554-.29-.554-.55V8.72c0-.254.226-.543.62-.65 1.095-.3 1.977-.996 2.614-1.708.635-.71 1.064-1.475 1.238-1.978.243-.7.407-1.768.482-2.85.025-.362.36-.594.667-.518l.262.066c.16.04.258.143.288.255a8.34 8.34 0 0 1-.145 4.725.5.5 0 0 0 .595.644l.003-.001.014-.003.058-.014a8.908 8.908 0 0 1 1.036-.157c.663-.06 1.457-.054 2.11.164.175.058.45.3.57.65.107.308.087.67-.266 1.022l-.353.353.353.354c.043.043.105.141.154.315.048.167.075.37.075.581 0 .212-.027.414-.075.582-.05.174-.111.272-.154.315l-.353.353.353.354c.047.047.109.177.005.488a2.224 2.224 0 0 1-.505.805l-.353.353.353.354c.006.005.041.05.041.17a.866.866 0 0 1-.121.416c-.165.288-.503.56-1.066.56z"/>
          </svg>
        </button>
        <span class="pl-2">{{ review.lgtm }}件</span>
      </div>
      <hr>
      <template v-for="reply in replys[review.ID]" :key="reply.ID">
        <div v-if="reply.lab_review_id==review.ID" class="border p-2 ml-2">
          <p>{{ reply.body }}</p>
          <span class="text-secondary m-0">返信者: {{ reply.user_id }} </span>
          <span class="text-secondary m-0 pl-3">返信日時: {{ reply.CreatedAt }} </span><br>
        </div>
      </template>
      <form action="" @submit.prevent="submitReply(review.ID)">
        <router-link v-if="!auth" class="pageLink d-inline form-control btn btn-outline-primary" to="/login">
          ログインして返信
        </router-link>
        <div v-else class="form-group ml-2">
          <button class="btn btn-outline-primary form-control w-100" @click="displayReplyForm(review.ID)" v-bind:id='"reply-disply-button-" + review.ID'>返信を追加</button>
          <div v-bind:id='"reply-form-" + review.ID' style="display: none;">
            <textarea v-model="replyBody[review.ID]" class="form-control p-1 my-4" placeholder='回答に返信' required />
            <button class="btn btn-outline-primary form-control w-100" type="submit">返信</button>
          </div>
        </div>
      </form>
    </div>
    <div>
      <h2 class="p-3">レビューする</h2>
      <form action="" @submit.prevent="submitReview">
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

</template>
 
<script>
import { computed, onMounted, ref } from 'vue'
import axios from 'axios';
import { useStore } from 'vuex';
import labData from '../data/lab-data.json'
export default {
  name: "LabReview",
  data() {
    const store = useStore()
    const auth = computed(() => store.state.auth)
    let lab = ref({})
    const reviewBody = ref("")
    const reviews = ref({})
    const replys = ref([])
    const replyBody = ref([])

    // テスト用
    // const reviews = []
    // reviews.push(
    //     {
    //         ID: 1,
    //         body: "ああああ",
    //         lgtm: 5,
    //         user_id: "aaa",
    //         CreatedAt: Date()
    //     }
    // )

    onMounted(async () => {
      try {
        // // ラボ名の取得
        for (const d of labData) {
          if (d.code === this.$route.params.professor) {
            lab.value = {
              name: d.name,
              code: d.code
            }
          }
        }

        const labReviewData = await axios.get(
          "/lab/reviews/" + this.$route.params.professor
        )
        reviews.value = labReviewData.data

        for (const d in labReviewData.data) {
          const labReplyData = await axios.get(
            "/lab/reply/" + labReviewData.data[d].ID
          )
          replys.value[labReviewData.data[d].ID] = labReplyData.data
        }
        console.log(replys.value[1])
        // 情報の取得
        // const answerData = await axios.get(
        //   "/answer/" + questionData.data.ID
        // )
        // answers.value = answerData.data

        // リプライ情報の取得
        // const replyData = await axios.get(
        //   "/reply/" + questionData.data.ID
        // )
        // replys.value = replyData.data

        // ログインしていない場合LGTMボタンをdisabledにする
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

    const submitReview = async () => {
      try {
        await axios.post("lab/review/post", {
          lab : this.$route.params.professor,
          lab_reviewer_id : localStorage.userID,
          body : reviewBody.value
        })
        // リロード
        // this.$router.go({path: this.$router.currentRoute.path, force: true})
      } catch (e) {
        console.log(e)
      }
      // try {
      //   const userData = await axios.get("user")
      //   await axios.post("answer/post", {
      //     parent_id : this.$route.params.question_id,
      //     user_id : userData.data.user_id,
      //     body : reviewBody.value
      //   })
      //   // リロード
      //   this.$router.go({path: this.$router.currentRoute.path, force: true})
      // } catch (e) {
      //   console.log(e)
      // }
    }

    const submitReply = async (id) => {
      try {
        await axios.post("lab/reply/post", {
          lab_review_id : String(id),
          user_id : localStorage.userID,
          body : replyBody.value[id]
        })
        // リロード
        this.$router.go({path: this.$router.currentRoute.path, force: true})
      } catch (e) {
        console.log(e)
      }
      // try {
      //   const userData = await axios.get("user")
      //   await axios.post("reply/post", {
      //     question_id : this.$route.params.question_id,
      //     parent_id : String(id),
      //     user_id : userData.data.user_id,
      //     body : replyBody.value[id]
      //   })
      //   // リロード
      //   this.$router.go({path: this.$router.currentRoute.path, force: true})
      // } catch (e) {
      //   console.log(e)
      // }
    }

    // 「返信を追加」ボタンを非表示にし, リプライフォームを表示する
    const displayReplyForm = (answerid) => {
      document.getElementById("reply-disply-button-" + answerid).style.display = "none";
      document.getElementById("reply-form-" + answerid).style.display = "block";
    }

    return {
      auth,
      lab,
      reviews,
      replys,
      reviewBody,
      replyBody,
      submitReview,
      submitReply,
      displayReplyForm,
    }
  }
};
</script>