import Vue from 'vue'
import Vuex from 'vuex'
import router from './../router'

Vue.use(Vuex)

const store =  new Vuex.Store({
    state: {
        userId: null,
        loggedIn: false,
        user: null
    },
    mutations: {
        clearUserId(state) {
            state.userId = null
        },
        setUserId(state, { userId }) {
            state.userId = userId
        },
        setLoggedIn(state, value) {
            switch (value) {
                case true:
                    state.loggedIn = true
                    break
                case false:
                    state.loggedIn = false
                    break
                default:
                    state.loggedIn = false
            }
        },
        setUser(state, user) {
            state.user = user
        }
    },
    actions: {
        fetchUser({ commit }, user) {
            commit('setLoggedIn', user !== null)
            if (user) {
                commit('setUser', {
                    displayName: user.displayName,
                    email: user.email
                })
            } else {
                commit('setUser', null)
            }
        },
        async checkUserId(context) {
            if (! localStorage.getItem("userID")) {
                context.commit('clearUserId')
            } else {
                let userId = localStorage.getItem("userID")
                context.commit('setUserId', { userId })
            }
        },
        setUserId({ _dispatch, commit }, { userId }) {
            localStorage.setItem("userID", userId)
            commit('setUserId', { userId })
        },
        clearUserId(context) {
            localStorage.clear()
            context.commit('clearUserId')
            router.push('/')
        }
    },
    getters: {
        userId(state) {
            return state.userId
        },
        user(state) {
            return state.user
        },
        loggedIn(state) {
            return state.user !== null
        }
    }
})

export default store