<template>
    <div class="navbar">
        <router-link to="/">Home</router-link>
        <router-link to="/shop">Shop</router-link>
        <router-link to="/about">About</router-link>
    </div>
</template>

<script setup>


</script>

<style scoped>

.navbar{
    display: grid;
    grid-template-columns: 1fr 1fr 1fr;
    padding: 5px;
    background-color: white;
    border-radius: 50px;
}

.navbar a{
    color: var(--black-color);
    text-transform: uppercase;
    text-align: center;
    grid-row: 1;
    padding: 10px;
}
.navbar>a:nth-child(1){
    grid-column: 1;
}
.navbar>a:nth-child(2){
    grid-column: 2;
}
.navbar>a:nth-child(3){
    grid-column: 3;
}
a.router-link-active{
    background-color: var(--black-color);
    color: white;
    border-radius: 50px;
}
</style>