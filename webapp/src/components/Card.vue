<template>
    <div class="card">
        <div class=card__animation>
            <div class="card__rank card__rank--upper">{{ rank }}</div>
            <div class="card__suit">{{ suit }}</div>
            <div class="card__rank card__rank--lower">{{ rank }}</div>
        </div>
        <div class="card__vote">{{ vote }}</div>
    </div>
</template>

<script>
export default {
    name: 'Card',
    props: ['rank', 'suit', 'vote'],
}
</script>

<style scoped>
.card {
    margin: 5px;
    float: left;
    position: relative;
    border: solid black;
    width: 65px;
    height: 90px;
    transition: all 0.2s ease-in;
}

.card.♠, .card.♣ {
    color: black;
}

.card.♥, .card.♦ {
    color: red;
}

.card__rank--upper {
    position: absolute;
    top: 3px;
    left: 3px;
}

.card__rank--lower {
    position: absolute;
    bottom: 3px;
    right: 3px;
}

.card__suit {
    font-size: 1.5em;
    position: absolute;
    top: 30%;
    left: calc(50% - 1ch);
}

.card__vote {
    font-size: 45px;
    position: absolute;
    bottom: -20px;
    text-shadow: 3px 2px black;
    font-weight: 900;
    color: cadetblue;
}

.card:hover {
    transform: translate(0px, -10px);
    cursor: pointer;
}
</style>
