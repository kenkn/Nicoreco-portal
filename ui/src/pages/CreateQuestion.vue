<template>
  <div>
    <h1 class="p-3">質問する</h1>
    <div class="p-3 border">
      <h2>講義名:{{ Subject }}</h2>
        <form action="" @submit.prevent="submit">
          <div class="form-group">
            <span>タイトル:</span>
            <input v-model="title" type="text" class="form-control w-100" required />
            <p class="m-0 py-2">本文:</p>
            <textarea v-model="body" class="form-control p-1 my-2 mb-4" rows="4" required />
            <button class="btn btn-outline-primary w-100" type="submit">質問を送信</button>
          </div>
        </form>

    </div>
  </div>
</template>
 
<script>
import { ref } from 'vue'
import axios from 'axios'
import { useRouter } from 'vue-router'
import subjectData from '../data/subject-data.json'
export default {
  name: "CreateQuestion",
  data() {
    let code = ""
    let Subject = ""
    const title = ref("")
    const body = ref("")
    const router = useRouter()
    // URLから講義codeを取得しjsonから講義名を取得
    for (const d of subjectData) {
      if (d.code === this.$route.params.subject) {
        Subject = d.name
        code = d.code
      }
    }

    const submit = async () => {
      // 確認(いらんかったら消してください) 
      if(window.confirm('送信します')) {
        await axios.post("/question/post", {
          questioner_id : localStorage.userID,
          subject       : this.$route.params.subject,
          title         : title.value,
          body          : body.value,
        })
        await router.push("/question/" + code)
      }
      else { return false; }
    }  
    return {
      Subject,
      title,
      body,
      submit
    }
  },
};
</script>
