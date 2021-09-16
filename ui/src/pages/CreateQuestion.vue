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
            <textarea v-model="Body" class="form-control p-1 my-2 mb-4" rows="4" required />
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
    const QuestionerID = "syouhei"
    const Title = ref("")
    const Body = ref("")
    const Lgtm = 0
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
      // apiに送信
      // await axios.post("register", {
      //   questioner_id    : QuestionerID.value,
      //   subject          : Subject,
      //   title            : Title.value,
      //   body             : Body.value,
      //   lgtm             : Lgtm.value
      // })
        await router.push("/question/" + code + "/")
      }
      else { return false; }
    }  
    return {
      Subject,
      Title,
      Body,
      submit
    }
  },
};
</script>