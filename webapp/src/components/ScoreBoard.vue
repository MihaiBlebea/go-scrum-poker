<template>
    <div>
        <h5>Score board</h5>
        <div class="row">
            <div class="col-md-3" v-for="(user, index) in users" :key="index">
                <div class="card mb-3">
                    <div class="card-body px-2 py-0 d-flex justify-content-start align-items-center">
                        <Vote :vote="user.vote.vote" />
                        <strong class="ms-3">{{ user.user.username }}</strong>
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
    props: ['users'],
    data: function() {
        return {
            //
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

