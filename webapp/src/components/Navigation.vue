<template>
    <nav class="navbar navbar-expand-lg navbar-light bg-light">
        <div class="container-fluid">
            <a class="navbar-brand" href="#">Scrum Poker</a>
            <div class="collapse navbar-collapse" id="navbarSupportedContent">
                <ul class="navbar-nav me-auto mb-2 mb-lg-0">
                    <li class="nav-item">
                        <router-link to="/lobby" class="nav-link">Lobby</router-link>
                    </li>
                </ul>
                <ul class="navbar-nav mb-2 mb-lg-0">
                    <li v-if="userIsLoggedIn" class="nav-item">
                        <a class="nav-link active" href="#">{{ user.displayName }}</a>
                    </li>
                    <li v-if="userIsLoggedIn" class="nav-item">
                        <a class="nav-link" href="#" v-on:click="logout">Logout</a>
                    </li>
                    <li v-if="!userIsLoggedIn" class="nav-item">
                        <router-link to="/login" class="nav-link">Login</router-link>
                    </li>
                    <li v-if="!userIsLoggedIn" class="nav-item">
                        <router-link to="/register" class="nav-link">Register</router-link>
                    </li>
                </ul>
            </div>
        </div>
    </nav>
</template>

<script>
import firebase from 'firebase'
import { mapGetters, mapActions } from 'vuex'

export default {
    data: function() {
        return {

        }
    },
    computed: {
        ...mapGetters([
            'user'
        ]),
        userIsLoggedIn: function() {
            return this.user !== null
        }
    },
    methods: {
        ...mapActions([
            'logout'
        ]),
        handleLogout: function() {
            this.logout().then(() => {
                this.$router.replace({ name: 'Login' })
            }).catch((err) => {
                console.error(err)
            })
        }
    }
}
</script>
