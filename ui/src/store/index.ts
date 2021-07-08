import { Commit, createStore } from 'vuex'

export default createStore({
  state: {
    auth: false
  },
  mutations: {
    setAuth: (state: { auth: boolean }, auth: boolean) => state.auth = auth
  },
  actions: {
    // param:
    // * auth - ログイン状態．ログイン済->true，ログアウト->false
    setAuth: ({ commit }: { commit: Commit }, auth: boolean) => commit('setAuth', auth)
  },
  modules: {
  }
})
