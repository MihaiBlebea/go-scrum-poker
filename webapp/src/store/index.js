import Vue from 'vue'
import Vuex from 'vuex'
import firebase from 'firebase'
import router from './../router'
import { api } from '../axios'

Vue.use(Vuex)

const store =  new Vuex.Store({
    state: {
        loggedIn: false,
        user: null,
        token: null
    },
    mutations: {
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
        },
        setToken(state, token) {
            state.token = token
        }
    },
    actions: {
        async handleAuthChange({ commit }, user) {
            commit('setLoggedIn', user !== null)
            if (user) {
                commit('setUser', {
                    displayName: user.displayName,
                    email: user.qemail
                })
                let token = await firebase
                    .auth()
                    .currentUser
                    .getIdToken(true)
                    
                commit('setToken', token)
            } else {
                commit('setUser', null)
                commit('setToken', null)
                router.replace({ name: 'Home' })
            }
        },
        register(context, { username, email, password }) {
            return new Promise((resolve, reject) => {
                firebase
                    .auth()
                    .createUserWithEmailAndPassword(email, password)
                    .then(data => {
                        data.user.updateProfile({
                            displayName: username
                        })

                        return firebase
                            .auth()
                            .currentUser
                            .getIdToken(true)
                    })
                    .then(token => {
                        return api
                            .post('/user', { 
                                username: username, 
                                email: email, 
                                token: token 
                            })
                        })
                    .then((result)=> {
                        resolve(result)
                    })
                    .catch(err => {
                        reject(err)
                    })
            })
        },
        login({ commit }, { email, password }) {
            return new Promise((resolve, reject) => {
                firebase
                    .auth()
                    .signInWithEmailAndPassword(email, password)
                    .then(data => {
                        resolve(data)
                    })
                    .catch(err => {
                        reject(err)
                    })
            })
        },
        async logout(_context) {
            await firebase
                .auth()
                .signOut()
        },
    },
    getters: {
        user(state) {
            return state.user
        },
        loggedIn(state) {
            return state.user !== null
        },
        token(state) {
            return state.token
        }
    }
})

export default store