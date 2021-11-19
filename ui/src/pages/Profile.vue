<template>
  <div>
    <h1 class="h3 mb-3 fw-normal">ユーザー情報確認・変更</h1>
    <VeeForm @submit="submit" :validation-schema="schema">
      <table class="table">
        <tbody>
          <tr>
            <th scope="row">Email</th>
            <td>{{email}}</td>
          </tr>
          <tr>
            <th scope="row">ユーザーID</th>
            <td>{{userID}}</td>
          </tr>
          <tr>
            <th scope="row">表示名</th>
            <td>
              <VeeField
                name="displayName"
                v-model="displayName"
                class="form-control"
                placeholder="表示名"
              />
              <ErrorMessage name="displayName" />
            </td>
          </tr>
          <tr>
            <th scope="row">学年</th>
            <td>
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
            </td>
          </tr>
          <tr>
            <th scope="row">パスワード</th>
            <td><router-link to="/forgot">パスワード変更はこちら</router-link></td>
          </tr>
        </tbody>
      </table>
      <p>{{ errMessage }}</p>
      <button v-if="!loading" class=" btn-lg btn-primary " type="submit">変更</button>
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
import { useStore } from 'vuex'
import { Form as VeeForm, Field as VeeField, ErrorMessage } from 'vee-validate';
import * as yup from 'yup';
export default {
  name: "Profile",
  components: {
    VeeForm,
    VeeField,
    ErrorMessage,
  },
  setup() {
    const email       = localStorage.email
    const userID      = localStorage.userID
    const displayName = ref(localStorage.displayName)
    const grade       = ref(localStorage.grade)
    const errMessage  = ref("")
    let loading       = false
    // バリデーション
    const schema = yup.object({
      displayName: yup.string().required("この項目は必須です"),
      grade: yup.string().required("学年を選択して下さい"),
    });
    // TODO 変更時の処理
    const submit = async () => {
      this.loading = true
      // try {
      // } catch (e) {
      //   this.loading = false
      //   errMessage.value = e.response.data.message
      // }
    }

    return {
      email,
      userID,
      displayName,
      grade,
      schema,
      errMessage,
      loading,
      submit
    }
  }
}
</script>
