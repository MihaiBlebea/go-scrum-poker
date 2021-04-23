<template>
    <div v-if="show" class="progress">
        <div 
            class="progress-bar progress-bar-striped bg-danger fill" 
            v-bind:style="style"
            aria-valuenow="100" 
            aria-valuemin="0" 
            aria-valuemax="100"
        ></div>
    </div>
</template>

<script>

export default {
    props: {
        duration: {
            type: Number,
            required: false,
            default: 2
        },
        show: {
            type: Boolean,
            required: false,
            default: false
        }
    },
    data: function() {
        return {
            style: {
                animationDuration: this.duration + 'ms'
            }
        }
    },
    watch: {
        show: function(newVal, _oldVal) {
            if (newVal) {
                setTimeout(()=> {
                    this.$emit('completed')
                }, this.duration)
            }
        }
    }
}
</script>

<style scoped>

.fill {
    animation-name: fillAnimation;
    animation-iteration-count: 1;
    animation-fill-mode: forwards;
}

@keyframes fillAnimation {
    from { width: 0%; }
    to { width: 100%; }
}
</style>
