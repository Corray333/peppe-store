<template>
    <div class="wrapper">
        <div class="cart">
            <h1>Your cart</h1>
            <div class="cart_container">
                <div class="cart_list">
                    <CartItem v-for="product in store.state.cart" :key="product.id" :cartItem="product" />
                </div>
                <div class="cart_controls">
                    <h2>Order</h2>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup>
import { useStore } from 'vuex'
import CartItem from '@/components/CartItem.vue'

const store = useStore()



</script>

<style scoped>

.wrapper{
    padding-top: 100px;
}
.cart_container{
    display: grid;
    grid-template-columns: 2fr 1fr;
    gap: var(--global-gap);
}
.cart_list{
    display: flex;
    flex-direction: column;
    gap: var(--global-gap);
}

</style>