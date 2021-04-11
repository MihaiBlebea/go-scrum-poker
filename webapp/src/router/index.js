import Vue from 'vue'
import Router from 'vue-router'
import 'bootstrap/dist/css/bootstrap.css'

import Auth from '@/pages/Auth'
import Room from '@/pages/Room'

Vue.use(Router)

export default new Router({
    routes: [
        {
            path: '/',
            name: 'Auth',
            component: Auth
        },
        {
            path: '/room/:id',
            name: 'Room',
            component: Room
        }
    ]
})
