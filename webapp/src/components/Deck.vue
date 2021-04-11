<template>
    <div>
        <div 
            v-for="(card, index) in deck" 
            :key="'card' + index" 
            v-on:click="voteHandler(card)"
        >
            <Card 
                :class="card.suit"
                :rank="card.rank"
                :suit="card.suit"
                :vote="card.vote"
            />
        </div>
    </div>
</template>

<script>
import axios from 'axios'
import Card from './Card'

export default {
    name: 'Deck',
    components: { Card },
    data() {
        return {
            deck: [],
            ranks: '2 3 4 5 6 7 8 9 10 J Q K A'.split(' '),
            suits: '♠ ♥ ♦ ♣'.split(' '),
            votes: []
        }
    },
    computed: {
        voteRanks: function() {
            if (this.votes.length === 0) {
                return []
            }

            return this.votes.map((vote)=> {
                switch (vote) {
                    case 1:
                    case 11:
                        return 'A'
                    case 12:
                        return 'J'
                    case 13:
                        return 'Q'
                    case 14:
                        return 'K'
                    default:
                        return vote.toString()
                }
            })
        }
    },
    methods: {
        buildDeck: function() {
            let ranks = this.voteRanks
            for (let i = 0; i < ranks.length; i++) {
                this.deck = this.deck.concat({
                    rank: ranks[i],
                    suit: this.getRandomSuit(),
                    vote: this.votes[i]
                })
            }
        },
        getRandomSuit: function() {
            return this.suits[this.random(0, 3)]
        },
        random: function(min, max) {
            min = Math.ceil(min)
            max = Math.floor(max)

            return Math.floor(Math.random() * (max - min + 1)) + min
        },
        obtainVotes: async function() {
            try {
                let result = await axios.get('http://localhost:8080/api/v1/votes')

                return result.data.votes
            } catch(err) {
                console.error(err)

                return []
            }
        },
        voteHandler: function(card) {
            console.log('voted')
            this.$emit('vote', { card })
        }
    },
    mounted: async function() {
        this.votes = await this.obtainVotes()
        this.buildDeck()

    }
}
</script>

