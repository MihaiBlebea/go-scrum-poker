<template>
    <div class="position-relative ms-3">
        <img
            v-bind:style="style"
            class="clip-circle" 
            src="https://lh3.googleusercontent.com/proxy/NjkxAJA1BL-eh6N7tZknvrMKG-OLVSbTXd0ukzs_E1LoQci4DzasMxsU1Tqoc_rALpOQWORxBeFUeHBZNp0La5-UYuZPcj7zxpduIFTxfgOZP_JVjAHwg5-CHIS_ZEBD-UDF2Vvi" />
        <div v-if="reveal" class="vote">{{ vote }}</div>
        <div class="text-center w-100">
            <strong>{{ username }}</strong>
        </div>
    </div>
</template>

<script>
export default {
    name: 'Vote',
    props: {
        vote: {
            type: Number,
            required: true
        },
        username: {
            type: String,
            required: true
        },
        reveal: {
            type: Boolean,
            required: false,
            default: false
        }
    },
    computed: {
        style: function() {
            if (this.vote === 0) {
                return {
                    filter: 'grayscale(100%)'
                }
            }
            
            return {}
        }
    }
}
</script>

<style scoped>
img {
    max-width: 150px;
    max-height: 150px;
}
.vote {
    font-size: 45px;
    text-shadow: 3px 2px black;
    font-weight: 900;
    color: cadetblue;

    position: absolute;
    bottom: 10px;
}
.clip-circle {
    clip-path: circle(60px at center);
}

</style>