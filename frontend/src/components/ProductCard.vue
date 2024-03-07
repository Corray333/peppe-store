<template>
    <div class="product_card">
        <img :src="`http://localhost:3000/images/products/${product[0].id +'-'+ colorPick}.jpeg`" alt="">
        <div class="card_controls">
            <div class="card_controrls_header">
                <h3><router-link :to="`/products/${product[0].id}`">{{product[0].name}}</router-link></h3>
                <p>{{price/100}} &#8381;</p>
            </div>
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
    </div>
</template>


<script setup>
import { watch, ref } from 'vue'

const props = defineProps(['product'])
const colorPick = ref()
const sizePick = ref()
const price = ref(0)
const colors = ref(new Set())
const sizes = ref(new Set())
watch(
    ()=>props.product,
    ()=>{
        if (props.product == undefined) return
        for (let type of props.product){
            if (type == undefined) return
            colors.value.add(type.color)
            sizes.value.add(type.size)
        }
        colorPick.value = props.product[0].color
        sizePick.value = props.product[0].size
    },
    {immediate: true, deep: true}
)
watch(
    ()=>colorPick,
    ()=>{
        if (props.product == undefined) return
        for (let type of props.product){
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
        if (props.product == undefined) return
        for (let type of props.product){
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
.product_card{
    min-width:calc((100% - 2em) / 3);
    padding: 25px;
    border-radius: 25px;
    height: fit-content;
    background-color: white;
    flex-grow: 1;
    overflow: hidden;
    position: relative;
}
.product_card>img{
    width: 100%;
    border-radius: 25px;
    transition: all 0.3s;
}
.product_card:hover >img{
    filter: blur(5px);
}
.product_card:hover > .card_controls{
    transform: translateY(0px);
}
.card_controls{
    position: absolute;
    bottom: 0;
    width: calc(100% - 50px);
    display: flex;
    flex-direction: column;
    gap: 15px;
    transition: all 0.3s;
    transform: translateY(200px);
    background-color: white;
    border-radius: 25px 25px 0 0;
    padding: 15px;
    box-shadow: 0 0 10px rgba(0,0,0,0.5)    ;
}
.card_controrls_header h3{
    text-overflow: ellipsis; 
    overflow: hidden; 
    white-space: nowrap;
    /* max-width: 180px; */
    width: 100%;
}
.card_controrls_header p{
    text-wrap: nowrap;
}
.card_controrls_header a{
    color: var(--black-color);
}
.card_controls>button{
    height: fit-content;
    padding: 10px 25px;
    border: none;
    background-color: var(--accent-color);
    border-radius: 50px;
    cursor: pointer;
    transition: all 0.3s;
}
.card_controls>button:hover{
    background-color: var(--black-color);
    color: white;
}

.card_controrls_header{
    display: flex;
    width: 100%;
    justify-content: space-between;
}

.sizes{
    display: flex;
    overflow: hidden;
    gap: 15px;
}
.sizes>*{
    padding: 5px 10px;
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
    overflow: hidden;
    gap: 15px;
}
.colors>*{
    height: 25px;
    width: 25px;
    border-radius: 25px;
    border: 1px solid rgba(0,0,0,0.5);
}

</style>