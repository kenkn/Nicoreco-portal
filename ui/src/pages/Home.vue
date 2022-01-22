<template>
  <div class="container">
    <template v-if="auth">
      <div class="mb-5 position-relative">
        <UserIcon :width="60" :height="60" class="d-inline float-left mr-3" />
        <h1 class="display-5 d-inline">
          ようこそ! {{displayName}}さん
        </h1>
      </div>
      <div class="d-flex justify-content-around text-center">
        <div v-for="menu in menus" v-bind:key="menu.id" class="btn position-relative" @click="currentTarget=menu.id">
          {{menu.name}}
          <div v-if="menu.id==currentTarget" class="position-absolute bottom-0 start-0 bg-danger w-100" style="height: 2px;"/>
        </div>
      </div>
      <hr>
      <!-- ロード画面 -->
      <Loader v-if="loading" />
      <!-- コンテンツ -->
      <div v-else>
        <!-- 投稿一覧(投稿がある場合) -->
        <div v-if="posts!=null">
          <div class="list-group container" ref="listGroup">
            <Post v-for="post in posts" :key="post.id" :post="post"></Post>
          </div>
        </div>
      
        <!-- 投稿がない場合 -->
        <div v-else>
          <div v-if="currentTarget=='question'">
            <h3 class="p-3">
              投稿している質問はありません
            </h3>
            <div class="mb-2 ml-3">
              <router-link to="/question">
                質問投稿はこちら
                <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-arrow-up-right-square" viewBox="0 0 16 16">
                  <path fill-rule="evenodd" d="M15 2a1 1 0 0 0-1-1H2a1 1 0 0 0-1 1v12a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1V2zM0 2a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2v12a2 2 0 0 1-2 2H2a2 2 0 0 1-2-2V2zm5.854 8.803a.5.5 0 1 1-.708-.707L9.243 6H6.475a.5.5 0 1 1 0-1h3.975a.5.5 0 0 1 .5.5v3.975a.5.5 0 1 1-1 0V6.707l-4.096 4.096z"/>
                </svg>
              </router-link>
              </div>
          </div>
          <div v-if="currentTarget=='review'">
            <h3 class="p-3">
              投稿しているレビューはありません
            </h3>
            <div class="mb-2 ml-3">
              <router-link to="/question">
                レビュー投稿はこちら
                <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-arrow-up-right-square" viewBox="0 0 16 16">
                  <path fill-rule="evenodd" d="M15 2a1 1 0 0 0-1-1H2a1 1 0 0 0-1 1v12a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1V2zM0 2a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2v12a2 2 0 0 1-2 2H2a2 2 0 0 1-2-2V2zm5.854 8.803a.5.5 0 1 1-.708-.707L9.243 6H6.475a.5.5 0 1 1 0-1h3.975a.5.5 0 0 1 .5.5v3.975a.5.5 0 1 1-1 0V6.707l-4.096 4.096z"/>
                </svg>
              </router-link>
            </div>
          </div>
        </div>
      </div>
    </template>

    <!-- ログインしていない場合 -->
    <template v-else>
      <div>
        <h1 class="display-5 mb-3">
          ログインしていません
        </h1>
        <div class="mb-2 ml-2">
          <router-link to="/login">
            ログインはこちら
            <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-arrow-up-right-square" viewBox="0 0 16 16">
              <path fill-rule="evenodd" d="M15 2a1 1 0 0 0-1-1H2a1 1 0 0 0-1 1v12a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1V2zM0 2a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2v12a2 2 0 0 1-2 2H2a2 2 0 0 1-2-2V2zm5.854 8.803a.5.5 0 1 1-.708-.707L9.243 6H6.475a.5.5 0 1 1 0-1h3.975a.5.5 0 0 1 .5.5v3.975a.5.5 0 1 1-1 0V6.707l-4.096 4.096z"/>
            </svg>
          </router-link>
        </div>
        <div class="mb-2 ml-2">
          <router-link to="/register">
            新規登録はこちら
            <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-arrow-up-right-square" viewBox="0 0 16 16">
              <path fill-rule="evenodd" d="M15 2a1 1 0 0 0-1-1H2a1 1 0 0 0-1 1v12a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1V2zM0 2a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2v12a2 2 0 0 1-2 2H2a2 2 0 0 1-2-2V2zm5.854 8.803a.5.5 0 1 1-.708-.707L9.243 6H6.475a.5.5 0 1 1 0-1h3.975a.5.5 0 0 1 .5.5v3.975a.5.5 0 1 1-1 0V6.707l-4.096 4.096z"/>
            </svg>
          </router-link>
        </div>
      </div>
    </template>
  </div>
</template>
 
<script>
import { ref, computed, watch, onMounted } from 'vue'
import axios from 'axios'
import { useStore } from 'vuex'
import Loader from "@/components/Loader"
import UserIcon from "@/components/UserIcon"
import Post from "@/components/Post"

export default {
  name: "Home",
  components: {
    UserIcon,
    Post,
    Loader
  },
  setup() {
    const loading     = ref(true) // ロード中であるか
    const store       = useStore()
    const auth        = computed(() => store.state.auth)
    const userId      = localStorage.userID
    const displayName = computed(() => store.state.displayName)
    const menus       = [
      // { "id": "home", "name": "ホーム" },
      { "id": "question", "name": "質問" },
      // { "id": "answer", "name": "回答" },
      { "id": "review", "name": "レビュー" },
      // { "id": "good", "name": "Good" },
    ]
    const currentTarget = ref("question")
    const posts     = ref([]) //質問・回答・レビューの一覧

    const getPosts = async (target) => {
      loading.value = true
      // 質問一覧を取得
      if(target=="question") {
        const questions = ref([]) //質問一覧
        try {
          const url = "questions/1061520060" // TODO 自分の質問を取得するurlに変更
          const { data } = await axios.get(url)
          questions.value = data
          // 取得した質問をpostに格納
          for(let question of questions.value) {
            let post = {
              id: question.ID,
              target: target,
              path: "/question/" + question.subject + "/" + question.ID,
              imagePath: require('@/assets/img/subject_icons/' + question.subject + '.png'),
              title: question.title,
              updatedAt: question.UpdatedAt,
              body: question.body,
              conversationCount: question.answer_count,
              lgtm: question.lgtm,
            }
            posts.value.push(post);
          }
        } catch (e) {
        console.log(e)
        }
      } else if(target=="review") {
       // レビュ一覧を取得 
        try {
          const reviews   = ref([]) //レビュー一覧
          const url = "/lab/reviews/y.ida" // TODO 自分のレビューを取得するurlに変更
          const { data } = await axios.get(url)
          reviews.value = data
          // 取得したレビューをpostに格納
          for(let review of reviews.value) {
            let post = {
              id: review.ID,
              target: target,
              path: "/lab/" + review.lab + "/" + review.ID,
              // imagePath: require('@/assets/img/lab_icons/' + review.labCourse + '.png'), // TODO lab.courseを追加したい
              title: review.title,
              updatedAt: review.UpdatedAt,
              body: review.body,
              // conversationCount: review.reply_count, //TODO レビューの場合は返信の数?
              lgtm: review.lgtm,
            }
            posts.value.push(post);
          }
        } catch (e) {
        console.log(e)
        }
      } else {
        posts.value = []
      }
      loading.value = false
    }

    // userIDが投稿したtargetを取得
    onMounted(async () => {
      getPosts(currentTarget.value)

    })
    // currentTargetが変更された場合データを取得
    watch(() => currentTarget.value, (newCurrentTarget) => {
      posts.value = []
      getPosts(newCurrentTarget)
    })

    return {
      loading,
      auth,
      userId,
      displayName,
      menus,
      currentTarget,
      posts
    }
  }
}
</script>
