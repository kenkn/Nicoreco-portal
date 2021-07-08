<template>
  <main class="form-forgot">
    <form @submit.prevent="forgot">
      <div v-if="notify.cls" :class="`alert alert-${notify.cls}`" role="alert">
        {{ notify.message }}
      </div>
      <h1 class="h3 mb-3 fw-normal">Please set your email</h1>
      <input
        v-model="email"
        type="email"
        class="form-control mb-3"
        placeholder="Email"
        required
      />

      <button class="w-100 btn-lg btn-primary" type="submit">ログイン</button>
    </form>
  </main>
</template>

<script>
import { reactive, ref } from "vue";
import axios from "axios";

export default {
  name: "Forgot",
  setup() {
    const email = ref("");
    // メール送信結果を格納
    const notify = reactive({
      cls: "",
      message: "",
    });
    const forgot = () => {
      try {
        // forgot apiへリクエスト送信
        axios.post("forgot", {
          email: email.value,
        });
        notify.cls = "success";
        notify.message = "Email was sent!!";
      } catch (e) {
        notify.cls = "danger";
        notify.message = "Email does not exit!!";
      }
    };

    return {
      email,
      notify,
      forgot,
    };
  },
};
</script>

<style>
.form-forgot {
  width: 100%;
  max-width: 330px;
  padding: 15px;
  margin: auto;
}

.form-forgot .form-control {
  position: relative;
  box-sizing: border-box;
  height: auto;
  padding: 10px;
  font-size: 16px;
}
.form-forgot .form-control:focus {
  z-index: 2;
}
.form-forgot input[type="email"] {
  margin-bottom: -1px;
  border-bottom-right-radius: 0;
  border-bottom-left-radius: 0;
}
.form-forgot input[type="password"] {
  margin-bottom: 10px;
  border-top-left-radius: 0;
  border-top-right-radius: 0;
}
</style>