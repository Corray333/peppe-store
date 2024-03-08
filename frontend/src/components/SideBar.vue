<template>
    <div class="sidebar">
        <h2>Filters</h2>
        <div class="filter">
            <p>Sort by:</p>
            <button id="dropdown_btn" type="button" @click="sortShow = !sortShow"><p>{{ sortType }}</p></button>
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
            <input type="text" name="" id="search" placeholder="T-Shirt..." v-model="searchQuery">
        </div>
        <div class="filter">
            <p>Price</p>
            <div class="prices">
                <input type="text" name="" id="" placeholder="min" v-model="minPrice">
                <input type="text" name="" id="" placeholder="max" v-model="maxPrice">
            </div>
        </div>
        <div class="filter">
            <p>Color</p>
            <div class="colors">
                <div 
                    class="color" 
                    v-for="(color, i) of colors" 
                    :key="i" 
                    :style="`background-color: ${color}; border:${colorsPicked.has(color) ? '2px solid var(--accent-color)':'2px solid rgba(0,0,0,0.2)'}`"
                    @click="colorsPicked.has(color) ? colorsPicked.delete(color) : colorsPicked.add(color)"
                ></div>
            </div>
        </div>
        <div class="filter">
            <p>Type</p>
            <button id="dropdown_btn" type="button" @click="typeShow = !typeShow"><p>{{ typePick }}</p></button>
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
        <div class="sizes"></div>
        <button type="button" @click="$emit('search', filters)">Search</button>
    </div>
</template>

<script setup>
import { ref, computed } from 'vue';

const colors = ref([
    'black',
    'white',
    'pink',
])

const searchQuery = ref('')
const minPrice = ref('0')
const maxPrice = ref('100000')
const sortType = ref('Min')
const sortShow = ref(false)
const typePick = ref('All')
const typeShow = ref(false)
const colorsPicked = ref(new Set())

const filters = computed(()=>{
    let f = "?"
    if(searchQuery.value){
        f += `searchQuery=${searchQuery.value}&`
    }
    if(minPrice.value){
        f += `minPrice=${parseInt(minPrice.value)*100}&`
    }
    if(maxPrice.value){
        f += `maxPrice=${parseInt(maxPrice.value)*100}&`
    }
    if(sortType.value){
        f += `sortType=${sortType.value}&`
    }
    if(typePick.value){
        f += `typeQuery=${typePick.value}&`
    }
    if(colorsPicked.value.size > 0){
        f += `color=${Array.from(colorsPicked.value).join(',')}&`
    }
    return f
})


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

#dropdown_btn{
    width: 100%;
    border: 1px solid var(--black-color);
    border-radius: 5px;
    background-color: white;
    z-index: 11;
    position: relative;
    text-align: left;
    padding: 0.5em;
}
#dropdown_btn:hover{
    color: var(--black-color);
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