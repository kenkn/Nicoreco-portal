<template>
  <div>
    <h1 class="p-3">{{ question.subject }}の質問</h1>
    <div id="question" class="p-3 border">
      <h2>{{ question.title }}</h2>
            <p>
        {{ question.Body }}
      </p>
      <span class="text-secondary m-0">質問者: {{ question.QuestionerID }} </span>
      <span class="text-secondary m-0 pl-3">質問日時: {{ question.created_at }} </span><br>
      <div class="mt-1">
        <button id="lgtm" class="btn btn-outline-primary">
          私も知りたい
        </button>
        <span class="pl-2">{{ question.Lgtm }}件</span>
      </div>
    </div>
    <h1 class="p-3">{{ question.answers }}件の回答</h1>
  
    <div v-for="answer in answers" :key="answer.pk" class="border p-3 mt-4">
      <p>{{answer.Body}}</p>
      <span class="text-secondary m-0">回答者: {{ answer.UserID }} </span>
      <span class="text-secondary m-0 pl-3">回答日時: {{ answer.created_at }} </span><br>
      <div class="mt-1">
        <button id="lgtm" class="btn btn-outline-primary">
          参考になった
        </button>
        <span class="pl-2">{{ question.Lgtm }}件</span>
      </div>
      <hr>
      <template v-for="reply in replys" :key="reply.pk">
        <div v-if="reply.ParentID==answer.pk" class="border p-2 ml-2">
          <p>{{ reply.Body }}</p>
          <span class="text-secondary m-0">返信者: {{ reply.UserID }} </span>
          <span class="text-secondary m-0 pl-3">返信日時: {{ reply.created_at }} </span><br>
          <div class="mt-1">
            <button id="lgtm" class="btn btn-outline-primary">
            参考になった
            </button>
            <span class="pl-2">{{ reply.Lgtm }}件</span>
          </div>
        </div>
      </template>
      <form action="" @submit.prevent="confirm">
        <div class="form-group ml-2">
          <textarea class="form-control p-1 my-4" placeholder='回答に返信' required />
          <button class="btn btn-outline-primary form-control w-100" type="submit">返信</button>

        </div>
      </form>
    </div>
    <div>
      <h2 class="p-3">回答する</h2>
      <form action="" @submit.prevent="confirm">
        <div class="form-group">
        <textarea class="form-control p-1 my-2 mb-4" rows="4" placeholder='回答を追加' required />
        <button class="btn btn-outline-primary w-100" type="submit">回答を送信</button>
        </div>
      </form>
    </div>

  </div>



</template>
 
<script>
export default {
  name: "Question",
  setup() {
    const question = {
      pk: 1,
      QuestionerID: 123456,
      subject: "高度ものづくり創成演習Ⅰ",
      title: "～について",
      Body: "質問の内容ああああああああああああああああああああああああああああああああああああああああああああああああああああああああああ",
      Lgtm: 3,
      answers: 5,
      created_at: Date.now(),
      updated_at: Date(1999, 11, 31, 23, 59, 59),
    }
    // 回答(仮置き)
    const answers = []
    answers.push({
      pk: 1,
      ParentID: "1",
      UserID: "tarou",
      Body: "アンサー1ああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああ",
      Lgtm: 5,
      created_at: Date(2021, 9, 14, 15, 12, 59),
    })
    answers.push({
      pk: 2,
      ParentID: "1",
      UserID: "jirou",
      Body: "アンサー2ああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああ",
      Lgtm: 100,
      created_at: Date(2021, 9, 18, 15, 10, 59),
    })
    // 回答へのリプライ(仮置き)
    const replys = []
    replys.push({
      pk: 1,
      ParentID: "1",
      UserID: "ripupu",
      Body: "アンサー1へのリプライ1",
      Lgtm: 5,
      created_at: Date(2021, 9, 14, 15, 12, 59),
    })
    replys.push({
      pk: 2,
      ParentID: "1",
      UserID: "ouou",
      Body: "アンサー1へのリプライ2",
      Lgtm: 100,
      created_at: Date(2021, 9, 18, 15, 10, 59),
    })
    replys.push({
      pk: 2,
      ParentID: "2",
      UserID: "jirou",
      Body: "アンサー2へのリプライ1",
      Lgtm: 100,
      created_at: Date(2021, 9, 18, 15, 10, 59),
    })

    // 送信しますかの確認
    const confirm = () => {
      if(window.confirm('送信します')) { return true; }
      else { return false; }
    }

    return {
      question,
      answers,
      replys,
      confirm
    }
  }
};
</script>