<template>
  <div class="container">
    <!-- ロード画面 -->
    <Loader v-if="loading" />
    <!-- コンテンツ -->
    <div v-else>
      <!-- 検索結果の投稿一覧(投稿がある場合) -->
      <div v-if="!(posts.length===0)">
        <div class="list-group" ref="listGroup">
          <Post v-for="post in posts" :key="post.id" :post="post"></Post>
        </div>
      </div>
    
      <!-- 投稿がない場合 -->
      <div v-else>
        <h1 class="p-b3 display-5">質問が見つかりませんでした</h1>
      </div>
    </div>
  </div>
</template>

<script>
import { useRoute } from 'vue-router'
import Post from "@/components/Post"
import { ref, onMounted, onBeforeUnmount } from 'vue'
import axios from 'axios'
import { useStore } from 'vuex'
import Loader from "@/components/Loader"

export default {
  name: "Search",
  components: {
    Post,
    Loader
  },
  setup() {
    const store = useStore()
    const route = useRoute()
    const loading   = ref(true) // ロード中であるか
    const posts     = ref([]) //質問・回答・レビューの一覧
    // TODO 検索処理に変更
    onMounted (async () => {
      loading.value = true
      // 質問一覧を取得
      const questions = ref([]) //質問一覧
      try {
        const url = "/search/" + route.query["q"] // TODO 自分の質問を取得するurlに変更
        const { data } = await axios.get(url)
        questions.value = data
        // 取得した質問をpostに格納
        for (let question of questions.value) {
          let post = {
            id: question.ID,
            target: 'question',
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
      loading.value = false
    })

    // 見出しの処理
    store.dispatch("setJumbotron", "検索結果")
    onBeforeUnmount(() =>
      store.dispatch("setJumbotron", "")
    )

    return {
      loading,
      posts
    }
  }
}
</script>
