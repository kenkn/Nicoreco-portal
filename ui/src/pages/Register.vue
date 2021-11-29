<template>
  <div class="form-register">
    <VeeForm @submit="submit" :validation-schema="schema">
      <h1 class="h3 mb-3 fw-normal">ユーザー登録</h1>
      <VeeField
        name="displayName"
        v-model="displayName"
        class="form-control"
        placeholder="表示名"
      />
      <ErrorMessage name="displayName" />

      <VeeField
        name="userID"
        v-model="userID"
        class="form-control"
        placeholder="ユーザID"
      />
      <ErrorMessage name="userID" />

      <VeeField
        name="grade"
        v-model="grade"
        as="select"
        class="form-control">
        <option disabled value="">学年</option>
        <option value="学部1年">学部1年</option>
        <option value="学部2年">学部2年</option>
        <option value="学部3年">学部3年</option>
        <option value="学部4年">学部4年</option>
        <option value="修士1年">修士1年</option>
        <option value="修士2年">修士2年</option>
        <option value="博士1年">博士1年</option>
        <option value="博士2年">博士2年</option>
        <option value="博士3年">博士3年</option>
      </VeeField>
      <ErrorMessage name="grade" as="p" class="m-0"/>

      <VeeField
        name="email"
        v-model="email"
        type="email"
        class="form-control d-inline"
        placeholder="山口大学メールアドレス"
      />
      <span class="d-inline lead borde">
        @yamaguchi-u.ac.jp
      </span>
      <ErrorMessage name="email" />


      <VeeField
        name="password"
        v-model="password"
        type="password"
        class="form-control"
        placeholder="パスワード"
      />
      <ErrorMessage name="password" />
      
      <VeeField
        name="passwordConfirm"
        v-model="passwordConfirm"
        type="password"
        class="form-control"
        placeholder="パスワード(確認)"
      />
      <ErrorMessage name="passwordConfirm" />
      <p>{{ errMessage }}</p>

      <button v-if="!loading" class="w-100 btn-lg btn-primary" type="submit">登録</button>
      <button v-else class="w-100 btn-lg btn-primary disabled" disabled>
        <div class="spinner-border" role="status">
          <span class="sr-only">Loading...</span>
        </div>
      </button>
    </VeeForm>
  </div>
</template>

<script>
import { ref } from 'vue'
import axios from 'axios'
import { useRouter } from 'vue-router'
import { Form as VeeForm, Field as VeeField, ErrorMessage } from 'vee-validate';
import * as yup from 'yup';
export default {
  name: "Register",
  components: {
    VeeForm,
    VeeField,
    ErrorMessage,
  },
  setup() {
    const router          = useRouter()
    const displayName     = ref("")
    const userID          = ref("")
    const grade           = ref("")
    const email           = ref("")
    const password        = ref("")
    const passwordConfirm = ref("")
    const errMessage      = ref("")
    const loading         = ref(false) 
    
    // バリデーション
    const schema = yup.object({
      displayName: yup.string().required("この項目は必須です"),
      userID: yup.string().required("この項目は必須です")
        .matches("^[0-9a-zA-Z]+$", "ユーザーIDは半角英数で入力してください").min(3, "ユーザIDは3文字以上で入力してください"),
      grade: yup.string().required("学年を選択して下さい"),
      email: yup.string().required("この項目は必須です"),
      password: yup.string().required("この項目は必須です").min(8, "パスワードは8文字以上で入力してください"),
      passwordConfirm: yup.string().required("この項目は必須です")
        .oneOf([yup.ref("password")], "パスワードが一致しません")
    });

    const submit = async () => {
      const emailYamaguchi = email.value + "@yamaguchi-u.ac.jp"
      loading.value = true
      try {
         await axios.post("register", {
          display_name     : displayName.value,
          user_id          : userID.value,
          grade            : grade.value,
          email            : emailYamaguchi,
          password         : password.value,
          password_confirm : passwordConfirm.value
        })
        // Login画面に戻る
        router.push("/login")
      } catch (e) {
        loading.value = false
        errMessage.value = e.response.data.message
      }
    }

    return {
      displayName,
      userID,
      grade,
      email,
      password,
      passwordConfirm,
      errMessage,
      schema,
      loading,
      submit
    }
  }
}
</script>

<style>
.form-register {
  width: 100%;
  max-width: 330px;
  padding: 15px;
  margin: auto;
}

.form-register .form-control {
  position: relative;
  box-sizing: border-box;
  height: auto;
  padding: 10px;
  font-size: 16px;
}
.form-register .form-control:focus {
  z-index: 2;
}
.form-register input[name="email"] {
  width: 40%;
  margin-bottom: -1px;
  border-bottom-right-radius: 0;
  border-bottom-left-radius: 0;
}
.form-register input[type="password"] {
  margin-bottom: 10px;
  border-top-left-radius: 0;
  border-top-right-radius: 0;
}
</style>