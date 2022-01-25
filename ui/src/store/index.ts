import { Commit, createStore } from 'vuex'

export default createStore({
  state: {
    // ナビバーで扱う情報は変更を監視しなければならないのでstoreにも記述
    auth: false,
    displayName: '',
    grade: '',
    jumbotron: '',
    isNotFound: false
  },
  mutations: {
    setAuth: (state, auth: boolean) => state.auth = auth,
    setDisplayName: (state, displayName: string) => state.displayName = displayName,
    setGrade: (state, grade: string) => state.grade = grade,
    setJumbotron: (state, jumbotron: string) => state.jumbotron = jumbotron,
    setIsNotFound: (state, isNotFound: boolean) => state.isNotFound = isNotFound
  },
  actions: {
    // param:
    // * auth - ログイン状態．ログイン済->true，ログアウト->false
    setAuth: ({ commit }: { commit: Commit }, auth: boolean) => commit('setAuth', auth),
    setDisplayName: ({ commit }: { commit: Commit }, displayName: string) => commit('setDisplayName', displayName),
    setGrade: ({ commit }: { commit: Commit }, grade: string) => commit('setGrade', grade),
    setJumbotron: ({ commit }: { commit: Commit }, jumbotron: string) => commit('setJumbotron', jumbotron),
    setIsNotFound: ({ commit }: { commit: Commit }, isNotFound: boolean) => commit('setIsNotFound', isNotFound)
  },
  modules: {
  }
})
