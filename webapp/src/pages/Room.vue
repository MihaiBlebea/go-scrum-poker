<template>
    <div class="container my-5">
        <div class="row">
            <div class="col">
                <button class="btn btn-primary" v-on:click="leaveRoom">Leave room</button>
                <h1>Room page</h1>
                <p>Room id: {{ roomId }}</p>
                <p>User id: {{ userId }}</p>
                <p>Turn: {{ turn }}</p>

                <ScoreBoard :users="users" />
                <Deck v-on:vote="sendVote" />

            </div>
        </div>
        <div class="row justify-content-center mt-5">
            <button class="btn btn-primary col-md-3" v-on:click="nextTurn">Next turn</button>
        </div>
    </div>
</template>

<script>
import { api } from '../axios'
import { mapGetters, mapActions } from 'vuex'
import Deck from './../components/Deck'
import ScoreBoard from './../components/ScoreBoard'

export default {
    components: { Deck, ScoreBoard },
    data() {
        return {
            socket: null,
            users: [],
            turn: 1
        }
    },
    computed: {
        ...mapGetters([
            'userId',
        ]),
        roomId: function() {
            return this.$route.params.id
        }
    },
    methods: {
        ...mapActions([
            'checkUserId',
            'clearUserId'
        ]),
        sendVote: async function(ev) {
            try {
                let result = await api.post('/vote', {
                    room_id: this.roomId,
                    user_id: this.userId,
                    vote: ev.card.vote
                })

                console.log(result.data)
            } catch(err) {
                console.error(err)
            }
        },
        getRoomState: async function() {
            try {
                let result = await api.get('/room?room_id=' + this.roomId)

                if (result.status !== 200) {
                    throw Error('Response status is not 200')
                }
                
                this.users = result.data.state.user_state
                this.turn = result.data.state.current_turn
            } catch(err) {
                this.users = []
                console.error(err)
            }
        },
        nextTurn: async function() {
            try {
                let result = await api.post('/room/next', {
                    room_id: this.roomId
                })

                if (result.status !== 200) {
                    throw Error('Response status is not 200')
                }

                this.turn = result.data.turn
            } catch(err) {
                console.error(err)
            }
        },
        leaveRoom: function() {
            this.clearUserId()
        },
        listen: function() {
            if (this.socket === null) {
                return
            }

            this.socket.onmessage = async (event)=> {
                await this.getRoomState()
            }
        }
    },
    mounted: async function() {
        await this.checkUserId()
        
        if (this.userId === null) {
            this.$router.push('/')
        }

        console.log("ROOM ID", this.$route.params.id)
        this.socket = new WebSocket('ws://localhost:8080/ws/' + this.$route.params.id)
        this.listen()

        await this.getRoomState()
    }
}
</script>
