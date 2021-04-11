<template>
    <div class="vote">{{ vote }}</div>
</template>

<script>
export default {
    name: 'Vote',
    props: ['vote']
}
</script>

<style scoped>
.vote {
    font-size: 45px;
    text-shadow: 3px 2px black;
    font-weight: 900;
    color: cadetblue;
}
</style>