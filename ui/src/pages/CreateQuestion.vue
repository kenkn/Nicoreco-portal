<template>
  <div>
    <h1 class="pb-3 display-5">質問する</h1>
    <div class="p-3 border border-dark bg-white rounded shadow-sm">
      <h2 class="display-6">講義名:{{ subjectName }}</h2>
        <VeeForm @submit="submit" :validation-schema="schema">
          <div class="form-group">
            <span>タイトル:</span>
            <VeeField name="title" v-model="title" class="form-control w-100" required />
            <ErrorMessage name="title" />
            <p class="m-0 py-2">本文:</p>
            <VeeField name="body" v-model="body" as="textarea" class="form-control p-1" rows="4" required />
            <ErrorMessage name="body" />
            <button class="btn btn-outline-primary w-100 mt-4" type="submit">質問を送信</button>
          </div>
        </VeeForm>

    </div>
  </div>
</template>
 
<script>
import { ref } from 'vue'
import axios from 'axios'
import { useRoute, useRouter } from 'vue-router'
import { useStore } from 'vuex'
import { Form as VeeForm, Field as VeeField, ErrorMessage } from 'vee-validate';
import * as yup from 'yup';
import subjectData from '../data/subject-data.json'
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
    const subjectCode = route.params.subject // 授業コード
    const subjectName = ref("") // 科目名
    const title       = ref("")
    const body        = ref("")

    // URLから科目を取得
    const subject = subjectData.find((subject) => subject.code == subjectCode)
    // subjectDataの中に一致するSubjectがない場合は404
    if(subject === undefined){
      store.dispatch("setIsNotFound", true)
    }
    else{
      subjectName.value =subject.name
    }

    // バリデーション
    const schema = yup.object({
      title: yup.string().required("この項目は必須です").max(50, "タイトルは50文字以下で入力してください"),
      body: yup.string().required("この項目は必須です"),
    });

    const submit = async () => {
      if(window.confirm('送信します')) {
        await axios.post("/question/post", {
          jwt           : localStorage.authToken,
          questioner_id : localStorage.userID,
          subject       : route.params.subject,
          title         : title.value,
          body          : body.value,
        })
        await router.push("/question/" + subjectCode)
      }
      else { return false; }
    }  
    
    return {
      subjectName,
      title,
      body,
      schema,
      submit
    }
  },
};
</script>
