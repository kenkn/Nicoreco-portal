import { Commit, createStore } from 'vuex'

export default createStore({
  state: {
    auth: false,
    displayName: ''
  },
  mutations: {
    setAuth: (state: { auth: boolean, displayName: string }, auth: boolean) => state.auth = auth,
    setDisplayName: (state: { auth: boolean, displayName: string }, displayName: string) => state.displayName = displayName
  },
  actions: {
    // param:
    // * auth - ログイン状態．ログイン済->true，ログアウト->false
    setAuth: ({ commit }: { commit: Commit }, auth: boolean) => commit('setAuth', auth),
    setDisplayName: ({ commit }: { commit: Commit }, displayName: string) => commit('setDisplayName', displayName)
  },
  modules: {
  }
})
