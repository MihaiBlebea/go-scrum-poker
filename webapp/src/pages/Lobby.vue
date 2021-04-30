<template>
    <div class="container my-5">
        <div class="row justify-content-center">
            <div class="col-md-6">
                <div class="card">
                    <div class="card-body">

                        <h5 class="card-title">Create Room</h5>
                        <div class="mb-3">
                            <label class="form-label">Room name</label>
                            <input type="text" class="form-control" v-model="roomName">
                        </div>

                        <div class="d-flex justify-content-center">
                            <button class="btn btn-primary" v-on:click="submit">Create room</button>
                        </div>
                    
                        <h5 class="card-title">{{ rightTitle }}</h5>

                        <div class="form-check form-switch">
                            <input class="form-check-input clickable" type="checkbox" v-model="mode">
                            <label class="form-check-label">Mode</label>
                        </div>

                        <div class="py-5">
                            <div v-if="mode">
                                <div class="input-group mb-3">
                                    <input type="text" class="form-control" v-model="roomId">
                                    <button class="btn btn-outline-secondary" type="button">Copy</button>
                                </div>
                            </div>
                            <div v-else>
                                <div class="mb-3">
                                    <label class="form-label">Room id</label>
                                    <input type="text" class="form-control" v-model="roomId">
                                </div>

                                <div class="d-flex justify-content-center">
                                    <button class="btn btn-primary" v-on:click="submit">Join</button>
                                </div>
                            </div>
                        </div>

                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
import { api } from '../axios'
import { mapActions, mapGetters } from 'vuex'

export default {
    data() {
        return {
            mode: true,
            roomName: '',
            roomId: '',
            username: '',
        }
    },
    computed: {
        ...mapGetters([
            'loggedIn',
            'user'
        ]),
        formTitle: function() {
            return this.mode ? 'Create room' : 'Join room'
        },
        rightTitle: function() {
            return this.mode ? 'Share room' : 'Join room'
        }
    },
    methods: {
        ...mapActions([
            'setUserId'
        ]),
        // Returns room id or throws error
        createRoom: async function() {
            console.log('TOKeN', this.token)
            let result = await api.post('/room', {
                name: this.roomName
            })

            if (result.status !== 200) {
                throw Error('Error during the api request')
            }

            return result.data.join_code
        },
        // Returns user id or throws error
        joinRoom: async function() {
            let result = await api.post('/room/' + this.roomId, {})

            if (result.status !== 200) {
                throw Error('Error during the api request')
            }

            return result.data.id
        },
        submit: async function() {
            try {
                // This will create room
                if (this.mode === true) {
                    this.roomId = await this.createRoom()
                    return
                }

                let userId = await this.joinRoom()

                this.setUserId({ userId: userId })

                this.$router.push('/room/' + this.roomId)
            } catch(err) {
                console.error(err)
            }
        }
    },
    mounted: function() {
        if (! this.loggedIn) {
            this.$router.replace({ name: 'Login' })
        }
    }
}
</script>

<style scoped>
.clickable {
    cursor: pointer;
}

.border-right {
    border-right: solid 1px rgba(0, 0, 0, 0.125);
}
</style>
