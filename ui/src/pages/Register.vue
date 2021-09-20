<template>
  <main class="form-register">
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

      <select
        v-model="grade"
        class="form-control"
        required
      >
        <option disabled value="">学年</option>
        <option>学部1年</option>
        <option>学部2年</option>
        <option>学部3年</option>
        <option>学部4年</option>
        <option>修士1年</option>
        <option>修士2年</option>
        <option>博士1年</option>
        <option>博士2年</option>
        <option>博士3年</option>
      </select>

      <VeeField
        name="email"
        v-model="email"
        type="email"
        class="form-control"
        placeholder="山口大学メールアドレス"
      />
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

      <button class="w-100 btn btn-lg btn-primary" type="submit">登録</button>
    </VeeForm>
  </main>
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
    const displayName     = ref("")
    const userID          = ref("")
    const grade           = ref("")
    const email           = ref("")
    const password        = ref("")
    const passwordConfirm = ref("")
    const router = useRouter()
    // バリデーション
    const schema = yup.object({
      displayName: yup.string().required("この項目は必須です"),
      userID: yup.string().required("この項目は必須です")
        .matches("^[0-9a-zA-Z]+$", "ユーザーIDは半角英数で入力してください").min(3, "ユーザIDは3文字以上で入力してください"),
      email: yup.string().required("この項目は必須です")
        .matches("^[0-9a-zA-Z]+@yamaguchi-u.ac.jp$", "山口大学のメールアドレスを入力してください"),
      password: yup.string().required("この項目は必須です").min(8, "パスワードは8文字以上で入力してください"),
      passwordConfirm: yup.string().required("この項目は必須です")
        .oneOf([yup.ref("password")], "パスワードが一致しません")
    });

    const submit = async () => {
      // Register apiへPOST
      await axios.post("register", {
        display_name     : displayName.value,
        user_id          : userID.value,
        grade            : grade.value,
        email            : email.value,
        password         : password.value,
        password_confirm : passwordConfirm.value
      })

      // Login画面に戻る
      await router.push("/login")
    }

    return {
      displayName,
      userID,
      grade,
      email,
      password,
      passwordConfirm,
      schema,
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
.form-register input[type="email"] {
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