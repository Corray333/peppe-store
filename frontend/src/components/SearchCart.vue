<template>
    <div class="search-cart">
        <div class="search-container">
            <input ref="inputVar" type="text" class="search" placeholder="Search...">
            <div class="button" @click="inputVar.classList.toggle('opened')"><SearchIcon/></div>
        </div>
        <router-link to="cart" class="button"><CartIcon/><p>{{ store.state.cart.length }}</p></router-link>
    </div>
</template>

<script setup>
import SearchIcon from './icons/SearchIcon.vue';
import CartIcon from './icons/CartIcon.vue';
import { ref } from 'vue'
import { useStore } from 'vuex'

const store = useStore()

const inputVar = ref(null)

</script>

<style scoped>

.search-cart{
    display: flex;
    align-items: center;
    gap: 15px;
    justify-content: flex-end;
}
.search-container{
    height: 50px;
    display: flex;
    position: relative;
}
.opened{
    max-width:250px;
}
input{
    height: 100%;
    padding: 15px;
    border-radius: 50px;
    border: none;
    font-size: 1em;
    width: 100%;
    max-width: 50px;
    transition: all 0.3s ease-in-out;
}
.button{
    border: none;
    background-color: white;
    display: flex;
    align-items: center;
    justify-content: center;
    border-radius: 50px;
    height: 50px;
    width: 50px;
    cursor: pointer;
}

.search-container>.button{
    position: absolute;
    right: 0;
}

a{
    position: relative;
}
a>p{
    position: absolute;
    height: 1.5em;
    width: 1.5em;
    display: flex;
    justify-content: center;
    align-items: center;
    bottom: -5px;
    left: -5px;
    background-color: white;
    border: 2px solid var(--accent-color);
    border-radius: 50px;
}


</style>