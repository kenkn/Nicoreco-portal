<template>
  <div>
    <h1 class="pb-3 display-5 d-md-none">レビューする</h1>
    <div class="p-3 border border-dark bg-white rounded shadow-sm">
      <h2 class="display-6">研究室名:{{ labName }}</h2>
        <VeeForm @submit="submitReview" :validation-schema="schema">
          <div class="form-group">
            <p class="m-0 py-2">レビュー:</p>
            <VeeField name="body" v-model="body" as="textarea" class="form-control p-1" rows="4" required />
            <ErrorMessage name="body" />
            <button class="btn btn-outline-primary w-100 mt-4" type="submit">レビューを送信</button>
          </div>
        </VeeForm>

    </div>
  </div>
</template>
 
<script>
import { ref, onBeforeUnmount } from 'vue'
import axios from 'axios'
import { useRoute, useRouter } from 'vue-router'
import { useStore } from 'vuex'
import { Form as VeeForm, Field as VeeField, ErrorMessage } from 'vee-validate';
import * as yup from 'yup';
import labData from '../data/lab-data.json'
export default {
  name: "CreateQuestion",
  components: {
    VeeForm,
    VeeField,
    ErrorMessage,
  },
  setup() {
    const store       = useStore()
    const route       = useRoute()
    const router      = useRouter()
    const labCode     = route.params.professor // ラボコード
    const labName = ref("") // 科目名
    const title       = ref("")
    const body        = ref("")

    // URLから研究室を取得
    const lab = labData.find((lab) => lab.code == labCode)
    // labDataの中に一致するlabがない場合は404
    if(lab === undefined){
      store.dispatch("setIsNotFound", true)
    }
    else{
      labName.value = lab.name
    }

    // バリデーション
    const schema = yup.object({
      body: yup.string().required("この項目は必須です"),
    });

    const submitReview = async () => {
      if(window.confirm('送信します')) {
        try {
          await axios.post("lab/review/post", {
            lab             : labCode,
            user_id : localStorage.userID,
            body            : body.value
          })
          // レビュー一覧ページへ
          router.push("/lab/" + labCode)
        } catch (e) {
          console.log(e)
        }
      }
    }

    // 見出しの処理
    store.dispatch("setJumbotron", labName.value + "のレビューを投稿")
    onBeforeUnmount(() =>
      store.dispatch("setJumbotron", "")
    )
    
    return {
      labName,
      body,
      schema,
      submitReview
    }
  },
};
</script>
