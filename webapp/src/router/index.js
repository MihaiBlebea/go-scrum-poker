import Vue from 'vue'
import Router from 'vue-router'
import 'bootstrap/dist/css/bootstrap.css'

import Lobby from '@/pages/Lobby'
import Room from '@/pages/Room'
import Register from '@/pages/Register'
import Login from '@/pages/Login'
import Home from '@/pages/Home'

Vue.use(Router)

export default new Router({
    routes: [
        {
            path: '/',
            name: 'Home',
            component: Home
        },
        {
            path: '/lobby',
            name: 'Lobby',
            component: Lobby
        },
        {
            path: '/register',
            name: 'Register',
            component: Register
        },
        {
            path: '/login',
            name: 'Login',
            component: Login
        },
        {
            path: '/room/:id',
            name: 'Room',
            component: Room
        }
    ]
})
