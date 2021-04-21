import Vue from 'vue'
import Router from 'vue-router'
import 'bootstrap/dist/css/bootstrap.css'

import Auth from '@/pages/Auth'
import Room from '@/pages/Room'
import Register from '@/pages/Register'
import Login from '@/pages/Login'

Vue.use(Router)

export default new Router({
    routes: [
        {
            path: '/',
            name: 'Auth',
            component: Auth
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
