import Vue from 'vue'
import App from './App'
import router from './router'
import store from './store'
import firebase from 'firebase/app'

Vue.config.productionTip = false

const configOptions = {
    apiKey: "AIzaSyBVoxxvQfBftbYsxqBNmQEvrRdcQ4PRTNw",
    authDomain: "scrumpoker-auth.firebaseapp.com",
    projectId: "scrumpoker-auth",
    storageBucket: "scrumpoker-auth.appspot.com",
    messagingSenderId: "714663771287",
    appId: "1:714663771287:web:85f41ee76902aec47ebdd2",
    measurementId: "G-QG630HYZ4X"
}

firebase.initializeApp(configOptions)

firebase.auth().onAuthStateChanged((user)=> {
    store.dispatch("handleAuthChange", user)
})

new Vue({
    el: '#app',
    router,
    components: { App },
    template: '<App/>',
    store: store,
})
