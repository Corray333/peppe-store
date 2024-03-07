<template>
    <div class="sidebar">
        <h2>Filters</h2>
        <div class="filter">
            <p>Sort by:</p>
            <button id="sort_btn" type="button" @click="sortShow = !sortShow"><p>{{ sortType }}</p></button>
            <div class="dropdown-wrapper">
                <div class="dropdown" :class="!sortShow ? 'dropdown_hidden' : ''">
                    <p @click="sortType='Min'; sortShow = !sortShow">Min</p>
                    <p @click="sortType='Max'; sortShow = !sortShow">Max</p>
                    <p @click="sortType='Popularity'; sortShow = !sortShow">Popularity</p>
                </div>
            </div>
        </div>
        <div class="filter">
            <p>Search</p>
            <input type="text" name="" id="search" placeholder="T-Shirt...">
        </div>
        <div class="filter">
            <p>Price</p>
            <div class="prices">
                <input type="text" name="" id="" placeholder="min">
                <input type="text" name="" id="" placeholder="max">
            </div>
        </div>
        <div class="filter">
            <p>Color</p>
            <div class="colors">
                <div class="color" v-for="(color, i) of colors" :key="i" :style="`background-color: ${color}`"></div>
            </div>
        </div>
        <div class="filter">
            <p>Type</p>
            <button id="sort_btn" type="button" @click="typeShow = !typeShow"><p>{{ typePick }}</p></button>
            <div class="dropdown-wrapper">
                <div class="dropdown" :class="!typeShow ? 'dropdown_hidden' : ''">
                    <p @click="typePick='All'; typeShow = !typeShow">All</p>
                    <p @click="typePick='Hoody'; typeShow = !typeShow">Hoody</p>
                    <p @click="typePick='T-Shirt'; typeShow = !typeShow">T-Shirt</p>
                    <p @click="typePick='Band'; typeShow = !typeShow">Band</p>
                </div>
            </div>
        </div>
        <p>Sizes</p>
    </div>
</template>

<script setup>
import { ref } from 'vue';

const colors = ref([
    'black',
    'white',
    'pink',
])

const minPrice = ref(0)
const maxPrice = ref(100000)
const sortType = ref('Min')
const sortShow = ref(false)
const typePick = ref('All')
const typeShow = ref(false)


</script>

<style scoped>
.sidebar {
    width: 100%;
    background-color: white;
    padding: 25px;
    gap: 10px;
    border-radius: 25px;
    display: flex;
    flex-direction: column;
    height: fit-content;
}

h2 {
    text-align: center;
}

.colors {
    display: flex;
    gap: 0.5em;
}

.color {
    height: 25px;
    width: 25px;
    border: 2px solid rgba(0, 0, 0, 0.2);
    border-radius: 50px;
}

input {
    border: 1px solid var(--black-color);
    border-radius: 5px;
    padding: 5px;
    width: 100%;
}

.prices {
    display: flex;
    gap: 10px;
}

#sort_btn{
    width: 100%;
    border: 1px solid var(--black-color);
    border-radius: 5px;
    background-color: white;
    z-index: 100;
    position: relative;
    text-align: left;
    padding: 0.5em;
}
.dropdown-wrapper{
    position: relative;
    width: 100%;
}
.dropdown{
    position: absolute;
    top: -5px;
    left: 0;
    width: 100%;
    background-color: white;
    border: 1px solid var(--black-color);
    border-radius: 0 0 5px 5px;
    display: flex;
    flex-direction: column;
    gap: 0.5em;
    z-index: 10;
    transition: all 0.3s;
    max-height: 200px;
    overflow: hidden;
}
.dropdown p{
    cursor: pointer;
    padding: 0 0.5em;
}
.dropdown p:hover{
    background-color: var(--black-color);
    color: white;
}
.dropdown_hidden{
    max-height: 0;
}

</style>