<template>
    <div class="wrapper">
        <section class="featured_products_block">
            <div class="featured_products_header">
                <p>Featured item</p>
                <div class="controls_block">
                    <button type="button" @click="counter--"><img src="../assets/icons/arrow-left.svg" alt=""></button>
                    <button type="button" @click="counter++"><img src="../assets/icons/arrow-right.svg" alt=""></button>
                </div>
            </div>
            <div class="featured-products-wrapper"> 
                <div class="featured_products" :style="`transform:translateX(calc((-100% - 1em) / 3 * ${counter}))`">
                <ProductCard v-for="(product, i) of products" :key="i" :product="product" />
            </div>
            </div>
        </section>
    </div>
</template>

<script setup>
import ProductCard from './ProductCard.vue'

import {ref , onMounted, watch} from 'vue'
import axios from 'axios'

const counter = ref(0)
watch(counter, () => {
    if(counter.value > products.value.length-3){
        counter.value = 0
    }
    if(counter.value < 0){
        counter.value = 2
    }
})

const products = ref([])
onMounted(async () => {
    const {data} = await axios.get('http://localhost:3001/products/featured')
    console.log(data)
    products.value = data
})

</script>

<style scoped>
.featured_products_wrapper{
    overflow: hidden;
}
.featured_products{
    width: 100%;
    display: flex;
    gap: 1em;
    transform: translateX(0);
    transition: transform 0.5s ease-in-out;
}
.featured_products_block{
    display: flex;
    flex-direction: column;
    gap: 25px;
}
.featured_products_header{
    display: flex;
    justify-content: space-between;
}
.featured_products_header>p{
    color: black;
    padding: 10px 25px;
    background-color: white;
    border-radius: 50px;
    width: fit-content;
    text-transform: uppercase;
    display: flex;
    align-items: center;
}

.controls_block{
    padding: 5px;
    display: flex;
    gap: 10px;
    background-color: white;
    border-radius: 50px;
}
.controls_block>button{
    height: 50px;
    width:  50px;
    border: none;
    background-color: var(--accent-color);
    border-radius: 50%;
    display: flex;
    align-items: center;
    justify-content: center;
    transition: all 0.3s;
}
.controls_block>button:hover{
    box-shadow: inset 0 0 7px rgba(0,0,0,0.5);
}
.controls_block>button:active{
    box-shadow: inset 0 0 13px rgba(0,0,0,0.5);
}


</style>