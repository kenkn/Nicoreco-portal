import { Commit, createStore } from 'vuex'

export default createStore({
  state: {
    auth: false,
    displayName: '',
    grade: '',
    isNotFound: false
  },
  mutations: {
    setAuth: (state, auth: boolean) => state.auth = auth,
    setDisplayName: (state, displayName: string) => state.displayName = displayName,
    setGrade: (state, grade: string) => state.grade = grade,
    setIsNotFound: (state: { auth: boolean, displayName: string, isNotFound: boolean }, isNotFound: boolean) => state.isNotFound = isNotFound
  },
  actions: {
    // param:
    // * auth - ログイン状態．ログイン済->true，ログアウト->false
    setAuth: ({ commit }: { commit: Commit }, auth: boolean) => commit('setAuth', auth),
    setDisplayName: ({ commit }: { commit: Commit }, displayName: string) => commit('setDisplayName', displayName),
    setGrade: ({ commit }: { commit: Commit }, grade: string) => commit('setGrade', grade),
    setIsNotFound: ({ commit }: { commit: Commit }, isNotFound: boolean) => commit('setIsNotFound', isNotFound)
  },
  modules: {
  }
})
