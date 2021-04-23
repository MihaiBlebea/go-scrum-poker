<template>
    <div>
        <h5>Score board</h5>
        <div class="row">
            <div class="col-md-3" v-for="(user, index) in sortedUsers" :key="index">
                <div class="card mb-3">
                    <div class="card-body px-2 py-0 d-flex justify-content-start align-items-center">
                        <Vote :vote="user.vote" />
                        <strong class="ms-3">{{ user.username }}</strong>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
import Vote from './Vote'

export default {
    name: 'ScoreBoard',
    components: {Vote},
    // props: ['users'],
    props: {
        users: {
            type: Array
        }
    },
    data: function() {
        return {
            //
        }
    },
    computed: {
        sortedUsers: function() {
            if (this.users === null || this.users.length === 0) {
                return
            }
            return this.users.sort(this.sort)
        }
    },
    methods: {
        sort: function(a, b) {
            return (a.username > b.username) ? 1 : ((b.username > a.username) ? -1 : 0)
        }
    }
}
</script>

<style scoped>
img {
    max-width: 50px;
    max-height: 50px;
    border-radius: 50%;
}
</style>