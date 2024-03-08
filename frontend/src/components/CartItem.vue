<template>
    <div class="cart_item" v-if="product != null">
        <div class="cart_item_info">
            <img :src="`http://localhost:3000/images/products/${product.id + '-' + product.color}.jpeg`" alt="">
            <div class="cart_item_info_text">
                <h3>{{ product.name }}</h3>
                <p>Size: {{ product.size }}</p>
                <p style="display: flex; gap: 1em;">Color:
                    <div class="color" :style="`background-color: ${product.color}`"></div>
                </p>
                <p>Availible: {{ product.quantity }}</p>
            </div>
        </div>
        <div class="cart_item_controls">
            <div class="counter">
                <button @click="cartItem.amount--; store.commit('saveCart')">-</button>
                <input type="text" v-model="cartItem.amount">
                <button @click="cartItem.amount++; store.commit('saveCart')">+</button>
            </div>
            <h3>{{ product.price / 100 }} &#8381;</h3>
            <button id="delete" @click="store.commit('removeFromCart', cartItem)"><img src="../assets/icons/bin.svg" alt=""></button>
        </div>
    </div>
</template>

<script setup>
import { ref, onBeforeMount, watch } from 'vue'
import { useStore } from 'vuex'
import axios from 'axios'

const props = defineProps(['cartItem'])
const product = ref(null)
const store = useStore()

watch(() => props.cartItem.amount, () => {
    if (props.cartItem.amount < 1) {
        props.cartItem.amount = 1
    }
    if (props.cartItem.amount > product.value.quantity) {
        props.cartItem.amount = product.value.quantity
    }
})

const load = async () => {
    const { data } = await axios.get(`http://localhost:3001/products/${props.cartItem.id}/type?size=${props.cartItem.size}&color=${props.cartItem.color}`)
    product.value = data
}
watch(
    ()=>store.state.cart,
    ()=> load(),
    {immediate: true, deep: true}
)



onBeforeMount(()=>load())



</script>

<style scoped>
.cart_item {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 1em;
    border-radius: 1em;
    background-color: white;
}

.cart_item_info {
    display: flex;
    align-items: center;
    gap: 1em;
}

.cart_item img {
    height: 10em
}

.color {
    height: 25px;
    width: 25px;
    border-radius: 25px;
    border: 1px solid rgba(0, 0, 0, 0.5);
}
.counter {
    display: flex;
    align-items: center;
}
.counter button {
    padding: 0.5em;
    width: 50px;
    height: 50px;
    border-radius: 50px 0 0 50px;
}
.counter button:last-child {
    border-radius: 0 50px 50px 0;
}
.counter input {
    width: 50px;
    height: 50px;
    text-align: center;
    border: none;
    background-color: rgba(135, 228, 100, 0.1);
}
.cart_item_controls{
    display: flex;
    flex-direction: column;
    gap: var(--global-gap);
    align-items: end;
}
#delete{
    height: 50px;
    width: 50px;
    display: flex;
    justify-content: center;
    align-items: center;
    background-color: transparent;
}
#delete:hover{
    background-color: transparent;
}
#delete>img{
    height: 100%;
}
</style>