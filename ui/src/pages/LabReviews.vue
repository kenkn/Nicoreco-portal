<template>
  <div class="container">
    <!-- ロード画面 -->
    <Loader v-if="loading"></Loader>
    <!-- コンテンツ -->
    <div v-else>
      <div>
        <h1 class="pb-3 d-inline-block display-5 d-md-none">{{ labName }}のレビュー一覧</h1>
        <router-link v-if="!auth" class="pageLink d-inline p-3" to="/login">
          <button type="button" class="btn btn-primary">ログインしてレビューする</button>
        </router-link>
        <router-link v-else class="pageLink d-inline p-3" :to="'/lab/' + labCode + '/create'">
          <button type="button" class="btn btn-primary">レビューする</button>
        </router-link>
      </div>
      <!-- 一覧ページへのリンク -->
      <div class="m-3">
        <router-link to="/lab">
         <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-arrow-left" viewBox="0 0 16 16">
           <path fill-rule="evenodd" d="M15 8a.5.5 0 0 0-.5-.5H2.707l3.147-3.146a.5.5 0 1 0-.708-.708l-4 4a.5.5 0 0 0 0 .708l4 4a.5.5 0 0 0 .708-.708L2.707 8.5H14.5A.5.5 0 0 0 15 8z"/>
         </svg>
         研究室一覧に戻る
        </router-link>
      </div>
      <!-- レビュー一覧(レビューがある場合) -->
      <div v-if="reviews!=null">
        <!-- list -->
        <div class="list-group">
          <router-link v-for="review in reviews" :key="review.ID" v-bind:to="$route.path  + '/' + review.ID" class="list-group-item list-group-item-action p-4">
            <div class="fw-bold text-break pr-3 pb-4">
              {{ review.body }}
            </div>
            <span class="badge bg-light text-dark text-muted mr-1">
              <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-chat-right-text" viewBox="0 0 16 16">
                <path d="M2 1a1 1 0 0 0-1 1v8a1 1 0 0 0 1 1h9.586a2 2 0 0 1 1.414.586l2 2V2a1 1 0 0 0-1-1H2zm12-1a2 2 0 0 1 2 2v12.793a.5.5 0 0 1-.854.353l-2.853-2.853a1 1 0 0 0-.707-.293H2a2 2 0 0 1-2-2V2a2 2 0 0 1 2-2h12z"/>
                <path d="M3 3.5a.5.5 0 0 1 .5-.5h9a.5.5 0 0 1 0 1h-9a.5.5 0 0 1-.5-.5zM3 6a.5.5 0 0 1 .5-.5h9a.5.5 0 0 1 0 1h-9A.5.5 0 0 1 3 6zm0 2.5a.5.5 0 0 1 .5-.5h5a.5.5 0 0 1 0 1h-5a.5.5 0 0 1-.5-.5z"/>
              </svg>
              <!-- TODO リプライの数を表示 -->
              {{ review.replyCount }}
            </span>
            <span class="badge bg-light text-dark text-muted">
              <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-hand-thumbs-up" viewBox="0 0 16 16">
                <path d="M8.864.046C7.908-.193 7.02.53 6.956 1.466c-.072 1.051-.23 2.016-.428 2.59-.125.36-.479 1.013-1.04 1.639-.557.623-1.282 1.178-2.131 1.41C2.685 7.288 2 7.87 2 8.72v4.001c0 .845.682 1.464 1.448 1.545 1.07.114 1.564.415 2.068.723l.048.03c.272.165.578.348.97.484.397.136.861.217 1.466.217h3.5c.937 0 1.599-.477 1.934-1.064a1.86 1.86 0 0 0 .254-.912c0-.152-.023-.312-.077-.464.201-.263.38-.578.488-.901.11-.33.172-.762.004-1.149.069-.13.12-.269.159-.403.077-.27.113-.568.113-.857 0-.288-.036-.585-.113-.856a2.144 2.144 0 0 0-.138-.362 1.9 1.9 0 0 0 .234-1.734c-.206-.592-.682-1.1-1.2-1.272-.847-.282-1.803-.276-2.516-.211a9.84 9.84 0 0 0-.443.05 9.365 9.365 0 0 0-.062-4.509A1.38 1.38 0 0 0 9.125.111L8.864.046zM11.5 14.721H8c-.51 0-.863-.069-1.14-.164-.281-.097-.506-.228-.776-.393l-.04-.024c-.555-.339-1.198-.731-2.49-.868-.333-.036-.554-.29-.554-.55V8.72c0-.254.226-.543.62-.65 1.095-.3 1.977-.996 2.614-1.708.635-.71 1.064-1.475 1.238-1.978.243-.7.407-1.768.482-2.85.025-.362.36-.594.667-.518l.262.066c.16.04.258.143.288.255a8.34 8.34 0 0 1-.145 4.725.5.5 0 0 0 .595.644l.003-.001.014-.003.058-.014a8.908 8.908 0 0 1 1.036-.157c.663-.06 1.457-.054 2.11.164.175.058.45.3.57.65.107.308.087.67-.266 1.022l-.353.353.353.354c.043.043.105.141.154.315.048.167.075.37.075.581 0 .212-.027.414-.075.582-.05.174-.111.272-.154.315l-.353.353.353.354c.047.047.109.177.005.488a2.224 2.224 0 0 1-.505.805l-.353.353.353.354c.006.005.041.05.041.17a.866.866 0 0 1-.121.416c-.165.288-.503.56-1.066.56z"/>
              </svg>
              {{ review.lgtm }}
            </span>
            <span>投稿日時: {{ FormatDate(review.CreatedAt) }}</span>
          </router-link>
        </div>
      </div>
      <!-- レビューがない場合 -->
      <div v-else>
        <h3 class="p-3">
          レビューはまだありません
        </h3>
      </div>
    </div>
  </div>
</template>
 
<script>
import { computed, ref, onMounted, onBeforeUnmount } from 'vue'
import { useRoute } from 'vue-router'
import { useStore } from 'vuex'
import axios from 'axios'
import labData from '../data/lab-data.json'
import Loader from "@/components/Loader"
import FormatDate from '@/functions/FormatDate'

export default {
  name: "labReviews",
  components: {
    Loader
  },
  setup() {
    const loading     = ref(true) // ロード中であるか(mountedの最後にロード画面を解除)
    const store       = useStore()
    const route       = useRoute()
    const auth        = computed(() => store.state.auth)
    const labCode     = route.params.professor // labコード
    const labName     = ref("")
    const reviews     = ref({})

    // URLから科目を取得
    const lab = labData.find((lab) => lab.code == labCode)
    // labDataの中に一致するlabがない場合は404
    if(lab === undefined){
      store.dispatch("setIsNotFound", true)
    }
    else{
      labName.value = lab.name
    }

    // 科目に対する質問一覧を取得
    onMounted(async () => {
      try {
        const url = "/lab/reviews/" + labCode
        const { data } = await axios.get(url)
        reviews.value = data
      } catch (e) {
        console.log(e)
      }
      // ロード画面を解除
      loading.value = false
    })

    // 見出しの処理
    store.dispatch("setJumbotron", labName.value + "のレビュー一覧")
    onBeforeUnmount(() =>
      store.dispatch("setJumbotron", "")
    )
    
    return {
      auth,
      labName,
      labCode,
      reviews,
      loading,
      FormatDate,
    }
  }
}
</script>
