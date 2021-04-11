import Vue from 'vue'
import Vuex from 'vuex'
import router from './../router'

Vue.use(Vuex)

const store =  new Vuex.Store({
    state: {
        userId: null
    },
    mutations: {
        clearUserId(state) {
            state.userId = null
        },
        setUserId(state, { userId }) {
            state.userId = userId
        }
    },
    actions: {
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
        }
    }
})

export default store