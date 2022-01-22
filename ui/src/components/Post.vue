<template>
  <router-link v-bind:to="post.path" class="list-group-item list-group-item-action d-flex position-relative" style="height: 200px;">
    <img :src="post.imagePath" alt="post-logo" class="border-right m-1 post-icon" width="180" height="180">
    <div class="m-2 position-relative w-100" style="overflow: hidden;">
      <div class="fw-bold h3" style="height:35px; white-space: nowrap; overflow: hidden; text-overflow: ellipsis;">
        {{ post.title }}
      </div>
      <div style="display: -webkit-box; -webkit-line-clamp: 3; -webkit-box-orient: vertical; overflow: hidden; text-overflow: ellipsis;">
        {{ post.body }}
      </div>
      
      <!-- サブ情報 -->
      <div class="position-absolute bottom-0 start-0 pb-1">
        <span class="badge bg-light text-dark text-muted mr-1">
          <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-chat-right-text" viewBox="0 0 16 16">
            <path d="M2 1a1 1 0 0 0-1 1v8a1 1 0 0 0 1 1h9.586a2 2 0 0 1 1.414.586l2 2V2a1 1 0 0 0-1-1H2zm12-1a2 2 0 0 1 2 2v12.793a.5.5 0 0 1-.854.353l-2.853-2.853a1 1 0 0 0-.707-.293H2a2 2 0 0 1-2-2V2a2 2 0 0 1 2-2h12z"/>
            <path d="M3 3.5a.5.5 0 0 1 .5-.5h9a.5.5 0 0 1 0 1h-9a.5.5 0 0 1-.5-.5zM3 6a.5.5 0 0 1 .5-.5h9a.5.5 0 0 1 0 1h-9A.5.5 0 0 1 3 6zm0 2.5a.5.5 0 0 1 .5-.5h5a.5.5 0 0 1 0 1h-5a.5.5 0 0 1-.5-.5z"/>
          </svg>
          {{ post.conversationCount }}
        </span>
        <span class="badge bg-light text-dark text-muted">
          <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-hand-thumbs-up" viewBox="0 0 16 16">
            <path d="M8.864.046C7.908-.193 7.02.53 6.956 1.466c-.072 1.051-.23 2.016-.428 2.59-.125.36-.479 1.013-1.04 1.639-.557.623-1.282 1.178-2.131 1.41C2.685 7.288 2 7.87 2 8.72v4.001c0 .845.682 1.464 1.448 1.545 1.07.114 1.564.415 2.068.723l.048.03c.272.165.578.348.97.484.397.136.861.217 1.466.217h3.5c.937 0 1.599-.477 1.934-1.064a1.86 1.86 0 0 0 .254-.912c0-.152-.023-.312-.077-.464.201-.263.38-.578.488-.901.11-.33.172-.762.004-1.149.069-.13.12-.269.159-.403.077-.27.113-.568.113-.857 0-.288-.036-.585-.113-.856a2.144 2.144 0 0 0-.138-.362 1.9 1.9 0 0 0 .234-1.734c-.206-.592-.682-1.1-1.2-1.272-.847-.282-1.803-.276-2.516-.211a9.84 9.84 0 0 0-.443.05 9.365 9.365 0 0 0-.062-4.509A1.38 1.38 0 0 0 9.125.111L8.864.046zM11.5 14.721H8c-.51 0-.863-.069-1.14-.164-.281-.097-.506-.228-.776-.393l-.04-.024c-.555-.339-1.198-.731-2.49-.868-.333-.036-.554-.29-.554-.55V8.72c0-.254.226-.543.62-.65 1.095-.3 1.977-.996 2.614-1.708.635-.71 1.064-1.475 1.238-1.978.243-.7.407-1.768.482-2.85.025-.362.36-.594.667-.518l.262.066c.16.04.258.143.288.255a8.34 8.34 0 0 1-.145 4.725.5.5 0 0 0 .595.644l.003-.001.014-.003.058-.014a8.908 8.908 0 0 1 1.036-.157c.663-.06 1.457-.054 2.11.164.175.058.45.3.57.65.107.308.087.67-.266 1.022l-.353.353.353.354c.043.043.105.141.154.315.048.167.075.37.075.581 0 .212-.027.414-.075.582-.05.174-.111.272-.154.315l-.353.353.353.354c.047.047.109.177.005.488a2.224 2.224 0 0 1-.505.805l-.353.353.353.354c.006.005.041.05.041.17a.866.866 0 0 1-.121.416c-.165.288-.503.56-1.066.56z"/>
          </svg>
          {{ post.lgtm }}
        </span>
        <span>投稿日時: {{ FormatDate(post.updatedAt) }}</span>
      </div>
    </div>
  </router-link>
</template>

<script>
import FormatDate from '@/functions/FormatDate'

export default {
  name: "MyPosts",
  props: {
    // post {
    //   id: String;
    //   target: String;
    //   path: String;
    //   imagePath: String;
    //   title: String;
    //   updatedAt: String;
    //   body: String;
    //   replyCount: Number;
    //   lgtm: Number;
    // }
    post: {
      type: Object
    },
  },
  setup() {
    return{
      FormatDate
    }
  }
}
</script>
<style scoped>
/* スマホ版 */
.post-icon {
  display: none;
}
/* PC版 */
/*画面サイズが768px(BootstrapのMedium)からはここを読み込む*/
@media screen and (min-width:768px) {
  .post-icon {
    display: inline-block;
  }
}
</style>