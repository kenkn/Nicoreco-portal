<template>
  <div>
    <h1 class="p-3">{{ lab.name }}のレビュー</h1>
    <h1 class="p-3">{{ reviews.length }}件のレビュー</h1>
  
    <div v-for="review in reviews" :key="review.ID" class="border p-3 mt-4">
      <p>{{ review.body }}</p>
      <span class="text-secondary m-0">レビュー者: {{ review.user_id }} </span>
      <span class="text-secondary m-0 pl-3">レビュー日時: {{ review.CreatedAt }} </span><br>
      <div class="mt-1">
        <button id="lgtm" class="btn btn-outline-primary">
          参考になった
        </button>
        <span class="pl-2">{{ review.lgtm }}件</span>
      </div>
      <hr>
      <template v-for="reply in replys" :key="reply.ID">
        <div v-if="reply.parent_id==review.ID" class="border p-2 ml-2">
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
    // const reviews = ref({})
    const replys = ref({})
    const replyBody = ref([])

    // テスト用
    const reviews = []
    reviews.push(
        {
            ID: 1,
            body: "ああああ",
            lgtm: 5,
            user_id: "aaa",
            CreatedAt: Date()
        }
    )

    onMounted(async () => {
      try {
        // ラボ名の取得
        for (const d of labData) {
          if (d.code === this.$route.params.professor) {
            lab.value = {
              name: d.name,
              code: d.code
            }
          }
        }

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
      } catch (e) {
        console.log(e)
      }
    })

    const submitReview = async () => {
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