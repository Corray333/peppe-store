<template>
  <div class="product-wrapper wrapper">
    <section class="product">
      <img :src="`http://localhost:3000/images/products/${product[0].id +'-' + colorPick}.jpeg`" alt="">
      <div class="product-info">
        <h1>{{product[0].name}}</h1>
        <p>{{product[0].description}}</p>
        <p>{{price/100}} &#8381;</p>
        <div class="card_controls_sizes">
            <p>Select size</p>
            <div class="sizes">
                <p v-for="(size, i) of sizes" :key="i" :style="`border:${size == sizePick ? '2px black solid':'2px solid rgba(0,0,0,0.2)' }`" @click="sizePick = size">{{ size }}</p>
            </div>
        </div>
        <div class="card_controls_colors">
            <p>Select color</p>
            <div class="colors">
                <div v-for="(color, i) of colors" :key="i" :style="`background-color: ${color}; border:${color == colorPick ? '2px black solid':'2px solid rgba(0,0,0,0.2)' }`" @click="colorPick = color"></div>
            </div>
        </div>
        <button type="button">Add to cart</button>
      </div>
    </section>
  </div>
</template>
<script setup>
import { useRoute } from 'vue-router'
import { onBeforeMount, ref, watch } from 'vue'
import axios from 'axios'

const route = useRoute()

const product = ref()

const colorPick = ref()
const sizePick = ref()
const price = ref(0)
const colors = ref(new Set())
const sizes = ref(new Set())

onBeforeMount(async () => {
    const {data} = await axios.get('http://localhost:3001/products/' + route.params.id)
    product.value = data
})
watch(
    ()=>product,
    ()=>{
        if (product.value == undefined) return
        for (let type of product.value){
            if (type == undefined) return
            colors.value.add(type.color)
            sizes.value.add(type.size)
        }
        colorPick.value = product.value[0].color
        sizePick.value = product.value[0].size
    },
    {immediate: true, deep: true}
)
watch(
    ()=>colorPick,
    ()=>{
        if (product.value == undefined) return
        for (let type of product.value){
            if (type == undefined) return
            if (type.color == colorPick.value && type.size == sizePick.value) {
                price.value = type.price
            }
        } 
    },
    {immediate: true, deep: true}
)
watch(
    ()=>sizePick,
    ()=>{
        if (product.value == undefined) return
        for (let type of product.value){
            if (type == undefined) return
            if (type.color == colorPick.value && type.size == sizePick.value) {
                price.value = type.price
            }
        } 
    },
    {immediate: true, deep: true}
)

</script>

<style scoped>
.product-wrapper{
  height: 100vh;
  display: flex;
  align-items: center;
}
.product{
  height: fit-content;
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 25px;
  align-items: start;
  background-color: white;
  border-radius: 25px;
  padding: 25px;
}
.product>img{
  width: 100%;
}
.product-info{
  display: flex;
  flex-direction: column;
  gap: 15px;
}

.sizes{
    display: flex;
    gap: 15px;
}
.sizes>*{
    padding: 5px 15px;
    border-radius: 15px;
    border: 1px solid rgba(0,0,0,0.5);
    color: rgba(0,0,0,0.5);
    cursor: pointer;
    transition: all 0.3s;
}
.sizes>*:hover{
    border: 1px solid var(--black-color);
    color: var(--black-color);
}

.colors{
    display: flex;
    gap: 15px;
}
.colors>*{
    height: 25px;
    width: 25px;
    border-radius: 25px;
    border: 1px solid rgba(0,0,0,0.5);
}

button{
    height: fit-content;
    padding: 10px 25px;
    border: none;
    background-color: var(--accent-color);
    border-radius: 50px;
    cursor: pointer;
    transition: all 0.3s;
}
button:hover{
    background-color: var(--black-color);
    color: white;
}
</style>