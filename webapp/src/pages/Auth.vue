<template>
    <div class="container my-5">
        <div class="row justify-content-center">
            <div class="col-md-6">
                <div class="card">
                    <div class="card-body">
                        <h5 class="card-title">{{ formTitle }}</h5>
                        
                        <div class="form-check form-switch">
                            <input class="form-check-input clickable" type="checkbox" v-model="mode">
                            <label class="form-check-label">Mode</label>
                        </div>
                        
                        <div class="mb-3" v-if="mode">
                            <label class="form-label">Create room</label>
                            <input type="text" class="form-control" v-model="roomName">
                        </div>
                        <div class="mb-3" v-else>
                            <label class="form-label">Join room</label>
                            <input type="text" class="form-control" v-model="roomId">
                        </div>
                        <div class="mb-3">
                            <label class="form-label">Username</label>
                            <input type="text" class="form-control" v-model="username">
                        </div>

                        <button class="btn btn-primary" v-on:click="submit">Create room</button>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
import { api } from './../axios'
import { mapActions } from 'vuex'

export default {
    data() {
        return {
            firebaseUi: null,
            mode: true,
            roomName: '',
            roomId: '',
            username: '',
        }
    },
    computed: {
        formTitle: function() {
            return this.mode ? 'Create room' : 'Join room'
        },
    },
    methods: {
        ...mapActions([
            'setUserId'
        ]),
        // Returns room id or throws error
        createRoom: async function() {
            let result = await api.post('/room', {
                room_name: this.roomName
            })

            if (result.status !== 200) {
                throw Error('Error during the api request')
            }

            return result.data.id
        },
        // Returns user id or throws error
        joinRoom: async function() {
            let result = await api.post('/room/user', {
                room_id: this.roomId,
                username: this.username
            })

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
        
    }
}
</script>

<style scoped>
.clickable {
    cursor: pointer;
}
</style>
