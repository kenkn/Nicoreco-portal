<template>
  <main class="col-12 col-xl-6">
    <h1 class="h3 mb-3 fw-normal">ユーザー情報</h1>
    <table class="table">
      <tbody>
        <tr>
          <th scope="row">表示名</th>
          <td>{{displayName}}</td>
        </tr>
        <tr>
          <th scope="row">ユーザーID</th>
          <td>{{userID}}</td>
        </tr>
        <tr>
          <th scope="row">学年</th>
          <td>{{grade}}
          </td>
        </tr>
        <tr>
          <th scope="row">Email</th>
          <td>{{email}}</td>
        </tr>
      </tbody>
    </table>
  </main>
</template>
 
<script>
export default {
  name: "Profile",
  setup() {
    const displayName = localStorage.displayName
    const userID = localStorage.userID
    const grade = localStorage.grade
    const email = localStorage.email

    return {
      displayName,
      userID,
      grade,
      email
    }
  }
}
</script>
