import { Commit, createStore } from 'vuex'

export default createStore({
  state: {
    auth: false,
    displayName: '',
    isNotFound: false
  },
  mutations: {
    setAuth: (state: { auth: boolean, displayName: string, isNotFound: boolean }, auth: boolean) => state.auth = auth,
    setDisplayName: (state: { auth: boolean, displayName: string, isNotFound: boolean }, displayName: string) => state.displayName = displayName,
    setIsNotFound: (state: { auth: boolean, displayName: string, isNotFound: boolean }, isNotFound: boolean) => state.isNotFound = isNotFound
  },
  actions: {
    // param:
    // * auth - ログイン状態．ログイン済->true，ログアウト->false
    setAuth: ({ commit }: { commit: Commit }, auth: boolean) => commit('setAuth', auth),
    setDisplayName: ({ commit }: { commit: Commit }, displayName: string) => commit('setDisplayName', displayName),
    setIsNotFound: ({ commit }: { commit: Commit }, isNotFound: boolean) => commit('setIsNotFound', isNotFound)
  },
  modules: {
  }
})
